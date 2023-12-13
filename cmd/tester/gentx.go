package cmd

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
	"github.com/dreamer-zq/turbo-tester/simple"
)

var (
	flagBatchSize      = "batch-size"
	flagConcurrent     = "concurrent"
	flagContract       = "contract"
	flagMaxThreads     = "max-threads"
	flagOutput         = "output"
	flagGasFeeCap      = "gas-fee-cap"
	flagGasTipCap      = "gas-tip-cap"
	flagGasLimit       = "gas-limit"
	flagPrivateKey     = "private-key"
	flagNonce          = "nonce"
	flagContractParams = "contract-method-params"
	flagContractMethod = "contract-method"
)

// GentxCmd returns a cobra Command for the "gentx" command.
//
// The command generates test data and outputs it to a CSV file.
// It takes no parameters and returns a pointer to a cobra.Command.
func GentxCmd(manager *simple.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gentx",
		Short: "Generate test data and output to cvs file",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadGlobalFlags(cmd, manager)
			if err != nil {
				return err
			}

			path, err := cmd.Flags().GetString(flagOutput)
			if err != nil {
				return err
			}

			tg, err := getGenerator(conf, cmd)
			if err != nil {
				return err
			}

			data, err := tg.Run()
			if err != nil {
				return err
			}
			return tester.SaveToCSV(path, data)
		},
	}

	addGenTxFlags(cmd)
	return cmd
}

func getGenerator(conf *GlobalConnfig, cmd *cobra.Command) (*tester.TxGenerator, error) {
	count, err := cmd.Flags().GetUint64(flagBatchSize)
	if err != nil {
		return nil, err
	}

	maxThreads, err := cmd.Flags().GetInt(flagMaxThreads)
	if err != nil {
		return nil, err
	}

	var opts []tester.Option
	gasLimit, err := cmd.Flags().GetUint64(flagGasLimit)
	if err != nil {
		return nil, err
	}
	opts = append(opts, tester.SetGasLimit(gasLimit))
	opts = append(opts, tester.SetBatchSize(count))

	gasFeeCap, err := cmd.Flags().GetInt64(flagGasFeeCap)
	if err != nil {
		return nil, err
	}
	opts = append(opts, tester.SetGasFeeCap(big.NewInt(gasFeeCap)))

	gasTipCap, err := cmd.Flags().GetInt64(flagGasTipCap)
	if err != nil {
		return nil, err
	}
	opts = append(opts, tester.SetGasTipCap(big.NewInt(gasTipCap)))

	privKey, err := cmd.Flags().GetString(flagPrivateKey)
	if err != nil {
		return nil, err
	}
	opts = append(opts, tester.SetPrivKey(privKey))

	nonce, err := cmd.Flags().GetInt64(flagNonce)
	if err != nil {
		return nil, err
	}
	opts = append(opts, tester.SetNonce(nonce))

	concurrent, err := cmd.Flags().GetBool(flagConcurrent)
	if err != nil {
		return nil, err
	}
	opts = append(opts, tester.SetConcurrent(concurrent))

	method, err := cmd.Flags().GetString(flagContractMethod)
	if err != nil {
		return nil, err
	}

	contractParams, err := cmd.Flags().GetStringSlice(flagContractParams)
	if err != nil {
		return nil, err
	}

	contractAddrStr, err := cmd.Flags().GetString(flagContract)
	if err != nil {
		return nil, err
	}
	contractAddr := common.HexToAddress(contractAddrStr)

	conf.contract.SetContractAddr(contractAddr)

	txBuilrder, err := conf.contract.GenTxBuilder(conf.client, method, contractParams)
	if err != nil {
		return nil, err
	}

	return tester.NewTxGenerator(
		conf.chainID,
		txBuilrder,
		tester.NewPool(maxThreads, "gentx"),
		opts...,
	), nil
}

func addGenTxFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagOutput, "", "csv file output path")
	cmd.MarkFlagRequired(flagOutput)
	addSendTxFlags(cmd)
}

func addTxFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagPrivateKey, "", "send the account private key for the transaction")
	cmd.Flags().Int64(flagNonce, 0, "user's nonce")
	cmd.Flags().Int64(flagGasFeeCap, 0, "gas fee cap to use for the 1559 transaction execution (nil = gas price oracle,fetch from chain)")
	cmd.Flags().Int64(flagGasTipCap, 0, "gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle,fetch from chain)")
	cmd.Flags().Uint64(flagGasLimit, 0, "gas limit to set for the transaction execution (0 = estimate,fetch from chain)")
}

func addSendTxFlags(cmd *cobra.Command) {
	addTxFlags(cmd)
	cmd.Flags().Uint64(flagBatchSize, 10, "number of transactions per batch")
	cmd.Flags().Bool(flagConcurrent, false, "whether to use concurrent mode,the number of concurrencies is the same as `data-count`")
	cmd.Flags().Int(flagMaxThreads, 100, "maximum number of threads")
	cmd.Flags().String(flagContractMethod, "", "the contract method name being tested")
	cmd.Flags().StringSlice(flagContractParams, []string{}, "the contract method params being tested")
	cmd.Flags().String(flagContract, "", "the contract address being tested")

	cmd.MarkFlagRequired(flagContract)
	cmd.MarkFlagRequired(flagContractMethod)
}
