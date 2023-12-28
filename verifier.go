package tester

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/exp/slog"
)

const (
	maxFailedCounter = 10
	period           = 10 * time.Second
)

type element struct {
	verify        Verify
	hash          string
	failedCounter int32
	next          time.Time
}

type record struct {
	Hash   string `csv:"hash"`
	Status string `csv:"status"`
}

// Verify is the verification function.
type Verify func() (success bool, err error)

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
		timer:  time.NewTicker(period),
		eth:    eth,
	}
}

// Add adds a new element to the Verifier's queue for verification.
//
// It takes a hash and a Verify function as parameters. If the Verifier is enabled,
// the function adds a new element to the queue with the provided hash and verify function.
// If the verify function is nil, it uses the defaultVerifyFn of the Verifier for verification.
// The element is added only if the Verifier is enabled.
func (v *Verifier) Add(hash common.Hash, verify Verify) {
	if v.enable {
		if verify == nil {
			verify = v.defaultVerifyFn(hash)
		}
		v.queue.Add(&element{
			verify: verify,
			hash:   hash.Hex(),
			next:   time.Now().Add(period),
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
		if ele.failedCounter >= maxFailedCounter {
			v.records = append(v.records, &record{
				Hash:   ele.hash,
				Status: "failed",
			})
			return true
		}
		now := time.Now()
		if now.Second() < ele.next.Second() {
			return false
		}
		success, err := ele.verify()
		if err != nil {
			ele.failedCounter++
			ele.next = ele.next.Add(time.Duration(ele.failedCounter) * period)
			return false
		}
		record := &record{
			Hash:   ele.hash,
			Status: "failed",
		}
		if success {
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

func (v *Verifier) defaultVerifyFn(hash common.Hash) Verify {
	return func() (success bool, err error) {
		receipt, err := v.eth.TransactionReceipt(context.Background(), hash)
		if err != nil || receipt == nil {
			return false, errors.New("get transaction receipt error")
		}
		return true, nil
	}
}
