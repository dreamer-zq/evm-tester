package simple

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	tester "github.com/dreamer-zq/evm-tester"
)

// Contract is an interface that defines the GenTxBuilder method.
type Contract interface {
	Deploy(auth *bind.TransactOpts, backend bind.ContractBackend, params []string) (common.Address, error)
	BindAddress(contractAddr common.Address)
	GenTxBuilder(conn *ethclient.Client, method string, params []string) (tester.CreateTx, error)
	MethodMap(conn *ethclient.Client) (map[string]Method, error)
}

// Method is an interface that defines the Call method.
type Method interface {
	FormatParams(params []string) ([]any, error)
	GenTx(opts *bind.TransactOpts, params ...any) (*types.Transaction,tester.Verifier ,error)
	Display() string
}

// MethodVerifiable is an interface that defines the Verifiers method.
type MethodVerifiable interface{
	Method
	GenVerifier(params []string) ([]tester.Verifier, error)
}
