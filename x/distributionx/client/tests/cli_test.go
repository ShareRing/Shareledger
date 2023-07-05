package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"
	"github.com/sharering/shareledger/testutil/network"
)

var (
	runOnce    = sync.Once{}
	genesisDir = "../../../../testutil/distributionx_genesis.json"
)

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	cfg.GenesisState = mustLoadTestGenesis(genesisDir)
	suite.Run(t, NewDistributionXIntegrationTestSuite(&cfg))
}

func mustLoadTestGenesis(dir string) map[string]json.RawMessage {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	bytes, err := os.ReadFile(absDir)
	if err != nil {
		panic(err)
	}
	var genesis tmtypes.GenesisDoc
	err = tmjson.Unmarshal(bytes, &genesis)
	if err != nil {
		panic(err)
	}
	var appState map[string]json.RawMessage
	err = json.Unmarshal(genesis.AppState, &appState)
	if err != nil {
		panic(err)
	}
	return appState

}
