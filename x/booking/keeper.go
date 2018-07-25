package booking

import (
	"encoding/json"
	"errors"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/types"
	utils "github.com/sharering/shareledger/utils"
	msg "github.com/sharering/shareledger/x/booking/messages"
	"bytes"
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

	assetBytes := assetStore.Get([]byte(msg.UUID))

	if assetBytes == nil {
		return types.Booking{}, errors.New("Asset not found")
	}

	var asset types.Asset

	derr := json.Unmarshal(assetBytes, &asset)

	if derr != nil {
		return types.Booking{}, errors.New("Asset decoding failed.")
	}

	if asset.Status == false {
		return types.Booking{}, errors.New("Asset is busy")
	}


	// Checking account

	accountBytes := accountStore.Get(msg.Renter)

	if accountBytes == nil {
		return types.Booking{}, errors.New("Renter not found")
	}

	var renter types.AppAccount

	rerr := json.Unmarshal(accountBytes, &renter)

	if rerr != nil {
		return types.Booking{}, errors.New("Renter decoding failed")
	}

	value := msg.Duration * asset.Fee

	if renter.Coins.Amount < value {
		return types.Booking{}, errors.New("Insufficient balance")
	}

	renter.Coins = renter.Coins.Minus(types.NewCoin("SHR", value))


	booking := types.NewBooking(bookingId,
		msg.Renter,
		msg.UUID,
		msg.Duration,
		false)

	bookingBytes, err1 := json.Marshal(booking)

	if err1 != nil {
		return types.Booking{}, fmt.Errorf("Booking encoding failed %s", err1.Error())
	}

	// Update asset status
	asset.Status = false

	assetBytes, err = json.Marshal(asset)
	if err != nil {
		return types.Booking{}, errors.New("Asset encoding failed.")
	}


	accountBytes, err = json.Marshal(renter)
	if err != nil {
		return types.Booking{}, errors.New("Renter encoding failed.")
	}


	assetStore.Set([]byte(asset.UUID), assetBytes)
	bookingStore.Set([]byte(booking.BookingID), bookingBytes)
	accountStore.Set(msg.Renter, accountBytes)

	return booking, nil

}



func (k Keeper) Complete(ctx sdk.Context, msg msg.MsgComplete) (types.Booking, error) {

	bookingStore := ctx.KVStore(k.bookingKey)
	assetStore := ctx.KVStore(k.assetKey)
	accountStore := ctx.KVStore(k.accountKey)


	// Checking booking

	bookingBytes := bookingStore.Get([]byte(msg.BookingID))

	if bookingBytes == nil {
		return types.Booking{}, errors.New("Booking not found")
	}

	var booking types.Booking
	err := json.Unmarshal(bookingBytes, &booking)

	if err != nil {
		return types.Booking{}, errors.New("Booking decoding failed")
	}

	if !bytes.Equal(msg.Renter, booking.Renter) {
		return types.Booking{}, errors.New("Mismatch Renter")
	}

	if booking.IsCompleted == true {
		return types.Booking{}, errors.New("This booking is already completed.")
	}




	// Check asset
	assetBytes := assetStore.Get([]byte(booking.UUID))
	if assetBytes == nil {
		return types.Booking{}, errors.New("Asset not found")
	}

	var asset types.Asset

	derr := json.Unmarshal(assetBytes, &asset)

	if derr != nil {
		return types.Booking{}, errors.New("Asset decoding failed.")
	}

	if asset.Status != false {
		return types.Booking{}, errors.New("Asset is not under renting period")
	}

	// Checking owner account

	accountBytes := accountStore.Get(asset.Creator)

	var owner types.AppAccount

	if accountBytes != nil {
		rerr := json.Unmarshal(accountBytes, &owner)


		if rerr != nil {
			return types.Booking{}, errors.New("Owner decoding failed")
		}
	} else {
		owner = types.AppAccount{
			Coins: types.NewCoin("SHR", 0),
		}
	}




	// Update owner balance
	value := booking.Duration * asset.Fee

	owner.Coins = owner.Coins.Plus(types.NewCoin("SHR", value))
	fmt.Printf("Owner balance: %s\n", owner.Coins)


	// Update Booking
	booking.IsCompleted = true
	bookingBytes, err1 := json.Marshal(booking)

	if err1 != nil {
		return types.Booking{}, fmt.Errorf("Booking encoding failed %s", err1.Error())
	}

	// Update asset status. Asset is available now
	asset.Status = true

	// Encoding variables
	assetBytes, err = json.Marshal(asset)
	if err != nil {
		return types.Booking{}, errors.New("Asset encoding failed.")
	}


	accountBytes, err = json.Marshal(owner)
	if err != nil {
		return types.Booking{}, errors.New("Owner encoding failed.")
	}


	assetStore.Set([]byte(asset.UUID), assetBytes)
	bookingStore.Set([]byte(booking.BookingID), bookingBytes)
	accountStore.Set(asset.Creator, accountBytes)

	return booking, nil

}