package simple

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
	"github.com/dreamer-zq/turbo-tester/simple/gen"
)

// ETicketSampler is a struct that implements the Sampler interface.
type ETicketSampler struct {
	contractAddr common.Address
}

// SetContract sets the contract address for the ETicketSampler.
//
// contractAddr: the address of the contract to be set.
func (tgs *ETicketSampler) SetContract(contractAddr common.Address) {
	tgs.contractAddr = contractAddr
}

// GenTxBuilder generates a CreateOrSendTx function for the ETicketSampler struct.
//
// It takes a *cobra.Command, *ethclient.Client, and common.Address as parameters.
// It returns a CreateOrSendTx function and an error.
func (tgs *ETicketSampler) GenTxBuilder(conn *ethclient.Client, method string, params []string) (tester.CreateTx, error) {
	methodMap, err := tgs.MethodMap(conn)
	if err != nil {
		return nil, err
	}

	m, ok := methodMap[method]
	if !ok {
		return nil, errors.New("invalid method")
	}
	p, err := m.FormatParams(params)
	if err != nil {
		return nil, err
	}

	return func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return m.GenTx(opts, p...)
	}, nil
}

// DeployContract deploys the ETicketSampler contract.
//
// It takes an authenticated transaction options and a contract backend as parameters.
// It returns the address of the deployed contract and an error if the deployment fails.
func (tgs *ETicketSampler) DeployContract(_ *cobra.Command, auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, error) {
	contractAddr, _, _, err := gen.DeployETicket(auth, backend)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to deploy contract")
	}
	return contractAddr, nil
}

// MethodMap returns a map of methods for the ETicketSampler type.
//
// No parameters.
// Returns a map of string keys to Method values.
func (tgs *ETicketSampler) MethodMap(conn *ethclient.Client) (map[string]Method, error) {
	ticker, err := gen.NewETicket(tgs.contractAddr, conn)
	if err != nil {
		return nil, err
	}

	return map[string]Method{
		"mint":      ETicketSamplerMintMethod{ticker},
	}, nil
}

// ETicketSamplerMintMethod is a struct that implements the Method interface.
type ETicketSamplerMintMethod struct {
	contract *gen.ETicket
}

// FormatParams formats the params for the ETicketSamplerMintMethod Go function.
//
// It takes in a slice of strings called params and returns a slice of interfaces and an error.
func (t ETicketSamplerMintMethod) FormatParams(params []string) ([]interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("invalid contract params")
	}

	to := common.HexToAddress(params[0])
	quantity,ok := new(big.Int).SetString(params[1], 10)
	if !ok {
		return nil, errors.New("invalid contract params quantity")
	}
	return []interface{}{to, quantity}, nil
}


// GenTx generates a transaction for the ETicketSamplerMintMethod Go function.
//
// It takes in the following parameter(s):
// - opts: a *bind.TransactOpts object representing the transaction options.
// - params: a variadic parameter that can take in any number of arguments.
//
// It returns a *types.Transaction object and an error.
func (t ETicketSamplerMintMethod) GenTx(opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, error) {
	if len(params) != 2 {
		return nil, errors.New("invalid contract params")
	}
	return t.contract.Mint(opts, params[0].(common.Address), params[1].(*big.Int))
}