package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sharering/shareledger/pkg/swap"
	"github.com/sharering/shareledger/x/gentlemint/client/tests"
	"os"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	netutilts "github.com/sharering/shareledger/testutil/network"
	electoralmoduletypes "github.com/sharering/shareledger/x/electoral/types"
	swapTypes "github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SwapIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	dir     string
	network *network.Network
}

func NewSwapIntegrationTestSuite(cfg network.Config) *SwapIntegrationTestSuite {
	return &SwapIntegrationTestSuite{cfg: cfg}
}

func (s *SwapIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for swap module")

	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir

	networkWithSchema(s.T(), &s.cfg)

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//override the keyring by our keyring information
	s.network.Validators[0].ClientCtx.Keyring = kb

	s.T().Log("setting up integration test suite successfully")
}

func (s *SwapIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "tearing down fail")
	s.network.Cleanup()
	s.T().Log("tearing down integration test suite")
}

func (s *SwapIntegrationTestSuite) TestDeposit() {
	type (
		Num struct {
			D int64
		}
	)
	type (
		TestCase struct {
			d         string
			iAmount   string
			txFee     int
			txCreator string
			oErr      error
			oRes      *sdk.TxResponse

			expectCreatorChange *Num
			expectModuleChange  *Num
		}
	)

	testSuite := []TestCase{
		{
			d:                   "deposit success",
			iAmount:             "10shr",
			txCreator:           netutilts.KeyAccount2,
			txFee:               10,
			oErr:                nil,
			oRes:                nil,
			expectCreatorChange: &Num{-20},
			expectModuleChange:  &Num{10},
		},
		{
			d:                   "deposit fail",
			iAmount:             "1000000000000000shr",
			txCreator:           netutilts.KeyAccount1,
			txFee:               10,
			oErr:                nil,
			expectCreatorChange: &Num{-10},
			expectModuleChange:  &Num{0},
			oRes:                &sdk.TxResponse{Code: sdkerrors.ErrInsufficientFunds.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {

			var balanceCreatorBeforeDeposit sdk.DecCoins
			var balanceModuleBeforeDeposit sdk.DecCoins

			if tc.expectModuleChange != nil {
				out, err := CmdFund(validatorCtx, netutilts.JSONFlag)
				if err != nil {
					s.Failf("query swap fund fail %s out %s", err.Error(), string(out.Bytes()))
				}
				fundRes := swapTypes.QueryBalanceResponse{}
				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)

				balanceModuleBeforeDeposit = sdk.NewDecCoinsFromCoins(*fundRes.GetBalance())

			}
			if tc.expectCreatorChange != nil {
				out := tests.CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator])
				balanceCreatorBeforeDeposit = sdk.NewDecCoinsFromCoins(out.Balances...)
			}

			out, err := CmdDeposit(validatorCtx,
				tc.iAmount,
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txFee),
			)
			if tc.oErr != nil {
				s.NotNilf(err, "this case need return error")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "deposit fail %s", out)
			}

			if tc.expectModuleChange != nil {
				var balanceModuleAfterDeposit sdk.DecCoins
				out, err = CmdFund(validatorCtx, netutilts.JSONFlag)
				if err != nil {
					s.Fail("query swap fund fail", err.Error())
				}

				fundRes := swapTypes.QueryBalanceResponse{}
				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
				if err != nil {
					s.T().Fatalf("can't unmarshal json %s", err)
				}
				balanceModuleAfterDeposit = sdk.NewDecCoinsFromCoins(*fundRes.GetBalance())

				s.Require().Equalf(balanceModuleBeforeDeposit.AmountOf(denom.Base).Add(sdk.NewDec(tc.expectModuleChange.D*denom.ShrExponent)), balanceModuleAfterDeposit.AmountOf(denom.Base), "module balance isn't equal")
			}
			if tc.expectCreatorChange != nil {
				out := tests.CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator])
				s.Require().Equalf(
					balanceCreatorBeforeDeposit.AmountOf(denom.Base).Add(sdk.NewDec(tc.expectCreatorChange.D*denom.ShrExponent)),
					sdk.NewDecFromInt(out.GetBalances().AmountOf(denom.Base)), "creator expect not equal")
			}
		})
	}
}

