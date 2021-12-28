package tests

import (
	testutil2 "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/electoral/client/tests"
	types2 "github.com/sharering/shareledger/x/gentlemint/types"
)

const MaximumSHRSupply = 2090000

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
		_ = s.network.WaitForNextBlock()
		balRes := types.QueryAllBalancesResponse{}

		s.NoErrorf(err, "load shr should not error %v", err)
		txResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "load shr must success %v", txResponse)
		out, err := testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyEmpty1])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)

		s.Equalf("100", balRes.GetBalances().AmountOf("shr").String(), "balance of user is not equal after load shr %s", balRes.GetBalances().String())

		out, err = testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyAuthority])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)

		s.Require().Equalf("9996", balRes.GetBalances().AmountOf("shr").String(), "authority balance after make transaction is not equal")

	})

	s.Run("load_shr_loader_isn't_authority", func() {
		stdOut, err := CmdLoadSHR(validatorCtx, s.network.Accounts[network.KeyEmpty1].String(), "100",
			network.MakeByAccount(network.KeyAccount1),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		s.NoErrorf(err, "load shr should not error %v", err)
		txResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code, "load shr mustn't success %v", txResponse)
	})

	s.Run("load_shr_but_supply_reach_to_limit", func() {
		stdOut, err := CmdLoadSHR(validatorCtx, s.network.Accounts[network.KeyEmpty1].String(), "4396000043",
			network.MakeByAccount(network.KeyAuthority),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
			network.SHRFee2(),
			network.JSONFlag(),
		)
		_ = s.network.WaitForNextBlock()
		s.NoErrorf(err, "load shr should not error %v", err)
		txResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerErrorCodeInvalidRequest, txResponse.Code, "load shr must not success %v", txResponse)

	})
}

func (s *GentlemintIntegrationTestSuite) TestBurnSHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("burn_shr_success", func() {

		balRes := types.QueryAllBalancesResponse{}
		out, err := testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyTreasurer])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)
		beforeBurnSHR := balRes.GetBalances().AmountOf("shr")
		stdOut, err := CmdBurnSHR(validatorCtx, "11", network.SHRFee2(), network.MakeByAccount(network.KeyTreasurer), network.BlockBroadcast(), network.SkipConfirmation())
		s.NoErrorf(err, "should not error %v", err)
		txResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "the response code should success %v", txResponse)
		_ = s.network.WaitForNextBlock()

		balRes = types.QueryAllBalancesResponse{}
		out, err = testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyTreasurer])
		s.NoErrorf(err, "query balance should not error %v", err)

		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		afterBurnSHR := balRes.GetBalances().AmountOf("shr")
		s.Require().NoError(err)

		expectAmount := beforeBurnSHR.Sub(sdk.NewInt(11 + 2))

		s.Equalf(expectAmount.String(), afterBurnSHR.String(), "the expect shr should be equal")
	})
}

func (s *GentlemintIntegrationTestSuite) TestBurnSHRP() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("burn_shrp_success", func() {

		balRes := types.QueryAllBalancesResponse{}
		out, err := testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyTreasurer])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)
		beforeBurnSHR := balRes.GetBalances().AmountOf("shrp")
		stdOut, err := CmdBurnSHRP(validatorCtx, "11", network.SHRFee2(), network.MakeByAccount(network.KeyTreasurer), network.BlockBroadcast(), network.SkipConfirmation())
		s.NoErrorf(err, "should not error %v", err)
		txResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "the response code should success %v", txResponse)
		_ = s.network.WaitForNextBlock()

		balRes = types.QueryAllBalancesResponse{}
		out, err = testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyTreasurer])
		s.NoErrorf(err, "query balance should not error %v", err)

		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		afterBurnSHR := balRes.GetBalances().AmountOf("shrp")
		s.Require().NoError(err)

		expectAmount := beforeBurnSHR.Sub(sdk.NewInt(11))

		s.Equalf(expectAmount.String(), afterBurnSHR.String(), "the expect shr should be equal")
	})
}

