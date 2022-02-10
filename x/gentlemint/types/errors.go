package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/gentlemint module sentinel errors
var (
	ErrPSHRSupplyExceeded = sdkerrors.Register(ModuleName, 41, "PShr supply exceeded")
)