func (s *SwapIntegrationTestSuite) TestWithDraw() {
	type (
		Num struct {
			D int64
		}
	)
	type (
		TestCase struct {
			d         string
			iAmount   string
			iReceiver string
			txFee     int
			txCreator string
			oErr      error
			oRes      *sdk.TxResponse

			expectReceiverChange *Num
			expectModuleChange   *Num
		}
	)

	testSuite := []TestCase{
		{
			d:                    "withdraw success",
			iAmount:              "10shr",
			txCreator:            netutilts.KeyTreasurer,
			iReceiver:            "shareledger1l5hkf2epa5xmvngnjaasj5dp9jp7ut6s9mrqve",
			txFee:                2,
			oErr:                 nil,
			oRes:                 &sdk.TxResponse{Code: sdkerrors.SuccessABCICode},
			expectReceiverChange: &Num{10},
			expectModuleChange:   &Num{-10},
		},
		{
			d:                    "deposit fail InsufficientFunds",
			iAmount:              "100000000000shr",
			iReceiver:            "shareledger1l5hkf2epa5xmvngnjaasj5dp9jp7ut6s9mrqve",
			txCreator:            netutilts.KeyTreasurer,
			txFee:                10,
			oErr:                 nil,
			expectReceiverChange: &Num{0},
			expectModuleChange:   &Num{0},
			oRes:                 &sdk.TxResponse{Code: sdkerrors.ErrInvalidRequest.ABCICode()},
		},
		{
			d:         "deposit fail creator isn't authority or treasure",
			iAmount:   "100000000000shr",
			iReceiver: "shareledger1l5hkf2epa5xmvngnjaasj5dp9jp7ut6s9mrqve",
			txCreator: netutilts.KeyAccount3,
			txFee:     10,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	_, _ = CmdDeposit(validatorCtx,
		"1000000000000nshr",
		netutilts.JSONFlag,
		netutilts.SkipConfirmation,
		netutilts.MakeByAccount(netutilts.KeyAuthority),
		netutilts.BlockBroadcast,
		netutilts.SHRFee(2),
	)

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			var balanceReceiverBeforeDeposit sdk.DecCoins
			var balanceModuleBeforeDeposit sdk.DecCoins

			if tc.expectModuleChange != nil {
				out, err := CmdFund(validatorCtx, netutilts.JSONFlag)
				if err != nil {
					s.Failf("query swap fund fail %s out %s", err.Error(), string(out.Bytes()))
				}
				fundRes := swapTypes.QueryBalanceResponse{}
				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)

				balanceModuleBeforeDeposit = sdk.NewDecCoinsFromCoins(*fundRes.GetBalance())

			}
			rAddr, _ := sdk.AccAddressFromBech32(tc.iReceiver)
			if tc.expectReceiverChange != nil {

				out := tests.CmdQueryBalance(s.T(), validatorCtx, rAddr)
				balanceReceiverBeforeDeposit = sdk.NewDecCoinsFromCoins(out.Balances...)
			}
			out, err := CmdWithdraw(validatorCtx,
				tc.iReceiver,
				tc.iAmount,
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txFee),
			)
			if tc.oErr != nil {
				s.NotNilf(err, "this case need return error")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "withdraw fail %s", out)
			}

			if tc.expectModuleChange != nil {
				var balanceModuleAfterDeposit sdk.DecCoins
				out, err = CmdFund(validatorCtx, netutilts.JSONFlag)
				if err != nil {
					s.Fail("query swap fund fail", err.Error())
				}

				fundRes := swapTypes.QueryBalanceResponse{}
				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
				if err != nil {
					s.T().Fatalf("can't unmarshal json %s", err)
				}
				balanceModuleAfterDeposit = sdk.NewDecCoinsFromCoins(*fundRes.GetBalance())
				s.Require().Equalf(
					balanceModuleBeforeDeposit.AmountOf(denom.Base).Add(sdk.NewDec(tc.expectModuleChange.D*denom.ShrExponent)).String(),
					balanceModuleAfterDeposit.AmountOf(denom.Base).String(), "module balance isn't equal")
			}
			if tc.expectReceiverChange != nil {
				out := tests.CmdQueryBalance(s.T(), validatorCtx, rAddr)
				s.Require().Equalf(
					balanceReceiverBeforeDeposit.AmountOf(denom.Base).Add(sdk.NewDec(tc.expectReceiverChange.D*denom.ShrExponent)).String(),
					sdk.NewDecFromInt(out.GetBalances().AmountOf(denom.Base)).String(), "receiver expect not equal")
			}
		})
	}
}

