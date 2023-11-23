package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

	tg := TxGenerator{
		contract: ticker,
		chainID:  chainID,
	}

	sender, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}

	rawTx := tg.GenTx(sender, big.NewInt(0), contractAddr)
	log.Println(rawTx)
}
