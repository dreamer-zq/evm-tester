package simple

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
)

// Sampler is an interface that defines the GenTxBuilder method.
type Sampler interface {
	DeployContract(cmd *cobra.Command,auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, error)
	GenTxBuilder(cmd *cobra.Command, conn *ethclient.Client) (tester.CreateTx, error)
	AddFlags(cmd *cobra.Command)
}

// GenTx generates a transaction function to redeem a ticket for a player.
//
// conn is an Ethereum client connection.
// contractAddr is the address of the ticket game contract.
// Returns a function that takes transaction options and a player address,
// and returns a transaction and an error.
// func GenTx(conn *ethclient.Client, contractAddr common.Address) tester.CreateTx {
// 	ticker, err := gen.NewTicketGame(contractAddr, conn)
// 	if err != nil {
// 		log.Fatalf("Failed to instantiate Storage contract: %v", err)
// 	}
// 	return func(opts *bind.TransactOpts) (*types.Transaction, error) {
// 		if len(params) != 1 {
// 			return nil, nil
// 		}
// 		player, ok := params[0].(common.Address)
// 		if !ok {
// 			return nil, nil
// 		}
// 		tokenURI := genTokenURI(opts.Nonce)
// 		if !ok {
// 			return nil, nil
// 		}
// 		return ticker.Redeem(opts, player, tokenURI)
// 	}
// }

func genTokenURI(senderNonce *big.Int) string {
	return fmt.Sprintf("http://redeem.io/%d", senderNonce.Int64())
}
