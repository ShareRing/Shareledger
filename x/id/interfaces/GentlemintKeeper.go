package interfaces

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IGentlemintKeeper interface {
	LoadCoins(ctx sdk.Context, toAddr sdk.AccAddress, amt sdk.Coins) error
}
