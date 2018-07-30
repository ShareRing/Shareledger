package booking

import (
	"fmt"
	"bytes"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	"github.com/google/go-cmp/cmp"

	"github.com/sharering/shareledger/types"
	utils "github.com/sharering/shareledger/utils"
	msg "github.com/sharering/shareledger/x/booking/messages"
	"github.com/sharering/shareledger/constants"
)

type Keeper struct {
	bookingKey sdk.StoreKey // key used to access the store from the context
	assetKey   sdk.StoreKey // asset key
	accountKey sdk.StoreKey // account key
	cdc        *wire.Codec
}


func NewKeeper(bookingKey sdk.StoreKey, assetKey sdk.StoreKey, accountKey sdk.StoreKey, cdc *wire.Codec) Keeper {
	return Keeper{
		bookingKey: bookingKey,
		assetKey:   assetKey,
		accountKey: accountKey,
		cdc:        cdc,
	}
}

//-----------------------------------------------

func (k Keeper) Book(ctx sdk.Context, msg msg.MsgBook) (types.Booking, error) {

	bookingStore := ctx.KVStore(k.bookingKey)
	assetStore := ctx.KVStore(k.assetKey)
	accountStore := ctx.KVStore(k.accountKey)

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


	// Checking account
	var renter types.AppAccount

	err = utils.Retrieve(accountStore, msg.Renter, &renter)

	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_RETRIEVAL,
			                               utils.ByteToString(msg.Renter),
			                               constants.STORE_BANK)
	}

	if renter == (types.AppAccount{}) {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_NOT_FOUND,
										   msg.Renter,
										   constants.STORE_ASSET)
	}


	// Calculate fee and deduce from Renter account
	value := msg.Duration * asset.Fee

	if renter.Coins.Amount < value {
		return types.Booking{}, fmt.Errorf(constants.BOOKING_INSUFFICIENT_BALANCE,
										  msg.Renter)
	}

	renter.Coins = renter.Coins.Minus(types.NewCoin(renter.Coins.Denom, value))


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

	err = utils.Store(accountStore, msg.Renter, renter)
	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_UPDATE,
										"types.Asset",
											constants.STORE_BANK)
	}

	return booking, nil

}



func (k Keeper) Complete(ctx sdk.Context, msg msg.MsgComplete) (types.Booking, error) {

	bookingStore := ctx.KVStore(k.bookingKey)
	assetStore := ctx.KVStore(k.assetKey)
	accountStore := ctx.KVStore(k.accountKey)


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

	var owner types.AppAccount

	err = utils.Retrieve(accountStore, asset.Creator, &owner)
	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_RETRIEVAL,
									   "types.AppAccount",
										   constants.STORE_BANK)
	}

	if owner == (types.AppAccount{}) {
		owner = types.NewDefaultAccount()
	}


	// Update owner balance
	value := booking.Duration * asset.Fee

	owner.Coins = owner.Coins.Plus(types.NewCoin("SHR", value))
	fmt.Printf("Owner balance: %s\n", owner.Coins)


	// Update Booking
	booking.IsCompleted = true

	// Update asset status. Asset is available now
	asset.Status = true


	err = utils.Store(assetStore, []byte(asset.UUID), asset)

	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_UPDATE,
								       "types.Asset",
										   constants.STORE_ASSET)
	}

	err = utils.Store(bookingStore, []byte(asset.UUID), booking)

	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_UPDATE,
								       "types.Booking",
										   constants.STORE_BOOKING)
	}

	err = utils.Store(accountStore, asset.Creator, owner)

	if err != nil {
		return types.Booking{}, fmt.Errorf(constants.ERROR_STORE_UPDATE,
								       "types.AppAccount",
										   constants.STORE_BANK)
	}

	return booking, nil

}