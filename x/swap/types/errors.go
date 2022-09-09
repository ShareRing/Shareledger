package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/swap module sentinel errors
var (
	ErrDuplicatedSwapIn = sdkerrors.Register(ModuleName, 1100, "the request in is processed in blockchain")
)
