package keeper

import (
	"fmt"

	"github.com/ShareRing/Shareledger/x/id/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		gmKeeper types.GentlemintKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	gmKeeper types.GentlemintKeeper,

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		gmKeeper: gmKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetIdByAddress(ctx sdk.Context, ownerAddr sdk.AccAddress) *types.ID {
	store := ctx.KVStore(k.storeKey)

	id := store.Get(types.IdAddressStateStoreKey(ownerAddr))

	if len(id) == 0 {
		// TODO
		return nil
	}

	ids := k.GetBaseID(ctx, append(types.IdStatePrefix, id...))
	rs := types.NewIDFromBaseID(string(id), &ids)
	return &rs
}

func (k Keeper) GetIdOnlyByAddress(ctx sdk.Context, ownerAddr sdk.AccAddress) []byte {
	store := ctx.KVStore(k.storeKey)

	id := store.Get(types.IdAddressStateStoreKey(ownerAddr))

	return id
}

func (k Keeper) GetIDByIdString(ctx sdk.Context, id string) *types.ID {
	ids := k.GetBaseID(ctx, types.IdStateStoreKey(id))
	rs := types.NewIDFromBaseID(id, &ids)
	return &rs
}

func (k Keeper) GetBaseID(ctx sdk.Context, id []byte) types.BaseID {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(id)

	if len(bz) == 0 {
		return types.BaseID{}
	}

	bId := types.MustUnmarshalBaseID(k.cdc, bz)
	return bId
}

func (k Keeper) SetID(ctx sdk.Context, id *types.ID) {
	store := ctx.KVStore(k.storeKey)

	// address -> id
	store.Set(types.IdAddressStateStoreKey(sdk.AccAddress(id.Data.OwnerAddress)), []byte(id.Id))

	// id -> {ID}
	basedId := id.ToBaseID()
	store.Set(types.IdStateStoreKey(id.Id), types.MustMarshalBaseID(k.cdc, &basedId))
}

// Check if an ID is existed or not. Then check the owner has id or not
func (k Keeper) IsExist(ctx sdk.Context, id *types.ID) bool {
	store := ctx.KVStore(k.storeKey)

	// Check owner id
	ids := store.Get(types.IdAddressStateStoreKey(sdk.AccAddress(id.Data.OwnerAddress)))

	if len(ids) != 0 {
		return true
	}

	// Check id
	bz := store.Get(types.IdStateStoreKey(id.Id))

	if len(bz) != 0 {
		return true
	}

	return false
}

func (k Keeper) IterateID(ctx sdk.Context, cb func(id types.ID) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.IdStatePrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		bId := types.MustUnmarshalBaseID(k.cdc, iterator.Value())
		idKey := iterator.Key()[len(types.IdStatePrefix):]
		id := types.NewIDFromBaseID(string(idKey), &bId)
		if cb(id) {
			break
		}
	}
}
