package simple

import (
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	tester "github.com/dreamer-zq/turbo-tester"
	"github.com/dreamer-zq/turbo-tester/simple/gen"
)

// TicketGameSampler is a struct that implements the Sampler interface.
type TicketGameSampler struct {
	contractAddr common.Address
}

// SetContractAddr sets the contract address for the TicketGameSampler.
//
// contractAddr: the address of the contract to be set.
func (tgs *TicketGameSampler) SetContractAddr(contractAddr common.Address) {
	tgs.contractAddr = contractAddr
}

// GenTxBuilder generates a CreateOrSendTx function for the TicketGameSampler struct.
//
// It takes a *cobra.Command, *ethclient.Client, and common.Address as parameters.
// It returns a CreateOrSendTx function and an error.
func (tgs *TicketGameSampler) GenTxBuilder(conn *ethclient.Client, method string, params []string) (tester.CreateTx, error) {
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

// Deploy deploys the TicketGame contract.
//
// It takes an authenticated transaction options and a contract backend as parameters.
// It returns the address of the deployed contract and an error if the deployment fails.
func (tgs *TicketGameSampler) Deploy(_ *cobra.Command, auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, error) {
	contractAddr, _, _, err := gen.DeployTicketGame(auth, backend)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to deploy contract")
	}
	return contractAddr, nil
}

// MethodMap returns a map of methods for the TicketGameSampler type.
//
// No parameters.
// Returns a map of string keys to Method values.
func (tgs *TicketGameSampler) MethodMap(conn *ethclient.Client) (map[string]Method, error) {
	ticker, err := gen.NewTicketGame(tgs.contractAddr, conn)
	if err != nil {
		return nil, err
	}

	abi, err := gen.TicketGameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return map[string]Method{
		"redeem":      TicketGameSamplerRedeemMethod{ticker,abi},
		"batchRedeem": TicketGameSamplerBatchRedeemMethod{ticker,abi},
	}, nil
}

// TicketGameSamplerRedeemMethod is a struct that implements the Method interface.
type TicketGameSamplerRedeemMethod struct {
	contract *gen.TicketGame
	abi      *abi.ABI
}

// FormatParams formats the params for the TicketGameSamplerRedeemMethod Go function.
//
// It takes in a slice of strings called params and returns a slice of interfaces and an error.
func (t TicketGameSamplerRedeemMethod) FormatParams(params []string) ([]interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("invalid contract params")
	}

	player := common.HexToAddress(params[0])
	tokenURI := "http://redeem.io/"
	return []interface{}{player, tokenURI}, nil
}

// GenTx generates a transaction for redeeming a ticket in the TicketGame contract.
//
// It takes in the following parameters:
//   - opts: The transaction options (bind.TransactOpts).
//   - params: The parameters required for redeeming the ticket.
//
// It returns a transaction object (*types.Transaction) and an error object (error).
func (t TicketGameSamplerRedeemMethod) GenTx(opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, error) {
	if len(params) != 2 {
		return nil, errors.New("invalid contract params")
	}
	player := params[0].(common.Address)
	tokenURI := params[1].(string)
	return t.contract.Redeem(opts, player, tokenURI)
}

// Display returns the string representation of the TicketGameSamplerRedeemMethod.
//
// No parameters.
// Returns a string.
func (t TicketGameSamplerRedeemMethod) Display() string {
	return t.abi.Methods["redeem"].String()
}

// TicketGameSamplerBatchRedeemMethod is a struct that implements the Method interface.
type TicketGameSamplerBatchRedeemMethod struct {
	contract *gen.TicketGame
	abi      *abi.ABI
}

// FormatParams formats the params for the TicketGameSamplerBatchRedeemMethod Go function.
//
// It takes in a slice of strings called params and returns a slice of interfaces and an error.
func (t TicketGameSamplerBatchRedeemMethod) FormatParams(params []string) ([]interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("invalid contract params")
	}

	player := common.HexToAddress(params[0])
	tokenURI := "http://redeem.io/"

	batchNum := 100
	players := make([]common.Address, 0, batchNum)
	tokenURIs := make([]string, 0, batchNum)
	for i := 0; i < batchNum; i++ {
		players = append(players, player)
		tokenURIs = append(tokenURIs, tokenURI)

	}
	return []interface{}{players, tokenURIs}, nil
}

// GenTx generates a transaction to redeem tickets in batches for the TicketGameSamplerBatchRedeemMethod contract.
//
// It takes the following parameters:
// - opts: The transaction options.
// - params: A variadic parameter list that represents the player addresses and token URIs.
//
// It returns a transaction object and an error.
func (t TicketGameSamplerBatchRedeemMethod) GenTx(opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, error) {
	if len(params) != 2 {
		return nil, errors.New("invalid contract params")
	}
	player := params[0].([]common.Address)
	tokenURI := params[1].([]string)
	return t.contract.BatchRedeem(opts, player, tokenURI)
}

// Display returns the string representation of the TicketGameSamplerBatchRedeemMethod.
//
// It does not take any parameters.
// It returns a string.
func (t TicketGameSamplerBatchRedeemMethod) Display() string {
	return t.abi.Methods["batchRedeem"].String()
}
