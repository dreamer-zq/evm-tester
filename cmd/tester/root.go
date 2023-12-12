package cmd

import (
	"context"
	"math/big"
	"strconv"

	"github.com/dreamer-zq/turbo-tester/simple"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	flagURL     = "url"
	flagChainID = "chain-id"
	flagName    = "name"
)

// NewRootCmd returns a new instance of the cobra.Command struct.
//
// This function does not take any parameters.
// It returns a pointer to a cobra.Command struct.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tester",
		Short: "Turbo tester app command",
	}

	manager := simple.NewManager()
	rootCmd.AddCommand(ListCmd(manager))
	rootCmd.AddCommand(NewContractCmd())
	return rootCmd
}

// GlobalConnfig represents a global config
type GlobalConnfig struct {
	chainID *big.Int
	url     string

	contract simple.Contract
	client   *ethclient.Client
}

func loadGlobalFlags(cmd *cobra.Command, manager *simple.Manager) (*GlobalConnfig, error) {
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

	return &GlobalConnfig{
		chainID:  big.NewInt(chainIDInt),
		url:      url,
		client:   ethclient.NewClient(rpcClient),
		contract: manager.GetContract(contractName),
	}, nil
}
