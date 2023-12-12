package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/dreamer-zq/turbo-tester/simple"
)

// NewContractCmd creates a new instance of the cobra.Command for the "contract" command.
//
// Returns a pointer to the cobra.Command.
func NewContractCmd() *cobra.Command {
	contractCmd := &cobra.Command{
		Use:   "contract",
		Short: "Turbo tester app command",
	}

	manager := simple.NewManager()
	contractCmd.AddCommand(DeployCmd(manager))
	contractCmd.AddCommand(MethodsCmd(manager))
	contractCmd.AddCommand(GentxCmd(manager))
	contractCmd.AddCommand(StartCmd(manager))

	contractCmd.PersistentFlags().String(flagURL, "", "turbo endpoint url")
	contractCmd.PersistentFlags().String(flagName, "eTicket", "contract name")
	contractCmd.PersistentFlags().Int64(flagChainID, 0, "turbo chain-id")
	contractCmd.MarkFlagRequired(flagURL)
	contractCmd.MarkFlagRequired(flagChainID)
	return contractCmd
}

// MethodsCmd creates a new Cobra command for the "methods" command.
//
// The manager parameter is a pointer to a simple.Manager object.
// It is used to manage the samplers.
//
// The function returns a pointer to a Cobra command object.
func MethodsCmd(manager *simple.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "methods",
		Short: "List all contract methods",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadGlobalFlags(cmd, manager)
			if err != nil {
				return err
			}

			m, err := conf.contract.MethodMap(conf.client)
			if err != nil {
				return err
			}

			var methodDesc []string
			for _, method := range m {
				methodDesc = append(methodDesc, method.Display())
			}
			fmt.Println(strings.Join(methodDesc, "\n"))
			return nil
		},
	}
	return cmd
}
