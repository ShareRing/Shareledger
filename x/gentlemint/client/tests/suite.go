package tests

import (
	"fmt"
	testutil2 "github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	netutilts "github.com/sharering/shareledger/testutil/network"
	gentleminttypes "github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/stretchr/testify/suite"
	"os"
	"strings"

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
		netutilts.SkipConfirmation,
		netutilts.BlockBroadcast,
		netutilts.SHRFee2,
	)
	s.Require().NoError(s.network.WaitForNextBlock())
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
	type (
		Num struct {
			D int
		}
	)
	testSuite := []struct {
		d                    string
		iLoadTarget          string
		iAmount              string
		txCreator            string
		txFee                int
		oErr                 error
		oRes                 *sdk.TxResponse
		expectTargetBalance  *Num
		expectCreatorBalance *Num
	}{
		{
			d:                    "load_shr_success",
			iLoadTarget:          netutilts.Accounts[netutilts.KeyEmpty1].String(),
			iAmount:              "100shr",
			txCreator:            netutilts.KeyAuthority,
			txFee:                2,
			oErr:                 nil,
			oRes:                 &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectTargetBalance:  &Num{98},
			expectCreatorBalance: &Num{9998},
		},
		{
			d:           "load_shr_but_supply_reach_to_limit",
			iLoadTarget: netutilts.Accounts[netutilts.KeyEmpty1].String(),
			iAmount:     "4396000043shr",
			txCreator:   netutilts.KeyAuthority,
			txFee:       2,
			oErr:        nil,
			oRes:        &sdk.TxResponse{Code: gentleminttypes.ErrBaseSupplyExceeded.ABCICode()},
		},
		{
			d:           "load_shr_loader_isn't_authority",
			iLoadTarget: netutilts.Accounts[netutilts.KeyEmpty1].String(),
			iAmount:     "100shr",
			txCreator:   netutilts.KeyAccount1,
			txFee:       2,
			oErr:        nil,
			oRes:        &sdk.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			stdOut, err := CmdLoadSHR(validatorCtx, tc.iLoadTarget,
				tc.iAmount,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txFee),
			)
			if tc.oErr != nil {
				s.Require().NotNilf(err, "this case require err")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "load shr must success %v", txResponse)
			}
			if tc.expectTargetBalance != nil {
				a, _ := sdk.AccAddressFromBech32(tc.iLoadTarget)
				balRes := CmdQueryBalance(s.T(), validatorCtx, a)
				s.Equalf(fmt.Sprintf("%d", int64(tc.expectTargetBalance.D)*denom.ShrExponent), balRes.GetBalances().AmountOf(denom.Base).String(), "balance of user is not equal after load shr %s", balRes.GetBalances().String())
			}
			if tc.expectCreatorBalance != nil {
				balRes := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator])
				s.Require().Equalf(fmt.Sprintf("%d", int64(tc.expectCreatorBalance.D)*denom.ShrExponent), balRes.GetBalances().AmountOf(denom.Base).String(), "authority balance after make transaction is not equal")
			}
		})
	}
}

func (s *GentlemintIntegrationTestSuite) TestBurnSHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	type (
		Num struct {
			D int
		}
	)
	testSuite := []struct {
		d                 string
		iAmount           string
		txCreator         string
		txFee             int
		oErr              error
		oRes              *sdk.TxResponse
		expectSubtractNum *Num
	}{
		{
			d:                 "burn_shr_success",
			iAmount:           "11shr",
			txCreator:         netutilts.KeyTreasurer,
			txFee:             2,
			oErr:              nil,
			oRes:              &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectSubtractNum: &Num{D: -13},
		},
	}

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			balResBeforeBurn := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator], netutilts.JSONFlag)
			stdOut, err := CmdBurnSHR(validatorCtx,
				tc.iAmount,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txFee),
			)
			if tc.oErr != nil {
				s.Require().NotNilf(err, "this case require err")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "load shr must success %v", txResponse)
			}

			if tc.expectSubtractNum != nil {
				balRes := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator])
				expectAmount := balResBeforeBurn.
					GetBalances().
					AmountOf(denom.Base).
					Add(sdk.NewInt(int64(tc.expectSubtractNum.D) * denom.ShrExponent))
				s.Equalf(expectAmount.String(), balRes.GetBalances().AmountOf(denom.Base).String(), "the expect shr should be equal")
			}
		})
	}
}

