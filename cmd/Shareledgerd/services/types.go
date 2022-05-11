package services

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	"time"
)

const flagConfigPath = "config"

var supportedTypes = map[string]struct{}{
	"in":  {},
	"out": {},
}

type processFunc func(ctx context.Context, network string) error

type Network struct {
	Signer   string `yaml:"signer"`
	Url      string `yaml:"url"`
	ChainId  int64  `yaml:"chainId"`
	Contract string `yaml:"contract"`
}

type RelayerConfig struct {
	Network      map[string]Network `yaml:"networks"`
	Type         string             `yaml:"type"`
	ScanInterval time.Duration      `yaml:"scanInterval"`
}

type Relayer struct {
	Config RelayerConfig
	Client client.Context

	qClient   swapmoduletypes.QueryClient
	msgClient swapmoduletypes.MsgClient
}
