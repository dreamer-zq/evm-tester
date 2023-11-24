package cmd

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
)

var (
	flagCount      = "count"
	flagConcurrent = "concurrent"
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
			url, err := cmd.Flags().GetString(flagURL)
			if err != nil {
				return err
			}
			conn, err := ethclient.Dial(url)
			if err != nil {
				return errors.New("failed to connect to the Ethereum client")
			}

			contractAddrStr, err := cmd.Flags().GetString(flagContract)
			if err != nil {
				return err
			}

			path, err := cmd.Flags().GetString(flagOutput)
			if err != nil {
				return err
			}

			count, err := cmd.Flags().GetInt32(flagCount)
			if err != nil {
				return err
			}

			concurrent, err := cmd.Flags().GetBool(flagConcurrent)
			if err != nil {
				return err
			}

			contractAddr := common.HexToAddress(contractAddrStr)
			chainID, err := conn.ChainID(context.Background())
			if err != nil {
				return err
			}

			maxThreads, err := cmd.Flags().GetInt(flagMaxThreads)
			if err != nil {
				return err
			}

			tg := tester.NewTxGenerator(
				chainID,
				genTx(conn, contractAddr),
				tester.NewPool(maxThreads),
			)

			var data []*tester.Payload
			if concurrent {
				data, err = tg.RandomBatchGenTxs(count, contractAddr)
				if err != nil {
					return err
				}
			} else {
				sender, err := crypto.GenerateKey()
				if err != nil {
					return err
				}

				data, err = tg.BatchGenTxs(sender, big.NewInt(0), count, contractAddr)
				if err != nil {
					return err
				}
			}
			return tester.SaveToCSV(path, data)
		},
	}
	cmd.Flags().Int32("count", 10, "the amount of data generated")
	cmd.Flags().Bool("concurrent", true, "whether to use concurrent mode,the number of concurrencies is the same as `data-count`")
	return cmd
}
