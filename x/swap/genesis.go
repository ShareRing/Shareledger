package swap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the request
	k.ImportRequest(ctx, genState.RequestList)

	// Set request count
	k.SetRequestCount(ctx, genState.RequestCount)
	// Set all the batch
	for _, elem := range genState.BatchList {
		k.SetBatch(ctx, elem)
	}

	// Set batch count
	k.SetBatchCount(ctx, genState.BatchCount)
	// Set all the format
	for _, elem := range genState.FormatList {
		k.SetSchema(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.IdList = k.GetAllId(ctx)
	genesis.RequestList = k.GetAllRequest(ctx)
	genesis.RequestCount = k.GetRequestCount(ctx)
	genesis.BatchList = k.GetAllBatch(ctx)
	genesis.BatchCount = k.GetBatchCount(ctx)
	genesis.FormatList = k.GetAllSchema(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
