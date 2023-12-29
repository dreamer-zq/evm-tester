package simple

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	tester "github.com/dreamer-zq/evm-tester"
	"github.com/dreamer-zq/evm-tester/simple/gen"
)

// ETicketSampler is a struct that implements the Sampler interface.
type ETicketSampler struct {
	contractAddr common.Address
}

// BindAddress sets the contract address for the ETicketSampler.
//
// contractAddr: the address of the contract to be set.
func (tgs *ETicketSampler) BindAddress(contractAddr common.Address) {
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

	return func(opts *bind.TransactOpts) (*types.Transaction, tester.Verifier, error) {
		p, err := m.FormatParams(params)
		if err != nil {
			return nil, nil, err
		}
		return m.GenTx(opts, p...)
	}, nil
}

// Deploy deploys the ETicketSampler contract.
//
// It takes an authenticated transaction options and a contract backend as parameters.
// It returns the address of the deployed contract and an error if the deployment fails.
func (tgs *ETicketSampler) Deploy(auth *bind.TransactOpts, backend bind.ContractBackend, params []string) (common.Address, error) {
	if len(params) != 4 {
		return common.Address{}, errors.New("invalid params")
	}
	infos := strings.Split(params[0], "|")
	rights := strings.Split(params[1], "|")

	validTime, ok := new(big.Int).SetString(params[2], 10)
	if !ok {
		return common.Address{}, errors.New("invalid contract params validTime")
	}
	contractAddr, _, _, err := gen.DeployETicket(auth, backend, infos, rights, nil, validTime)
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

	abi, err := gen.ETicketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return map[string]Method{
		"mint":             &ETicketSamplerMintMethod{ticker, abi, nil},
		"safeTransferFrom": ETicketSamplerTranferMethod{ticker, abi},
	}, nil
}

// ETicketSamplerMintMethod is a struct that implements the Method interface.
type ETicketSamplerMintMethod struct {
	contract  *gen.ETicket
	abi       *abi.ABI
	tokenNext *big.Int
}

// FormatParams formats the params for the ETicketSamplerMintMethod Go function.
//
// It takes in a slice of strings called params and returns a slice of interfaces and an error.
func (t *ETicketSamplerMintMethod) FormatParams(params []string) ([]any, error) {
	if len(params) != 2 {
		return nil, errors.New("invalid contract params")
	}

	to := common.HexToAddress(params[0])
	from, ok := new(big.Int).SetString(params[1], 10)
	if !ok {
		return nil, errors.New("invalid contract params from")
	}

	tokenID := from
	if t.tokenNext != nil {
		tokenID = t.tokenNext
	}
	return []any{to, tokenID}, nil
}

// GenTx generates a transaction for the ETicketSamplerMintMethod Go function.
//
// It takes in the following parameter(s):
// - opts: a *bind.TransactOpts object representing the transaction options.
// - params: a variadic parameter that can take in any number of arguments.
//
// It returns a *types.Transaction object and an error.
func (t *ETicketSamplerMintMethod) GenTx(opts *bind.TransactOpts, params ...any) (*types.Transaction, tester.Verifier, error) {
	if len(params) != 2 {
		return nil, nil, errors.New("invalid contract params")
	}
	tokenID := params[1].(*big.Int)
	tx, err := t.contract.Mint(opts, params[0].(common.Address), tokenID)
	if err != nil {
		return nil, nil, err
	}
	t.tokenNext = new(big.Int).Add(big.NewInt(tokenID.Int64()), new(big.Int).SetUint64(1))
	return tx, nil, nil
}

// Display returns a string representation of the ETicketSamplerMintMethod.
//
// It does not take any parameters.
// It returns a string.
func (t *ETicketSamplerMintMethod) Display() string {
	return t.abi.Methods["mint"].String()
}

// ETicketSamplerTranferMethod is a struct that implements the Method interface.
type ETicketSamplerTranferMethod struct {
	contract *gen.ETicket
	abi      *abi.ABI
}

// FormatParams formats the params for the ETicketSamplerTranferMethod Go function.
//
// It takes in a slice of strings called params and returns a slice of interfaces and an error.
func (t ETicketSamplerTranferMethod) FormatParams(params []string) ([]any, error) {
	if len(params) != 3 {
		return nil, errors.New("invalid contract params")
	}
	from := common.HexToAddress(params[0])
	to := common.HexToAddress(params[1])
	tokenID, ok := new(big.Int).SetString(params[2], 10)
	if !ok {
		return nil, errors.New("invalid contract params tokenID")
	}
	return []any{from, to, tokenID}, nil
}

// GenTx generates a transaction for the ETicketSamplerTranferMethod Go function.
//
// It takes in the following parameter(s):
// - opts: a *bind.TransactOpts object representing the transaction options.
// - params: a variadic parameter that can take in any number of arguments.
//
// It returns a *types.Transaction object and an error.
func (t ETicketSamplerTranferMethod) GenTx(opts *bind.TransactOpts, params ...any) (*types.Transaction, tester.Verifier, error) {
	if len(params) != 3 {
		return nil, nil, errors.New("invalid contract params")
	}
	tx, err := t.contract.SafeTransferFrom(opts, params[0].(common.Address), params[1].(common.Address), params[2].(*big.Int))
	if err != nil {
		return nil, nil, err
	}
	return tx, nil, nil
}

// Display returns a string representing the ETicketSamplerTranferMethod.
//
// No parameters.
// Returns a string.
func (t ETicketSamplerTranferMethod) Display() string {
	return t.abi.Methods["safeTransferFrom"].String()
}
