package services

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"time"

	"github.com/sharering/shareledger/cmd/Shareledgerd/services/subscriber"

	"github.com/sharering/shareledger/cmd/Shareledgerd/services/database"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
)

const flagConfigPath = "config"

var supportedTypes = map[string]struct{}{
	"in":          {},
	"out":         {},
	"approver-in": {},
}

type processFunc func(ctx context.Context, network string) error

type Network struct {
	Signer        string `yaml:"signer"`
	Exponent      int    `yaml:"exponent"`
	Url           string `yaml:"url"`
	ChainId       int64  `yaml:"chainId"`
	SwapContract  string `yaml:"swapContract"` // swap contract address
	TokenContract string `yaml:"tokenContract"`
	SwapTopic     string `yaml:"swapTopic"` // swap topic
	TransferTopic string `yaml:"transferTopic"`
	Retry         Retry  `yaml:"retry"`
}

type Retry struct {
	IntervalRetry   time.Duration `yaml:"intervalRetry"`
	RetryPercentage float64       `yaml:"retryPercentage"`
}

type RelayerConfig struct {
	Network        map[string]Network `yaml:"networks"`
	Type           string             `yaml:"type"`
	AutoApprove    bool               `yaml:"autoApprove"`
	ScanInterval   time.Duration      `yaml:"scanInterval"`
	MongoURI       string             `yaml:"mongoURI"`
	DbName         string             `yaml:"dbName"`
	CollectionName string             `yaml:"collectionName"`
}

type Relayer struct {
	Config RelayerConfig
	db     database.DBRelayer
	events map[string]subscriber.Service

	cmd                *cobra.Command
	qClient            swapmoduletypes.QueryClient
	clientTx           client.Context
	preRunBroadcastTxs []tx.PreRunBroadcastTx
}

func (r *Relayer) Validate() error {
	if _, found := supportedTypes[r.Config.Type]; !found {
		return errors.New(fmt.Sprintf("type, %s, relayer is not supported", r.Config.Type))
	}
	return nil
}