func (s *SwapIntegrationTestSuite) TestCancel() {
	type (
		swapArg struct {
			wAmount string
			id      string
		}
		Num struct {
			D int64
		}

		feeSchema struct {
			net    string
			inFee  string
			outFee string
		}

		// cancelArgs mean we will take the element from list of swap request id [f:t]
		cancelArgs struct {
			f int //from
			t int // to
		}

		cancelSuite struct {
			description    string
			initSwapOut    []swapArg
			initFeeSchema  []feeSchema
			approve        int // we will approve first approve number request ID result of initSwapOut func
			cancelArg      cancelArgs
			iTxCreatorSwap string

			iTxCreatorCancel string

			iTxFee string
			oErr   error
			oRes   *sdk.TxResponse

			expectCreatorChange *Num
			expectModuleChange  *Num
		}
	)
	cliCtx := s.network.Validators[0].ClientCtx
	cancelCase := []cancelSuite{
		{
			description: "In case cancel success",
			initSwapOut: []swapArg{
				{
					wAmount: "20shr",
				},
			},

			approve: 0,
			cancelArg: cancelArgs{
				f: 0,
				t: 1,
			},
			iTxCreatorSwap:      netutilts.KeyAccount1,
			iTxCreatorCancel:    netutilts.KeyAccount1,
			iTxFee:              netutilts.SHRFee2,
			oRes:                &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectCreatorChange: &Num{D: -4}, //expect the creator just must pay 4shr for txn fee
			expectModuleChange:  &Num{D: 0},  // expect module amount doest change cause
		},
		{
			description: "In case can not cancel cause there are some swap request was approved",
			initSwapOut: []swapArg{
				{
					wAmount: "20shr",
				},
				{
					wAmount: "30shr",
				},
				{
					wAmount: "10shr",
				},
			},
			approve: 2,
			cancelArg: cancelArgs{
				f: 0,
				t: 2,
			},
			iTxCreatorSwap:      netutilts.KeyAccount1,
			iTxCreatorCancel:    netutilts.KeyAccount1,
			iTxFee:              netutilts.SHRFee2,
			oRes:                &sdk.TxResponse{Code: sdkerrors.ErrInvalidRequest.ABCICode()},
			expectCreatorChange: &Num{D: -128}, //expect the creator just must pay 4shr for txn fee //swap fee constantly 20shr for eth network
			expectModuleChange:  &Num{D: 120},  // expect module amount doest change cause
		},
		{
			description: "In case can not cancel cause the tx creator isn't owner of this request",
			initSwapOut: []swapArg{
				{
					wAmount: "20shr",
				},
				{
					wAmount: "30shr",
				},
				{
					wAmount: "10shr",
				},
			},
			approve: 0,
			cancelArg: cancelArgs{
				f: 0,
				t: 2,
			},
			iTxCreatorSwap:      netutilts.KeyAccount1,
			iTxCreatorCancel:    netutilts.KeyAccount2,
			iTxFee:              netutilts.SHRFee2,
			oRes:                &sdk.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
			expectCreatorChange: &Num{D: -126}, //expect the creator just must pay 4shr for txn fee
			expectModuleChange:  &Num{D: 120},  // expect module amount doest change cause
		},
	}

	for _, ts := range cancelCase {
		s.Run(ts.description, func() {
			_ = s.network.WaitForNextBlock()
			var swapIDs []string

			var balanceCreatorBefore sdk.DecCoins
			var balanceModuleBefore sdk.DecCoins
			if ts.expectModuleChange != nil {
				out, err := CmdFund(cliCtx, netutilts.JSONFlag)
				if err != nil {
					s.Failf("query swap fund fail %s out %s", err.Error(), string(out.Bytes()))
				}
				fundRes := swapTypes.QueryBalanceResponse{}
				err = cliCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
				balanceModuleBefore = sdk.NewDecCoinsFromCoins(*fundRes.GetBalance())

			}
			if ts.expectCreatorChange != nil {
				out := tests.CmdQueryBalance(s.T(), cliCtx, netutilts.Accounts[ts.iTxCreatorSwap])
				balanceCreatorBefore = sdk.NewDecCoinsFromCoins(out.Balances...)
			}

			for _, sw := range ts.initSwapOut {
				out, err := CmdOut(cliCtx,
					"0x7b9039bd633411b48a5a5c4262b5b1a16546d209",
					"eth",
					sw.wAmount,
					ts.iTxFee,
					netutilts.SkipConfirmation,
					netutilts.BlockBroadcast,
					netutilts.MakeByAccount(ts.iTxCreatorSwap))
				if err != nil {
					s.Fail("fail when init the swap out request", err)
				}
				txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
				if txRes.Code != netutilts.ShareLedgerSuccessCode {
					s.Fail("fail when init the swap out request %s", txRes.String())
				}
				log := netutilts.ParseRawLogGetEvent(s.T(), txRes.RawLog)[0]
				attr := log.Events.GetEventByType(s.T(), swapTypes.EventTypeCreateRequest)
				swapIDs = append(swapIDs, attr.Get(s.T(), swapTypes.EventAttrSwapId).Value)
			}
			if ts.approve > 0 {
				appIDs := strings.Join(swapIDs[0:ts.approve], ",")
				out, err := CmdApprove(
					cliCtx,
					netutilts.KeyTreasurer,
					appIDs,
					"eth",
					netutilts.SHRFee2,
					netutilts.SkipConfirmation,
					netutilts.BlockBroadcast,
					netutilts.MakeByAccount(netutilts.KeyApproverRelayer))
				if err != nil {
					s.Fail("fail when approve the swap out request", err)
				}
				txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
				if txRes.Code != netutilts.ShareLedgerSuccessCode {
					s.Fail("fail when init the swap out request %s", txRes.String())
				}

			}

			var cancelIDs string
			cancelIDs = strings.Join(swapIDs[ts.cancelArg.f:ts.cancelArg.t], ",")
			out, err := CmdCancel(cliCtx,
				cancelIDs,
				ts.iTxFee,
				netutilts.MakeByAccount(ts.iTxCreatorCancel),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast)
			if err != nil {
				s.Failf("fail when cancel request", "%s", err)
			}
			txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
			if ts.oErr != nil {
				s.Require().NotNilf(err, "this case require err")
			}
			if ts.oRes != nil {
				if txRes.Code != ts.oRes.Code {
					s.Fail("fail when cancel request", "require cancel code must equal with test case", txRes.String())
				}
			}

			_ = s.network.WaitForNextBlock()
			if ts.expectModuleChange != nil {
				var balanceModuleAfterCancel sdk.DecCoins
				out, err = CmdFund(cliCtx, netutilts.JSONFlag)
				if err != nil {
					s.Fail("query swap fund fail", err.Error())
				}

				fundRes := swapTypes.QueryBalanceResponse{}
				err = cliCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
				if err != nil {
					s.T().Fatalf("can't unmarshal json %s", err)
				}
				balanceModuleAfterCancel = sdk.NewDecCoinsFromCoins(*fundRes.GetBalance())

				s.Require().Equalf(balanceModuleBefore.AmountOf(denom.Base).Add(sdk.NewDec(ts.expectModuleChange.D*denom.ShrExponent)).String(), balanceModuleAfterCancel.AmountOf(denom.Base).String(), "module balance isn't equal")
			}
			if ts.expectCreatorChange != nil {
				out := tests.CmdQueryBalance(s.T(), cliCtx, netutilts.Accounts[ts.iTxCreatorSwap])

				s.Require().Equalf(
					balanceCreatorBefore.AmountOf(denom.Base).Add(sdk.NewDec(ts.expectCreatorChange.D*denom.ShrExponent)),
					sdk.NewDecFromInt(out.GetBalances().AmountOf(denom.Base)), "creator expect not equal")
			}

		})

	}
}

