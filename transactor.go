package tester

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/exp/slog"
)

// Result represents the result of a transaction.
type Result struct {
	TotalFailedTxCount int64
	TotalTxCount       atomic.Int64
	StartTime          time.Time
	EndTime            time.Time
	MinResponseTime    int64
	MaxResponseTime    int64
}

// Println prints the result.
//
// It prints the following information:
// - FailedTx: the total number of failed transactions.
// - TotalTx: the total number of transactions.
// - TotalTime: the total time taken for all transactions.
// - MinResponseTime: the minimum response time among all transactions.
// - MaxResponseTime: the maximum response time among all transactions.
func (r *Result) Println() {
	fmt.Println(
		"FailedTx:", r.TotalFailedTxCount,
		"TotalTx:", r.TotalTxCount.Load(),
		"TotalTime:", time.Duration(r.EndTime.Sub(r.StartTime))*time.Nanosecond,
		"MinResponseTime:", time.Duration(r.MinResponseTime)*time.Nanosecond,
		"MaxResponseTime:", time.Duration(r.MaxResponseTime)*time.Nanosecond,
	)
}

// TransactorOpts is a function that takes in a pointer to a Transactor object
type TransactorOpts func(*Transactor) *Transactor

// SetEndTime sets the end time for a TransactorOpts function.
//
// endTime: the end time to be set (int64).
// Returns: a function that sets the end time for a Transactor (TransactorOpts).
func SetEndTime(endTime time.Time) TransactorOpts {
	return func(t *Transactor) *Transactor {
		t.endTime = endTime
		return t
	}
}

// SetSync sets the sync flag of a Transactor.
//
// sync: a boolean indicating whether the Transactor should sync.
// Returns: a TransactorOpts function that sets the sync flag of a Transactor.
func SetSync(sync bool) TransactorOpts {
	return func(t *Transactor) *Transactor {
		t.sync = sync
		return t
	}
}


// SetTotalBatch sets the total batch value for the TransactorOpts.
//
// totalBatch: The total batch value to be set.
// Returns: The modified TransactorOpts.
func SetTotalBatch(totalBatch int64) TransactorOpts {
	return func(t *Transactor) *Transactor {
		t.totalBatch = totalBatch
		return t
	}
}

// Transactor is a struct that can be used to send transactions.
type Transactor struct {
	eth            *ethclient.Client
	pool           *Pool
	sync           bool
	rs             *Result
	totalBatch     int64
	batchNo        atomic.Int64
	endTime        time.Time
	gen            *TxGenerator
	payloads       chan []*Payload
	exit           chan int
	contractParams []interface{}
	mu             sync.Mutex
}

// NewTransactor creates a new Transactor instance.
//
// It takes in an ethclient.Client pointer, a Pool pointer, and a TxGenerator pointer as parameters.
// It returns a pointer to a Transactor.
func NewTransactor(eth *ethclient.Client, maxConcurrentNum int, gen *TxGenerator, opts ...TransactorOpts) *Transactor {
	transactor := &Transactor{
		eth:      eth,
		pool:     NewPool(maxConcurrentNum, "transactor"),
		rs:       &Result{},
		gen:      gen,
		payloads: make(chan []*Payload, 10),
		exit:     make(chan int),
	}
	for _, opt := range opts {
		transactor = opt(transactor)
	}
	return transactor
}

// Run runs the Transactor.
//
// It starts the producer and consumer goroutines to handle transaction processing.
// The function waits for a signal on the exit channel. If the signal is received,
// it prints statistics about the transaction processing, closes the transaction pool,
// and returns.
func (t *Transactor) Run() {
	go t.listenExit()
	go t.produceTx()
	go t.consumeTx()

	select {
	case <-t.exit:
		t.pool.Close()
		t.rs.Println()
		return
	}
}

// Stop stops the Transactor.
//
// No parameters.
// No return types.
func (t *Transactor) Stop() {
	t.exit <- 1
}

func (t *Transactor) produceTx() {
	for {
		// totalTxs is a counter that keeps track of the total number of transactions sent
		if t.canFinish() {
			t.gen.pool.Close()
			return
		}
		payloads, err := t.gen.Run()
		if err != nil {
			slog.Error("Failed to generate transaction", "err", err)
			continue
		}
		slog.Info("produce transactions", "count", len(payloads))
		t.batchNo.Add(1)
		t.payloads <- payloads
	}
}

func (t *Transactor) consumeTx() {
	for payloads := range t.payloads {
		slog.Info("consume transactions", "count", len(payloads))
		t.batchSendTx(context.Background(), payloads)
	}
}

func (t *Transactor) listenExit() {
	tick := time.NewTicker(1 * time.Second)
	for {
		<-tick.C
		if t.canFinish() {
			slog.Info("stop transactor")
			t.Stop()
			return
		}
	}
}

func (t *Transactor) canFinish() bool {
	batchNo := t.batchNo.Load()
	now := time.Now()

	slog.Info("can finish transactor ?",
		"totalBatch", t.totalBatch,
		"current", batchNo,
		"now", now,
		"endTime", t.endTime,
	)
	// totalTxs is a counter that keeps track of the total number of transactions sent
	if t.totalBatch > 0 && t.totalBatch <= batchNo {
		return true
	}

	// endTime is a counter that keeps track of the total number of transactions sent
	if !t.endTime.IsZero() && now.Second() >= t.endTime.Second() {
		return true
	}
	return false
}

// BatchSendTx sends a batch of transactions using the provided context and transaction objects.
func (t *Transactor) batchSendTx(ctx context.Context, payloads []*Payload) {
	if t.sync {
		t.batchSendTxBySync(ctx, payloads)
		return
	}

	for _, payload := range payloads {
		txCopy := *payload.Tx
		t.pool.Submit(func() {
			begin := time.Now()
			err := t.eth.SendTransaction(ctx, &txCopy)
			t.Count(err, time.Since(begin).Nanoseconds())
		})
	}
	// t.pool.Finish()
}

func (t *Transactor) batchSendTxBySync(ctx context.Context, payloads []*Payload) {
	for _, payload := range payloads {
		begin := time.Now()
		err := t.eth.SendTransaction(ctx, payload.Tx)
		t.Count(err, time.Since(begin).Nanoseconds())
	}
}

// Count updates the transaction statistics based on the given error and duration.
//
// It takes an error as a parameter to determine if the transaction was successful or not.
// The function also takes an integer value representing the duration of the transaction.
func (t *Transactor) Count(err error, took int64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.rs.MinResponseTime > took {
		t.rs.MinResponseTime = took
	}
	if t.rs.MaxResponseTime < took {
		t.rs.MaxResponseTime = took
	}

	t.rs.TotalTxCount.Add(1)
	if err != nil {
		t.rs.TotalFailedTxCount++
	}

	if t.rs.StartTime.IsZero() {
		t.rs.StartTime = time.Now()
	}
	t.rs.EndTime = time.Now()
}
