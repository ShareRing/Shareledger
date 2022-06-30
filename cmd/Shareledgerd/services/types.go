package services

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

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
	Url           string `yaml:"rpc_url"`
	ChainId       int64  `yaml:"chain_id"`
	SwapContract  string `yaml:"swap_contract_address"` // swap contract address
	TokenContract string `yaml:"token_contract_address"`
	SwapTopic     string `yaml:"swap_topic"` // swap topic
	TransferTopic string `yaml:"transfer_topic"`
	Retry         Retry  `yaml:"retry"`
}

type Retry struct {
	IntervalRetry   time.Duration `yaml:"interval_retry"`
	RetryPercentage float64       `yaml:"retry_percentage"`
}

type RelayerConfig struct {
	Network      map[string]Network `yaml:"network"`
	Type         string             `yaml:"type"`
	AutoApprove  bool               `yaml:"auto_approve"`
	ScanInterval time.Duration      `yaml:"scan_interval"`
	MongoURI     string             `yaml:"mongo_url"`
	DbName       string             `yaml:"database_name"`
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
