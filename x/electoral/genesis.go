package electoral

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the accState
	for _, elem := range genState.AccStateList {
		k.SetAccState(ctx, elem)
	}
	// Set if defined
	if genState.Authority != nil {
		k.SetAuthority(ctx, *genState.Authority)
	}
	// Set if defined
	if genState.Treasurer != nil {
		k.SetTreasurer(ctx, *genState.Treasurer)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.AccStateList = k.GetAllAccState(ctx)
	// Get all authority
	authority, found := k.GetAuthority(ctx)
	if found {
		genesis.Authority = &authority
	}
	// Get all treasurer
	treasurer, found := k.GetTreasurer(ctx)
	if found {
		genesis.Treasurer = &treasurer
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
