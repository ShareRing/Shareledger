package tests

import (
	"fmt"
	testutil2 "github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	netutilts "github.com/sharering/shareledger/testutil/network"
	types2 "github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/stretchr/testify/suite"
	"os"

	"github.com/sharering/shareledger/x/electoral/client/tests"
)

const MaximumSHRSupply = 2090000

type GentlemintIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
	dir     string
}

func NewGentlemintIntegrationTestSuite(cf network.Config) *GentlemintIntegrationTestSuite {
	return &GentlemintIntegrationTestSuite{
		cfg: cf,
	}
}

func (s *GentlemintIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for document module")

	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//override the keyring by our keyring information
	s.network.Validators[0].ClientCtx.Keyring = kb

	s.T().Log("setting up document data....")

	//Enroll ACCOUNT_OPERATOR
	out, _ := tests.ExCmdEnrollLoader(
		s.network.Validators[0].ClientCtx,
		s.T(),
		netutilts.Accounts[netutilts.KeyLoader].String(),
		netutilts.MakeByAccount(netutilts.KeyAuthority),
		netutilts.SkipConfirmation(),
		netutilts.BlockBroadcast(),
		netutilts.SHRFee2(),
	)
	_ = s.network.WaitForNextBlock()
	res := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init operator fail %v", res.String())

	s.T().Log("setting up integration test suite successfully")

}
func (s *GentlemintIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir))
	s.network.Cleanup()
	s.T().Log("tearing down integration test suite")
}

func (s *GentlemintIntegrationTestSuite) TestLoadSHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("load_shr_success", func() {
		balResBeforeLoad := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyEmpty1])
		fee := 2
		stdOut, err := CmdLoadSHR(validatorCtx, netutilts.Accounts[netutilts.KeyEmpty1].String(),
			"100shr",
			netutilts.MakeByAccount(netutilts.KeyAuthority),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
			netutilts.SHRFee(fee),
		)
		_ = s.network.WaitForNextBlock()
		balRes := types.QueryAllBalancesResponse{}

		s.NoErrorf(err, "load shr should not error %v", err)
		txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txResponse.Code, "load shr must success %v", txResponse)
		balRes = CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyEmpty1])
		s.NoErrorf(err, "query balance should not error %v", err)
		expectSHR := (100 - int64(fee)) * denom.ShrExponent
		s.T().Logf("the balance before load and after load %s, %s", balResBeforeLoad, balRes)
		s.Equalf(fmt.Sprintf("%d", expectSHR), balRes.GetBalances().AmountOf(denom.Base).String(), "balance of user is not equal after load shr %s", balRes.GetBalances().String())

		balRes = CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyAuthority])
		s.NoErrorf(err, "query balance should not error %v", err)
		expectSHR = 9998 * denom.ShrExponent
		s.Require().Equalf(fmt.Sprintf("%d", expectSHR), balRes.GetBalances().AmountOf(denom.Base).String(), "authority balance after make transaction is not equal")

	})

	s.Run("load_shr_loader_isn't_authority", func() {
		stdOut, err := CmdLoadSHR(validatorCtx, netutilts.Accounts[netutilts.KeyEmpty1].String(), "100shr",
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
			netutilts.SHRFee2(),
		)
		s.NoErrorf(err, "load shr should not error %v", err)
		txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerErrorCodeUnauthorized, txResponse.Code, "load shr mustn't success %v", txResponse)
	})

	s.Run("load_shr_but_supply_reach_to_limit", func() {
		stdOut, err := CmdLoadSHR(validatorCtx, netutilts.Accounts[netutilts.KeyEmpty1].String(), "4396000043shr",
			netutilts.MakeByAccount(netutilts.KeyAuthority),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
			netutilts.SHRFee2(),
			netutilts.JSONFlag(),
		)
		_ = s.network.WaitForNextBlock()
		s.NoErrorf(err, "load shr should not error %v", err)
		txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerErrorCodeMaxSupply, txResponse.Code, "load shr must not success %v", txResponse)

	})
}

func (s *GentlemintIntegrationTestSuite) TestBurnSHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("burn_shr_success", func() {

		balRes := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyTreasurer], netutilts.JSONFlag())

		beforeBurnSHR := balRes.GetBalances().AmountOf(denom.Base)
		stdOut, err := CmdBurnSHR(validatorCtx, "11shr", netutilts.SHRFee2(), netutilts.MakeByAccount(netutilts.KeyTreasurer), netutilts.BlockBroadcast(), netutilts.SkipConfirmation())
		s.NoErrorf(err, "should not error %v", err)
		txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txResponse.Code, "the response code should success %v", txResponse)
		_ = s.network.WaitForNextBlock()

		balRes = CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyTreasurer])
		afterBurnSHR := balRes.GetBalances().AmountOf(denom.Base)
		s.Require().NoError(err)

		expectAmount := beforeBurnSHR.Sub(sdk.NewInt((11 + 2) * denom.ShrExponent))

		//s.Equal(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(expectAmount*denom.USDExponent)), sdk.NewCoin(denom.BaseUSD, accBalance.Balances.AmountOf(denom.BaseUSD)), accBalance.Balances)

		s.Equalf(expectAmount.String(), afterBurnSHR.String(), "the expect shr should be equal")
	})
}

