package tests

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	denom "github.com/sharering/shareledger/x/utils/demo"

	testutil2 "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/asset/client/tests"
)

type BookingIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	bookingIDOfAccount3 string
	bookingID2          string
}

func NewBookingIntegrationTestSuite(cf network.Config) *BookingIntegrationTestSuite {
	return &BookingIntegrationTestSuite{
		cfg: cf,
	}
}

func (s *BookingIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for booking module")

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.T().Log("setting up asset data..")

	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee := "3"

	_, err = tests.ExCmdCreateAsset(s.network.Validators[0].ClientCtx, assetID, assetHash, assetStatus, SHRPFee,
		network.SHRFee10(),
		network.JSONFlag(),
		network.MakeByAccount(network.KeyTreasurer),
	)
	s.NoError(err)
	_ = s.network.WaitForNextBlock()

	s.T().Log("setting up integration test suite successfully")
	_ = s.network.WaitForNextBlock()

}
func (s *BookingIntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
}

func (s *BookingIntegrationTestSuite) TestCreateBooking() {
	validatorCtx := s.network.Validators[0].ClientCtx
	s.Run("create_booking_with_total_free_asset_should_be_success", func() {
		stdOut, err := ExCmdCreateBooking(validatorCtx, "1eb07acc-6c2d-4148-889f-61752c49a4b3", "2",
			network.SHRFee10(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount1))

		s.NoErrorf(err, "fail to create the booking err %+v", err)
		txnResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txnResponse.Code, "the txn response %s ", stdOut.String())

		logs := network.ParseRawLogGetEvent(s.T(), txnResponse.RawLog)
		l := logs[0]
		attr := l.Events.GetEventByType(s.T(), "BookingStart")
		bookingID := attr.Get(s.T(), "BookingId")
		s.bookingID2 = bookingID.Value
		//validate booking
		stdOut, err = ExCmdCGetBooking(validatorCtx, bookingID.Value)
		s.NoErrorf(err, "fail to get the booking err %+v", err)
		booking := BookingJsonUnmarshal(s.T(), stdOut.Bytes())
		s.Equal(s.network.Accounts[network.KeyAccount1].String(), booking.Booker, "booker address isn't equal")
		s.Equal(int64(2), booking.Duration)
		s.Equal(false, booking.IsCompleted)
		//validate the asset
		stdOut, err = tests.ExCmdGetAsset(validatorCtx, "1eb07acc-6c2d-4148-889f-61752c49a4b3")
		s.NoError(err, "fail to get the asset data")
		asset := tests.AssetJsonUnmarshal(s.T(), stdOut.Bytes())
		s.Equal(false, asset.Status)

		//validate the owner of booking
		accByte, err := testutil2.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyAccount1])
		s.NoError(err)
		accBalance := network.BalanceJsonUnmarshal(s.T(), accByte.Bytes())
		//default shrp 100
		s.Equal(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(94*denom.USDExponent)), sdk.NewCoin(denom.BaseUSD, accBalance.Balances.AmountOf(denom.BaseUSD)), accBalance.Balances)
		//s.Equal(fmt.Sprintf("%d", 94), accBalance.Balances.AmountOf("shrp").String())
	})

	s.Run("create_booking_with_same_asset_it_should_be_got_error", func() {

		stdOut, err := ExCmdCreateBooking(validatorCtx, "1eb07acc-6c2d-4148-889f-61752c49a4b3", "2",
			network.SHRFee10(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount2),
		)
		s.NoError(err, "fail to make txn err %+v", err)
		txnResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerBookingAssetAlreadyBooked, txnResponse.Code, "the txn response %s ", txnResponse.String())

		//validate the asset
		stdOut, err = tests.ExCmdGetAsset(validatorCtx, "1eb07acc-6c2d-4148-889f-61752c49a4b3")
		s.NoError(err, "fail to get the asset data")
		asset := tests.AssetJsonUnmarshal(s.T(), stdOut.Bytes())
		s.Equal(false, asset.Status)

		//validate the owner of booking
		accByte, err := testutil2.QueryBalancesExec(validatorCtx, s.network.Accounts[network.KeyAccount2])
		s.NoError(err)
		accBalance := network.BalanceJsonUnmarshal(s.T(), accByte.Bytes())
		//default shrp 100

		s.Equal(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(100*denom.USDExponent)), sdk.NewCoin(denom.BaseUSD, accBalance.Balances.AmountOf(denom.BaseUSD)), accBalance.Balances)

		//s.Equal(fmt.Sprintf("%d", 100), accBalance.Balances.AmountOf("shrp").String())
	})

}

func (s *BookingIntegrationTestSuite) TestCompleteBooking() {

	validatorCtx := s.network.Validators[0].ClientCtx

	_, err := tests.ExCmdCreateAsset(validatorCtx, "assetID-1", "assetHash", "true", "2",
		network.SHRFee10(),
		network.JSONFlag(),
		network.MakeByAccount(network.KeyTreasurer),
	)
	s.NoError(err)
	_ = s.network.WaitForNextBlock()
	stdOut2, err := ExCmdCreateBooking(s.network.Validators[0].ClientCtx, "assetID-1", "2",
		network.SHRFee10(),
		network.JSONFlag(),
		network.MakeByAccount(network.KeyAccount3),
	)
	s.NoErrorf(err, "init fail to create the booking err %+v", err)
	txnResponse := network.ParseStdOut(s.T(), stdOut2.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode, txnResponse.Code, "the txn create booking response %s ", stdOut2.String())

	logs := network.ParseRawLogGetEvent(s.T(), txnResponse.RawLog)
	l := logs[0]
	attr := l.Events.GetEventByType(s.T(), "BookingStart")
	bookingID := attr.Get(s.T(), "BookingId")
	bookingID2 := bookingID.Value

	s.Run("complete_booking_but_txn_creator_isn't_owner_booking", func() {

		stdOut, err := ExCmdCCompleteBooking(validatorCtx, bookingID2,
			network.SHRFee6(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount1))

		s.NoError(err, "fail to make txn err %+v", err)
		txnResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerBookingBookerIsNotOwner, txnResponse.Code, "the txn response %s ", txnResponse.String())

		stdOut, err = ExCmdCGetBooking(validatorCtx, bookingID2)
		s.NoErrorf(err, "fail to get the booking err %+v", err)
		booking := BookingJsonUnmarshal(s.T(), stdOut.Bytes())

		s.Equal(false, booking.IsCompleted)
		//validate the asset
		stdOut, err = tests.ExCmdGetAsset(validatorCtx, "assetID-1")
		s.NoError(err, "fail to get the asset data")
		asset := tests.AssetJsonUnmarshal(s.T(), stdOut.Bytes())
		s.Equal(false, asset.Status)
	})

	s.Run("complete_booking_success", func() {

		stdOut, err := ExCmdCCompleteBooking(validatorCtx, bookingID2, network.SHRFee6(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount3),
		)
		s.NoError(err, "fail to make txn err %+v", err)
		txnResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txnResponse.Code, "the txn response %s ", txnResponse.String())

		stdOut, err = ExCmdCGetBooking(validatorCtx, bookingID2)
		s.NoErrorf(err, "fail to get the booking err %+v", err)
		booking := BookingJsonUnmarshal(s.T(), stdOut.Bytes())

		s.Equal(true, booking.IsCompleted)
		//validate the asset
		stdOut, err = tests.ExCmdGetAsset(validatorCtx, "assetID-1")
		s.NoError(err, "fail to get the asset data")
		asset := tests.AssetJsonUnmarshal(s.T(), stdOut.Bytes())
		s.Equal(true, asset.Status)

	})

}
