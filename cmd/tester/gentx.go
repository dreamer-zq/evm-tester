package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/dreamer-zq/turbo-tester/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
)

// GentxCmd returns a cobra Command for the "gentx" command.
//
// The command generates test data and outputs it to a CSV file.
// It takes no parameters and returns a pointer to a cobra.Command.
func GentxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gentx",
		Short: "Generate test data and output to cvs file",
		RunE: func(cmd *cobra.Command, args []string) error {
			url,err := cmd.Flags().GetString(flagURL)
			if err != nil {
				return err
			}
			conn, err := ethclient.Dial(url)
			if err != nil {
				log.Fatalf("Failed to connect to the Ethereum client: %v", err)
			}

			contractAddrStr,err := cmd.Flags().GetString(flagContract)
			if err != nil {
				return err
			}

			path,err := cmd.Flags().GetString(flagOutput)
			if err != nil {
				return err
			}

			contractAddr := common.HexToAddress(contractAddrStr)
			chainID,err := conn.ChainID(context.Background())
			if err != nil {
				return err
			}

			maxThreads,err := cmd.Flags().GetInt(flagMaxThreads)
			if err != nil {
				return err
			}

			tg := tester.NewTxGenerator(
				chainID,
				genTx(conn,contractAddr),
				tester.NewPool(maxThreads),
			)

			sender, err := crypto.GenerateKey()
			if err != nil {
				return err
			}

			data,err := tg.BatchGenTxs(sender, big.NewInt(0),100, contractAddr)
			if err != nil {
				return err
			}
			return tester.SaveToCSV(path, data)
		},
	}
	return cmd
}

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
