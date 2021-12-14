package electoral

import (
	"github.com/ShareRing/Shareledger/x/electoral/keeper"
	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the accState
	for _, elem := range genState.AccStateList {
		k.SetAccState(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.AccStateList = k.GetAllAccState(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
