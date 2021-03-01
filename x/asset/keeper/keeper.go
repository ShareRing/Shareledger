package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/asset/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// Gets the entire Whois metadata struct for a name
func (k Keeper) GetAsset(ctx sdk.Context, uuid string) types.Asset {
	store := ctx.KVStore(k.storeKey)

	if !k.IsAssetPresent(ctx, uuid) {
		return types.NewAsset()
	}

	bz := store.Get([]byte(uuid))

	var asset types.Asset

	k.cdc.MustUnmarshalBinaryBare(bz, &asset)

	return asset
}

func (k Keeper) SetAsset(ctx sdk.Context, uuid string, asset types.Asset) {
	if asset.Creator.Empty() || len(asset.UUID) == 0 {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(uuid), k.cdc.MustMarshalBinaryBare(asset))
}

func (k Keeper) SetAssetStatus(ctx sdk.Context, uuid string, status bool) {
	asset := k.GetAsset(ctx, uuid)
	asset.Status = status
	k.SetAsset(ctx, uuid, asset)
}

func (k Keeper) DeleteAsset(ctx sdk.Context, uuid string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(uuid))
}

func (k Keeper) GetAssetsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) IterateAssets(ctx sdk.Context, cb func(a types.Asset) bool) {
	iterator := k.GetAssetsIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var asset types.Asset
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &asset)
		if cb(asset) {
			break
		}
	}
}

func (k Keeper) IsAssetPresent(ctx sdk.Context, uuid string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(uuid))
}
