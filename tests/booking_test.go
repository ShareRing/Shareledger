package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBookingCrateBooking(t *testing.T)  {
	t.Parallel()
	f := InitFixturesKeySeedModuleInitCoinSHRP(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)

	//Create the booking asset
	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee:= "3"
	//Done the booking asset
	maxSHR := f.QueryTotalSupplyOf("shr")
	maxSHRP := f.QueryTotalSupplyOf("shrp")

	t.Logf("max SHR %v",maxSHR)
	t.Logf("max SHRP %v",maxSHRP)

	_,stdOut,_ := f.ExecuteAssetCreate(assetHash,assetID,assetStatus,SHRPFee,keyUser1)

	txnResponse:=ParseStdOut(t,stdOut)

	require.Equal(t, ShareLedgerSuccessCode,txnResponse.Code)

	_,stdOut,_ = f.ExecuteBookingBook(assetID,2,keyUser2)
	tests.WaitForNextHeightTM(f.Port)
	txnResponse=ParseStdOut(t,stdOut)
	require.Equal(t, ShareLedgerSuccessCode,txnResponse.Code)

	logs := ParseRawLogGetEvent(t,txnResponse.RawLog)

	l := logs[0]
	attr := l.Events.GetEventByType(t,"BookingStart")
	bookingID := attr.Get(t,"BookingId")

	t.Logf("the txn reponse afrer booking %s",txnResponse)

	//Validate booking
	bookingAfterCreate := f.ExecuteBookingGetBook(bookingID.Value)
	require.NotNil(t, bookingAfterCreate,"can't get booking data")
	require.Equal(t,assetID,bookingAfterCreate.UUID,"uuid of booking doesn't match asset uuid")
	require.Equal(t,int64(2),bookingAfterCreate.Duration,"booking duration doesn't match")

	//validate booker account
	booker := f.QueryAccount(bookingAfterCreate.Booker)
	require.NotNil(t, booker)
	//user with init 20SHRP after 2 day booking they must spend 6SHRP. 20-6 = 14
	require.Equal(t,14,booker.GetCoins().AmountOf("shrp") )

	//validate booking asset
	assetAfterRent := f.ExecuteAssetGet(assetID,keyUser1)
	require.NotNil(t, bookingAfterCreate,"can't get asset data")
	require.Equal(t, false,assetAfterRent.Status,"asset after rent must be false to prevent double renter")

	f.Cleanup()
}


func TestBookingCrateDoubleBooking(t *testing.T)  {
	t.Parallel()
	f := InitFixturesKeySeedModuleInitCoinSHRP(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)

	//Create the booking asset
	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee:= "3"
	//Done the booking asset
	maxSHR := f.QueryTotalSupplyOf("shr")
	maxSHRP := f.QueryTotalSupplyOf("shrp")

	t.Logf("max SHR %v",maxSHR)
	t.Logf("max SHRP %v",maxSHRP)

	_,stdOut,_ := f.ExecuteAssetCreate(assetHash,assetID,assetStatus,SHRPFee,keyUser1)

	txnResponse:=ParseStdOut(t,stdOut)

	require.Equal(t, ShareLedgerSuccessCode,txnResponse.Code)

	_,stdOut,_ = f.ExecuteBookingBook(assetID,2,keyUser2)
	tests.WaitForNextHeightTM(f.Port)
	txnResponse=ParseStdOut(t,stdOut)
	require.Equal(t, ShareLedgerSuccessCode,txnResponse.Code)

	_,stdOut,_ = f.ExecuteBookingBook(assetID,5,keyUser1)
	tests.WaitForNextHeightTM(f.Port)
	txnResponse=ParseStdOut(t,stdOut)
	require.Equal(t, ShareLedgerBookingAssetAlreadyBooked,txnResponse.Code)

	f.Cleanup()
}