func (s *SwapIntegrationTestSuite) TestReject() {
	type (
		swapArg struct {
			wAmount string
			wFee    string
			id      string
		}
		Num struct {
			D int64
		}

		// cancelArgs mean we will take the element from list of swap request id [f:t]
		rejectArgs struct {
			f int //from
			t int // to
		}

		cancelSuite struct {
			description string
			initSwapOut []swapArg
			approve     int // we will approve first approve number request ID result of initSwapOut func
			rejectArg   rejectArgs

			iTxCreatorSwap   string
			iTxCreatorReject string

			iTxFee string
			oErr   error
			oRes   *sdk.TxResponse

			expectCreatorChange *Num
			expectModuleChange  *Num
		}
	)
	cliCtx := s.network.Validators[0].ClientCtx
	cancelCase := []cancelSuite{
		{
			description: "In case reject success",
			initSwapOut: []swapArg{
				{
					wAmount: "20shr",
				},
			},
			approve: 0,
			rejectArg: rejectArgs{
				f: 0,
				t: 1,
			},
			iTxCreatorSwap:      netutilts.KeyAccount1,
			iTxCreatorReject:    netutilts.KeyApproverRelayer,
			iTxFee:              netutilts.SHRFee2,
			oRes:                &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectCreatorChange: &Num{D: -2},
			expectModuleChange:  &Num{D: 0},
		},
		{
			description: "In case can not reject cause there are some swap request was approved",
			initSwapOut: []swapArg{
				{
					wAmount: "20shr",
				},
				{
					wAmount: "30shr",
				},
				{
					wAmount: "10shr",
				},
			},
			approve: 2,
			rejectArg: rejectArgs{
				f: 0,
				t: 2,
			},
			iTxCreatorSwap:      netutilts.KeyAccount1,
			iTxCreatorReject:    netutilts.KeyApproverRelayer,
			iTxFee:              netutilts.SHRFee2,
			oRes:                &sdk.TxResponse{Code: sdkerrors.ErrInvalidRequest.ABCICode()},
			expectCreatorChange: &Num{D: -126},
			expectModuleChange:  &Num{D: 120}, // expect module amount doest change cause
		},
		{
			description: "In case can not reject cause the tx creator isn't authorizer",
			initSwapOut: []swapArg{
				{
					wAmount: "20shr",
				},
				{
					wAmount: "30shr",
				},
				{
					wAmount: "10shr",
				},
			},
			approve: 0,
			rejectArg: rejectArgs{
				f: 0,
				t: 2,
			},
			iTxCreatorSwap:      netutilts.KeyAccount1,
			iTxCreatorReject:    netutilts.KeyAccount2,
			iTxFee:              netutilts.SHRFee2,
			oRes:                &sdk.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
			expectCreatorChange: nil, //expect the creator just must pay 4shr for txn fee
			expectModuleChange:  nil, // expect module amount doest change cause
		},
	}

	for _, ts := range cancelCase {
		s.Run(ts.description, func() {
			_ = s.network.WaitForNextBlock()
			var swapIDs []string

			var balanceCreatorBefore sdk.DecCoins
			var balanceModuleBefore sdk.DecCoins
			if ts.expectModuleChange != nil {
				out, err := CmdFund(cliCtx, netutilts.JSONFlag)
				if err != nil {
					s.Failf("query swap fund fail %s out %s", err.Error(), string(out.Bytes()))
				}
				fundRes := swapTypes.QueryBalanceResponse{}
				err = cliCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)

				balanceModuleBefore = sdk.NewDecCoinsFromCoins(*fundRes.GetBalance())
			}
			if ts.expectCreatorChange != nil {
				out := tests.CmdQueryBalance(s.T(), cliCtx, netutilts.Accounts[ts.iTxCreatorSwap])
				balanceCreatorBefore = sdk.NewDecCoinsFromCoins(out.Balances...)

			}
			for _, sw := range ts.initSwapOut {
				out, err := CmdOut(cliCtx,
					"0x7b9039bd633411b48a5a5c4262b5b1a16546d209",
					"eth",
					sw.wAmount,
					ts.iTxFee,
					netutilts.SkipConfirmation,
					netutilts.BlockBroadcast,
					netutilts.MakeByAccount(ts.iTxCreatorSwap))
				if err != nil {
					s.Fail("fail when init the swap out request", err)
				}
				txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
				if txRes.Code != netutilts.ShareLedgerSuccessCode {
					s.Fail("fail when init the swap out request %s", txRes.String())
				}
				log := netutilts.ParseRawLogGetEvent(s.T(), txRes.RawLog)[0]
				attr := log.Events.GetEventByType(s.T(), swapTypes.EventTypeCreateRequest)
				swapIDs = append(swapIDs, attr.Get(s.T(), swapTypes.EventAttrSwapId).Value)
			}
			if ts.approve > 0 {
				appIDs := strings.Join(swapIDs[0:ts.approve], ",")
				out, err := CmdApprove(
					cliCtx,
					netutilts.KeyTreasurer,
					appIDs,
					"eth",
					netutilts.SHRFee2,
					netutilts.SkipConfirmation,
					netutilts.BlockBroadcast,
					netutilts.MakeByAccount(netutilts.KeyApproverRelayer))
				if err != nil {
					s.Fail("fail when init the swap out request", err)
				}
				txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
				if txRes.Code != netutilts.ShareLedgerSuccessCode {
					s.Fail("fail when init the swap out request %s", txRes.String())
				}

			}

			var cancelIDs string
			cancelIDs = strings.Join(swapIDs[ts.rejectArg.f:ts.rejectArg.t], ",")
			out, err := CmdReject(cliCtx,
				cancelIDs,
				ts.iTxFee,
				netutilts.MakeByAccount(ts.iTxCreatorReject),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast)

			if ts.oErr != nil {
				s.Require().NotNilf(err, "this case require err")
			}
			txRes := netutilts.ParseStdOut(s.T(), out.Bytes())

			if ts.oRes != nil {
				s.Require().Equalf(ts.oRes.Code, txRes.Code, "cancel error code not equal")
				if txRes.Code != ts.oRes.Code {
				}
			}

			_ = s.network.WaitForNextBlock()
			if ts.expectModuleChange != nil {
				var balanceModuleAfterCancel sdk.DecCoins
				out, err = CmdFund(cliCtx, netutilts.JSONFlag)
				if err != nil {
					s.Fail("query swap fund fail", err.Error())
				}

				fundRes := swapTypes.QueryBalanceResponse{}
				err = cliCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
				if err != nil {
					s.T().Fatalf("can't unmarshal json %s", err)
				}
				balanceModuleAfterCancel = sdk.NewDecCoinsFromCoins(*fundRes.GetBalance())

				s.Require().Equalf(balanceModuleBefore.AmountOf(denom.Base).Add(sdk.NewDec(ts.expectModuleChange.D*denom.ShrExponent)).String(), balanceModuleAfterCancel.AmountOf(denom.Base).String(), "module balance isn't equal")
			}
			if ts.expectCreatorChange != nil {
				out := tests.CmdQueryBalance(s.T(), cliCtx, netutilts.Accounts[ts.iTxCreatorSwap])
				s.Require().Equalf(
					balanceCreatorBefore.AmountOf(denom.Base).Add(sdk.NewDec(ts.expectCreatorChange.D*denom.ShrExponent)),
					sdk.NewDecFromInt(out.GetBalances().AmountOf(denom.Base)), "creator expect not equal")
			}

		})

	}
}

