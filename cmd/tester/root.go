package cmd

import "github.com/spf13/cobra"

var (
	flagURL        = "url"
	flagContract   = "contract"
	flagMaxThreads = "max-threads"
	flagOutput     = "output"
)

// NewRootCmd returns a new instance of the cobra.Command struct.
//
// This function does not take any parameters.
// It returns a pointer to a cobra.Command struct.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tester",
		Short: "Turbo tester app command",
	}
	rootCmd.AddCommand(GentxCmd())
	rootCmd.PersistentFlags().String(flagURL, "", "evm endpoint url")
	rootCmd.PersistentFlags().String(flagContract, "", "the contract address being tested")
	rootCmd.PersistentFlags().Int(flagMaxThreads, 100, "maximum number of threads")
	rootCmd.PersistentFlags().String(flagOutput, "", "csv file output path")

	rootCmd.MarkFlagRequired(flagURL)
	rootCmd.MarkFlagRequired(flagContract)
	rootCmd.MarkFlagRequired(flagOutput)
	return rootCmd
}
