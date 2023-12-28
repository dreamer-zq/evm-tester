package simple

import (
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	tester "github.com/dreamer-zq/evm-tester"
	"github.com/dreamer-zq/evm-tester/data/db"
	"github.com/dreamer-zq/evm-tester/simple/gen"
)

var (
	_ Contract = &POAPSampler{}

	pageSize = 300
)

// POAPSampler is a struct that implements the Sampler interface.
type POAPSampler struct {
	contractAddr common.Address
}

// Deploy implements Contract.
func (poap *POAPSampler) Deploy(auth *bind.TransactOpts, backend bind.ContractBackend, params []string) (common.Address, error) {
	panic("unimplemented")
}

// GenTxBuilder implements Contract.
func (poap *POAPSampler) GenTxBuilder(conn *ethclient.Client, method string, params []string) (tester.CreateTx, error) {
	methodMap, err := poap.MethodMap(conn)
	if err != nil {
		return nil, err
	}

	m, ok := methodMap[method]
	if !ok {
		return nil, errors.New("invalid method")
	}

	_, err = db.Connect(params[0])
	if err != nil {
		return nil, err
	}
	contractParams := params[1:]

	return func(opts *bind.TransactOpts) (*types.Transaction, tester.Verify, error) {
		p, err := m.FormatParams(contractParams)
		if err != nil {
			return nil, nil, err
		}
		return m.GenTx(opts, p...)
	}, nil
}

// MethodMap implements Contract.
func (poap *POAPSampler) MethodMap(conn *ethclient.Client) (map[string]Method, error) {
	contract, err := gen.NewPOAP(poap.contractAddr, conn)
	if err != nil {
		return nil, err
	}

	abi, err := gen.POAPMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return map[string]Method{
		"batchMint": &POAPSamplerBatchMintMethod{contract, abi, 1},
	}, nil
}

// SetContractAddr implements Contract.
func (poap *POAPSampler) SetContractAddr(contractAddr common.Address) {
	poap.contractAddr = contractAddr
}

// POAPSamplerBatchMintMethod is a struct that implements the Method interface.
type POAPSamplerBatchMintMethod struct {
	contract *gen.POAP
	abi      *abi.ABI
	page     int
}

// FormatParams formats the params for the POAPSamplerBatchMintMethod Go function.
func (t *POAPSamplerBatchMintMethod) FormatParams(params []string) ([]interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("invalid contract params")
	}

	tokenID, ok := new(big.Int).SetString(params[0], 10)
	if !ok {
		return nil, errors.New("invalid contract params tokenID")
	}
	offset := (t.page - 1) * pageSize
	addrs, err := db.Accounts{}.AddressPageQuery(pageSize, offset)
	if err != nil {
		return nil, err
	}
	slog.Info("build transaction", "tokenID", tokenID, "offset", offset, "page", t.page, "pageSize", pageSize, "len", len(addrs))
	return []interface{}{addrs.Address(), tokenID}, nil
}

// GenTx generates a transaction for the POAPSamplerBatchMintMethod Go function.
func (t *POAPSamplerBatchMintMethod) GenTx(opts *bind.TransactOpts, params ...interface{}) (*types.Transaction, tester.Verify, error) {
	if len(params) != 2 {
		return nil, nil, errors.New("invalid contract params")
	}
	addrs := params[0].([]common.Address)
	if len(addrs) == 0 {
		return nil, nil, tester.ErrExit
	}
	tokenID := new(big.Int).Set(params[1].(*big.Int))

	tx, err := t.contract.BatchMint(opts, addrs, tokenID)
	if err != nil {
		return nil, nil, err
	}
	
	var ids []*big.Int
	for i := 0; i < len(addrs); i++ {
		ids = append(ids, tokenID)
	}

	veirfy := func() (bool, error) {
		balances,err := t.contract.BalanceOfBatch(&bind.CallOpts{}, addrs, ids)
		if err != nil {
			return false, err
		}
		
		for i, balance := range balances {
			bal := balance.Int64()
			if bal == 0 {
				return false, nil	
			}
			if bal > 1 {
				slog.Info("duplicate airdrop", "address", addrs[i], "balance", bal,"tokenID", tokenID)
				return false, nil
			}
		}
		return true, nil
	}
	t.page++
	return tx, veirfy, nil
}

// Display implements Method.
func (t *POAPSamplerBatchMintMethod) Display() string {
	return t.abi.Methods["batchMint"].String()
}
