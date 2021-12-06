package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/gentlemint module sentinel errors
var (
	ErrSHRSupplyExceeded = sdkerrors.Register(ModuleName, 1, "SHR supply exceeded")
)
var (
	ErrSenderIsNotAccountOperator = "Sender is not account operator"
)
