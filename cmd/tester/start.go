package cmd

import (
	"time"

	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
	"github.com/dreamer-zq/turbo-tester/simple"
)

var (
	flagTotalBatch = "run-total-batch"
	flagRunPeriod  = "run-period"
	flagUserNum    = "run-user-num"
	flagSegment    = "run-segment"
	flagSendMode   = "send-mode"
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
func StartCmd(manager *simple.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Generate test data and send to the blockchain",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadGlobalFlags(cmd, manager)
			if err != nil {
				return err
			}

			generator, err := getGenerator(conf, cmd)
			if err != nil {
				return err
			}

			totalBatch, err := cmd.Flags().GetInt64(flagTotalBatch)
			if err != nil {
				return err
			}

			sendModeStr, err := cmd.Flags().GetString(flagSendMode)
			if err != nil {
				return err
			}
			sendMode, err := tester.ParseSendMode(sendModeStr)
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
				generator,
				tester.SetTotalBatch(totalBatch),
				tester.SetEndTime(endTime),
				tester.SetSendMode(sendMode),
			)
			transactor.Run()
			return nil
		},
	}
	addSendTxFlags(cmd)
	cmd.Flags().Int(flagUserNum, 0, "maximum number of concurrent users")
	cmd.Flags().Duration(flagRunPeriod, 0, "stress test execution time,eg: 5m")
	cmd.Flags().Bool(flagSegment, false, "whether to enable segmented statistics requires run-total-batch to be greater than 1")
	cmd.Flags().Int64(flagTotalBatch, 0, "total production batches, and `--run-period`, choose one of the two,`totalTxs = totalBatch * count`")
	cmd.Flags().String(flagSendMode, "parallel", "transaction sending mode, `oneByOne`,`parallel` ,`segment` or `batch`")
	return cmd
}
