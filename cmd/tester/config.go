package cmd

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/dreamer-zq/turbo-tester/simple"
)

// GlobalConfig represents a global config
type GlobalConfig struct {
	chainID *big.Int
	url     string

	contract simple.Contract
	client   *ethclient.Client
}

func loadGlobalFlags(cmd *cobra.Command, manager *simple.Manager) (*GlobalConfig, error) {
	url, err := cmd.Flags().GetString(flagURL)
	if err != nil {
		return nil, err
	}

	contractName, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return nil, err
	}

	chainIDInt, err := cmd.Flags().GetInt64(flagChainID)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpc.DialOptions(context.Background(), url, rpc.WithHeader("X-Chain", strconv.FormatInt(chainIDInt, 10)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the Ethereum client")
	}

	return &GlobalConfig{
		chainID:  big.NewInt(chainIDInt),
		url:      url,
		client:   ethclient.NewClient(rpcClient),
		contract: manager.GetContract(contractName),
	}, nil
}

// TransactionConfig represents a transaction config
type TransactionConfig struct {
	gasLimit  uint64
	gasFeeCap *big.Int
	gasTipCap *big.Int
	nonce     int64
	privKey   *ecdsa.PrivateKey

	contractAddr         common.Address
	contractMethod       string
	contractMethodParams []string
	batchSize            uint64
}

func loadTransactionFlags(cmd *cobra.Command, client *ethclient.Client) (*TransactionConfig, error) {
	batchSize, err := cmd.Flags().GetUint64(flagBatchSize)
	if err != nil {
		return nil, err
	}

	gasLimit, err := cmd.Flags().GetUint64(flagGasLimit)
	if err != nil {
		return nil, err
	}

	gasFeeCap, err := cmd.Flags().GetInt64(flagGasFeeCap)
	if err != nil {
		return nil, err
	}

	gasTipCap, err := cmd.Flags().GetInt64(flagGasTipCap)
	if err != nil {
		return nil, err
	}

	privKeyStr, err := cmd.Flags().GetString(flagPrivateKey)
	if err != nil {
		return nil, err
	}

	var privKey *ecdsa.PrivateKey
	if privKeyStr != "" {
		privKeyStr = strings.TrimPrefix(privKeyStr, "0x")
		privKey, err = crypto.HexToECDSA(privKeyStr)
		if err != nil {
			return nil, err
		}
	}

	nonce, err := cmd.Flags().GetInt64(flagNonce)
	if err != nil {
		return nil, err
	}

	if privKey != nil {
		senderAddr := crypto.PubkeyToAddress(privKey.PublicKey)
		if nonce == 0 {
			nonceAct, err := client.NonceAt(context.Background(), senderAddr, nil)
			if err != nil {
				return nil, err
			}
			nonce = int64(nonceAct)
		}
	}

	method, err := cmd.Flags().GetString(flagContractMethod)
	if err != nil {
		return nil, err
	}

	contractParams, err := cmd.Flags().GetStringSlice(flagContractParams)
	if err != nil {
		return nil, err
	}

	contractAddrStr, err := cmd.Flags().GetString(flagContract)
	if err != nil {
		return nil, err
	}
	contractAddr := common.HexToAddress(contractAddrStr)

	return &TransactionConfig{
		gasLimit:             gasLimit,
		gasFeeCap:            big.NewInt(gasFeeCap),
		gasTipCap:            big.NewInt(gasTipCap),
		nonce:                nonce,
		privKey:              privKey,
		contractAddr:         contractAddr,
		contractMethod:       method,
		contractMethodParams: contractParams,
		batchSize:            batchSize,
	}, nil
}
