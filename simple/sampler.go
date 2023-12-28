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
	SetContractAddr(contractAddr common.Address)
	GenTxBuilder(conn *ethclient.Client, method string, params []string) (tester.CreateTx, error)
	MethodMap(conn *ethclient.Client) (map[string]Method, error)
}

// Method is an interface that defines the Call method.
type Method interface {
	FormatParams(params []string) ([]interface{}, error)
	GenTx(opts *bind.TransactOpts, params ...interface{}) (*types.Transaction,tester.Verify ,error)
	Display() string
}