func (s *GentlemintIntegrationTestSuite) TestBurnSHRP() {
	validatorCtx := s.network.Validators[0].ClientCtx

	type (
		Num struct {
			D int
		}
	)
	testSuite := []struct {
		d                 string
		iAmount           string
		txCreator         string
		txFee             int
		oErr              error
		oRes              *sdk.TxResponse
		expectSubtractNum *Num
	}{
		{
			d:                 "burn_shrp_success",
			iAmount:           "11shrp",
			txCreator:         netutilts.KeyTreasurer,
			txFee:             2,
			oErr:              nil,
			oRes:              &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectSubtractNum: &Num{D: -11},
		},
	}

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			balResBeforeBurn := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator], netutilts.JSONFlag)
			stdOut, err := CmdBurnSHR(validatorCtx,
				tc.iAmount,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txFee),
			)
			if tc.oErr != nil {
				s.Require().NotNilf(err, "this case require err")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "load shr must success %v", txResponse)
			}

			if tc.expectSubtractNum != nil {
				balRes := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator])
				expectAmount := balResBeforeBurn.
					GetBalances().
					AmountOf(denom.BaseUSD).
					Add(sdk.NewInt(int64(tc.expectSubtractNum.D) * denom.USDExponent))
				s.Equalf(expectAmount.String(), balRes.GetBalances().AmountOf(denom.BaseUSD).String(), "the expect shr should be equal")
			}
		})
	}

}

func (s *GentlemintIntegrationTestSuite) TestBuySHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	type (
		Num struct {
			D float64
		}
	)
	testSuite := []struct {
		d                   string
		iAmount             string
		txCreator           string
		txFee               int
		oErr                error
		oRes                *sdk.TxResponse
		expectChangeNumSHR  *Num
		expectChangeNumSHRP *Num
	}{
		{
			d:                   "buy_shr_success",
			iAmount:             "211",
			txCreator:           netutilts.KeyAccount4,
			txFee:               2,
			oErr:                nil,
			oRes:                &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectChangeNumSHR:  &Num{D: 209},
			expectChangeNumSHRP: &Num{D: -1.06},
		},
	}

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			balResBeforeBuy := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator], netutilts.JSONFlag)
			stdOut, err := CmdBuySHR(validatorCtx,
				tc.iAmount,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txFee),
			)
			if tc.oErr != nil {
				s.Require().NotNilf(err, "this case require err")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "load shr must success %v", txResponse)
			}
			balResAfterBuy := CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator])
			if tc.expectChangeNumSHR != nil {
				expectSHR := balResBeforeBuy.GetBalances().AmountOf(denom.Base).Add(sdk.NewInt(int64(tc.expectChangeNumSHR.D) * denom.ShrExponent))
				s.Equalf(expectSHR.String(), balResAfterBuy.GetBalances().AmountOf(denom.Base).String(), "shr amount must be equal")
			}
			if tc.expectChangeNumSHRP != nil {
				shrpSub := tc.expectChangeNumSHRP.D * float64(denom.USDExponent)
				expectSHRP := balResBeforeBuy.GetBalances().AmountOf(denom.BaseUSD).Add(sdk.NewInt(int64(shrpSub)))
				s.Equalf(expectSHRP.String(), balResAfterBuy.GetBalances().AmountOf(denom.BaseUSD).String(), "shrp amount must be equal")
			}
		})
	}

}

