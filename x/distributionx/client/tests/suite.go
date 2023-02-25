package tests

import (
	"os"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	netutilts "github.com/sharering/shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
)

type DistributionxIntergrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	dir     string
	network *network.Network
}

func NewDistributionxIntergrationTestSuite(cf *network.Config) *DistributionxIntergrationTestSuite {
	return &DistributionxIntergrationTestSuite{
		Suite:   suite.Suite{},
		cfg:     network.Config{},
		dir:     "",
		network: &network.Network{},
	}
}

func (s *DistributionxIntergrationTestSuite) SetupSuite() {
	s.T().Log("setting up intergration test suite for distributionx module")

	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	// override the keyring
	s.network.Validators[0].ClientCtx.Keyring = kb
	s.T().Log("setting up intergration test suite successfully")
}

func (s *DistributionxIntergrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "cleanup test case fails")
	s.network.Cleanup()
	s.T().Log("tearing down intergration test suite")
}

func (s *DistributionxIntergrationTestSuite) TestCase1() {

}
func (s *DistributionxIntergrationTestSuite) TestCase2() {

}
