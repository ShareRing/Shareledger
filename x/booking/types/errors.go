package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/booking module sentinel errors
var (
	ErrAssetDoesNotExist      = sdkerrors.Register(ModuleName, 41, "asset does not exist")
	ErrIllegalAssetRate       = sdkerrors.Register(ModuleName, 42, "asset is negative")
	ErrAssetAlreadyBooked     = sdkerrors.Register(ModuleName, 43, "asset is already booked")
	ErrUnableToGenerateBookID = sdkerrors.Register(ModuleName, 44, "unable to generate bookID")
	ErrInvalidBooking         = sdkerrors.Register(ModuleName, 45, "booking field values are invalid")
	ErrNotBookerOfAsset       = sdkerrors.Register(ModuleName, 46, "Signer didn't booking this booking")
	ErrAssetNotBooked         = sdkerrors.Register(ModuleName, 47, "Asset is not booked")
	ErrBookingIsCompleted     = sdkerrors.Register(ModuleName, 48, "Booking is completed")
	ErrUUIDMismatch           = sdkerrors.Register(ModuleName, 49, "uuid of booking and asset not matched")
	ErrBookingDoesNotExist    = sdkerrors.Register(ModuleName, 50, "Booking does not exist")
)
