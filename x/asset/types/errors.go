package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/asset module sentinel errors
var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1, "Asset does not exist")
	ErrAssetExist       = sdkerrors.Register(ModuleName, 2, "Asset already exist")
)
