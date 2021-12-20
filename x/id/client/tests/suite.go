package tests

import (
	"github.com/ShareRing/Shareledger/testutil/network"
	"github.com/ShareRing/Shareledger/x/electoral/client/tests"
	"github.com/stretchr/testify/suite"
)

type IDIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewIDIntegrationTestSuite(cfg network.Config) *IDIntegrationTestSuite {
	return &IDIntegrationTestSuite{cfg: cfg}
}

func (s *IDIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for booking module")

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.T().Log("setting up document data....")

	out := tests.ExCmdEnrollAccountOperator(
		s.network.Validators[0].ClientCtx,
		s.T(),
		s.network.Accounts[network.KeyOperator].String(),
		network.MakeByAccount(network.KeyAuthority),
		network.SkipConfirmation(),
		network.BlockBroadcast(),
		network.SHRFee2(),
	)
	_ = s.network.WaitForNextBlock()
	res := network.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode, res.Code, "init operator fail %v", res.String())

	out = tests.ExCmdEnrollDocIssuer(
		s.network.Validators[0].ClientCtx,
		s.T(),
		s.network.Accounts[network.KeyAccount1].String(),
		network.SHRFee2(),
		network.MakeByAccount(network.KeyOperator),
		network.SkipConfirmation(),
	)
	_ = s.network.WaitForNextBlock()
	res = network.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode, res.Code, "init doc issuer fail %v", res.String())

	s.T().Log("setting up integration test suite successfully")
}
func (s *IDIntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
}
