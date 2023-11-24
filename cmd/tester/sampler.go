package cmd

import (
	"fmt"
	"log"
	"math/big"

	"github.com/dreamer-zq/turbo-tester/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	tester "github.com/dreamer-zq/turbo-tester"
)

func genTx(conn *ethclient.Client, contractAddr common.Address) tester.CreateOrSendTx {
	ticker, err := gen.NewTicketGame(contractAddr, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}
	return func(opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, error) {
		if len(params) != 1 {
			return nil, nil
		}
		player, ok := params[0].(common.Address)
		if !ok {
			return nil, nil
		}
		tokenURI := genTokenURI(opts.Nonce)
		if !ok {
			return nil, nil
		}
		return ticker.Redeem(opts, player, tokenURI)
	}
}

func genTokenURI(senderNonce *big.Int) string {
	return fmt.Sprintf("http://redeem.io/%d", senderNonce.Int64())
}