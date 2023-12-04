package simple

import (
	"fmt"
	"log"
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
	"github.com/dreamer-zq/turbo-tester/simple/gen"
)

var (
	flagContractParams = "contract-method-params"
	flagContract       = "contract"
)

// TicketGameSampler is a struct that implements the Sampler interface.
type TicketGameSampler struct{}

// GenTxBuilder generates a CreateOrSendTx function for the TicketGameSampler struct.
//
// It takes a *cobra.Command, *ethclient.Client, and common.Address as parameters.
// It returns a CreateOrSendTx function and an error.
func (tgs TicketGameSampler) GenTxBuilder(cmd *cobra.Command, conn *ethclient.Client) (tester.CreateTx, error) {
	contractAddrStr, err := cmd.Flags().GetString(flagContract)
	if err != nil {
		return nil, err
	}
	contractAddr := common.HexToAddress(contractAddrStr)

	contractParams, err := cmd.Flags().GetStringSlice(flagContractParams)
	if err != nil {
		return nil, err
	}

	if len(contractParams) != 1 {
			return nil, errors.New("invalid contract params")
	}

	ticker, err := gen.NewTicketGame(contractAddr, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}

	return func(opts *bind.TransactOpts) (*types.Transaction, error) {
		player := common.HexToAddress(contractParams[0])
		tokenURI := genTokenURI(opts.Nonce)
		return ticker.Redeem(opts, player, tokenURI)
	}, nil
}

// AddFlags adds flags to the given command.
//
// Parameters:
// - cmd: a pointer to a cobra.Command object.
//
// Return type: None.
func (tgs TicketGameSampler) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringSlice(flagContractParams, []string{}, "the contract method params being tested")
	cmd.Flags().String(flagContract, "", "the contract address being tested")

	cmd.MarkFlagRequired(flagContract)
}

// DeployContract deploys the TicketGame contract.
//
// It takes an authenticated transaction options and a contract backend as parameters.
// It returns the address of the deployed contract and an error if the deployment fails.
func (tgs TicketGameSampler) DeployContract(_ *cobra.Command, auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, error) {
	contractAddr, _, _, err := gen.DeployTicketGame(auth, backend)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to deploy contract")
	}
	return contractAddr, nil
}

func genTokenURI(senderNonce *big.Int) string {
	return fmt.Sprintf("http://redeem.io/%d", senderNonce.Int64())
}
