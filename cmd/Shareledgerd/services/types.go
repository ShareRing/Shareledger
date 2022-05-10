package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	crypto2 "github.com/sharering/shareledger/pkg/crypto"
	swaputils "github.com/sharering/shareledger/pkg/swap"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	"time"
)

type BatchDetail struct {
	Batch      swapmoduletypes.Batch
	Requests   []swapmoduletypes.Request
	SignSchema swapmoduletypes.SignSchema
}

func (b BatchDetail) Validate() error {
	if len(b.Requests) == 0 {
		return fmt.Errorf("requests is empty")
	}
	if len(b.SignSchema.Schema) == 0 {
		return fmt.Errorf("schema is empty")
	}
	return nil
}

func (b BatchDetail) Digest() (common.Hash, error) {
	var hash common.Hash
	var signFormatData apitypes.TypedData
	if err := json.Unmarshal([]byte(b.SignSchema.Schema), &signFormatData); err != nil {
		return hash, err
	}
	data, err := swaputils.BuildTypedData(signFormatData, b.Requests)
	if err != nil {
		return hash, err
	}
	hash, err = crypto2.Keccak256HashEIP712(data)
	return hash, err
}

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