func (s *GentlemintIntegrationTestSuite) TestBuyCent() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("buy_cent_success", func() {
		stdOut, err := CmdBuyCent(validatorCtx, "10", network.SkipConfirmation(), network.BlockBroadcast(), network.SHRFee2(), network.MakeByAccount(network.KeyAccount2))
		s.NoErrorf(err, "should no error %v", err)
		txRes := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txRes.Code, "should no error %v", txRes.String())
		_ = s.network.WaitForNextBlock()
		balRes := types.QueryAllBalancesResponse{}
		out, err := testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyAccount2])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)
		s.Equalf("1000", balRes.GetBalances().AmountOf("cent").String(), "cent amount must be equal")
		s.Equalf("9998", balRes.GetBalances().AmountOf("shr").String(), "shr amount must be equal")
		s.Equalf("90", balRes.GetBalances().AmountOf("shrp").String(), "shrp must be equal")

	})

	s.Run("buy_cent_not_insufficient_shrp", func() {
		stdOut, err := CmdBuyCent(validatorCtx, "1121000000", network.SkipConfirmation(), network.BlockBroadcast(), network.SHRFee2(), network.MakeByAccount(network.KeyEmpty3))
		s.NoErrorf(err, "should no error %v", err)
		txRes := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerCoinInsufficient, txRes.Code, "should got error %v", txRes.String())

	})
}

func (s *GentlemintIntegrationTestSuite) TestBuySHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("buy_shr_success", func() {
		stdOut, err := CmdBuySHR(validatorCtx, "211",
			network.SkipConfirmation(),
			network.BlockBroadcast(), network.SHRFee2(),
			network.MakeByAccount(network.KeyAccount4))

		s.NoErrorf(err, "should no error %v", err)
		txRes := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txRes.Code, "should no error %v", txRes.String())
		_ = s.network.WaitForNextBlock()
		balRes := types.QueryAllBalancesResponse{}
		out, err := testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyAccount4])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)
		s.Equalf("94", balRes.GetBalances().AmountOf("cent").String(), "cent amount must be equal")
		s.Equalf("10209", balRes.GetBalances().AmountOf("shr").String(), "shr amount must be equal")
		s.Equalf("98", balRes.GetBalances().AmountOf("shrp").String(), "shrp must be equal")

	})

	s.Run("buy_shr_not_insufficient_shrp", func() {
		stdOut, err := CmdBuySHR(validatorCtx, "2223",
			network.SkipConfirmation(),
			network.BlockBroadcast(),
			network.SHRFee2(),
			network.MakeByAccount(network.KeyEmpty1))
		s.NoErrorf(err, "should no error %v", err)
		txRes := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerCoinInsufficient, txRes.Code, "should got error %v", stdOut.String())

	})
}

func (s *GentlemintIntegrationTestSuite) TestLoadSHRP() {
	validatorCtx := s.network.Validators[0].ClientCtx
	s.Run("load_shrp_success", func() {
		balRes := types.QueryAllBalancesResponse{}
		out, err := testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyAccount1])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)
		s.T().Logf("before load %s", balRes.GetBalances())
		out, err = CmdLoadSHRP(validatorCtx, s.network.Accounts[network.KeyAccount1].String(), "129",
			network.SkipConfirmation(),
			network.BlockBroadcast(),
			network.MakeByAccount(network.KeyLoader),
			network.SHRFee2())
		s.NoErrorf(err, "should not error %v", err)
		txnResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txnResponse.Code, "the txn must be success %v", txnResponse.String())

		s.T().Logf("make txn address %v", s.network.Accounts[network.KeyLoader].String())
		s.T().Logf("make account1 address %v", s.network.Accounts[network.KeyAccount1].String())

		_ = s.network.WaitForNextBlock()
		balRes = types.QueryAllBalancesResponse{}
		out, err = testutil.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyAccount1])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)

		s.Equalf("229", balRes.GetBalances().AmountOf("shrp").String(), "shrp should equal")
		s.Equalf("9999", balRes.GetBalances().AmountOf("shr").String(), "shr should equal")
	})
}

