package booking

import (
	"bytes"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	"github.com/google/go-cmp/cmp"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	utils "github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	msg "github.com/sharering/shareledger/x/booking/messages"
)

type Keeper struct {
	bookingKey sdk.StoreKey // key used to access the store from the context
	assetKey   sdk.StoreKey // asset key
	//accountKey sdk.StoreKey // account key
	accountMapper auth.AccountMapper // account mapper
	cdc           *wire.Codec
}

func NewKeeper(bookingKey sdk.StoreKey, assetKey sdk.StoreKey, am auth.AccountMapper, cdc *wire.Codec) Keeper {
	return Keeper{
		bookingKey:    bookingKey,
		assetKey:      assetKey,
		accountMapper: am,
		//accountKey: accountKey,
		cdc: cdc,
	}
}

//-----------------------------------------------

func (k Keeper) Book(ctx sdk.Context, msg msg.MsgBook) (types.Booking, error) {

	bookingStore := ctx.KVStore(k.bookingKey)
	assetStore := ctx.KVStore(k.assetKey)

	bookingId, err := utils.GenUUID(msg)

	if err != nil {
		return types.Booking{}, fmt.Errorf("bookingID generation failed %s", err.Error())
	}

	// Checking asset

	var asset types.Asset

	err = utils.Retrieve(assetStore, []byte(msg.UUID), &asset)

	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_RETRIEVAL,
			"types.Asset",
			constants.STORE_BOOKING)
	}

	if cmp.Equal(asset, (types.Asset{})) {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_NOT_FOUND,
			msg.UUID,
			constants.STORE_BOOKING)
	}

	if asset.Status == false {
		return types.Booking{}, fmt.Errorf(constants.BOOKING_ASSET_RENTED,
			asset.UUID)
	}

	renterAcc := k.accountMapper.GetAccount(ctx, msg.Renter)
	if renterAcc == nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_RETRIEVAL,
			utils.ByteToString(msg.Renter),
			constants.STORE_BANK)
	}

	// Calculate fee and deduce from Renter account
	value := msg.Duration * asset.Fee

	renterCoins := renterAcc.GetCoins()

	renterCoinsAfter := renterCoins.Minus(types.NewCoin(constants.BOOKING_DENOM, value))

	if !renterCoinsAfter.IsNotNegative() {
		return types.Booking{}, fmt.Errorf(constants.BOOKING_INSUFFICIENT_BALANCE,
			msg.Renter)
	}

	booking := types.NewBooking(bookingId,
		msg.Renter,
		msg.UUID,
		msg.Duration,
		false)

	// Update asset status
	asset.Status = false

	err = utils.Store(bookingStore, []byte(booking.BookingID), booking)
	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_UPDATE,
			"types.Booking",
			constants.STORE_BOOKING)
	}

	err = utils.Store(assetStore, []byte(asset.UUID), asset)
	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_UPDATE,
			"types.Asset",
			constants.STORE_ASSET)
	}

	renterAcc.SetCoins(renterCoinsAfter)
	k.accountMapper.SetAccount(ctx, renterAcc)

	return booking, nil

}

func (k Keeper) Complete(ctx sdk.Context, msg msg.MsgComplete) (types.Booking, error) {

	bookingStore := ctx.KVStore(k.bookingKey)
	assetStore := ctx.KVStore(k.assetKey)
	//accountStore := ctx.KVStore(k.accountKey)

	// Checking booking
	var booking types.Booking

	err := utils.Retrieve(bookingStore, []byte(msg.BookingID), &booking)
	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_RETRIEVAL,
			"types.Booking",
			constants.STORE_BOOKING)
	}

	if !bytes.Equal(msg.Renter, booking.Renter) {
		return types.Booking{}, fmt.Errorf(constants.BOOKING_MISMATCH_RENTER,
			utils.ByteToString(booking.Renter),
			utils.ByteToString(msg.Renter))
	}

	if booking.IsCompleted == true {
		return types.Booking{}, fmt.Errorf(constants.BOOKING_COMPLETED_ERROR,
			booking.BookingID)
	}

	// Check asset
	var asset types.Asset

	err = utils.Retrieve(assetStore, []byte(booking.UUID), &asset)
	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_RETRIEVAL,
			"types.Asset",
			constants.STORE_ASSET)
	}

	if asset.Status != false {
		return types.Booking{}, fmt.Errorf(constants.BOOKING_ASSET_NOT_RENTED,
			asset.UUID)
	}

	// Checking owner account
	ownerAccount := k.accountMapper.GetAccount(ctx, asset.Creator)

	if ownerAccount == nil {
		ownerAccount = auth.NewSHRAccountWithAddress(asset.Creator)
	}

	ownerCoins := ownerAccount.GetCoins()

	// Update owner balance
	value := booking.Duration * asset.Fee

	ownerCoinsAfter := ownerCoins.Plus(types.NewCoin(constants.BOOKING_DENOM, value))
	fmt.Printf("Owner balance: %s\n", ownerCoinsAfter)

	// Update Booking
	booking.IsCompleted = true

	// Update asset status. Asset is available now
	asset.Status = true

	// Save asset detail
	err = utils.Store(assetStore, []byte(asset.UUID), asset)

	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_UPDATE,
			"types.Asset",
			constants.STORE_ASSET)
	}

	// Save booking detail
	err = utils.Store(bookingStore, []byte(asset.UUID), booking)

	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_UPDATE,
			"types.Booking",
			constants.STORE_BOOKING)
	}

	// Save new balance to owner
	ownerAccount.SetCoins(ownerCoinsAfter)
	k.accountMapper.SetAccount(ctx, ownerAccount)

	return booking, nil

}
