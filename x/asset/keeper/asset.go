package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/asset/types"
)

func (k Keeper) GetAsset(ctx sdk.Context, uuid string) (types.Asset, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UUIDKeyPrefix))

	if !k.IsAssetExist(ctx, uuid) {
		return types.Asset{}, false
	}

	bz := store.Get([]byte(uuid))

	var asset types.Asset

	k.cdc.MustUnmarshalLengthPrefixed(bz, &asset)

	return asset, true
}

func (k Keeper) SetAsset(ctx sdk.Context, uuid string, asset types.Asset) {
	if len(asset.Creator) == 0 || len(asset.UUID) == 0 {
		return
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UUIDKeyPrefix))

	store.Set([]byte(uuid), k.cdc.MustMarshalLengthPrefixed(&asset))
}

func (k Keeper) SetAssetStatus(ctx sdk.Context, uuid string, status bool) {
	asset, found := k.GetAsset(ctx, uuid)
	if !found {
		return
	}

	asset.Status = status
	k.SetAsset(ctx, uuid, asset)
}

func (k Keeper) DeleteAsset(ctx sdk.Context, uuid string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UUIDKeyPrefix))
	store.Delete([]byte(uuid))
}

func (k Keeper) GetAssetsIterator(ctx sdk.Context) sdk.Iterator {
	return sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UUIDKeyPrefix))
}

func (k Keeper) IterateAssets(ctx sdk.Context, cb func(a types.Asset) bool) {
	iterator := k.GetAssetsIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var asset types.Asset
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &asset)
		if cb(asset) {
			break
		}
	}
}

func (k Keeper) IsAssetExist(ctx sdk.Context, uuid string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UUIDKeyPrefix))
	return store.Has([]byte(uuid))
}
