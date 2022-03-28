package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/document module sentinel errors
var (
	ErrDocNotExisted      = sdkerrors.Register(ModuleName, 41, "Doc does not exist")
	ErrDocExisted         = sdkerrors.Register(ModuleName, 42, "Doc existed")
	ErrDocInvalidData     = sdkerrors.Register(ModuleName, 43, "Invalid data")
	ErrorNotIssuer        = sdkerrors.Register(ModuleName, 44, "Not doc issuer")
	ErrDocRevoked         = sdkerrors.Register(ModuleName, 45, "Doc already revoked")
	ErrHolderIDNotExisted = sdkerrors.Register(ModuleName, 46, "Holder ID not exist")
)