func (s *GentlemintIntegrationTestSuite) TestBurnSHRP() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("burn_shrp_success", func() {

		balRes := types.QueryAllBalancesResponse{}
		balRes = CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyTreasurer])

		beforeBurnSHRP := balRes.GetBalances().AmountOf(denom.BaseUSD)
		stdOut, err := CmdBurnSHRP(validatorCtx, "11shrp", netutilts.SHRFee2(), netutilts.MakeByAccount(netutilts.KeyTreasurer), netutilts.BlockBroadcast(), netutilts.SkipConfirmation())
		s.NoErrorf(err, "should not error %v", err)
		txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txResponse.Code, "the response code should success %v", txResponse)
		_ = s.network.WaitForNextBlock()

		balRes = types.QueryAllBalancesResponse{}
		balRes = CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyTreasurer])
		s.NoErrorf(err, "query balance should not error %v", err)

		afterBurnSHR := balRes.GetBalances().AmountOf(denom.BaseUSD)
		s.Require().NoError(err)

		expectAmount := beforeBurnSHRP.Sub(sdk.NewInt(11 * denom.USDExponent))

		s.Equalf(expectAmount.String(), afterBurnSHR.String(), "the expect shrp should be equal")
	})
}

func (s *GentlemintIntegrationTestSuite) TestBuySHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("buy_shr_success", func() {
		stdOut, err := CmdBuySHR(validatorCtx, "211",
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(), netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyAccount4))

		s.NoErrorf(err, "should no error %v", err)
		txRes := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txRes.Code, "should no error %v", txRes.String())
		_ = s.network.WaitForNextBlock()
		balRes := types.QueryAllBalancesResponse{}
		out, err := testutil.QueryBalancesExec(validatorCtx, netutilts.Accounts[netutilts.KeyAccount4])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)

		expectSHR := 10209 * denom.ShrExponent

		s.Equalf(fmt.Sprintf("%d", expectSHR), balRes.GetBalances().AmountOf(denom.Base).String(), "shr amount must be equal")
		s.Equalf(fmt.Sprintf("%d", 9894), balRes.GetBalances().AmountOf(denom.BaseUSD).String(), "shrp must be equal")

	})

	s.Run("buy_shr_not_insufficient_shrp", func() {
		stdOut, err := CmdBuySHR(validatorCtx, "2223",
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyEmpty1))
		s.NoErrorf(err, "should no error %v", err)
		txRes := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerCoinInsufficient, txRes.Code, "should got error %v", stdOut.String())

	})
}

func (s *GentlemintIntegrationTestSuite) TestLoadSHRP() {
	validatorCtx := s.network.Validators[0].ClientCtx
	s.Run("load_shrp_success", func() {
		balRes := types.QueryAllBalancesResponse{}
		out, err := testutil.QueryBalancesExec(validatorCtx, netutilts.Accounts[netutilts.KeyAccount1])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)
		s.T().Logf("before load %s", balRes.GetBalances())
		out, err = CmdLoadSHRP(validatorCtx, netutilts.Accounts[netutilts.KeyAccount1].String(), "129shrp",
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
			netutilts.MakeByAccount(netutilts.KeyLoader),
			netutilts.SHRFee2())
		s.NoErrorf(err, "should not error %v", err)
		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "the txn must be success %v", txnResponse.String())

		s.T().Logf("make txn address %v", netutilts.Accounts[netutilts.KeyLoader].String())
		s.T().Logf("make account1 address %v", netutilts.Accounts[netutilts.KeyAccount1].String())

		_ = s.network.WaitForNextBlock()
		balRes = types.QueryAllBalancesResponse{}
		out, err = testutil.QueryBalancesExec(validatorCtx, netutilts.Accounts[netutilts.KeyAccount1])
		s.NoErrorf(err, "query balance should not error %v", err)
		err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
		s.Require().NoError(err)
		shrpExpect := 229 * denom.USDExponent
		shrExpect := 9998 * denom.ShrExponent
		s.Equalf(fmt.Sprintf("%d", shrpExpect), balRes.GetBalances().AmountOf(denom.BaseUSD).String(), "shrp should equal")
		s.Equalf(fmt.Sprintf("%d", shrExpect), balRes.GetBalances().AmountOf(denom.Base).String(), "shr should equal")
	})
}

