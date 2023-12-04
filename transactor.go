package tester

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/exp/slog"
)

// Result represents the result of a transaction.
type Result struct {
	Batch              int64
	TotalFailedTxCount int64
	TotalTxCount       atomic.Int64
	StartTime          time.Time
	EndTime            time.Time
	MinResponseTime    int64
	MaxResponseTime    int64
}

// BatchResult represents the result of a batch of transactions.
type BatchResult struct {
	batchNo  int64
	payloads []*Payload
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

// SetSegmentStat sets the segmentStat field of the TransactorOpts struct.
//
// It takes a boolean value segmentStat as a parameter and returns a TransactorOpts function.
func SetSegmentStat(segmentStat bool) TransactorOpts {
	return func(t *Transactor) *Transactor {
		t.segmentStat = segmentStat
		return t
	}
}

// Transactor is a struct that can be used to send transactions.
type Transactor struct {
	eth  *ethclient.Client
	pool *Pool
	sync bool

	totalBatch int64
	batchNo    atomic.Int64
	produceTxs atomic.Int64
	endTime    time.Time
	gen        *TxGenerator
	batch      chan *BatchResult
	mu         sync.Mutex

	segmentStat bool
	rs          *Result
	segments    map[int64]*Result

	producerExit atomic.Bool

	exit chan int
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
		batch:    make(chan *BatchResult, 10),
		exit:     make(chan int),
		segments: make(map[int64]*Result),
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
		t.printResult()
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
		if t.stopProducer() {
			t.gen.pool.Close()
			return
		}
		payloads, err := t.gen.Run()
		if err != nil {
			slog.Error("Failed to generate transaction", "err", err)
			continue
		}
		batchNo := t.batchNo.Load()
		slog.Info("produce transactions", "count", len(payloads), "batchNo", batchNo)

		t.batchNo.Add(1)
		t.produceTxs.Add(int64(len(payloads)))
		t.batch <- &BatchResult{
			batchNo:  batchNo,
			payloads: payloads,
		}
	}
}

func (t *Transactor) consumeTx() {
	for batch := range t.batch {
		slog.Info("consume transactions", "batchNo", batch.batchNo)
		t.batchSendTx(context.Background(), batch)
	}
}

func (t *Transactor) listenExit() {
	tick := time.NewTicker(1 * time.Second)
	for {
		<-tick.C
		t.stopProducer()
		t.stopConsumer()
	}
}

func (t *Transactor) stopProducer() bool {
	if t.producerExit.Load() {
		return true
	}

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
		t.producerExit.Store(true)
		return true
	}

	// endTime is a counter that keeps track of the total number of transactions sent
	if !t.endTime.IsZero() && now.Second() >= t.endTime.Second() {
		t.producerExit.Store(true)
		return true
	}
	return false
}

func (t *Transactor) stopConsumer() {
	if t.producerExit.Load() && t.rs.TotalTxCount.Load() == t.produceTxs.Load() {
		t.Stop()
	}
}

// BatchSendTx sends a batch of transactions using the provided context and transaction objects.
func (t *Transactor) batchSendTx(ctx context.Context, batch *BatchResult) {
	if t.sync {
		t.batchSendTxBySync(ctx, batch)
		return
	}

	for _, payload := range batch.payloads {
		txCopy := *payload.Tx
		t.pool.Submit(func() {
			begin := time.Now()
			err := t.eth.SendTransaction(ctx, &txCopy)
			t.Count(batch.batchNo, err, time.Since(begin).Nanoseconds())
		})
	}
	segmentStat := true
	if segmentStat {
		t.pool.Finish()
	}
}

func (t *Transactor) batchSendTxBySync(ctx context.Context, batch *BatchResult) {
	for _, payload := range batch.payloads {
		begin := time.Now()
		err := t.eth.SendTransaction(ctx, payload.Tx)
		t.Count(batch.batchNo, err, time.Since(begin).Nanoseconds())
	}
}

// Count updates the transaction statistics based on the given error and duration.
//
// It takes an error as a parameter to determine if the transaction was successful or not.
// The function also takes an integer value representing the duration of the transaction.
func (t *Transactor) Count(batchNo int64, err error, took int64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	count := func(rs *Result, err error, took int64) {
		rs.Batch = batchNo
		if rs.MinResponseTime > took {
			rs.MinResponseTime = took
		}
		if rs.MaxResponseTime < took {
			rs.MaxResponseTime = took
		}

		rs.TotalTxCount.Add(1)
		if err != nil {
			rs.TotalFailedTxCount++
		}

		if rs.StartTime.IsZero() {
			rs.StartTime = time.Now()
		}
		rs.EndTime = time.Now()
	}

	// totalTxs is a counter that keeps track of the total number of transactions sent
	count(t.rs, err, took)
	// segmented statistics of the results of each batch
	if t.totalBatch > 1 && t.segmentStat {
		rs, ok := t.segments[batchNo]
		if !ok {
			rs = &Result{}
			t.segments[batchNo] = rs
		}
		count(rs, err, took)
	}
}

func (t *Transactor) printResult() {
	header := []string{"BatchNo", "Sample", "Fail", "Transaction/s", "TotalTime", "MinResponseTime", "MaxResponseTime", "AvgResponseTime"}
	formatResult := func(rs *Result) []string {
		totalTxCount := rs.TotalTxCount.Load()
		totalTime := rs.EndTime.Sub(rs.StartTime)
		row := []string{
			strconv.FormatInt(rs.Batch, 10),
			strconv.FormatInt(totalTxCount, 10),
			strconv.FormatInt(rs.TotalFailedTxCount, 10),
			strconv.FormatFloat(float64(totalTxCount-rs.TotalFailedTxCount)/totalTime.Seconds(), 'f', 6, 64),
			rs.EndTime.Sub(rs.StartTime).String(),
			(time.Duration(rs.MinResponseTime) * time.Nanosecond).String(),
			(time.Duration(rs.MaxResponseTime) * time.Nanosecond).String(),
			(time.Duration((rs.MaxResponseTime+rs.MinResponseTime)/2) * time.Nanosecond).String(),
		}
		return row
	}
	if t.totalBatch > 1 && t.segmentStat {
		fmt.Println("Output segmented statistics:")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(header)
		table.SetAutoFormatHeaders(false)

		for batchNo := int64(0); batchNo < t.totalBatch; batchNo++ {
			rs, ok := t.segments[batchNo]
			if !ok {
				continue
			}
			table.Append(formatResult(rs))
		}
		table.Render()
	}

	fmt.Println("Output total statistics:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoFormatHeaders(false)
	table.Append(formatResult(t.rs))
	table.Render()
}
