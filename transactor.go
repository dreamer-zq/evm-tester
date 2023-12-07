package tester

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/exp/slog"
)

const (
	// OneByOne represents sending transactions one by one.
	OneByOne SendMode = "oneByOne"
	// Parallel represents sending transactions in parallel.
	Parallel SendMode = "parallel"
	// Segment represents sending transactions in segments.
	Segment SendMode = "segment"
	// Batch represents sending transactions in batches.
	Batch SendMode = "batch"
)

// SendMode represents the mode of sending transactions.
type SendMode string

// ParseSendMode parses the input string and returns the corresponding SendMode
// constant if it matches one of the defined modes. Otherwise, it returns an
// error.
//
// Parameters:
// - mode: The input string to be parsed.
//
// Return types:
// - SendMode: The corresponding SendMode constant.
// - error: An error if the input string does not match any defined modes.
func ParseSendMode(mode string) (SendMode, error) {
	switch mode {
	case string(OneByOne):
		return OneByOne, nil
	case string(Parallel):
		return Parallel, nil
	case string(Segment):
		return Segment, nil
	case string(Batch):
		return Batch, nil
	default:
		return "", fmt.Errorf("invalid send mode: %s", mode)
	}
}

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

// SetSendMode sets the send mode for the TransactorOpts.
//
// Parameters:
// - sendMode: the send mode to be set.
//
// Returns:
// - a function that sets the send mode and returns the Transactor.
func SetSendMode(sendMode SendMode) TransactorOpts {
	return func(t *Transactor) *Transactor {
		t.sendMode = sendMode
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

	sendMode SendMode
	rs       *Result
	segments map[int64]*Result

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
		batch:    make(chan *BatchResult, 100),
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
		t.sendTx(context.Background(), batch)
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

// sendTx sends a batch of transactions using the provided context and transaction objects.
func (t *Transactor) sendTx(ctx context.Context, batch *BatchResult) {
	defer slog.Info("current executed transaction information",
		"totalBatch", t.totalBatch,
		"currentBatch", batch.batchNo,
		"totalTx", t.rs.TotalTxCount.Load(),
		"failedTx", t.rs.TotalFailedTxCount,
		"startTime", t.rs.StartTime,
		"pool", t.pool.Stat(),
	)
	switch t.sendMode {
	case OneByOne:
		t.sendTxSync(ctx, batch)
		break
	case Segment:
		t.sendTxSegment(ctx, batch)
		break
	case Parallel:
		t.sendTxParallel(ctx, batch)
		break
	case Batch:
		t.batchSendTxs(ctx, batch)
		break
	}
}

func (t *Transactor) sendTxParallel(ctx context.Context, batch *BatchResult) {
	for _, payload := range batch.payloads {
		txCopy := *payload.Tx
		t.pool.Submit(func() {
			begin := time.Now()
			err := t.eth.SendTransaction(ctx, &txCopy)
			t.tally(batch.batchNo, err, time.Since(begin).Nanoseconds())
		})
	}
}

func (t *Transactor) sendTxSegment(ctx context.Context, batch *BatchResult) {
	for _, payload := range batch.payloads {
		txCopy := *payload.Tx
		t.pool.Submit(func() {
			begin := time.Now()
			err := t.eth.SendTransaction(ctx, &txCopy)
			t.tally(batch.batchNo, err, time.Since(begin).Nanoseconds())
		})
	}
	t.pool.Finish()
}

func (t *Transactor) sendTxSync(ctx context.Context, batch *BatchResult) {
	for _, payload := range batch.payloads {
		begin := time.Now()
		err := t.eth.SendTransaction(ctx, payload.Tx)
		t.tally(batch.batchNo, err, time.Since(begin).Nanoseconds())
	}
}

func (t *Transactor) batchSendTxs(ctx context.Context, batch *BatchResult) {
	elems := make([]rpc.BatchElem, 0, len(batch.payloads))
	for _, payload := range batch.payloads {
		data, _ := payload.Tx.MarshalBinary()
		elems = append(elems, rpc.BatchElem{
			Method: "eth_sendRawTransaction",
			Args:   []interface{}{hexutil.Encode(data)},
			Result: new(string),
		})
	}
	t.pool.Submit(func() {
		begin := time.Now()
		err := t.eth.Client().BatchCallContext(ctx, elems)
		if err != nil {
			return
		}
		took := time.Since(begin).Nanoseconds()
		for _, elem := range elems {
			t.tally(batch.batchNo, elem.Error, took)
		}
	})
}

// tally updates the transaction statistics based on the given error and duration.
//
// It takes an error as a parameter to determine if the transaction was successful or not.
// The function also takes an integer value representing the duration of the transaction.
func (t *Transactor) tally(batchNo int64, err error, took int64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	count := func(rs *Result, err error, took int64) {
		rs.Batch = batchNo
		if rs.MinResponseTime > took || rs.MinResponseTime == 0 {
			rs.MinResponseTime = took
		}
		if rs.MaxResponseTime < took {
			rs.MaxResponseTime = took
		}

		rs.TotalTxCount.Add(1)
		if err != nil {
			slog.Error("failed to send transaction", "err", err, "batchNo", batchNo)
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
	if t.totalBatch > 1 && t.sendMode == Segment {
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
	if t.totalBatch > 1 && t.sendMode == Segment {
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
