package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/asset module sentinel errors
var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 41, "Asset does not exist")
	ErrAssetExist       = sdkerrors.Register(ModuleName, 42, "Asset already exist")
)
