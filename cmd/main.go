package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dreamer-zq/turbo-tester/gen"
)

func main() {
	contractAddr := common.HexToAddress("0x547bD9C389686441d9a56Db1DaffF505bC216073")
	chainID := big.NewInt(1223)
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	ticker, err := gen.NewTicketGame(contractAddr, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}

	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}

	// Create an authorized transactor and call the store function
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth.NoSend = true
	auth.Nonce = big.NewInt(0)

	rawTransaction, err := ticker.Redeem(auth, contractAddr, "tokenURI string")
	if err != nil {
		log.Fatalf("Failed to call Redeem: %v", err)
	}
	txbz, err := rawTransaction.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to call Redeem: %v", err)
	}
	log.Println(hexutil.Bytes(txbz).String())
}