func (s *GentlemintIntegrationTestSuite) TestLoadSHRP() {
	validatorCtx := s.network.Validators[0].ClientCtx

	type (
		Num struct {
			D int
		}
	)
	testSuite := []struct {
		d                       string
		iLoadTarget             string
		iAmount                 string
		txCreator               string
		txFee                   int
		oErr                    error
		oRes                    *sdk.TxResponse
		expectBalanceSHRChange  *Num
		expectBalanceSHRPChange *Num
	}{
		{
			d:                       "load_shrp_success",
			iLoadTarget:             netutilts.Accounts[netutilts.KeyAccount3].String(),
			iAmount:                 "129shrp",
			txCreator:               netutilts.KeyLoader,
			txFee:                   2,
			oErr:                    nil,
			oRes:                    &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectBalanceSHRChange:  &Num{-2},
			expectBalanceSHRPChange: &Num{129},
		},
	}
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			a, _ := sdk.AccAddressFromBech32(tc.iLoadTarget)
			balResBeforeLoad := CmdQueryBalance(s.T(), validatorCtx, a)
			stdOut, err := CmdLoadSHRP(validatorCtx, tc.iLoadTarget,
				tc.iAmount,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txFee),
			)
			if tc.oErr != nil {
				s.Require().NotNilf(err, "this case require err")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "load shr must success %v", txResponse)
			}
			a, _ = sdk.AccAddressFromBech32(tc.iLoadTarget)
			balResAfterLoad := CmdQueryBalance(s.T(), validatorCtx, a)
			if tc.expectBalanceSHRPChange != nil {
				shrEx := balResBeforeLoad.GetBalances().AmountOf(denom.BaseUSD).Add(sdk.NewInt(int64(tc.expectBalanceSHRPChange.D) * denom.USDExponent))
				s.Equalf(shrEx.String(), balResAfterLoad.GetBalances().AmountOf(denom.BaseUSD).String(), "shrp should equal")

			}
			if tc.expectBalanceSHRChange != nil {
				shrEx := balResBeforeLoad.GetBalances().AmountOf(denom.Base).Add(sdk.NewInt(int64(tc.expectBalanceSHRChange.D) * denom.ShrExponent))
				s.Equalf(shrEx.String(), balResAfterLoad.GetBalances().AmountOf(denom.Base).String(), "shr should equal")
			}
		})
	}

}

