package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sdistribution module sentinel errors
var (
	ErrInvalidParams = sdkerrors.Register(ModuleName, 2, "invalid params")
)