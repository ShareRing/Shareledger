package services

import (
	"context"
	"time"

	"github.com/sharering/shareledger/cmd/Shareledgerd/services/subscriber"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/database"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
)

const flagConfigPath = "config"

var supportedTypes = map[string]struct{}{
	"in":  {},
	"out": {},
}

type processFunc func(ctx context.Context, network string) error

type Network struct {
	Signer                              string `yaml:"signer"`
	Exponent                            int    `yaml:"exponent"`
	Url                                 string `yaml:"url"`
	ChainId                             int64  `yaml:"chainId"`
	SwapContract                        string `yaml:"contract"` // swap contract address
	TokenContract                       string `yaml:"tokenContract"`
	LastScannedTransferEventBlockNumber int64  `yaml:"lastScannedTransferEventBlockNumber"`
	LastScannedSwapEventBlockNumber     int64  `yaml:"lastScannedSwapEventBlockNumber"`
	SwapTopic                           string `yaml:"swapTopic"` // swap topic
	TransferTopic                       string `yaml:"transferTopic"`
	Retry                               Retry  `yaml:"retry"`
}

type Retry struct {
	IntervalRetry   time.Duration `yaml:"intervalRetry"`
	MaxRetry        int           `yaml:"maxRetry"`
	RetryPercentage float64       `yaml:"retryPercentage"`
}

type RelayerConfig struct {
	Network        map[string]Network `yaml:"networks"`
	Type           string             `yaml:"type"`
	ScanInterval   time.Duration      `yaml:"scanInterval"`
	MongoURI       string             `yaml:"mongoURI"`
	DbName         string             `yaml:"dbName"`
	CollectionName string             `yaml:"collectionName"`
}

type Relayer struct {
	Config RelayerConfig
	Client client.Context
	db     database.DBRelayer
	events map[string]subscriber.Service

	qClient   swapmoduletypes.QueryClient
	msgClient swapmoduletypes.MsgClient
}
