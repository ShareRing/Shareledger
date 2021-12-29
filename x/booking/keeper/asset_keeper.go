package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	assetTypes "github.com/sharering/shareledger/x/asset/types"
)

func (k Keeper) SetAssetStatus(ctx sdk.Context, uuid string, status bool) {
	k.assetKeeper.SetAssetStatus(ctx, uuid, status)
}

func (k Keeper) GetAsset(ctx sdk.Context, uuid string) (assetTypes.Asset, bool) {
	asset, found := k.assetKeeper.GetAsset(ctx, uuid)
	if !found {
		return assetTypes.Asset{}, false
	}

	return asset, true
}
