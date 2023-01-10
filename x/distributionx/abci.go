package distributionx

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/keeper"
	"github.com/sharering/shareledger/x/distributionx/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker sets the proposer for determining distribution during endblock
// and distribute rewards for the previous block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	k.AllocateTokens(ctx)
	// reset counter & re-calculate builderList
	params := k.GetParams(ctx)
	if req.Header.Height > 0 && req.Header.Height%int64(params.BuilderWindows) == 0 {
		allBuilderCount := k.GetAllBuilderCount(ctx)
		var counter uint64 = 0
		for _, builderCount := range allBuilderCount {
			counter += 1
			if builderCount.Count >= params.TxThreshold {
				k.SetBuilderList(ctx, types.BuilderList{
					Id:              counter,
					ContractAddress: builderCount.Index,
				})
			}
			builderCount.Count = 0
			k.SetBuilderCount(ctx, builderCount)
		}

		k.SetBuilderListCount(ctx, counter)
	}
}
