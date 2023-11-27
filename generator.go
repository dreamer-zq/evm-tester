package tester

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

// Payload is a struct that contains the raw transaction and the chain ID.
type Payload struct {
	RawTx   string `csv:"raw_tx"`
	ChainID string `csv:"chain_id"`
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
// CreateOrSendTx is a function type that can create or send transactions.
type CreateOrSendTx func(opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, error)

// TxGenerator generates transactions for the TicketGame contract.
type TxGenerator struct {
	chainID        *big.Int
	createOrSendTx CreateOrSendTx
	pool           *Pool
	gasFeeCap      *big.Int // Gas fee cap to use for the 1559 transaction execution (nil = gas price oracle)
	gasTipCap      *big.Int // Gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle)
	gasLimit       uint64   // Gas limit to set for the transaction execution (0 = estimate)
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
	createOrSendTx CreateOrSendTx,
	pool *Pool,
	options ...Option,
) *TxGenerator {
	tg := &TxGenerator{
		chainID:        chainID,
		createOrSendTx: createOrSendTx,
		pool:           pool,
	}
	for _, option := range options {
		tg = option(tg)
	}
	return tg 
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
func (tg *TxGenerator) GenTx(sender *ecdsa.PrivateKey, senderNonce *big.Int, params ...interface{}) (*Payload, error) {
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

	rawTransaction, err := tg.createOrSendTx(auth, params...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create or send transaction")
	}

	txbz, err := rawTransaction.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal transaction")
	}
	return &Payload{
		RawTx:   hexutil.Bytes(txbz).String(),
		ChainID: tg.chainID.String(),
	}, nil
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
func (tg *TxGenerator) BatchGenTxs(sender *ecdsa.PrivateKey, senderNonce *big.Int, batchSize int32, params ...interface{}) ([]*Payload, error) {
	txs := make([]*Payload, 0, batchSize)
	for i := 0; i < int(batchSize); i++ {
		tx, err := tg.GenTx(sender, senderNonce, params...)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate transaction")
		}
		txs = append(txs, tx)
		senderNonce.Add(senderNonce, big.NewInt(1))
	}
	tg.pool.Close()
	return txs, nil
}

// RandomGenTx generates a random transaction for the given player address.
//
// player: the address of the player.
// Returns: the hex string representation of the generated transaction.
func (tg *TxGenerator) RandomGenTx(params ...interface{}) (*Payload, error) {
	sender, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}
	senderNonce := big.NewInt(0)
	return tg.GenTx(sender, senderNonce, params...)
}

// RandomBatchGenTxs generates a batch of random transactions.
//
// It takes in the batchSize parameter, which specifies the number of transactions to generate in the batch.
// The player parameter is used to specify the address of the player associated with the transactions.
//
// The function returns a slice of strings, which represents the generated transactions.
func (tg *TxGenerator) RandomBatchGenTxs(batchSize int32, params ...interface{}) ([]*Payload, error) {
	txs := make([]*Payload, 0, batchSize)
	for i := 0; i < int(batchSize); i++ {
		tg.pool.Submit(func() {
			tx, err := tg.RandomGenTx(params...)
			if err != nil {
				log.Fatalf("Failed to generate transaction: %v", err)
			}
			txs = append(txs, tx)
		})
	}
	tg.pool.Close()
	return txs, nil
}
