package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/booking module sentinel errors
var (
	ErrAssetDoesNotExist      = sdkerrors.New(ModuleName, 1, "asset does not exist")
	ErrIllegalAssetRate       = sdkerrors.New(ModuleName, 2, "asset is negative")
	ErrAssetAlreadyBooked     = sdkerrors.New(ModuleName, 3, "asset is already booked")
	ErrUnableToGenerateBookID = sdkerrors.New(ModuleName, 4, "unable to generate bookID")
	ErrInvalidBooking         = sdkerrors.New(ModuleName, 5, "booking field values are invalid")
	ErrNotBookerOfAsset       = sdkerrors.New(ModuleName, 6, "Signer didn't booking this booking")
	ErrAssetNotBooked         = sdkerrors.New(ModuleName, 7, "Asset is not booked")
	ErrBookingIsCompleted     = sdkerrors.New(ModuleName, 8, "Booking is completed")
	ErrUUIDMismatch           = sdkerrors.New(ModuleName, 9, "uuid of booking and asset not matched")
)