func (s *SwapIntegrationTestSuite) TestApprove() {
	cliCtx := s.network.Validators[0].ClientCtx

	swapInfoResOut, err := CmdSearch(cliCtx, "pending", fmt.Sprintf("--ids=3000,3001"), netutilts.JSONFlag)
	if err != nil {
		s.Failf("get swap fail fail", "err %s res %s", err, swapInfoResOut.String())
	}

	swapRes := swapTypes.QuerySwapResponse{}

	err = cliCtx.Codec.UnmarshalJSON(swapInfoResOut.Bytes(), &swapRes)
	if err != nil {
		s.Fail("fail to get back swap", "err %s out %s", err, swapInfoResOut.String())
	}

	out, err := CmdApprove(cliCtx, netutilts.KeyAccountTestSign, "3000,3001", "eth", netutilts.SkipConfirmation, netutilts.BlockBroadcast, netutilts.MakeByAccount(netutilts.KeyApproverRelayer))
	if err != nil {
		s.Fail("approve fail", "out %s", out.String())
	}
	txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
	if txRes.Code != netutilts.ShareLedgerSuccessCode {
		s.Fail("fail when approve request", "response %s", txRes.String())
	}

	log := netutilts.ParseRawLogGetEvent(s.T(), txRes.RawLog)[0]
	attr := log.Events.GetEventByType(s.T(), swapTypes.EventTypeApproveRequests)
	batchID := attr.Get(s.T(), swapTypes.EventAttrBatchId).Value

	outReader, err := CmdGetBatches(cliCtx, fmt.Sprintf("--ids=%s", batchID), netutilts.JSONFlag)
	if err != nil {
		s.Fail("fail to get back batch", "err %s", err)
	}

	batchesRes := swapTypes.QueryBatchesResponse{}

	err = cliCtx.Codec.UnmarshalJSON(outReader.Bytes(), &batchesRes)
	if err != nil {
		s.Fail("fail to get back batch", "err %s", err)
	}

	//re calculate digest
	schemaResOut, err := CmdGetSchema(cliCtx, "eth")
	if err != nil {
		s.Fail("fail to schema back batch", "err %s", err)
	}
	schemaRes := swapTypes.QuerySchemaResponse{}
	err = cliCtx.Codec.UnmarshalJSON(schemaResOut.Bytes(), &schemaRes)
	if err != nil {
		s.Fail("fail to get schema batch", "err %s", err)
	}

	signDetail := swap.NewSignDetail(swapRes.Swaps, schemaRes.GetSchema())
	digest, err := signDetail.Digest()
	if err != nil {
		s.Fail("fail to get digest", "err %s", err)
	}
	b := batchesRes.GetBatches()[0]

	ks := keyring.NewKeyRingETH(cliCtx.Keyring)
	_, npk, err := ks.Sign(netutilts.KeyAccountTestSign, digest.Bytes())
	if err != nil {
		s.Fail("fail to get signer", "err %s", err)
	}

	sig, _ := hexutil.Decode(b.Signature)
	s.Equalf(true, npk.VerifySignature(digest.Bytes(), sig), "verify sign fail")
	s.Equalf("0xd9a5705095d8c83fc051fde2dda2e47fb81d16ee23f11f9322c0656e6020ee9001f73fd951ad0d1d7d36a59900d2bd481c47477d01fbd1a2a6da5c7f6d78129d1c", b.GetSignature(), "eip sign not same")
	s.Equalf("0xb63b8aa6f75b29271051d9069070d0555f4e6cdaf35e72d69ffcb366a4d47a08", digest.String(), "digest not equal")
}

