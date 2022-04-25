package tests

import (
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	netutilts "github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/gentlemint/client/tests"
	swapTypes "github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/stretchr/testify/suite"
	"os"
	"strings"
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
	s.T().Log("setting up integration test suite for booking module")

	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//override the keyring by our keyring information
	s.network.Validators[0].ClientCtx.Keyring = kb

	//s.seedingSwapRequest(7, nil, []string{"1", "2", "3"}, nil)

	s.T().Log("setting up integration test suite successfully")
}

//seedingSwapRequest this function helps us to prepare the swapping request for testing
// numberRq is total request that uou want to create
// pending is array of requestID that you want it become pending. this mean you do nothing this id
// approved is array of requestID that you want this become approved status by calling approve func
// rejected same about but that is used for rejected....
func (s *SwapIntegrationTestSuite) seedingSwapRequest(numberRq int, pending, approved, rejected []string) {
	cliCtx := s.network.Validators[0].ClientCtx
	for i := 0; i < numberRq; i++ {
		_, err := CmdOut(
			cliCtx,
			"0x7b9039bd633411b48a5a5c4262b5b1a16546d209",
			"etherium",
			"10shr",
			"2shr",
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.SHRFee2,
			netutilts.SkipConfirmation,
			netutilts.SyncBroadcast)
		if err != nil {
			s.Fail("crate swap out fail %s", err)
		}
	}
	_ = s.network.WaitForNextBlock()

	approvedIds := strings.Join(approved, ",")
	out, err := CmdApprove(
		cliCtx,
		"some_random_hash",
		approvedIds, netutilts.SHRFee2,
		netutilts.SkipConfirmation,
		netutilts.BlockBroadcast)
	if err != nil {
		s.Fail("crate swap out fail %s", err)
	}
	txOut := netutilts.ParseStdOut(s.T(), out.Bytes())
	if txOut.Code != netutilts.ShareLedgerSuccessCode {
		s.T().Error("fail to approve")
	}
}

func (s *SwapIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "tearing down fail")
	s.network.Cleanup()
	s.T().Log("tearing down integration test suite")
}

//func (s *SwapIntegrationTestSuite) TestDeposit() {
//	type (
//		Num struct {
//			D int64
//		}
//	)
//	type (
//		TestCase struct {
//			d         string
//			iAmount   string
//			txFee     int
//			txCreator string
//			oErr      error
//			oRes      *sdk.TxResponse
//
//			expectCreatorChange *Num
//			expectModuleChange  *Num
//		}
//	)
//
//	testSuite := []TestCase{
//		{
//			d:                   "deposit success",
//			iAmount:             "10shr",
//			txCreator:           netutilts.KeyAccount2,
//			txFee:               10,
//			oErr:                nil,
//			oRes:                nil,
//			expectCreatorChange: &Num{-20},
//			expectModuleChange:  &Num{10},
//		},
//		{
//			d:                   "deposit fail",
//			iAmount:             "1000000000000000shr",
//			txCreator:           netutilts.KeyAccount1,
//			txFee:               10,
//			oErr:                nil,
//			expectCreatorChange: &Num{-10},
//			expectModuleChange:  &Num{0},
//			oRes:                &sdk.TxResponse{Code: sdkerrors.ErrInsufficientFunds.ABCICode()},
//		},
//	}
//	validatorCtx := s.network.Validators[0].ClientCtx
//	for _, tc := range testSuite {
//		s.Run(tc.d, func() {
//
//			var balanceCreatorBeforeDeposit sdk.DecCoins
//			var balanceModuleBeforeDeposit sdk.DecCoins
//
//			if tc.expectModuleChange != nil {
//				out, err := CmdFund(validatorCtx, netutilts.JSONFlag)
//				if err != nil {
//					s.Failf("query swap fund fail %s out %s", err.Error(), string(out.Bytes()))
//				}
//				fundRes := swapTypes.QueryFundResponse{}
//				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
//
//				for _, c := range fundRes.Available {
//					if c.Denom == denom.Shr {
//						balanceModuleBeforeDeposit = sdk.NewDecCoins(*c)
//					}
//				}
//
//			}
//			if tc.expectCreatorChange != nil {
//				out := tests.CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator])
//				balanceCreatorBeforeDeposit = sdk.NewDecCoinsFromCoins(out.Balances...)
//			}
//
//			out, err := CmdDeposit(validatorCtx,
//				tc.iAmount,
//				netutilts.JSONFlag,
//				netutilts.SkipConfirmation,
//				netutilts.MakeByAccount(tc.txCreator),
//				netutilts.BlockBroadcast,
//				netutilts.SHRFee(tc.txFee),
//			)
//			if tc.oErr != nil {
//				s.NotNilf(err, "this case need return error")
//			}
//			if tc.oRes != nil {
//				txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
//				s.Equalf(tc.oRes.Code, txResponse.Code, "deposit fail %s", out)
//			}
//
//			if tc.expectModuleChange != nil {
//				var balanceModuleAfterDeposit sdk.DecCoins
//				out, err = CmdFund(validatorCtx, netutilts.JSONFlag)
//				if err != nil {
//					s.Fail("query swap fund fail", err.Error())
//				}
//
//				fundRes := swapTypes.QueryFundResponse{}
//				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
//				if err != nil {
//					s.T().Fatalf("can't unmarshal json %s", err)
//				}
//				for _, c := range fundRes.Available {
//					if c.Denom == denom.Shr {
//						balanceModuleAfterDeposit = sdk.NewDecCoins(*c)
//					}
//				}
//				s.Require().Equalf(balanceModuleBeforeDeposit.AmountOf(denom.Shr).Add(sdk.NewDec(tc.expectModuleChange.D)), balanceModuleAfterDeposit.AmountOf(denom.Shr), "module balance isn't equal")
//			}
//			if tc.expectCreatorChange != nil {
//				out := tests.CmdQueryBalance(s.T(), validatorCtx, netutilts.Accounts[tc.txCreator])
//				s.Require().Equalf(
//					balanceCreatorBeforeDeposit.AmountOf(denom.Base).Add(sdk.NewDec(tc.expectCreatorChange.D*denom.ShrExponent)),
//					sdk.NewDecFromInt(out.GetBalances().AmountOf(denom.Base)), "creator expect not equal")
//			}
//		})
//	}
//}

