package cmd

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
)

var (
	flagCount      = "count"
	flagConcurrent = "concurrent"
	flagContract   = "contract"
	flagMaxThreads = "max-threads"
	flagOutput     = "output"
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
			conf,err := loadGlobalFlags(cmd)
			if err != nil {
				return err
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
			maxThreads, err := cmd.Flags().GetInt(flagMaxThreads)
			if err != nil {
				return err
			}

			tg := tester.NewTxGenerator(
				conf.chainID,
				genTx(conf.client, contractAddr),
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
	cmd.Flags().Int32(flagCount, 10, "the amount of data generated")
	cmd.Flags().Bool(flagConcurrent, true, "whether to use concurrent mode,the number of concurrencies is the same as `data-count`")
	cmd.Flags().String(flagContract, "", "the contract address being tested")
	cmd.Flags().Int(flagMaxThreads, 100, "maximum number of threads")
	cmd.Flags().String(flagOutput, "", "csv file output path")
	
	cmd.MarkFlagRequired(flagContract)
	cmd.MarkFlagRequired(flagOutput)
	return cmd
}
