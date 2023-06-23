//go:build e2e

package id

import (
	"fmt"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/id/types"
	"github.com/stretchr/testify/suite"
	"path/filepath"
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	cfg.NumValidators = 1
	return &E2ETestSuite{cfg: cfg}
}

var id1 = types.Id{
	Id: "Id1",
	Data: &types.BaseID{
		IssuerAddress: "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak",
		BackupAddress: "BackupAddress",
		OwnerAddress:  "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr",
		ExtraData:     "ExtraData",
	},
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for id module")
	// the nodeDir, and moniker hard code at here in cosmos-sdk:
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:398
	// So just reuse it
	rootDir := s.T().TempDir()
	moniker := fmt.Sprintf("node%d", s.cfg.NumValidators-1)
	// TestingGenesis should use the same KeyringDir as validator KeyringDir
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:400
	nodeDir := filepath.Join(rootDir, moniker, "simcli")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg, nodeDir, moniker)
	s.Require().NotNil(kr)

	// update Id module state
	idGenesis := types.GenesisState{
		IDs: []*types.Id{&id1},
	}
	s.cfg.GenesisState[types.ModuleName] = s.cfg.Codec.MustMarshalJSON(&idGenesis)

	s.network = network.New(s.T(), rootDir, s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
	s.Require().NoError(s.network.WaitForNextBlock())
}
