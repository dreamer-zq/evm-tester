package tester

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/exp/slog"
)

const maxFailedCounter = 10

type element struct {
	hash          common.Hash
	failedCounter atomic.Int32
}

type record struct {
	Hash   string `csv:"hash"`
	Status string `csv:"status"`
}

// Verifier is a struct that verifies the hashes in the queue.
type Verifier struct {
	enable  bool
	queue   *Queue[*element]
	timer   *time.Ticker
	eth     *ethclient.Client
	records []*record
}

// NewVerifier creates a new Verifier instance.
//
// enable: a boolean indicating whether the Verifier is enabled.
// eth: an instance of ethclient.Client used for interacting with the Ethereum blockchain.
// Returns a pointer to the newly created Verifier instance.
func NewVerifier(enable bool, eth *ethclient.Client) *Verifier {
	return &Verifier{
		enable: enable,
		queue:  NewQueue[*element](),
		timer:  time.NewTicker(10 * time.Second),
		eth:    eth,
	}
}

// Add adds a hash to the Verifier.
//
// The parameter `hash` is the hash to be added to the Verifier.
func (v *Verifier) Add(hash common.Hash) {
	if v.enable {
		v.queue.Add(&element{
			hash: hash,
		})
	}
}

// Start verifies the Verifier.
//
// parallelable is a boolean indicating whether the verification is parallelizable.
// It does not return anything.
func (v *Verifier) Start(parallelable bool) {
	if !v.enable {
		return
	}

	validate := func(ele *element) bool {
		if ele.failedCounter.Load() >= maxFailedCounter {
			v.records = append(v.records, &record{
				Hash:   ele.hash.String(),
				Status: "failed",
			})
			return true
		}
		receipt, err := v.eth.TransactionReceipt(context.Background(), ele.hash)
		if err != nil || receipt == nil {
			ele.failedCounter.Add(1)
			return false
		}
		record := &record{
			Hash:   ele.hash.String(),
			Status: "failed",
		}
		if receipt.Status == types.ReceiptStatusSuccessful {
			record.Status = "success"
		}
		v.records = append(v.records, record)
		return true
	}
	for range v.timer.C {
		slog.Info("verify transactions", "left", v.queue.Length())
		if !parallelable {
			v.queue.Iterate(validate)
		} else {
			v.queue.IterateParallel(validate)
		}
	}
}

// Finish checks if the Verifier has finished processing.
//
// It returns true if the Verifier is not enabled or if the queue length is zero,
// otherwise it returns false.
func (v *Verifier) Finish(total int64) bool {
	if !v.enable {
		return true
	}
	if v.queue.IsEmpty() && int64(len(v.records)) == total {
		v.timer.Stop()
		SaveToCSV("./result.csv", v.records)
		return true
	}
	return false
}