func (s *GentlemintIntegrationTestSuite) TestSendSHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	var (
		stdOut      testutil2.BufferWriter
		err         error
		txnResponse sdk.TxResponse
	)

	stdOut, err = CmdLoadSHR(validatorCtx, netutilts.Accounts[netutilts.KeyEmpty4].String(), "100shr",
		netutilts.SkipConfirmation(),
		netutilts.MakeByAccount(netutilts.KeyAuthority),
		netutilts.BlockBroadcast(),
		netutilts.SHRFee2(),
	)
	s.NoErrorf(err, "load shr error %v", err)
	txnResponse = netutilts.ParseStdOut(s.T(), stdOut.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "txn response got error %v", txnResponse.String())
	_ = s.network.WaitForNextBlock()

	s.Run("send_shr_success", func() {
		stdOut, err = CmdSendSHR(validatorCtx, netutilts.Accounts[netutilts.KeyEmpty3].String(), "1220shr",
			netutilts.JSONFlag(),
			netutilts.SHRFee4(),
			netutilts.BlockBroadcast(),
			netutilts.SkipConfirmation(),
			netutilts.MakeByAccount(netutilts.KeyMillionaire))
		s.NoErrorf(err, "send shr error %v", err)
		txnResponse = netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "txn response got error %v", txnResponse.String())
		_ = s.network.WaitForNextBlock()
		balanceKeyEmpty1 := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyEmpty3])
		balanceKeyMillionaire := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyMillionaire])

		emptyExpect := 1220 * denom.ShrExponent
		millionExpect := 998776 * denom.ShrExponent

		s.Equalf(fmt.Sprintf("%d", emptyExpect), balanceKeyEmpty1.GetBalances().AmountOf(denom.Base).String(), "balance of receiver no equal")
		s.Equalf(fmt.Sprintf("%d", millionExpect), balanceKeyMillionaire.GetBalances().AmountOf(denom.Base).String(), "balance of sender no equal")

	})
	s.Run("send_shr_fail_insufficient_shr", func() {

		stdOut, err = CmdSendSHR(validatorCtx, netutilts.Accounts[netutilts.KeyEmpty2].String(), "900997shr",
			netutilts.JSONFlag(),
			netutilts.SHRFee4(),
			netutilts.BlockBroadcast(),
			netutilts.SkipConfirmation(),
			netutilts.MakeByAccount(netutilts.KeyEmpty4))
		s.NoErrorf(err, "send shr error %v", err)
		txnResponse = netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerCoinInsufficient, txnResponse.Code, "txn response got error %v", txnResponse.String())
		_ = s.network.WaitForNextBlock()
		balanceKeyEmpty2 := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyEmpty2])
		balanceKeyEmpty4 := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyEmpty4])

		empty4Expect := 94 * denom.ShrExponent
		s.Equalf("0", balanceKeyEmpty2.GetBalances().AmountOf(denom.Base).String(), "balance of receiver no equal")
		s.Equalf(fmt.Sprintf("%d", empty4Expect), balanceKeyEmpty4.GetBalances().AmountOf(denom.Base).String(), "balance of sender no equal")

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
		stdOut, err = CmdSendSHRP(validatorCtx, netutilts.Accounts[netutilts.KeyEmpty5].String(), "123shrp",
			netutilts.SkipConfirmation(),
			netutilts.SHRFee4(),
			netutilts.BlockBroadcast(),
			netutilts.MakeByAccount(netutilts.KeyMillionaire))

		s.NoErrorf(err, "send shrp error %v", err)
		txnResponse = netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "txn response got error %v", txnResponse.String())
		_ = s.network.WaitForNextBlock()

		balanceKeyEmpty5 := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyEmpty5])
		balanceKeyMillionaire := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[netutilts.KeyMillionaire])

		expectCoinEmpty := 123 * denom.USDExponent
		expectSHR := 999877 * denom.USDExponent
		s.Equalf(fmt.Sprintf("%d", expectCoinEmpty), balanceKeyEmpty5.GetBalances().AmountOf(denom.BaseUSD).String(), "balance of receiver no equal")
		s.Equalf(fmt.Sprintf("%d", expectSHR), balanceKeyMillionaire.GetBalances().AmountOf(denom.BaseUSD).String(), "balance of sender no equal")
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
		stdOut, err = CmdSetExchangeRate(validatorCtx, "12", netutilts.SHRFee2(), netutilts.MakeByAccount(netutilts.KeyTreasurer), netutilts.SkipConfirmation(), netutilts.BlockBroadcast())
		s.NoErrorf(err, "set exchange rate error %v", err)
		txnResponse = netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "txn response got error %v", txnResponse.String())
		_ = s.network.WaitForNextBlock()
		_ = s.network.WaitForNextBlock()
		stdOut, err = CmdGetExchangeRate(validatorCtx, netutilts.JSONFlag())
		s.NoErrorf(err, "get exchange rate error %v", err)

		exchangeRate := types2.QueryExchangeRateResponse{}
		err = validatorCtx.Codec.UnmarshalJSON(stdOut.Bytes(), &exchangeRate)
		s.NoErrorf(err, "should no error %v", err)
		s.Equalf(fmt.Sprintf("%d", 12), exchangeRate.GetRate(), "the rate is not equal")

	})
}
