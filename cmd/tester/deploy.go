package cmd

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	"github.com/dreamer-zq/turbo-tester/simple"
)

// DeployCmd returns a new instance of the `cobra.Command` struct for the `deploy` command.
//
// No parameters.
// Returns a pointer to a `cobra.Command` struct.
func DeployCmd(manager *simple.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := loadGlobalFlags(cmd, manager)
			if err != nil {
				return err
			}

			senderPrivateKey, err := crypto.GenerateKey()
			// Create an authorized transactor and call the store function
			auth, err := bind.NewKeyedTransactorWithChainID(senderPrivateKey, conf.chainID)
			if err != nil {
				return errors.Wrap(err, "failed to create authorized transactor")
			}

			contractAddr, err := conf.simpler.DeployContract(cmd, auth, conf.client)
			if err != nil {
				return errors.Wrap(err, "failed to deploy contract")
			}

			fmt.Println("ContractAddr", contractAddr.Hex())
			fmt.Println("DeployerAddr", crypto.PubkeyToAddress(senderPrivateKey.PublicKey).Hex())
			fmt.Println("DeployerPriv", hexutil.Bytes(crypto.FromECDSA(senderPrivateKey)).String())
			return nil
		},
	}
	return cmd
}
