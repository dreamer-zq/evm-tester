package simple

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"time"

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

	return func(opts *bind.TransactOpts) (*types.Transaction, tester.Verifier, error) {
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
		"batchMint": &POAPSamplerBatchMintMethod{contract, abi, 1, 0},
	}, nil
}

// BindAddress implements Contract.
func (poap *POAPSampler) BindAddress(contractAddr common.Address) {
	poap.contractAddr = contractAddr
}

// POAPSamplerBatchMintMethod is a struct that implements the Method interface.
type POAPSamplerBatchMintMethod struct {
	contract *gen.POAP
	abi      *abi.ABI
	page     int
	total    int
}

// FormatParams formats the params for the POAPSamplerBatchMintMethod Go function.
func (t *POAPSamplerBatchMintMethod) FormatParams(params []string) ([]any, error) {
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
	return []any{addrs.Address(), tokenID, offset}, nil
}

// GenTx generates a transaction for the POAPSamplerBatchMintMethod Go function.
func (t *POAPSamplerBatchMintMethod) GenTx(opts *bind.TransactOpts, params ...any) (*types.Transaction, tester.Verifier, error) {
	if len(params) != 3 {
		return nil, nil, errors.New("invalid contract params")
	}
	addrs := params[0].([]common.Address)
	if len(addrs) == 0 {
		slog.Info("airdrop end", "page", t.page, "total", t.total)
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

	id := fmt.Sprintf("%s-%d", tx.Hash().Hex(), params[2].(int))
	veirfy := t.genVerifier(id, addrs, ids)
	t.page++
	t.total += len(addrs)
	return tx, veirfy, nil
}

// GenVerifier generates a verifier for the POAPSamplerBatchMintMethod type.
//
// It takes an array of strings called params as a parameter.
// It returns an array of tester.Verify and an error.
func (t *POAPSamplerBatchMintMethod) GenVerifier(params []string) ([]tester.Verifier, error) {
	var (
		page      = 1
		verifiers []tester.Verifier
	)
	_, err := db.Connect(params[0])
	if err != nil {
		return nil, err
	}

	tokenID, ok := new(big.Int).SetString(params[1], 10)
	if !ok {
		return nil, errors.New("invalid contract params tokenID")
	}
	for {
		offset := (page - 1) * pageSize
		addrs, err := db.Accounts{}.AddressPageQuery(pageSize, offset)
		if err != nil {
			return nil, err
		}
		if len(addrs) == 0 {
			break
		}

		var ids []*big.Int
		for i := 0; i < len(addrs); i++ {
			ids = append(ids, tokenID)
		}
		id := fmt.Sprintf("offset-%d", offset)
		verifiers = append(verifiers, t.genVerifier(id, addrs.Address(), ids))
		page++
	}
	return verifiers, nil
}

// Display returns a string representing the "batchMint" method of the POAPSamplerBatchMintMethod type.
//
// No parameters.
// Returns a string.
func (t *POAPSamplerBatchMintMethod) Display() string {
	return t.abi.Methods["batchMint"].String()
}

func (t *POAPSamplerBatchMintMethod) genVerifier(id string, addrs []common.Address, ids []*big.Int) tester.Verifier {
	verify := func() (bool, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		
		balances, err := t.contract.BalanceOfBatch(&bind.CallOpts{
			Context: ctx,
		}, addrs, ids)
		if err != nil {
			slog.Info("call BalanceOfBatch failed", "id", id,"err", err)
			return false, err
		}

		for i, balance := range balances {
			bal := balance.Int64()
			if bal == 0 {
				return false, nil
			}
			if bal > 1 {
				slog.Info("duplicate airdrop", "address", addrs[i], "balance", bal, "tokenID", ids[i].String())
				return false, nil
			}
		}
		return true, nil
	}
	return tester.NewGenericVerifier(id, verify)
}
