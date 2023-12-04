package simple

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
)

// Sampler is an interface that defines the GenTxBuilder method.
type Sampler interface {
	DeployContract(cmd *cobra.Command,auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, error)
	GenTxBuilder(cmd *cobra.Command, conn *ethclient.Client) (tester.CreateTx, error)
	AddFlags(cmd *cobra.Command)
}
