package cmd

import (
	"github.com/dreamer-zq/evm-tester/simple"
	"github.com/spf13/cobra"
)

var (
	flagURL     = "url"
	flagChainID = "chain-id"
	flagName    = "contract-name"
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

	manager := simple.NewManager()
	rootCmd.AddCommand(ListCmd(manager))
	rootCmd.AddCommand(NewContractCmd())
	return rootCmd
}