package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	netutilts "github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/asset/types"
	bookingtypes "github.com/sharering/shareledger/x/booking/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"os"

	testutil2 "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/x/asset/client/tests"
)

type BookingIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
	dir     string

	bookingIDOfAccount3 string
	bookingID2          string

	//a map with key is booking name and the value is booking ID
	eBooking map[string]string
}

func NewBookingIntegrationTestSuite(cf network.Config) *BookingIntegrationTestSuite {
	return &BookingIntegrationTestSuite{
		cfg: cf,
	}
}

func (s *BookingIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for booking module")

	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//override the keyring by our keyring information
	s.network.Validators[0].ClientCtx.Keyring = kb

	s.T().Log("setting up asset data..")
	s.setupTestMaterial()
	s.T().Log("setting up integration test suite successfully")
	s.Require().NoError(s.network.WaitForNextBlock())

}
func (s *BookingIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "fail to cleanup")
	s.network.Cleanup()
	s.T().Log("tearing down integration test suite")
}

func (s *BookingIntegrationTestSuite) TestCreateBooking() {
	validatorCtx := s.network.Validators[0].ClientCtx
	type TestCase struct {
		description string
		iAssetID    string
		iDuration   string
		iTxnFee     int
		iTxnCreator string
		oErr        error
		oAsset      *types.Asset
		oRes        *sdk.TxResponse
		oBooking    *bookingtypes.Booking
		oAccBalance sdk.Coins
	}
	//each test case should be use difference account to avoid conflict the balance
	testSuite := []TestCase{
		{
			//This test case wil run with account 1 has 100shrp
			description: "create the booking with total free asset",
			iAssetID:    "free_asset",
			iDuration:   "2",
			iTxnFee:     10,
			iTxnCreator: netutilts.KeyAccount1,
			oErr:        nil,
			oAsset: &types.Asset{
				UUID:   "free_asset",
				Status: false,
			},
			oRes: &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oBooking: &bookingtypes.Booking{
				Booker:      netutilts.Accounts[netutilts.KeyAccount1].String(),
				UUID:        "free_asset",
				Duration:    2,
				IsCompleted: false,
			},
			oAccBalance: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(94*denom.USDExponent))),
		},
		{
			//This test case wil run with account 1 has 100shrp
			description: "create the booking with not free asset",
			iAssetID:    "not_free_asset",
			iDuration:   "2",
			iTxnFee:     10,
			iTxnCreator: netutilts.KeyAccount2,
			oErr:        nil,
			oAsset: &types.Asset{
				UUID:   "not_free_asset",
				Status: false,
			},
			oRes:        &sdk.TxResponse{Code: bookingtypes.ErrAssetAlreadyBooked.ABCICode()},
			oBooking:    nil,
			oAccBalance: sdk.Coins{},
		},
	}

	for _, tc := range testSuite {
		s.Run(tc.description, func() {
			stdOut, err := ExCmdCreateBooking(validatorCtx, tc.iAssetID, tc.iDuration,
				netutilts.SHRFee(tc.iTxnFee),
				netutilts.JSONFlag,
				netutilts.MakeByAccount(tc.iTxnCreator))
			if tc.oErr != nil {
				s.NotNilf(err, "case %s require error this step", tc.description)
			}

			txnResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())

			if tc.oRes != nil {
				s.Equal(tc.oRes.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.description, txnResponse.String()))
			}
			if tc.oBooking != nil {
				logs := netutilts.ParseRawLogGetEvent(s.T(), txnResponse.RawLog)
				l := logs[0]
				attr := l.Events.GetEventByType(s.T(), "BookingStart")
				bookingID := attr.Get(s.T(), "BookingId")

				stdOut, err = ExCmdCGetBooking(validatorCtx, bookingID.Value)
				s.NoErrorf(err, "fail to get the booking err %+v", err)
				booking := BookingJsonUnmarshal(s.T(), stdOut.Bytes())
				s.Equal(tc.oBooking.Booker, booking.Booker, "booker address isn't equal")
				s.Equal(tc.oBooking.Duration, booking.Duration)
				s.Equal(tc.oBooking.IsCompleted, booking.IsCompleted)
			}

			if tc.oAsset != nil {
				stdOut, err = tests.ExCmdGetAsset(validatorCtx, tc.iAssetID)
				s.NoError(err, "fail to get the asset data")
				asset := tests.AssetJsonUnmarshal(s.T(), stdOut.Bytes())
				s.Equal(tc.oAsset.Status, asset.Status)

			}
			if !tc.oAccBalance.IsZero() {
				//validate the owner of booking
				accByte, err := testutil2.QueryBalancesExec(validatorCtx, netutilts.Accounts[tc.iTxnCreator])
				s.NoError(err)
				accBalance := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())
				//default shrp 100
				s.Equal(tc.oAccBalance.AmountOf(denom.BaseUSD), accBalance.Balances.AmountOf(denom.BaseUSD), accBalance.Balances)
			}

		})
	}

}

