package tester

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/olekukonko/tablewriter"
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

type tallyItem struct {
	txHash  common.Hash
	batchNo int64
	err     error
	verify  Verify
	took    int64
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

	totalBatch int64
	batchNo    atomic.Int64
	produceTxs atomic.Int64
	endTime    time.Time
	gen        *TxGenerator
	batch      chan *BatchResult
	tallyCh    chan *tallyItem
	mu         sync.Mutex
	verifer    *Verifier

	sendMode SendMode
	rs       *Result
	segments map[int64]*Result

	producerExit atomic.Bool
	consumerExit atomic.Bool

	exit chan int
}

// NewTransactor creates a new Transactor instance.
//
// It takes in an ethclient.Client pointer, a Pool pointer, and a TxGenerator pointer as parameters.
// It returns a pointer to a Transactor.
func NewTransactor(eth *ethclient.Client, maxConcurrentNum int, gen *TxGenerator, enable bool, opts ...TransactorOpts) *Transactor {
	transactor := &Transactor{
		eth:      eth,
		pool:     NewPool(maxConcurrentNum, "transactor"),
		rs:       &Result{},
		gen:      gen,
		batch:    make(chan *BatchResult, 1000),
		tallyCh:  make(chan *tallyItem, 5000),
		exit:     make(chan int),
		segments: make(map[int64]*Result),
	}
	for _, opt := range opts {
		transactor = opt(transactor)
	}
	transactor.verifer = NewVerifier(enable, transactor.eth)
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
	go t.startTally()
	go t.consumeTx()
	go t.verifer.Start(t.sendMode == Parallel)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-t.exit:
		t.Exit()
		return
	case <-sigs: // wait for sigs to quit exit
		t.Exit()
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

// Exit closes the batch and tallyCh channels, closes the pool, and prints the result.
//
// No parameters.
// No return type.
func (t *Transactor) Exit() {
	t.pool.Close()
	t.printResult()
	close(t.batch)
	close(t.tallyCh)
}

func (t *Transactor) produceTx() {
	for {
		// totalTxs is a counter that keeps track of the total number of transactions sent
		if t.stopProducer() {
			t.gen.pool.Close()
			return
		}
		payloads, exit, err := t.gen.Run()
		if err != nil {
			slog.Error("failed to generate transactions", "err", err)
			break
		}

		if exit {
			t.producerExit.Store(true)
		}

		if len(payloads) == 0 {
			break
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

func (t *Transactor) startTally() {
	for item := range t.tallyCh {
		if item.err == nil {
			t.verifer.Add(item.txHash, item.verify)
		}
		t.tally(item.batchNo, item.err, item.took)
	}
}

func (t *Transactor) listenExit() {
	tick := time.NewTicker(1 * time.Second)
	for {
		<-tick.C
		t.stopProducer()
		t.stopConsumer()
		t.stopVerifier()
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
	if t.producerExit.Load() &&
		t.rs.TotalTxCount.Load() == t.produceTxs.Load() &&
		len(t.tallyCh) == 0 {
		t.consumerExit.Store(true)
	}
}

func (t *Transactor) stopVerifier() {
	if t.producerExit.Load() &&
		t.consumerExit.Load() &&
		t.verifer.Finish(t.rs.TotalTxCount.Load()) {
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
		t.sendTxsSync(ctx, batch)
		break
	case Segment:
		t.sendTxsSegment(ctx, batch)
		break
	case Parallel:
		t.sendTxsParallel(ctx, batch)
		break
	case Batch:
		t.sendTxsBatch(ctx, batch)
		break
	}
}

func (t *Transactor) sendTxsParallel(ctx context.Context, batch *BatchResult) {
	for _, payload := range batch.payloads {
		txCopy := *payload.Tx
		t.pool.Submit(func() {
			begin := time.Now()
			err := t.eth.SendTransaction(ctx, &txCopy)
			t.tallyCh <- &tallyItem{txCopy.Hash(), batch.batchNo, err, payload.VerifyFn, time.Since(begin).Nanoseconds()}
		})
	}
}

func (t *Transactor) sendTxsSegment(ctx context.Context, batch *BatchResult) {
	for _, payload := range batch.payloads {
		txCopy := *payload.Tx
		t.pool.Submit(func() {
			begin := time.Now()
			err := t.eth.SendTransaction(ctx, &txCopy)
			t.tallyCh <- &tallyItem{txCopy.Hash(), batch.batchNo, err, payload.VerifyFn, time.Since(begin).Nanoseconds()}
		})
	}
	t.pool.Finish()
}

func (t *Transactor) sendTxsSync(ctx context.Context, batch *BatchResult) {
	for _, payload := range batch.payloads {
		begin := time.Now()
		err := t.eth.SendTransaction(ctx, payload.Tx)
		t.tallyCh <- &tallyItem{payload.Tx.Hash(), batch.batchNo, err, payload.VerifyFn, time.Since(begin).Nanoseconds()}
	}
}

func (t *Transactor) sendTxsBatch(ctx context.Context, batch *BatchResult) {
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
		for i, elem := range elems {
			hash := batch.payloads[i].Tx.Hash()
			verify := batch.payloads[i].VerifyFn
			t.tallyCh <- &tallyItem{hash, batch.batchNo, elem.Error, verify, took}
		}
	})
	if !t.gen.concurrent {
		t.pool.Finish()
	}
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
