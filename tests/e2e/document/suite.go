//go:build e2e

package document

import (
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/document/types"
	"github.com/stretchr/testify/suite"
)

var firstDoc = types.Document{
	Holder:  "USER-1",
	Issuer:  "shareledger19l9teyc2znfv630sv9gzjc92xurzxcers75xud",
	Proof:   "testProof",
	Data:    "testData",
	Version: 0,
}

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger document module")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg)
	docGenesis := types.GenesisState{
		Documents: []*types.Document{&firstDoc},
	}
	docGenesisBz, err := s.cfg.Codec.MarshalJSON(&docGenesis)
	s.Require().NoError(err)
	s.cfg.GenesisState[types.ModuleName] = docGenesisBz

	s.network = network.New(s.T(), s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
	s.Require().NoError(s.network.WaitForNextBlock())
}
func (s *E2ETestSuite) TearDownSuite() {
	s.network.Cleanup()
}
