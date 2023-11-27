package cmd

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
	"github.com/dreamer-zq/turbo-tester/simple"
)

var (
	flagCount      = "count"
	flagConcurrent = "concurrent"
	flagContract   = "contract"
	flagMaxThreads = "max-threads"
	flagOutput     = "output"
	flagGasFeeCap  = "gas-fee-cap"
	flagGasTipCap  = "gas-tip-cap"
	flagGasLimit   = "gas-limit"
	flagPrivateKey = "private-key"
	flagNonce      = "nonce"
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
			conf, err := loadGlobalFlags(cmd)
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

			var opts []tester.Option
			gasLimit, err := cmd.Flags().GetUint64(flagGasLimit)
			if err != nil {
				return err
			}
			opts = append(opts, tester.SetGasLimit(gasLimit))

			gasFeeCap, err := cmd.Flags().GetInt64(flagGasFeeCap)
			if err != nil {
				return err
			}
			opts = append(opts, tester.SetGasFeeCap(big.NewInt(gasFeeCap)))

			gasTipCap, err := cmd.Flags().GetInt64(flagGasTipCap)
			if err != nil {
				return err
			}
			opts = append(opts, tester.SetGasTipCap(big.NewInt(gasTipCap)))

			tg := tester.NewTxGenerator(
				conf.chainID,
				simple.GenTx(conf.client, contractAddr),
				tester.NewPool(maxThreads),
				opts...,
			)

			privKey, err := cmd.Flags().GetString(flagPrivateKey)
			if err != nil {
				return err
			}

			var data []*tester.Payload
			switch {
			case concurrent:
				data, err = tg.RandomBatchGenTxs(count, contractAddr)
				if err != nil {
					return err
				}
				break
			case privKey != "":
				sender, err := crypto.HexToECDSA(privKey)
				if err != nil {
					return err
				}
				nonce, err := cmd.Flags().GetInt64(flagNonce)
				if err != nil {
					return err
				}
				data, err = tg.BatchGenTxs(sender, big.NewInt(nonce), count, contractAddr)
				if err != nil {
					return err
				}
				break
			default:
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
	cmd.Flags().String(flagPrivateKey, "", "send the account private key for the transaction")
	cmd.Flags().Int64(flagNonce, 0, "user's nonce")
	cmd.Flags().Int64(flagGasFeeCap, 0, "gas fee cap to use for the 1559 transaction execution (nil = gas price oracle,fetch from chain)")
	cmd.Flags().Int64(flagGasTipCap, 0, "gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle,fetch from chain)")
	cmd.Flags().Uint64(flagGasLimit, 0, "gas limit to set for the transaction execution (0 = estimate,fetch from chain)")

	cmd.MarkFlagRequired(flagContract)
	cmd.MarkFlagRequired(flagOutput)
	return cmd
}