//func (s *SwapIntegrationTestSuite) TestWithDraw() {
//	type (
//		Num struct {
//			D int64
//		}
//	)
//	type (
//		TestCase struct {
//			d         string
//			iAmount   string
//			iReceiver string
//			txFee     int
//			txCreator string
//			oErr      error
//			oRes      *sdk.TxResponse
//
//			expectReceiverChange *Num
//			expectModuleChange   *Num
//		}
//	)
//
//	testSuite := []TestCase{
//		{
//			d:                    "withdraw success",
//			iAmount:              "10shr",
//			txCreator:            netutilts.KeyTreasurer,
//			iReceiver:            "shareledger1l5hkf2epa5xmvngnjaasj5dp9jp7ut6s9mrqve",
//			txFee:                2,
//			oErr:                 nil,
//			oRes:                 &sdk.TxResponse{Code: sdkerrors.SuccessABCICode},
//			expectReceiverChange: &Num{10},
//			expectModuleChange:   &Num{-10},
//		},
//		{
//			d:                    "deposit fail InsufficientFunds",
//			iAmount:              "100000000000shr",
//			iReceiver:            "shareledger1l5hkf2epa5xmvngnjaasj5dp9jp7ut6s9mrqve",
//			txCreator:            netutilts.KeyTreasurer,
//			txFee:                10,
//			oErr:                 nil,
//			expectReceiverChange: &Num{0},
//			expectModuleChange:   &Num{0},
//			oRes:                 &sdk.TxResponse{Code: sdkerrors.ErrInsufficientFunds.ABCICode()},
//		},
//		{
//			d:         "deposit fail creator isn't authority or treasure",
//			iAmount:   "100000000000shr",
//			iReceiver: "shareledger1l5hkf2epa5xmvngnjaasj5dp9jp7ut6s9mrqve",
//			txCreator: netutilts.KeyAccount3,
//			txFee:     10,
//			oErr:      nil,
//			oRes:      &sdk.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
//		},
//	}
//	validatorCtx := s.network.Validators[0].ClientCtx
//	for _, tc := range testSuite {
//		s.Run(tc.d, func() {
//			var balanceReceiverBeforeDeposit sdk.DecCoins
//			var balanceModuleBeforeDeposit sdk.DecCoins
//
//			if tc.expectModuleChange != nil {
//				out, err := CmdFund(validatorCtx, netutilts.JSONFlag)
//				if err != nil {
//					s.Failf("query swap fund fail %s out %s", err.Error(), string(out.Bytes()))
//				}
//				fundRes := swapTypes.QueryFundResponse{}
//				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
//
//				for _, c := range fundRes.Available {
//					if c.Denom == denom.Shr {
//						balanceModuleBeforeDeposit = sdk.NewDecCoins(*c)
//					}
//				}
//
//			}
//			rAddr, _ := sdk.AccAddressFromBech32(tc.iReceiver)
//			if tc.expectReceiverChange != nil {
//
//				out := tests.CmdQueryBalance(s.T(), validatorCtx, rAddr)
//				balanceReceiverBeforeDeposit = sdk.NewDecCoinsFromCoins(out.Balances...)
//			}
//			out, err := CmdWithdraw(validatorCtx,
//				tc.iReceiver,
//				tc.iAmount,
//				netutilts.JSONFlag,
//				netutilts.SkipConfirmation,
//				netutilts.MakeByAccount(tc.txCreator),
//				netutilts.BlockBroadcast,
//				netutilts.SHRFee(tc.txFee),
//			)
//			if tc.oErr != nil {
//				s.NotNilf(err, "this case need return error")
//			}
//			if tc.oRes != nil {
//				txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
//				s.Equalf(tc.oRes.Code, txResponse.Code, "deposit fail %s", out)
//			}
//
//			if tc.expectModuleChange != nil {
//				var balanceModuleAfterDeposit sdk.DecCoins
//				out, err = CmdFund(validatorCtx, netutilts.JSONFlag)
//				if err != nil {
//					s.Fail("query swap fund fail", err.Error())
//				}
//
//				fundRes := swapTypes.QueryFundResponse{}
//				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
//				if err != nil {
//					s.T().Fatalf("can't unmarshal json %s", err)
//				}
//				for _, c := range fundRes.Available {
//					if c.Denom == denom.Shr {
//						balanceModuleAfterDeposit = sdk.NewDecCoins(*c)
//					}
//				}
//				s.Require().Equalf(
//					balanceModuleBeforeDeposit.AmountOf(denom.Shr).Add(sdk.NewDec(tc.expectModuleChange.D)).String(),
//					balanceModuleAfterDeposit.AmountOf(denom.Shr).String(), "module balance isn't equal")
//			}
//			if tc.expectReceiverChange != nil {
//				out := tests.CmdQueryBalance(s.T(), validatorCtx, rAddr)
//				s.Require().Equalf(
//					balanceReceiverBeforeDeposit.AmountOf(denom.Base).Add(sdk.NewDec(tc.expectReceiverChange.D*denom.ShrExponent)).String(),
//					sdk.NewDecFromInt(out.GetBalances().AmountOf(denom.Base)).String(), "receiver expect not equal")
//			}
//		})
//	}
//}

