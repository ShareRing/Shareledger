package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SendCoinsFromModuleToAccount(
	ctx sdk.Context,
	senderModule string,
	recipientAddr sdk.AccAddress,
	amt sdk.Coins,
) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}

func (k Keeper) SendCoinsFromAccountToModule(
	ctx sdk.Context,
	senderAddr sdk.AccAddress,
	recipientModule string,
	amt sdk.Coins,
) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}