func (s *BookingIntegrationTestSuite) TestCompleteBooking() {

	validatorCtx := s.network.Validators[0].ClientCtx

	type TestCase struct {
		description string
		iAssetID    string
		iBookingID  string
		iTxnFee     int
		iTxnCreator string
		oErr        error
		oAsset      *types.Asset
		oRes        *sdk.TxResponse
		oBooking    *bookingtypes.Booking
	}

	testSuite := []TestCase{
		{
			description: "complete the booking success",
			iBookingID:  s.eBooking["pre_booking_1"],
			iTxnFee:     6,
			iAssetID:    "asset_pre_booking_1",
			iTxnCreator: netutilts.KeyAccount3,
			oErr:        nil,
			oAsset: &types.Asset{
				UUID:   "asset_pre_booking_1",
				Status: true,
			},
			oRes:     &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oBooking: &bookingtypes.Booking{IsCompleted: true},
		},
		{
			description: "complete the booking without permission",
			iBookingID:  s.eBooking["pre_booking_2"],
			iTxnFee:     6,
			iAssetID:    "asset_pre_booking_2",
			iTxnCreator: netutilts.KeyAccount4,
			oErr:        nil,
			oAsset: &types.Asset{
				UUID:   "asset_pre_booking_2",
				Status: false,
			},
			oRes:     &sdk.TxResponse{Code: bookingtypes.ErrNotBookerOfAsset.ABCICode()},
			oBooking: &bookingtypes.Booking{IsCompleted: false},
		},
	}

	for _, tc := range testSuite {
		stdOut, err := ExCmdCCompleteBooking(validatorCtx, tc.iBookingID,
			netutilts.SHRFee(tc.iTxnFee),
			netutilts.JSONFlag,
			netutilts.MakeByAccount(tc.iTxnCreator))
		if tc.oErr != nil {
			s.NotNilf(err, "case %s require error this step", tc.description)
		}

		if tc.oRes != nil {
			txnResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
			s.Equalf(tc.oRes.Code, txnResponse.Code, "the txn create booking response %s ", stdOut.String())
		}
		if tc.oBooking != nil {
			stdOut, err = ExCmdCGetBooking(validatorCtx, tc.iBookingID)
			s.NoErrorf(err, "fail to get the booking err %+v", err)
			booking := BookingJsonUnmarshal(s.T(), stdOut.Bytes())
			s.Equal(tc.oBooking.IsCompleted, booking.IsCompleted)
		}
		if tc.oAsset != nil {
			stdOut, err = tests.ExCmdGetAsset(validatorCtx, tc.iAssetID)
			s.NoError(err, "fail to get the asset data")
			asset := tests.AssetJsonUnmarshal(s.T(), stdOut.Bytes())
			s.Equal(tc.oAsset.Status, asset.Status)
		}
	}
}

func (s *BookingIntegrationTestSuite) setupTestMaterial() {

	s.eBooking = make(map[string]string)
	type (
		ExistedAsset struct {
			AssetID     string
			AssetHash   string
			AssetStatus string
			SHRPFee     string
			MakeBy      string
		}
		ExistedBooking struct {
			BookingName string
			AssetID     string
			Duration    string
			MakeBy      string
		}
	)

	listAsset := []ExistedAsset{
		{
			AssetID:     "free_asset",
			AssetHash:   "cc6f58bd1ada876f0a4941ad579908eda726d6da",
			SHRPFee:     "3",
			AssetStatus: "true",
			MakeBy:      netutilts.KeyTreasurer,
		},
		{
			AssetID:     "not_free_asset",
			AssetHash:   "e787c3c3-54e3-4ef0-9be9-d10b2d0616ae",
			SHRPFee:     "3",
			AssetStatus: "false",
			MakeBy:      netutilts.KeyTreasurer,
		},
		{
			AssetID:     "asset_pre_booking_1",
			AssetHash:   "07c31f4e-76ef-46d6-9900-74b6e5c864e4",
			SHRPFee:     "3",
			AssetStatus: "true",
			MakeBy:      netutilts.KeyTreasurer,
		},
		{
			AssetID:     "asset_pre_booking_2",
			AssetHash:   "44e116e9-df07-4bac-9bfa-a470a8dd4cdc",
			SHRPFee:     "3",
			AssetStatus: "true",
			MakeBy:      netutilts.KeyTreasurer,
		},
	}

	listBooking := []ExistedBooking{
		{
			BookingName: "pre_booking_1",
			AssetID:     "asset_pre_booking_1",
			Duration:    "2",
			MakeBy:      netutilts.KeyAccount3,
		},

		{
			BookingName: "pre_booking_2",
			AssetID:     "asset_pre_booking_2",
			Duration:    "2",
			MakeBy:      netutilts.KeyAccount3,
		},
	}

	for _, la := range listAsset {
		_, err := tests.ExCmdCreateAsset(s.network.Validators[0].ClientCtx, la.AssetID, la.AssetHash, la.AssetStatus, la.SHRPFee,
			netutilts.SHRFee10,
			netutilts.JSONFlag,
			netutilts.MakeByAccount(la.MakeBy),
		)
		s.NoError(err)
		s.Require().NoError(s.network.WaitForNextBlock())
	}
	for _, lb := range listBooking {
		stdOut, err := ExCmdCreateBooking(s.network.Validators[0].ClientCtx, lb.AssetID, lb.Duration,
			netutilts.SHRFee(10),
			netutilts.JSONFlag,
			netutilts.MakeByAccount(lb.MakeBy))

		s.NoError(err)
		txnResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
		//s.T().Log(stdOut.String())
		logs := netutilts.ParseRawLogGetEvent(s.T(), txnResponse.RawLog)
		l := logs[0]
		attr := l.Events.GetEventByType(s.T(), "BookingStart")
		bookingID := attr.Get(s.T(), "BookingId")

		s.eBooking[lb.BookingName] = bookingID.Value
		s.Require().NoError(s.network.WaitForNextBlock())

	}

}