func (s *GentlemintIntegrationTestSuite) TestSendSHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	var (
		stdOut      testutil2.BufferWriter
		err         error
		txnResponse sdk.TxResponse
	)

	stdOut, err = CmdLoadSHR(validatorCtx, s.network.Accounts[network.KeyEmpty4].String(), "100",
		network.SkipConfirmation(),
		network.MakeByAccount(network.KeyAuthority),
		network.BlockBroadcast(),
		network.SHRFee2(),
	)
	s.NoErrorf(err, "load shr error %v", err)
	txnResponse = network.ParseStdOut(s.T(), stdOut.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode, txnResponse.Code, "txn response got error %v", txnResponse.String())
	_ = s.network.WaitForNextBlock()

	s.Run("send_shr_success", func() {
		stdOut, err = CmdSendSHR(validatorCtx, s.network.Accounts[network.KeyEmpty3].String(), "1220",
			network.JSONFlag(),
			network.SHRFee4(),
			network.BlockBroadcast(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyMillionaire))
		s.NoErrorf(err, "send shr error %v", err)
		txnResponse = network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txnResponse.Code, "txn response got error %v", txnResponse.String())
		_ = s.network.WaitForNextBlock()
		balanceKeyEmpty1 := CmdQueryBalance(s.T(), validatorCtx, s.network.Accounts[network.KeyEmpty3])
		balanceKeyMillionaire := CmdQueryBalance(s.T(), validatorCtx, s.network.Accounts[network.KeyMillionaire])

		s.Equalf("1220", balanceKeyEmpty1.GetBalances().AmountOf("shr").String(), "balance of receiver no equal")
		s.Equalf("998776", balanceKeyMillionaire.GetBalances().AmountOf("shr").String(), "balance of sender no equal")

	})
	s.Run("send_shr_fail_insufficient_shr", func() {
		stdOut, err = CmdSendSHR(validatorCtx, s.network.Accounts[network.KeyEmpty2].String(), "97",
			network.JSONFlag(),
			network.SHRFee4(),
			network.BlockBroadcast(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyEmpty4))
		s.NoErrorf(err, "send shr error %v", err)
		txnResponse = network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerCoinInsufficient, txnResponse.Code, "txn response got error %v", txnResponse.String())
		_ = s.network.WaitForNextBlock()
		balanceKeyEmpty2 := CmdQueryBalance(s.T(), validatorCtx, s.network.Accounts[network.KeyEmpty2])
		balanceKeyEmpty4 := CmdQueryBalance(s.T(), validatorCtx, s.network.Accounts[network.KeyEmpty4])

		s.Equalf("0", balanceKeyEmpty2.GetBalances().AmountOf("shr").String(), "balance of receiver no equal")
		s.Equalf("96", balanceKeyEmpty4.GetBalances().AmountOf("shr").String(), "balance of sender no equal")

	})

}

func (s *GentlemintIntegrationTestSuite) TestSendSHRP() {
	var (
		stdOut      testutil2.BufferWriter
		err         error
		txnResponse sdk.TxResponse
	)
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("send_shrp_success", func() {
		stdOut, err = CmdSendSHRP(validatorCtx, s.network.Accounts[network.KeyEmpty5].String(), "123",
			network.SkipConfirmation(),
			network.SHRFee4(),
			network.BlockBroadcast(),
			network.MakeByAccount(network.KeyMillionaire))

		s.NoErrorf(err, "send shrp error %v", err)
		txnResponse = network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txnResponse.Code, "txn response got error %v", txnResponse.String())
		_ = s.network.WaitForNextBlock()

		balanceKeyEmpty5 := CmdQueryBalance(s.T(), validatorCtx, s.network.Accounts[network.KeyEmpty5])
		balanceKeyMillionaire := CmdQueryBalance(s.T(), validatorCtx, s.network.Accounts[network.KeyMillionaire])

		s.Equalf("123", balanceKeyEmpty5.GetBalances().AmountOf("shrp").String(), "balance of receiver no equal")
		s.Equalf("999877", balanceKeyMillionaire.GetBalances().AmountOf("shrp").String(), "balance of sender no equal")
	})
}

func (s *GentlemintIntegrationTestSuite) TestSetExchangeRate() {

	var (
		stdOut      testutil2.BufferWriter
		err         error
		txnResponse sdk.TxResponse
	)
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("set_exchange_rate", func() {
		stdOut, err = CmdSetExchangeRate(validatorCtx, "12", network.SHRFee2(), network.MakeByAccount(network.KeyTreasurer), network.SkipConfirmation(), network.BlockBroadcast())
		s.NoErrorf(err, "set exchange rate error %v", err)
		txnResponse = network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txnResponse.Code, "txn response got error %v", txnResponse.String())
		_ = s.network.WaitForNextBlock()
		_ = s.network.WaitForNextBlock()
		stdOut, err = CmdGetExchangeRate(validatorCtx, network.JSONFlag())
		s.NoErrorf(err, "get exchange rate error %v", err)

		exchangeRate := types2.QueryGetExchangeRateResponse{}
		err = validatorCtx.Codec.UnmarshalJSON(stdOut.Bytes(), &exchangeRate)
		s.NoErrorf(err, "should no error %v", err)
		s.Equalf(float64(12), exchangeRate.GetRate(), "the rate is not equal")

	})
}
