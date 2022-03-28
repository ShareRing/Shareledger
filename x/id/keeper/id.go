package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/id/types"
)

func (k Keeper) SetID(ctx sdk.Context, id *types.Id) {
	baseStore := ctx.KVStore(k.storeKey)

	// address -> id
	addressStore := prefix.NewStore(baseStore, types.KeyPrefix(types.AddressKeyPrefix))
	a, _ := sdk.AccAddressFromBech32(id.Data.OwnerAddress)
	addressStore.Set(a, []byte(id.Id))

	// id -> {ID}
	basedId := id.ToBaseID()
	idStore := prefix.NewStore(baseStore, types.KeyPrefix(types.IDKeyPrefix))
	idStore.Set([]byte(id.Id), types.MustMarshalBaseID(k.cdc, &basedId))
}

func (k Keeper) GetBaseID(ctx sdk.Context, id []byte) (types.BaseID, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IDKeyPrefix))
	bz := store.Get(id)

	if len(bz) == 0 {
		return types.BaseID{}, false
	}

	bid := types.MustUnmarshalBaseID(k.cdc, bz)
	return bid, true
}

func (k Keeper) GetFullIDByAddress(ctx sdk.Context, ownerAddr sdk.AccAddress) (*types.Id, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKeyPrefix))

	id := store.Get(ownerAddr)
	if len(id) == 0 {
		// TODO
		return nil, false
	}

	bid, found := k.GetBaseID(ctx, id)
	if !found {
		return nil, false
	}

	rs := types.NewIDFromBaseID(string(id), &bid)
	return &rs, true
}

func (k Keeper) GetIdByAddress(ctx sdk.Context, ownerAddr sdk.AccAddress) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKeyPrefix))

	id := store.Get(ownerAddr)

	return id
}

func (k Keeper) GetFullIDByIDString(ctx sdk.Context, id string) (*types.Id, bool) {
	bid, found := k.GetBaseID(ctx, []byte(id))
	if !found {
		return nil, false
	}
	rs := types.NewIDFromBaseID(id, &bid)
	return &rs, true
}

// Check if an ID is existed or not. Then check the owner has id or not
func (k Keeper) IsExist(ctx sdk.Context, id *types.Id) bool {
	baseStore := ctx.KVStore(k.storeKey)

	// Check owner id
	addressStore := prefix.NewStore(baseStore, types.KeyPrefix(types.AddressKeyPrefix))
	a, _ := sdk.AccAddressFromBech32(id.Data.OwnerAddress)
	idBytes := addressStore.Get(a)

	if len(idBytes) != 0 {
		return true
	}

	// Check id
	idStore := prefix.NewStore(baseStore, types.KeyPrefix(types.IDKeyPrefix))
	bz := idStore.Get([]byte(id.Id))

	return len(bz) != 0
}

func (k Keeper) IterateID(ctx sdk.Context, cb func(id types.Id) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.IDKeyPrefix))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		bId := types.MustUnmarshalBaseID(k.cdc, iterator.Value())
		idKey := iterator.Key()[len(types.IDKeyPrefix):]
		id := types.NewIDFromBaseID(string(idKey), &bId)
		if cb(id) {
			break
		}
	}
}
