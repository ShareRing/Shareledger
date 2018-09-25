package pos

import (
	abci "github.com/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, k Keeper) (ValidatorUpdates []abci.Validator) {
	pool := k.GetPool(ctx)

	// Process Validator Provisions
	// blockTime := ctx.BlockHeader().Time // XXX assuming in seconds, confirm
	// if pool.InflationLastTime+blockTime >= 3600 {
	// 	pool.InflationLastTime = blockTime
	// 	pool = k.processProvisions(ctx)
	// }

	// save the params
	k.setPool(ctx, pool)

	// reset the intra-transaction counter
	k.setIntraTxCounter(ctx, 0)

	// calculate validator set changes
	ValidatorUpdates = k.getTendermintUpdates(ctx)
	k.clearTendermintUpdates(ctx)
	return
}
