package simple

import (
	"log/slog"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	tester "github.com/dreamer-zq/turbo-tester"
	"github.com/dreamer-zq/turbo-tester/simple/gen"
)

// ETicketSampler is a struct that implements the Sampler interface.
type ETicketSampler struct {
	contractAddr common.Address
}

// SetContractAddr sets the contract address for the ETicketSampler.
//
// contractAddr: the address of the contract to be set.
func (tgs *ETicketSampler) SetContractAddr(contractAddr common.Address) {
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

	return func(opts *bind.TransactOpts) (*types.Transaction, error) {
		p, err := m.FormatParams(params)
		if err != nil {
			return nil, err
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
	totalSupply, ok := new(big.Int).SetString(params[2], 10)
	if !ok {
		return common.Address{}, errors.New("invalid contract params totalSupply")
	}

	validTime, ok := new(big.Int).SetString(params[3], 10)
	if !ok {
		return common.Address{}, errors.New("invalid contract params validTime")
	}
	contractAddr, _, _, err := gen.DeployETicket(auth, backend, infos, rights, totalSupply, validTime)
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
		"mint":             ETicketSamplerMintMethod{ticker, abi},
		"safeTransferFrom": ETicketSamplerTranferMethod{ticker, abi},
		"multicall":        &ETicketSamplerMulticallMethod{ticker, abi, nil},
	}, nil
}

// ETicketSamplerMintMethod is a struct that implements the Method interface.
type ETicketSamplerMintMethod struct {
	contract *gen.ETicket
	abi      *abi.ABI
}

// FormatParams formats the params for the ETicketSamplerMintMethod Go function.
//
// It takes in a slice of strings called params and returns a slice of interfaces and an error.
func (t ETicketSamplerMintMethod) FormatParams(params []string) ([]interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("invalid contract params")
	}

	to := common.HexToAddress(params[0])
	quantity, ok := new(big.Int).SetString(params[1], 10)
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

// Display returns a string representation of the ETicketSamplerMintMethod.
//
// It does not take any parameters.
// It returns a string.
func (t ETicketSamplerMintMethod) Display() string {
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
func (t ETicketSamplerTranferMethod) FormatParams(params []string) ([]interface{}, error) {
	if len(params) != 3 {
		return nil, errors.New("invalid contract params")
	}
	from := common.HexToAddress(params[0])
	to := common.HexToAddress(params[1])
	tokenID, ok := new(big.Int).SetString(params[2], 10)
	if !ok {
		return nil, errors.New("invalid contract params tokenID")
	}
	return []interface{}{from, to, tokenID}, nil
}

// GenTx generates a transaction for the ETicketSamplerTranferMethod Go function.
//
// It takes in the following parameter(s):
// - opts: a *bind.TransactOpts object representing the transaction options.
// - params: a variadic parameter that can take in any number of arguments.
//
// It returns a *types.Transaction object and an error.
func (t ETicketSamplerTranferMethod) GenTx(opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, error) {
	if len(params) != 3 {
		return nil, errors.New("invalid contract params")
	}
	return t.contract.SafeTransferFrom(opts, params[0].(common.Address), params[1].(common.Address), params[2].(*big.Int))
}

// Display returns a string representing the ETicketSamplerTranferMethod.
//
// No parameters.
// Returns a string.
func (t ETicketSamplerTranferMethod) Display() string {
	return t.abi.Methods["safeTransferFrom"].String()
}

// ETicketSamplerMulticallMethod is a struct that implements the Method interface.
type ETicketSamplerMulticallMethod struct {
	contract  *gen.ETicket
	abi       *abi.ABI
	tokenNext *big.Int
}

// FormatParams is a function that takes in a slice of strings and returns a slice of interfaces and an error.
//
// It formats the provided parameters and returns them in the desired format.
//
// Params:
// - params: a slice of strings representing the parameters to be formatted.
//
// Returns:
// - []interface{}: a slice of interfaces representing the formatted parameters.
// - error: an error indicating any issues that occurred during the formatting process.
func (t *ETicketSamplerMulticallMethod) FormatParams(params []string) ([]interface{}, error) {
	if len(params) == 0 {
		return nil, errors.New("invalid contract params")
	}

	var datas [][]byte
	switch params[0] {
	case "mint":
		to := common.HexToAddress(params[1])
		amount, err := strconv.ParseInt(params[2], 10, 64)
		if err != nil {
			return nil, errors.New("invalid contract params amount")
		}

		var (
			tokenIDFrom *big.Int
			ok          bool
		)
		if t.tokenNext != nil {
			tokenIDFrom = big.NewInt(t.tokenNext.Int64())
		} else {
			tokenIDFrom, ok = new(big.Int).SetString(params[3], 10)
			if !ok {
				return nil, errors.New("invalid contract params tokenIdFrom")
			}
		}
		tokenIDTo := big.NewInt(tokenIDFrom.Int64() + amount)
		for i := tokenIDFrom; i.Cmp(tokenIDTo) < 0; i.Add(i, big.NewInt(1)) {
			data, err := t.abi.Pack("mint", to, i)
			if err != nil {
				return nil, err
			}
			datas = append(datas, data)
		}
		t.tokenNext = big.NewInt(tokenIDTo.Int64() + 1)
		slog.Info("mint erc721 token", "tokenNext", t.tokenNext.Int64(), "tokenIdFrom", tokenIDFrom.Int64(), "tokenIDTo", tokenIDTo.Int64())
	}
	return []interface{}{datas}, nil
}

// GenTx generates a transaction for the ETicketSamplerMulticallMethod.
//
// opts: The transaction options.
// params: The parameters for the transaction.
// returns: The generated transaction and any error that occurred.
func (t *ETicketSamplerMulticallMethod) GenTx(opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, error) {
	if len(params) != 1 {
		return nil, errors.New("invalid contract params")
	}
	return t.contract.Multicall(opts, params[0].([][]byte))
}

// Display returns a string representing the "multicall" method of the ETicketSamplerMulticallMethod type.
//
// No parameters.
// Returns a string.
func (t *ETicketSamplerMulticallMethod) Display() string {
	return "function multicall(string method, string[] methodParams) returns(bytes[] results)"
}