func TestBookingCompleteBooking(t *testing.T)  {
	t.Parallel()
	f := InitFixturesKeySeedModuleInitCoinSHRP(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)

	//Create the booking asset
	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee:= "3"
	//Done the booking asset
	maxSHR := f.QueryTotalSupplyOf("shr")
	maxSHRP := f.QueryTotalSupplyOf("shrp")

	t.Logf("max SHR %v",maxSHR)
	t.Logf("max SHRP %v",maxSHRP)

	_,stdOut,_ := f.ExecuteAssetCreate(assetHash,assetID,assetStatus,SHRPFee,keyUser1)

	txnResponse:=ParseStdOut(t,stdOut)

	require.Equal(t, ShareLedgerSuccessCode,txnResponse.Code)

	_,stdOut,_ = f.ExecuteBookingBook(assetID,2,keyUser2)
	tests.WaitForNextHeightTM(f.Port)
	txnResponse=ParseStdOut(t,stdOut)
	require.Equal(t, ShareLedgerSuccessCode,txnResponse.Code)

	logs := ParseRawLogGetEvent(t,txnResponse.RawLog)

	l := logs[0]
	attr := l.Events.GetEventByType(t,"BookingStart")
	bookingID := attr.Get(t,"BookingId")


	_,stdOut,_ = f.ExecuteBookingComplete(bookingID.Value,keyUser2)
	txnResponse=ParseStdOut(t,stdOut)

	require.Equal(t, ShareLedgerSuccessCode,txnResponse.Code)

	//Validate booking
	bookingAfterCreate := f.ExecuteBookingGetBook(bookingID.Value)
	require.NotNil(t, bookingAfterCreate,"can't get booking data")
	require.Equal(t,assetID,bookingAfterCreate.UUID,"uuid of booking doesn't match asset uuid")
	require.Equal(t,true,bookingAfterCreate.IsCompleted,"booking now must be completed")


	assetAfterCompleteBooking := f.ExecuteAssetGet(assetID,keyUser1)
	require.NotNil(t, bookingAfterCreate,"can't get asset data")
	require.Equal(t, true,assetAfterCompleteBooking.Status,"asset after complete must be able for next rent")

	f.Cleanup()
}

func TestBookingCompleteBookingUserNotOwnerOfBooking(t *testing.T)  {
	t.Parallel()
	f := InitFixturesKeySeedModuleInitCoinSHRP(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)

	//Create the booking asset
	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee:= "3"
	//Done the booking asset
	maxSHR := f.QueryTotalSupplyOf("shr")
	maxSHRP := f.QueryTotalSupplyOf("shrp")

	t.Logf("max SHR %v",maxSHR)
	t.Logf("max SHRP %v",maxSHRP)

	_,stdOut,_ := f.ExecuteAssetCreate(assetHash,assetID,assetStatus,SHRPFee,keyUser1)

	txnResponse:=ParseStdOut(t,stdOut)

	require.Equal(t, ShareLedgerSuccessCode,txnResponse.Code)

	_,stdOut,_ = f.ExecuteBookingBook(assetID,2,keyUser2)
	tests.WaitForNextHeightTM(f.Port)
	txnResponse=ParseStdOut(t,stdOut)
	require.Equal(t, ShareLedgerSuccessCode,txnResponse.Code)

	logs := ParseRawLogGetEvent(t,txnResponse.RawLog)

	l := logs[0]
	attr := l.Events.GetEventByType(t,"BookingStart")
	bookingID := attr.Get(t,"BookingId")


	_,stdOut,_ = f.ExecuteBookingComplete(bookingID.Value,keyUser1)
	txnResponse=ParseStdOut(t,stdOut)

	require.Equal(t, ShareLedgerBookingBookerIsNotOwner,txnResponse.Code)

	//Validate booking
	bookingAfterCreate := f.ExecuteBookingGetBook(bookingID.Value)
	require.NotNil(t, bookingAfterCreate,"can't get booking data")
	require.Equal(t,assetID,bookingAfterCreate.UUID,"uuid of booking doesn't match asset uuid")
	require.Equal(t,false,bookingAfterCreate.IsCompleted,"booking still not completed yet")


	assetAfterCompleteBooking := f.ExecuteAssetGet(assetID,keyUser1)
	require.NotNil(t, bookingAfterCreate,"can't get asset data")
	require.Equal(t, false,assetAfterCompleteBooking.Status,"the asset wasn't released")

	f.Cleanup()
}