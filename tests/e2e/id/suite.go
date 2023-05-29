//go:build e2e

package id

import (
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/id/types"
	"github.com/stretchr/testify/suite"
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

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg)

	// update Id module state
	genesisState := s.cfg.GenesisState
	var idGenesis types.GenesisState
	s.Require().NoError(s.cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &idGenesis))
	idGenesis.IDs = []*types.Id{&id1}
	idGenesisBz, err := s.cfg.Codec.MarshalJSON(&idGenesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = idGenesisBz
	s.cfg.GenesisState = genesisState

	s.network = network.New(s.T(), s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestSuite) TearDownSuite() {
}
