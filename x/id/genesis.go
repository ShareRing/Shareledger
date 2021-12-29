package id

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/id/keeper"
	"github.com/sharering/shareledger/x/id/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for _, id := range genState.IDs {
		k.SetID(ctx, id)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	cb := func(id types.Id) bool {
		genesis.IDs = append(genesis.IDs, &id)
		return false
	}

	k.IterateID(ctx, cb)

	return genesis
}
