package cmd

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
)

var (
	flagTotalTxs  = "run-total-txs"
	flagSync      = "run-sync"
	flagRunPeriod = "run-period"
	flagUserNum   = "run-user-num"
)

// StartCmd generates a cobra command for sending transaction.
//
// The function does the following:
// - Creates a new cobra command.
// - Sets the command's use and short description.
// - Defines a RunE function that handles the command's execution.
// - Retrieves the contract address from the command flag.
// - Retrieves the output path from the command flag.
// - Retrieves the generator from the command.
// - Runs the generator with the contract address.
// Returns the generated cobra command.
func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
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
			contractAddr := common.HexToAddress(contractAddrStr)

			tg, err := getGenerator(conf, cmd)
			if err != nil {
				return err
			}

			totalTxs, err := cmd.Flags().GetInt64(flagTotalTxs)
			if err != nil {
				return err
			}

			sync, err := cmd.Flags().GetBool(flagSync)
			if err != nil {
				return err
			}

			runPeriod, err := cmd.Flags().GetDuration(flagRunPeriod)
			if err != nil {
				return err
			}

			var endTime time.Time
			if runPeriod > 0 {
				endTime = time.Now().Add(runPeriod)
			}

			userNum, err := cmd.Flags().GetInt(flagUserNum)
			if err != nil {
				return err
			}

			transactor := tester.NewTransactor(
				conf.client,
				userNum,
				tg,
				tester.SetTotalTxs(totalTxs),
				tester.SetSync(sync),
				tester.SetEndTime(endTime),
			)
			transactor.Run(contractAddr)
			return nil
		},
	}
	addGenTxFlags(cmd)
	cmd.Flags().Int(flagUserNum, 0, "maximum number of concurrent users")
	cmd.Flags().Duration(flagRunPeriod, 0, "stress test execution time,eg: 5m")
	cmd.Flags().Bool(flagSync, false, "whether transaction execution is in synchronous mode")
	cmd.Flags().Int64(flagTotalTxs, 0, "the total number of transactions executed, and `--run-period`, choose one of the two")
	return cmd
}
