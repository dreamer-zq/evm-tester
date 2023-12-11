package cmd

import (
	"fmt"
	"strings"

	"github.com/dreamer-zq/turbo-tester/simple"
	"github.com/spf13/cobra"
)

// ListCmd returns a new cobra command for the "list" command.
//
// The command lists all samplers (contract) and their details.
// It takes a manager object as a parameter.
// It returns a pointer to a cobra.Command object.
func ListCmd(manager *simple.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all contracts",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(strings.Join(manager.ListSamplers(), ","))
			return nil
		},
	}
	return cmd
}
