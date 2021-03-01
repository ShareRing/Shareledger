package asset

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	Assets []types.Asset
}

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, a := range data.Assets {
		keeper.SetAsset(ctx, a.UUID, a)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var assets []types.Asset
	cb := func(a types.Asset) bool {
		assets = append(assets, a)
		return false
	}
	k.IterateAssets(ctx, cb)
	return GenesisState{
		Assets: assets,
	}
}