func networkWithSchema(t *testing.T, cf *network.Config) {
	t.Helper()
	var gen = swapTypes.GenesisState{}
	require.NoError(t, cf.Codec.UnmarshalJSON(cf.GenesisState[swapTypes.ModuleName], &gen))

	elector := electoralmoduletypes.GenesisState{}
	require.NoError(t, cf.Codec.UnmarshalJSON(cf.GenesisState[electoralmoduletypes.ModuleName], &elector))

	elector.AccStateList = []electoralmoduletypes.AccState{
		{
			Key:     "approver" + netutilts.Accounts[netutilts.KeyApproverRelayer].String(),
			Address: netutilts.Accounts[netutilts.KeyApproverRelayer].String(),
			Status:  "active",
		},
		{
			Key:     "relayer" + netutilts.Accounts[netutilts.KeyApproverRelayer].String(),
			Address: netutilts.Accounts[netutilts.KeyApproverRelayer].String(),
			Status:  "active",
		},
	}
	bufEl, err := cf.Codec.MarshalJSON(&elector)
	require.NoError(t, err)

	feeInCoin := sdk.NewCoin(denom.Base, sdk.NewInt(10000000000))
	feeOutCoin := sdk.NewCoin(denom.Base, sdk.NewInt(20000000000))

	gen.Schemas = []swapTypes.Schema{
		{
			Network:          "eth",
			ContractExponent: 2,
			Fee: &swapTypes.Fee{
				In:  &feeInCoin,
				Out: &feeOutCoin,
			},
			Schema: `{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x3","verifyingContract":"0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560","salt":""}}`,
		},
		{
			Network: "bsc",
			Fee: &swapTypes.Fee{
				In:  &feeInCoin,
				Out: &feeOutCoin,
			},
			ContractExponent: 2,
			Schema:           `{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x3","verifyingContract":"0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560","salt":""}}`,
		},
	}
	gen.RequestCount = 4000

	gen.Requests = []swapTypes.Request{
		{
			Id:       3000,
			DestAddr: "0x97b98d335c28f9ad9c123e344a78f00c84146431",
			Amount: sdk.Coin{
				Denom:  denom.Base,
				Amount: sdk.NewInt(3455000000000), //3455 shr
			},
			Status:     "pending",
			SrcNetwork: "shareledger",
			BatchId:    0,
		}, {
			Id:       3001,
			DestAddr: "0x97b98d335c28f9ad9c123e344a78f00c84146431",
			Amount: sdk.Coin{
				Denom:  denom.Base,
				Amount: sdk.NewInt(6733000000000), //6733 shr
			},
			Status:     "pending",
			SrcNetwork: "shareledger",
			BatchId:    0,
		},
	}

	buf, err := cf.Codec.MarshalJSON(&gen)
	require.NoError(t, err)
	cf.GenesisState[swapTypes.ModuleName] = buf
	cf.GenesisState[electoralmoduletypes.ModuleName] = bufEl
}
