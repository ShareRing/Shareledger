package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrSHRSupplyExceeded   = sdkerrors.Register(ModuleName, 1, "SHR supply exceeded")
	ErrInvalidExchangeRate = sdkerrors.New(ModuleName, 2, "Invalid exchange rate")
	ErrDoesNotExist        = sdkerrors.New(ModuleName, 3, "Does not exist")

	ErrSenderIsNotAuthority       = "Sender is not authority"
	ErrSenderIsNotAccountOperator = "Sender is not account operator"

	ErrSenderIsNotIssuer   = "Sender is not document issuer"
	ErrSenderIsNotIdSigner = "Sender is not id signer"
)
