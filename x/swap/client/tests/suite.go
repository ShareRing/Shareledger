package tests

import (
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	netutilts "github.com/sharering/shareledger/testutil/network"
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
		TestCase struct {
			d         string
			iAmount   string
			txFee     int
			txCreator string
			oErr      error
			oRes      *sdk.TxResponse
		}
	)

	testSuite := []TestCase{
		{
			d:         "deposit success",
			iAmount:   "10shr",
			txCreator: netutilts.KeyAccount2,
			txFee:     10,
			oErr:      nil,
			oRes:      nil,
		},
		{
			d:         "deposit fail",
			iAmount:   "100000000000shr",
			txCreator: netutilts.KeyEmpty1,
			txFee:     10,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: sdkerrors.ErrInsufficientFunds.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
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
		})
	}
}

func (s *SwapIntegrationTestSuite) TestWithDraw() {
	type (
		TestCase struct {
			d         string
			iAmount   string
			iReceiver string
			txFee     int
			txCreator string
			oErr      error
			oRes      *sdk.TxResponse
		}
	)

	testSuite := []TestCase{
		{
			d:         "withdraw success",
			iAmount:   "10shr",
			txCreator: netutilts.KeyTreasurer,
			iReceiver: "shareledger1l5hkf2epa5xmvngnjaasj5dp9jp7ut6s9mrqve",
			txFee:     10,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: sdkerrors.SuccessABCICode},
		},
		{
			d:         "deposit fail InsufficientFunds",
			iAmount:   "100000000000shr",
			iReceiver: "shareledger1l5hkf2epa5xmvngnjaasj5dp9jp7ut6s9mrqve",
			txCreator: netutilts.KeyTreasurer,
			txFee:     10,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: sdkerrors.ErrInsufficientFunds.ABCICode()},
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
		})
	}
}
