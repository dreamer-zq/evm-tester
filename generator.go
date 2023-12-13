package tester

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"net/http"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

// Payload is a struct that contains the raw transaction and the chain ID.
type Payload struct {
	Tx      *types.Transaction `csv:"-"`
	RawTx   string             `csv:"raw_tx"`
	ChainID string             `csv:"chain_id"`
}

// Option is a function type that can be used to configure the TxGenerator.
type Option func(*TxGenerator) *TxGenerator

// SetGasFeeCap sets the gas fee cap for the TxGenerator.
//
// Parameters:
// - gasFeeCap: A pointer to a big.Int representing the gas fee cap.
//
// Returns:
// - An Option function that sets the gas fee cap for the TxGenerator.
func SetGasFeeCap(gasFeeCap *big.Int) Option {
	return func(tg *TxGenerator) *TxGenerator {
		tg.gasFeeCap = gasFeeCap
		return tg
	}
}

// SetGasTipCap sets the gas tip cap option for the TxGenerator.
//
// gasTipCap: A pointer to a big.Int representing the gas tip cap.
// Returns: An Option function that sets the gas tip cap for the TxGenerator.
func SetGasTipCap(gasTipCap *big.Int) Option {
	return func(tg *TxGenerator) *TxGenerator {
		tg.gasTipCap = gasTipCap
		return tg
	}
}

// SetGasLimit sets the gas limit option for the TxGenerator.
func SetGasLimit(gasLimit uint64) Option {
	return func(tg *TxGenerator) *TxGenerator {
		tg.gasLimit = gasLimit
		return tg
	}
}

// SetBatchSize sets the batch size for the TxGenerator.
//
// It takes a batchSize parameter of type uint64 and returns an Option.
func SetBatchSize(batchSize uint64) Option {
	return func(tg *TxGenerator) *TxGenerator {
		tg.batchSize = batchSize
		return tg
	}
}

// SetConcurrent returns an Option function that sets the concurrent flag of a TxGenerator object.
//
// Parameters:
//
//	concurrent - a boolean value indicating whether the TxGenerator should be executed concurrently.
//
// Returns:
//
//	An Option function that sets the concurrent flag and returns the modified TxGenerator object.
func SetConcurrent(concurrent bool) Option {
	return func(tg *TxGenerator) *TxGenerator {
		tg.concurrent = concurrent
		return tg
	}
}

// SetNonce sets the nonce value for the TxGenerator.
//
// Parameter:
// - nonce: the nonce value to set.
//
// Return:
// - *TxGenerator: the updated TxGenerator.
func SetNonce(nonce int64) Option {
	return func(tg *TxGenerator) *TxGenerator {
		tg.nonce = nonce
		return tg
	}
}

// SetPrivKey sets the private key for the TxGenerator.
//
// privKey: the private key string.
// Returns: the TxGenerator with the updated private key.
func SetPrivKey(privKey string) Option {
	return func(tg *TxGenerator) *TxGenerator {
		tg.privKey = privKey
		return tg
	}
}

// CreateTx is a function type that can create or send transactions.
type CreateTx func(opts *bind.TransactOpts) (*types.Transaction, error)

// TxGenerator generates transactions for the TicketGame contract.
type TxGenerator struct {
	chainID    *big.Int
	createTx   CreateTx
	pool       *Pool
	gasFeeCap  *big.Int // Gas fee cap to use for the 1559 transaction execution (nil = gas price oracle)
	gasTipCap  *big.Int // Gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle)
	gasLimit   uint64   // Gas limit to set for the transaction execution (0 = estimate)
	batchSize  uint64
	privKey    string
	nonce      int64
	concurrent bool
}

// NewTxGenerator initializes a new instance of the TxGenerator struct.
//
// It takes the chainID, createOrSendTx, and pool as parameters.
// The chainID is a pointer to a big.Int type, representing the chain ID.
// The createOrSendTx is a function type that can be used to create or send transactions.
// The pool is a pointer to the Pool struct, representing a transaction pool.
//
// It returns a pointer to the TxGenerator struct.
func NewTxGenerator(
	chainID *big.Int,
	createTx CreateTx,
	pool *Pool,
	options ...Option,
) *TxGenerator {
	tg := &TxGenerator{
		chainID:  chainID,
		createTx: createTx,
		pool:     pool,
	}
	for _, option := range options {
		tg = option(tg)
	}
	return tg
}

