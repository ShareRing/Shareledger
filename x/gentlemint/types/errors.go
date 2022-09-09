package types

// DONTCOVER

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

// x/gentlemint module sentinel errors
var (
	ErrBaseSupplyExceeded = sdkerrors.Register(ModuleName, 41, fmt.Sprintf("%v supply exceeded", denom.Base))
)
