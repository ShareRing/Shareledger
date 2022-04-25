package tests

import (
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	netutilts "github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/gentlemint/client/tests"
	swapTypes "github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/stretchr/testify/suite"
	"os"
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
				fundRes := swapTypes.QueryFundResponse{}
				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)

				for _, c := range fundRes.Available {
					if c.Denom == denom.Shr {
						balanceModuleBeforeDeposit = sdk.NewDecCoins(*c)
					}
				}

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

				fundRes := swapTypes.QueryFundResponse{}
				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
				if err != nil {
					s.T().Fatalf("can't unmarshal json %s", err)
				}
				for _, c := range fundRes.Available {
					if c.Denom == denom.Shr {
						balanceModuleAfterDeposit = sdk.NewDecCoins(*c)
					}
				}
				s.Require().Equalf(balanceModuleBeforeDeposit.AmountOf(denom.Shr).Add(sdk.NewDec(tc.expectModuleChange.D)), balanceModuleAfterDeposit.AmountOf(denom.Shr), "module balance isn't equal")
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
			oRes:                 &sdk.TxResponse{Code: sdkerrors.ErrInsufficientFunds.ABCICode()},
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
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			var balanceReceiverBeforeDeposit sdk.DecCoins
			var balanceModuleBeforeDeposit sdk.DecCoins

			if tc.expectModuleChange != nil {
				out, err := CmdFund(validatorCtx, netutilts.JSONFlag)
				if err != nil {
					s.Failf("query swap fund fail %s out %s", err.Error(), string(out.Bytes()))
				}
				fundRes := swapTypes.QueryFundResponse{}
				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)

				for _, c := range fundRes.Available {
					if c.Denom == denom.Shr {
						balanceModuleBeforeDeposit = sdk.NewDecCoins(*c)
					}
				}

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
				s.Equalf(tc.oRes.Code, txResponse.Code, "deposit fail %s", out)
			}

			if tc.expectModuleChange != nil {
				var balanceModuleAfterDeposit sdk.DecCoins
				out, err = CmdFund(validatorCtx, netutilts.JSONFlag)
				if err != nil {
					s.Fail("query swap fund fail", err.Error())
				}

				fundRes := swapTypes.QueryFundResponse{}
				err = validatorCtx.Codec.UnmarshalJSON(out.Bytes(), &fundRes)
				if err != nil {
					s.T().Fatalf("can't unmarshal json %s", err)
				}
				for _, c := range fundRes.Available {
					if c.Denom == denom.Shr {
						balanceModuleAfterDeposit = sdk.NewDecCoins(*c)
					}
				}
				s.Require().Equalf(
					balanceModuleBeforeDeposit.AmountOf(denom.Shr).Add(sdk.NewDec(tc.expectModuleChange.D)).String(),
					balanceModuleAfterDeposit.AmountOf(denom.Shr).String(), "module balance isn't equal")
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
