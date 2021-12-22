package tests

import (
	"github.com/ShareRing/Shareledger/testutil/network"
	"github.com/ShareRing/Shareledger/x/electoral/client/tests"
	"github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
)

type GentlemintIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewGentlemintIntegrationTestSuite(cf network.Config) *GentlemintIntegrationTestSuite {
	return &GentlemintIntegrationTestSuite{
		cfg: cf,
	}
}

func (s *GentlemintIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for document module")

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.T().Log("setting up document data....")

	//Enroll ACCOUNT_OPERATOR
	out, _ := tests.ExCmdEnrollLoader(
		s.network.Validators[0].ClientCtx,
		s.T(),
		s.network.Accounts[network.KeyLoader].String(),
		network.MakeByAccount(network.KeyAuthority),
		network.SkipConfirmation(),
		network.BlockBroadcast(),
		network.SHRFee2(),
	)
	_ = s.network.WaitForNextBlock()
	res := network.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode, res.Code, "init operator fail %v", res.String())

	s.T().Log("setting up integration test suite successfully")

}
func (s *GentlemintIntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
}

func (s *GentlemintIntegrationTestSuite) TestLoadSHR() {
	validatorCtx := s.network.Validators[0].ClientCtx
	s.Run("load_shr_success", func() {
		stdOut, err := CmdLoadSHR(validatorCtx, s.network.Accounts[network.KeyEmpty1].String(), "100",
			network.MakeByAccount(network.KeyAuthority),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		s.NoErrorf(err, "load shr should not error %v", err)
		txResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "load shr must success %v", txResponse)

		out, err := testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyEmpty1])
		s.NoErrorf(err, "query balance should not error %v", err)
		balRes := types.QueryAllBalancesResponse{}
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)

		s.Equalf("100", balRes.GetBalances().AmountOf("shr").String(), "balance of user is not equal after load shr %s", balRes.GetBalances().String())

	})
}
