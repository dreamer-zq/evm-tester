package cmd

import (
	"errors"
	"time"

	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/evm-tester"
	"github.com/dreamer-zq/evm-tester/simple"
)

// VerifyCmd returns a new instance of the "verify" command.
//
// The "verify" command is used to generate test data and send it to the blockchain.
// It takes a manager of type `*simple.Manager` as a parameter.
// This function returns a pointer to a `*cobra.Command` object.
func VerifyCmd(manager *simple.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "perform some validation logic",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadGlobalFlags(cmd, manager)
			if err != nil {
				return err
			}
			ccf := &ContractCallConfig{}
			if err := ccf.load(cmd); err != nil {
				return err
			}
			conf.contract.SetContractAddr(ccf.addr)

			methods, err := conf.contract.MethodMap(conf.client)
			if err != nil {
				return err
			}

			methodVerifiable, ok := methods[ccf.method].(simple.MethodVerifiable)
			if !ok {
				return errors.New("the method is not verifiable")
			}

			verifiers, err := methodVerifiable.GenVerifier(ccf.methodParams)
			verifierManager := tester.NewVerifierManager(true, conf.client)
			for _, v := range verifiers {
				verifierManager.Add(v)
			}
			go verifierManager.Start(true)
			for {
				if verifierManager.Finish(int64(len(verifiers))) {
					break
				}
				time.Sleep(5 * time.Second)
			}
			return nil
		},
	}
	ContractCallConfig{}.bindFlags(cmd)
	return cmd
}
