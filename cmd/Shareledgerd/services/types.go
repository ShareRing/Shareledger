package services

import (
	"context"
	"time"

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
	Signer           string `yaml:"signer"`
	Url              string `yaml:"url"`
	ChainId          int64  `yaml:"chainId"`
	Contract         string `yaml:"contract"`
	LastScannedBlock int64  `yaml:"lastScannedBlock"`
	Topic            string `yaml:"topic"`
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

	qClient   swapmoduletypes.QueryClient
	msgClient swapmoduletypes.MsgClient
}
