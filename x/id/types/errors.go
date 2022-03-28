package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/id module sentinel errors
var (
	ErrIdNotExisted    = sdkerrors.Register(ModuleName, 41, "Id does not exist")
	ErrIdExisted       = sdkerrors.Register(ModuleName, 42, "Id existed")
	InvalidData        = sdkerrors.Register(ModuleName, 43, "Invalid data")
	ErrWrongBackupAddr = sdkerrors.Register(ModuleName, 44, "Wrong backup address")
	ErrOwnerHasID      = sdkerrors.Register(ModuleName, 45, "This address already own an ID")
	ErrNotOwner        = sdkerrors.Register(ModuleName, 46, "No permission to update this ID")
)