func (s *GentlemintIntegrationTestSuite) TestSendSHR() {
	validatorCtx := s.network.Validators[0].ClientCtx

	var (
		stdOut      testutil2.BufferWriter
		err         error
		txnResponse sdk.TxResponse
	)
	type (
		Num struct {
			D int64
		}
	)
	testSuite := []struct {
		d                       string
		iSender                 string
		iReceiver               string
		iAmount                 string
		txFee                   int
		txCreator               string
		oErr                    error
		oRes                    *sdk.TxResponse
		expectSenderSHRChange   *Num
		expectReceiverSHRChange *Num
	}{
		{
			d:                       "send_shr_success",
			iSender:                 netutilts.Accounts[netutilts.KeyMillionaire].String(),
			iReceiver:               netutilts.Accounts[netutilts.KeyEmpty3].String(),
			iAmount:                 "1220shr",
			txFee:                   4,
			txCreator:               netutilts.KeyMillionaire,
			oErr:                    nil,
			oRes:                    &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectSenderSHRChange:   &Num{D: -1224},
			expectReceiverSHRChange: &Num{D: 1220},
		},
		{
			d:                       "send_shr_fail_insufficient_shr",
			iSender:                 netutilts.Accounts[netutilts.KeyAccount2].String(),
			iReceiver:               netutilts.Accounts[netutilts.KeyEmpty3].String(),
			iAmount:                 "10022000shr",
			txFee:                   4,
			txCreator:               netutilts.KeyAccount2,
			oErr:                    nil,
			oRes:                    &sdk.TxResponse{Code: sdkerrors.ErrInsufficientFunds.ABCICode()},
			expectSenderSHRChange:   &Num{D: -4},
			expectReceiverSHRChange: &Num{D: 0},
		},
	}

	for _, tc := range testSuite {
		senderAddr, _ := sdk.AccAddressFromBech32(tc.iSender)
		senderBalanceBeforeSend := CmdQueryBalance(s.T(), validatorCtx, senderAddr)
		receiverAddr, _ := sdk.AccAddressFromBech32(tc.iReceiver)
		receiverBalanceBeforeSend := CmdQueryBalance(s.T(), validatorCtx, receiverAddr)

		stdOut, err = CmdSendSHR(validatorCtx, tc.iReceiver, tc.iAmount,
			netutilts.JSONFlag,
			netutilts.SHRFee(tc.txFee),
			netutilts.BlockBroadcast,
			netutilts.SkipConfirmation,
			netutilts.MakeByAccount(tc.txCreator))

		if tc.oErr != nil {
			s.Require().NotNilf(err, "this case require error")
		}

		if tc.oRes != nil {
			txnResponse = netutilts.ParseStdOut(s.T(), stdOut.Bytes())
			s.Equalf(tc.oRes.Code, txnResponse.Code, "txn response got error %v", txnResponse.String())
		}
		senderAddr, _ = sdk.AccAddressFromBech32(tc.iSender)
		senderBalanceAfterSend := CmdQueryBalance(s.T(), validatorCtx, senderAddr)
		receiverAddr, _ = sdk.AccAddressFromBech32(tc.iReceiver)
		receiverBalanceAfterSend := CmdQueryBalance(s.T(), validatorCtx, receiverAddr)
		if tc.expectReceiverSHRChange != nil {
			expect := receiverBalanceBeforeSend.GetBalances().AmountOf(denom.Base).Add(sdk.NewInt(tc.expectReceiverSHRChange.D * denom.ShrExponent))
			s.Require().Equalf(expect.String(), receiverBalanceAfterSend.GetBalances().AmountOf(denom.Base).String(), "receiver shr balance isn't equal")
		}

		if tc.expectSenderSHRChange != nil {
			expect := senderBalanceBeforeSend.GetBalances().AmountOf(denom.Base).Add(sdk.NewInt(tc.expectSenderSHRChange.D * denom.ShrExponent))
			s.Require().Equalf(expect.String(), senderBalanceAfterSend.GetBalances().AmountOf(denom.Base).String(), "sender shr balance isn't equal")
		}

	}

}

