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
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
