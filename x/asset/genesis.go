package asset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/asset/keeper"
	"github.com/sharering/shareledger/x/asset/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for _, a := range genState.Assets {
		k.SetAsset(ctx, a.GetUUID(), *a)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport saffolding # genesis/module/export
	cb := func(a types.Asset) bool {
		genesis.Assets = append(genesis.Assets, &a)
		return false
	}
	k.IterateAssets(ctx, cb)

	return genesis
}