func (s *SwapIntegrationTestSuite) TestCancel() {
	type (
		swapArg struct {
			wAmount string
			wFee    string
			id      string
		}
		Num struct {
			D int64
		}
		cancelArgs struct {
			ids []string
		}
		cancelSuite struct {
			description string
			initSwapOut []swapArg
			approve     int // we will approve first approve number request ID result of initSwapOut func
			cancelArg   cancelArgs
			iTxCreator  string
			iTxFee      string
			oErr        error
			oRes        *sdk.TxResponse

			expectCreatorChange *Num
			expectModuleChange  *Num
		}
	)
	cliCtx := s.network.Validators[0].ClientCtx
	cancelCase := []cancelSuite{
		{
			description: "init case",
			initSwapOut: []swapArg{
				{
					wAmount: "20shr",
					wFee:    "5shr",
				},
			},
			approve:             0,
			iTxCreator:          netutilts.KeyAccount1,
			iTxFee:              netutilts.SHRFee2,
			oRes:                &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			expectCreatorChange: &Num{D: 21},  // expect refund 25shr - 2shr swap fee, 2shr cancel fee  = 21
			expectModuleChange:  &Num{D: -25}, // expect return 25shr to tx creator
		},
	}

	for _, ts := range cancelCase {
		s.Run(ts.description, func() {
			var swapIDs []string
			for _, sw := range ts.initSwapOut {
				out, err := CmdOut(cliCtx,
					"0x7b9039bd633411b48a5a5c4262b5b1a16546d209",
					"etherium",
					sw.wAmount,
					sw.wFee,
					ts.iTxFee,
					netutilts.SkipConfirmation,
					netutilts.BlockBroadcast,
					netutilts.MakeByAccount(ts.iTxCreator))
				if err != nil {
					s.Fail("fail when init the swap out request", err)
				}
				txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
				if txRes.Code != netutilts.ShareLedgerSuccessCode {
					s.Fail("fail when init the swap out request %s", txRes.String())
				}
				log := netutilts.ParseRawLogGetEvent(s.T(), txRes.RawLog)[0]
				attr := log.Events.GetEventByType(s.T(), "swap_out")
				swapIDs = append(swapIDs, attr.Get(s.T(), "swap_id").Value)
			}
			if ts.approve > 0 {
				appIDs := strings.Join(swapIDs[0:ts.approve], ",")
				out, err := CmdApprove(
					cliCtx,
					"some_random_hash",
					appIDs,
					netutilts.SHRFee2,
					netutilts.SkipConfirmation,
					netutilts.BlockBroadcast,
					netutilts.MakeByAccount(netutilts.KeyTreasurer))
				if err != nil {
					s.Fail("fail when init the swap out request", err)
				}
				txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
				if txRes.Code != netutilts.ShareLedgerSuccessCode {
					s.Fail("fail when init the swap out request %s", txRes.String())
				}

			}

			var balanceCreatorBeforeCancel sdk.DecCoins
			var balanceModuleBeforeCancel sdk.DecCoins

			if ts.expectModuleChange != nil {
				out, err := CmdFund(cliCtx, netutilts.JSONFlag)
				if err != nil {
					s.Failf("query swap fund fail %s out %s", err.Error(), string(out.Bytes()))
				}
				fundRes := swapTypes.QueryFundResponse{}
				err = cliCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)

				for _, c := range fundRes.Available {
					if c.Denom == denom.Shr {
						balanceModuleBeforeCancel = sdk.NewDecCoins(*c)
					}
				}

			}
			if ts.expectCreatorChange != nil {
				out := tests.CmdQueryBalance(s.T(), cliCtx, netutilts.Accounts[ts.iTxCreator])
				balanceCreatorBeforeCancel = sdk.NewDecCoinsFromCoins(out.Balances...)
			}

			cancelIDs := strings.Join(swapIDs[:ts.approve], ",")

			out, err := CmdCancel(cliCtx,
				cancelIDs,
				ts.iTxFee,
				netutilts.MakeByAccount(ts.iTxCreator),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast)
			if err != nil {
				s.Fail("fail when cancel request", err)
			}
			txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
			if txRes.Code != netutilts.ShareLedgerSuccessCode {
				s.Fail("fail when cancel request %s", txRes.String())
			}

			if ts.expectModuleChange != nil {
				var balanceModuleAfterDeposit sdk.DecCoins
				out, err = CmdFund(cliCtx, netutilts.JSONFlag)
				if err != nil {
					s.Fail("query swap fund fail", err.Error())
				}

				fundRes := swapTypes.QueryFundResponse{}
				err = cliCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
				if err != nil {
					s.T().Fatalf("can't unmarshal json %s", err)
				}
				for _, c := range fundRes.Available {
					if c.Denom == denom.Shr {
						balanceModuleAfterDeposit = sdk.NewDecCoins(*c)
					}
				}
				s.Require().Equalf(balanceModuleBeforeCancel.AmountOf(denom.Shr).Add(sdk.NewDec(ts.expectModuleChange.D)), balanceModuleAfterDeposit.AmountOf(denom.Shr), "module balance isn't equal")
			}
			if ts.expectCreatorChange != nil {
				out := tests.CmdQueryBalance(s.T(), cliCtx, netutilts.Accounts[ts.iTxCreator])
				s.Require().Equalf(
					balanceCreatorBeforeCancel.AmountOf(denom.Base).Add(sdk.NewDec(ts.expectCreatorChange.D*denom.ShrExponent)),
					sdk.NewDecFromInt(out.GetBalances().AmountOf(denom.Base)), "creator expect not equal")
			}

		})

	}
}
