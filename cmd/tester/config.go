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

	"github.com/dreamer-zq/evm-tester/simple"
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
	gasLimit   uint64
	gasFeeCap  *big.Int
	gasTipCap  *big.Int
	nonce      int64
	privKey    *ecdsa.PrivateKey
	noSign     bool
	batchSize  uint64
	callConfig *ContractCallConfig
}

func (tc *TransactionConfig) load(cmd *cobra.Command, client *ethclient.Client) (err error) {
	tc.batchSize, err = cmd.Flags().GetUint64(flagBatchSize)
	if err != nil {
		return err
	}

	tc.gasLimit, err = cmd.Flags().GetUint64(flagGasLimit)
	if err != nil {
		return err
	}

	gasFeeCap, err := cmd.Flags().GetInt64(flagGasFeeCap)
	if err != nil {
		return err
	}
	tc.gasFeeCap = big.NewInt(gasFeeCap)

	gasTipCap, err := cmd.Flags().GetInt64(flagGasTipCap)
	if err != nil {
		return err
	}
	tc.gasTipCap = big.NewInt(gasTipCap)

	privKeyStr, err := cmd.Flags().GetString(flagPrivateKey)
	if err != nil {
		return err
	}

	if privKeyStr != "" {
		privKeyStr = strings.TrimPrefix(privKeyStr, "0x")
		tc.privKey, err = crypto.HexToECDSA(privKeyStr)
		if err != nil {
			return err
		}
	}

	tc.nonce, err = cmd.Flags().GetInt64(flagNonce)
	if err != nil {
		return err
	}

	tc.noSign, err = cmd.Flags().GetBool(flagNoSign)
	if err != nil {
		return err
	}

	if tc.privKey != nil {
		senderAddr := crypto.PubkeyToAddress(tc.privKey.PublicKey)
		if tc.nonce == 0 {
			nonceAct, err := client.NonceAt(context.Background(), senderAddr, nil)
			if err != nil {
				return err
			}
			tc.nonce = int64(nonceAct)
		}
	}

	var callConfig = &ContractCallConfig{}
	if err := callConfig.load(cmd); err != nil {
		return err
	}
	tc.callConfig = callConfig
	return nil
}

// ContractCallConfig represents a contract config
type ContractCallConfig struct {
	addr         common.Address
	method       string
	methodParams []string
}

func (cc *ContractCallConfig) load(cmd *cobra.Command) (err error) {
	cc.methodParams, err = cmd.Flags().GetStringSlice(flagContractParams)
	if err != nil {
		return err
	}

	contractAddrStr, err := cmd.Flags().GetString(flagContract)
	if err != nil {
		return err
	}
	cc.addr = common.HexToAddress(contractAddrStr)
	cc.method, err = cmd.Flags().GetString(flagContractMethod)
	if err != nil {
		return err
	}
	return nil
}

func (cc ContractCallConfig) bindFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagContractMethod, "", "the contract method name being tested")
	cmd.Flags().StringSlice(flagContractParams, []string{}, "the contract method params being tested")
	cmd.Flags().String(flagContract, "", "the contract address being tested")

	cmd.MarkFlagRequired(flagContract)
	cmd.MarkFlagRequired(flagContractMethod)
}
