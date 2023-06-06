package asset

import (
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/asset/types"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

var asset1 = &types.Asset{
	Creator: "shareledger19ac3d6cwqwpzvaxr4xv9kfduwtyswad88fjgw4",
	Hash:    []byte{123},
	UUID:    "UUID",
	Status:  true,
	Rate:    10,
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger asset module")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg)
	assetGenesis := &types.GenesisState{
		Assets: []*types.Asset{asset1},
	}
	s.cfg.GenesisState[types.ModuleName] = s.cfg.Codec.MustMarshalJSON(assetGenesis)

	s.network = network.New(s.T(), s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
	s.NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestSuite) TearDownSuite() {
	s.network.Cleanup()
}
