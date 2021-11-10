package asset

import (
	"github.com/ShareRing/Shareledger/x/asset/keeper"
	"github.com/ShareRing/Shareledger/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, a := range genState.Assets {
		k.SetAsset(ctx, a.UUID, *a)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var assets []*types.Asset
	cb := func(a types.Asset) bool {
		assets = append(assets, &a)
		return false
	}
	k.IterateAssets(ctx, cb)
	return &types.GenesisState{
		Assets: assets,
	}
}

// type GenesisState struct {
// 	Assets []types.Asset
// }

func NewGenesisState() types.GenesisState {
	return types.GenesisState{}
}

// TODO
func ValidateGenesis(data types.GenesisState) error {
	return nil
}

func DefaultGenesisState() types.GenesisState {
	return types.GenesisState{}
}
