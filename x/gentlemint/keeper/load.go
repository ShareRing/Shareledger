package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) LoadCoins(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coins) error {
	return k.loadCoins(ctx, address, coin)
}
