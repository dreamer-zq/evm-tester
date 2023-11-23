package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/dreamer-zq/turbo-tester/gen"
)

// TxGenerator generates transactions for the TicketGame contract.
type TxGenerator struct {
	chainID  *big.Int
	contract *gen.TicketGame
	pool     *Pool
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
func (tg *TxGenerator) GenTx(sender *ecdsa.PrivateKey, senderNonce *big.Int, player common.Address) string {
	// Create an authorized transactor and call the store function
	auth, err := bind.NewKeyedTransactorWithChainID(sender, tg.chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth.NoSend = true
	auth.Nonce = senderNonce

	rawTransaction, err := tg.contract.Redeem(auth, player, genTokenURI(senderNonce))
	if err != nil {
		log.Fatalf("Failed to call Redeem: %v", err)
	}

	txbz, err := rawTransaction.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to call MarshalBinary: %v", err)
	}
	return hexutil.Bytes(txbz).String()
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
func (tg *TxGenerator) BatchGenTxs(sender *ecdsa.PrivateKey, senderNonce *big.Int, batchSize int32, player common.Address) []string {
	txs := make([]string, 0, batchSize)
	for i := 0; i < int(batchSize); i++ {
		txs = append(txs, tg.GenTx(sender,senderNonce, player))
		senderNonce.Add(senderNonce, big.NewInt(1))
	}
	tg.pool.Close()
	return txs
}

// RandomGenTx generates a random transaction for the given player address.
//
// player: the address of the player.
// Returns: the hex string representation of the generated transaction.
func (tg *TxGenerator) RandomGenTx(player common.Address) string {
	sender, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}
	senderNonce := big.NewInt(0)
	return tg.GenTx(sender,senderNonce, player)
}

// RandomBatchGenTxs generates a batch of random transactions.
//
// It takes in the batchSize parameter, which specifies the number of transactions to generate in the batch.
// The player parameter is used to specify the address of the player associated with the transactions.
//
// The function returns a slice of strings, which represents the generated transactions.
func (tg *TxGenerator) RandomBatchGenTxs(batchSize int32, player common.Address) []string {
	txs := make([]string, 0, batchSize)
	for i := 0; i < int(batchSize); i++ {
		tg.pool.Submit(func() {
			txs = append(txs, tg.RandomGenTx(player))
		})
	}
	tg.pool.Close()
	return txs
}

func genTokenURI(senderNonce *big.Int) string {
	return fmt.Sprintf("http://redeem.io/%d", senderNonce.Int64())
}
