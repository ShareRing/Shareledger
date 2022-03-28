package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/document/types"
)

func (k Keeper) SetDoc(ctx sdk.Context, doc *types.Document) {
	baseStore := ctx.KVStore(k.storeKey)
	detailStore := prefix.NewStore(baseStore, types.KeyPrefix(types.DocDetailKeyPrefix))
	detailStore.Set(doc.GetKeyDetailState(), types.MustMarshalDocDetailState(k.cdc, doc.GetDetailState()))

	basicStore := prefix.NewStore(baseStore, types.KeyPrefix(types.DocBasicKeyPrefix))
	basicStore.Set(doc.GetKeyBasicState(), types.MustMarshalDocBasicState(k.cdc, doc.GetBasicState()))
}

func (k Keeper) GetDoc(ctx sdk.Context, queryDoc types.Document) (types.Document, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DocDetailKeyPrefix))
	bz := store.Get(queryDoc.GetKeyDetailState())

	if len(bz) == 0 {
		return types.Document{}, false
	}

	ds := types.MustUnmarshalDocDetailState(k.cdc, bz)

	queryDoc.Version = ds.Version
	queryDoc.Data = ds.Data
	return queryDoc, true
}

func (k Keeper) GetDocByProof(ctx sdk.Context, queryDoc types.Document) (types.Document, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DocBasicKeyPrefix))
	bz := store.Get(queryDoc.GetKeyBasicState())
	if len(bz) == 0 {
		return types.Document{}, false
	}

	bs := types.MustUnmarshalDocBasicState(k.cdc, bz)
	queryDoc.Holder = bs.Holder
	queryDoc.Issuer = bs.Issuer

	result, found := k.GetDoc(ctx, queryDoc)
	if !found {
		return types.Document{}, false
	}

	return result, true
}

func (k Keeper) IterateAllDocsOfAHolder(ctx sdk.Context, holderId string, cb func(doc types.Document) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DocDetailKeyPrefix))

	queryDoc := types.Document{Holder: holderId}
	it := sdk.KVStorePrefixIterator(store, queryDoc.GetKeyDetailOfHolder())

	defer it.Close()
	for ; it.Valid(); it.Next() {

		doc := types.MustMarshalFromDetailRawState(k.cdc, it.Key(), it.Value())
		if cb(*doc) {
			break
		}
	}
}

func (k Keeper) IterateDocs(ctx sdk.Context, cb func(doc types.Document) bool) {
	store := ctx.KVStore(k.storeKey)

	it := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.DocDetailKeyPrefix))

	defer it.Close()
	for ; it.Valid(); it.Next() {

		doc := types.MustMarshalFromDetailRawState(k.cdc, it.Key(), it.Value())
		if cb(*doc) {
			break
		}
	}
}

func (k Keeper) IterateAllDocOfHolderByIssuer(ctx sdk.Context, holder string, issuer string, cb func(doc types.Document) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DocDetailKeyPrefix))

	queryDoc := types.Document{Holder: holder, Issuer: issuer}
	it := sdk.KVStorePrefixIterator(store, queryDoc.GetKeyDetailHolderAndIssuer())

	defer it.Close()
	for ; it.Valid(); it.Next() {
		doc := types.MustMarshalFromDetailRawState(k.cdc, it.Key(), it.Value())
		if cb(*doc) {
			break
		}
	}
}

func (k Keeper) IsIDExist(ctx sdk.Context, id string) bool {
	_, found := k.idKeeper.GetFullIDByIDString(ctx, id)
	return found
}
