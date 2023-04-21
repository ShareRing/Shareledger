package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/distributionx module sentinel errors
var (
	ErrInvalidParams = errorsmod.Register(ModuleName, 2, "invalid params")
)
