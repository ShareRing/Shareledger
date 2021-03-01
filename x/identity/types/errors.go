package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrIdAlreadyExists = sdkerrors.New(ModuleName, 1, "Id already exists")
	ErrIdNotExists     = sdkerrors.New(ModuleName, 2, "Id does not exists")
)