func (s *GentlemintIntegrationTestSuite) TestSendSHRP() {
	var (
		stdOut      testutil2.BufferWriter
		err         error
		txnResponse sdk.TxResponse
	)
	validatorCtx := s.network.Validators[0].ClientCtx

	type (
		Num struct {
			D int64
		}
	)
	testSuite := []struct {
		d                        string
		iSender                  string
		iReceiver                string
		iAmount                  string
		txFee                    int
		txCreator                string
		oErr                     error
		oRes                     *sdk.TxResponse
		expectSenderSHRPChange   *Num
		expectReceiverSHRPChange *Num
	}{
		{
			d:                        "send_shrp_success",
			iSender:                  netutilts.Accounts[netutilts.KeyMillionaire].String(),
			iReceiver:                netutilts.Accounts[netutilts.KeyEmpty5].String(),
			iAmount:                  "123shrp",
			txFee:                    4,
			txCreator:                netutilts.KeyMillionaire,
			oErr:                     nil,
			oRes:                     &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectSenderSHRPChange:   &Num{D: -123},
			expectReceiverSHRPChange: &Num{D: 123},
		},
	}

	for _, tc := range testSuite {
		senderAddr, _ := sdk.AccAddressFromBech32(tc.iSender)
		senderBalanceBeforeSend := CmdQueryBalance(s.T(), validatorCtx, senderAddr)
		receiverAddr, _ := sdk.AccAddressFromBech32(tc.iReceiver)
		receiverBalanceBeforeSend := CmdQueryBalance(s.T(), validatorCtx, receiverAddr)

		stdOut, err = CmdSendSHRP(validatorCtx, tc.iReceiver, tc.iAmount,
			netutilts.JSONFlag,
			netutilts.SHRFee(tc.txFee),
			netutilts.BlockBroadcast,
			netutilts.SkipConfirmation,
			netutilts.MakeByAccount(tc.txCreator))

		if tc.oErr != nil {
			s.Require().NotNilf(err, "this case require error")
		}

		if tc.oRes != nil {
			txnResponse = netutilts.ParseStdOut(s.T(), stdOut.Bytes())
			s.Equalf(tc.oRes.Code, txnResponse.Code, "txn response got error %v", txnResponse.String())
		}
		senderAddr, _ = sdk.AccAddressFromBech32(tc.iSender)
		senderBalanceAfterSend := CmdQueryBalance(s.T(), validatorCtx, senderAddr)
		receiverAddr, _ = sdk.AccAddressFromBech32(tc.iReceiver)
		receiverBalanceAfterSend := CmdQueryBalance(s.T(), validatorCtx, receiverAddr)
		if tc.expectReceiverSHRPChange != nil {
			expect := receiverBalanceBeforeSend.GetBalances().AmountOf(denom.BaseUSD).Add(sdk.NewInt(tc.expectReceiverSHRPChange.D * denom.USDExponent))
			s.Require().Equalf(expect.String(), receiverBalanceAfterSend.GetBalances().AmountOf(denom.BaseUSD).String(), "receiver shrp balance isn't equal")
		}

		if tc.expectSenderSHRPChange != nil {
			expect := senderBalanceBeforeSend.GetBalances().AmountOf(denom.BaseUSD).Add(sdk.NewInt(tc.expectSenderSHRPChange.D * denom.USDExponent))
			s.Require().Equalf(expect.String(), senderBalanceAfterSend.GetBalances().AmountOf(denom.BaseUSD).String(), "sender shrp balance isn't equal")
		}

	}

}

func (s *GentlemintIntegrationTestSuite) TestSetExchangeRate() {

	var (
		stdOut      testutil2.BufferWriter
		err         error
		txnResponse sdk.TxResponse
	)
	validatorCtx := s.network.Validators[0].ClientCtx

	testSuite := []struct {
		d         string
		iRate     string
		txFee     int
		txCreator string
		oErr      error
		oRes      *sdk.TxResponse
		oRate     string
	}{
		{
			d:         "set_exchange_rate",
			iRate:     "12",
			txFee:     2,
			txCreator: netutilts.KeyTreasurer,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oRate:     "12",
		},
	}

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			stdOut, err = CmdSetExchangeRate(validatorCtx,
				tc.iRate, netutilts.SHRFee(tc.txFee),
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast)
			if tc.oErr != nil {
				s.Require().NotNilf(err, "error is required in this case")
			}
			if tc.oRes != nil {
				txnResponse = netutilts.ParseStdOut(s.T(), stdOut.Bytes())
				s.Equalf(tc.oRes.Code, txnResponse.Code, "txn response got error %v", txnResponse.String())
			}
			if strings.TrimSpace(tc.oRate) != "" {
				stdOut, err = CmdGetExchangeRate(validatorCtx, netutilts.JSONFlag)
				exchangeRate := gentleminttypes.QueryExchangeRateResponse{}
				err = validatorCtx.Codec.UnmarshalJSON(stdOut.Bytes(), &exchangeRate)
				s.NoErrorf(err, "should no error %v", err)
				s.Equalf(tc.oRate, exchangeRate.GetRate(), "the rate is not equal")
			}
		})
	}
}
