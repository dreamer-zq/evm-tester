package cmd

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	"github.com/dreamer-zq/turbo-tester/simple"
)

var (
	flagContractConstructorParams = "contract-constructor-params"
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

			constructorParams, err := cmd.Flags().GetStringSlice(flagContractConstructorParams)
			if err != nil {
				return err
			}

			gasLimit, err := cmd.Flags().GetUint64(flagGasLimit)
			if err != nil {
				return err
			}

			gasFeeCap, err := cmd.Flags().GetInt64(flagGasFeeCap)
			if err != nil {
				return err
			}

			gasTipCap, err := cmd.Flags().GetInt64(flagGasTipCap)
			if err != nil {
				return err
			}

			privKey, err := cmd.Flags().GetString(flagPrivateKey)
			if err != nil {
				return err
			}

			nonce, err := cmd.Flags().GetInt64(flagNonce)
			if err != nil {
				return err
			}

			var senderPrivateKey *ecdsa.PrivateKey
			if privKey != "" {
				privKey = strings.TrimPrefix(privKey, "0x")
				senderPrivateKey, err = crypto.HexToECDSA(privKey)
				if err != nil {
					return err
				}
			} else {
				senderPrivateKey, err = crypto.GenerateKey()
				if err != nil {
					return err
				}
			}
			// Create an authorized transactor and call the store function
			auth, err := bind.NewKeyedTransactorWithChainID(senderPrivateKey, conf.chainID)
			if err != nil {
				return errors.Wrap(err, "failed to create authorized transactor")
			}
			auth.GasFeeCap = big.NewInt(gasFeeCap)
			auth.GasTipCap = big.NewInt(gasTipCap)
			auth.GasLimit = gasLimit
			auth.Nonce = big.NewInt(nonce)
			contractAddr, err := conf.contract.Deploy(auth, conf.client, constructorParams)
			if err != nil {
				return errors.Wrap(err, "failed to deploy contract")
			}

			fmt.Println("ContractAddr", contractAddr.Hex())
			fmt.Println("DeployerAddr", crypto.PubkeyToAddress(senderPrivateKey.PublicKey).Hex())
			fmt.Println("DeployerPriv", hexutil.Bytes(crypto.FromECDSA(senderPrivateKey)).String())
			return nil
		},
	}
	addTxFlags(cmd)
	cmd.Flags().StringSlice(flagContractConstructorParams, []string{}, "the contract constructor params")
	return cmd
}
