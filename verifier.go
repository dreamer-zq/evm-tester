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
	id       string
	verifier Verifier

	failedCounter int32
	next          time.Time
}

type record struct {
	ID     string `csv:"id"`
	Status string `csv:"status"`
}

// Verify is the verification function.
type Verify func() (success bool, err error)

// Verifier is the interface that verifies the hashes in the queue.
type Verifier interface {
	ID() string
	Verify() (success bool, err error)
}

// GenericVerifier is a struct that implements the Verifier interface.
type GenericVerifier struct {
	id     string
	verify Verify
}

// TxVerify verifies the transaction receipt for a given hash.
//
// It takes an `eth` parameter of type `*ethclient.Client` which represents the Ethereum client.
// The `hash` parameter of type `common.Hash` represents the hash of the transaction.
//
// It returns a `GenericVerifier` which is used to verify the transaction receipt.
func TxVerify(eth *ethclient.Client, hash common.Hash) GenericVerifier {
	verify := func() (bool, error) {
		receipt, err := eth.TransactionReceipt(context.Background(), hash)
		if err != nil || receipt == nil {
			return false, errors.New("get transaction receipt error")
		}
		return true, nil
	}
	return NewGenericVerifier(hash.Hex(), verify)
}

// NewGenericVerifier creates a new instance of GenericVerifier.
//
// Parameters:
// - id: The identifier for the verifier.
// - verify: The verification function.
//
// Returns:
// - GenericVerifier: The newly created instance of GenericVerifier.
func NewGenericVerifier(id string, verify Verify) GenericVerifier {
	return GenericVerifier{id, verify}
}

// ID returns the ID of the GenericVerifier.
//
// No parameters.
// Returns a string.
func (gv GenericVerifier) ID() string {
	return gv.id
}

// Verify verifies the GenericVerifier.
//
// It returns a boolean value indicating whether the verification was successful or not,
// and an error if any error occurred during the verification process.
func (gv GenericVerifier) Verify() (bool, error) {
	return gv.verify()
}

// VerifierManager is a struct that verifies the hashes in the queue.
type VerifierManager struct {
	enable  bool
	queue   *Queue[*element]
	timer   *time.Ticker
	eth     *ethclient.Client
	records []*record
	running bool
}

// NewVerifierManager creates a new Verifier instance.
//
// enable: a boolean indicating whether the Verifier is enabled.
// eth: an instance of ethclient.Client used for interacting with the Ethereum blockchain.
// Returns a pointer to the newly created Verifier instance.
func NewVerifierManager(enable bool, eth *ethclient.Client) *VerifierManager {
	return &VerifierManager{
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
func (vm *VerifierManager) Add(verifier Verifier, txHash *common.Hash) {
	if vm.enable {
		if verifier == nil && txHash != nil {
			verifier = TxVerify(vm.eth, *txHash)
		}

		vm.queue.Add(&element{
			verifier: verifier,
			id:       verifier.ID(),
			next:     time.Now().Add(period),
		})
	}
}

// Start starts the VerifierManager.
//
// It validates the elements in the VerifierManager and updates the records accordingly.
// It uses a timer to periodically execute the validation logic.
// If the VerifierManager is disabled, the function returns immediately.
// The function has no parameters and does not return any values.
func (vm *VerifierManager) Start(parallel bool) {
	if !vm.enable {
		return
	}

	validate := func(ele *element) bool {
		if ele.failedCounter >= maxFailedCounter {
			vm.records = append(vm.records, &record{
				ID:     ele.id,
				Status: "failed",
			})
			return true
		}
		now := time.Now()
		if now.Second() < ele.next.Second() {
			return false
		}
		success, err := ele.verifier.Verify()
		if err != nil {
			ele.failedCounter++
			ele.next = ele.next.Add(time.Duration(ele.failedCounter) * period)
			return false
		}
		record := &record{
			ID:     ele.id,
			Status: "failed",
		}
		if success {
			record.Status = "success"
		}
		vm.records = append(vm.records, record)
		return true
	}

	for range vm.timer.C {
		if vm.running {
			continue
		}
		vm.running = true

		slog.Info("execute to verify logic", "left", vm.queue.Length())
		vm.queue.Iterate(validate)
		vm.running = false
	}
}

// Finish checks if the Verifier has finished processing.
//
// It returns true if the Verifier is not enabled or if the queue length is zero,
// otherwise it returns false.
func (vm *VerifierManager) Finish(total int64) bool {
	if !vm.enable {
		return true
	}
	slog.Info("verify finished?", "left", vm.queue.Length(), "total", total, "records", len(vm.records))
	if vm.queue.IsEmpty() && int64(len(vm.records)) == total {
		vm.timer.Stop()
		SaveToCSV("./result.csv", vm.records)
		return true
	}
	return false
}