// Run runs the TxGenerator.
//
// It generates a batch of transactions based on the TxGenerator's configuration.
// If the TxGenerator is concurrent, it calls the RandomBatchGenTxs method to generate the transactions.
// If the TxGenerator has a private key, it calls the BatchGenTxs method to generate the transactions using the private key.
// If neither of the above conditions are met, it generates a new private key and calls the BatchGenTxs method to generate the transactions.
// It returns the generated transactions and any error that occurred.
func (tg *TxGenerator) Run() ([]*Payload, error) {
	var (
		data []*Payload
		err  error
	)
	switch {
	case tg.concurrent:
		data, err = tg.RandomBatchGenTxs()
		if err != nil {
			return nil, err
		}
		break
	case tg.privKey != "":
		privKey := strings.TrimPrefix(tg.privKey, "0x")
		sender, err := crypto.HexToECDSA(privKey)
		if err != nil {
			return nil, err
		}
		data, err = tg.BatchGenTxs(sender, big.NewInt(tg.nonce))
		if err != nil {
			return nil, err
		}
		break
	default:
		sender, err := crypto.GenerateKey()
		if err != nil {
			return nil, err
		}

		data, err = tg.BatchGenTxs(sender, big.NewInt(0))
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

// GenTx generates a transaction using the provided sender's private key, sender nonce, and player address.
//
// Parameters:
// - sender: The private key of the sender.
// - senderNonce: The nonce of the sender.
// - player: The address of the player.
//
// Returns:
// - The hexadecimal representation of the generated transaction.
func (tg *TxGenerator) GenTx(sender *ecdsa.PrivateKey, senderNonce *big.Int) (*Payload, error) {
	rawTransaction, err := tg.genTx(sender, senderNonce)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create or send transaction")
	}

	txbz, err := rawTransaction.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal transaction")
	}
	return &Payload{
		Tx:      rawTransaction,
		RawTx:   hexutil.Bytes(txbz).String(),
		ChainID: tg.chainID.String(),
	}, nil
}

func (tg *TxGenerator) genTx(sender *ecdsa.PrivateKey, senderNonce *big.Int) (*types.Transaction, error) {
	// Create an authorized transactor and call the store function
	auth, err := bind.NewKeyedTransactorWithChainID(sender, tg.chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create authorized transactor")
	}
	auth.NoSend = true
	auth.Nonce = senderNonce
	auth.GasLimit = tg.gasLimit
	auth.GasTipCap = tg.gasTipCap
	auth.GasFeeCap = tg.gasFeeCap

	header := make(http.Header)
	header.Add("X-Chain", tg.chainID.String())
	auth.Context = rpc.NewContextWithHeaders(context.Background(), header)

	rawTransaction, err := tg.createTx(auth)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create or send transaction")
	}
	return rawTransaction, nil
}

// BatchGenTxs generates a batch of transactions using the given sender's private key, sender's nonce, batch size, and player address.
//
// Parameters:
// - sender: The sender's private key for signing the transactions.
// - senderNonce: The nonce value of the sender's account.
// - batchSize: The number of transactions to generate in the batch.
// - player: The address of the player to include in the transactions.
//
// Return:
// - []string: The generated transactions as a slice of strings.
func (tg *TxGenerator) BatchGenTxs(sender *ecdsa.PrivateKey, senderNonce *big.Int) ([]*Payload, error) {
	txs := make([]*Payload, 0, tg.batchSize)
	for i := uint64(0); i < tg.batchSize; i++ {
		tx, err := tg.GenTx(sender, senderNonce)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate transaction")
		}
		txs = append(txs, tx)
		senderNonce.Add(senderNonce, big.NewInt(1))
	}
	tg.pool.Close()
	tg.nonce = senderNonce.Int64() // Update the nonce value
	return txs, nil
}

// RandomGenTx generates a random transaction for the given player address.
//
// player: the address of the player.
// Returns: the hex string representation of the generated transaction.
func (tg *TxGenerator) RandomGenTx() (*Payload, error) {
	sender, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}
	senderNonce := big.NewInt(0)
	return tg.GenTx(sender, senderNonce)
}

// RandomBatchGenTxs generates a batch of random transactions.
//
// It takes in the batchSize parameter, which specifies the number of transactions to generate in the batch.
// The player parameter is used to specify the address of the player associated with the transactions.
//
// The function returns a slice of strings, which represents the generated transactions.
func (tg *TxGenerator) RandomBatchGenTxs() ([]*Payload, error) {
	txs := make([]*Payload, 0, tg.batchSize)
	mu := sync.Mutex{}
	for i := uint64(0); i < tg.batchSize; i++ {
		tg.pool.Submit(func() {
			tx, err := tg.RandomGenTx()
			if err != nil {
				log.Fatalf("Failed to generate transaction: %v", err)
			}
			mu.Lock()
			txs = append(txs, tx)
			mu.Unlock()
		})
	}
	tg.pool.Finish()
	return txs, nil
}
