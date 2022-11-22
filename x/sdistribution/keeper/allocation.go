package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// AllocateTokens handles distribution of the collected fees
func (k *Keeper) Allocation(ctx sdk.Context) {
	logger := k.Logger(ctx)
	logger.Debug("Start")

	// fetch and clear the collected fees for distribution, since this is
	// called in BeginBlock, collected fees will be from the previous block

}
