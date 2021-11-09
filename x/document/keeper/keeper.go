package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/ShareRing/Shareledger/x/document/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetDoc(ctx sdk.Context, doc *types.Document) {
	store := ctx.KVStore(k.storeKey)

	// Doc detail
	store.Set(doc.GetKeyDetailState(), types.MustMarshalDocDetailState(k.cdc, doc.GetDetailState()))

	// Doc basic for easy query
	store.Set(doc.GetKeyBasicState(), types.MustMarshalDocBasicState(k.cdc, doc.GetBasicState()))
}

func (k Keeper) GetDocByProof(ctx sdk.Context, queryDoc types.Document) types.Document {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(queryDoc.GetKeyBasicState())
	if len(bz) == 0 {
		return types.Document{}
	}

	dbs := types.MustUnmarshalDocBasicState(k.cdc, bz)
	queryDoc.Holder = dbs.Holder
	queryDoc.Issuer = dbs.Issuer

	return k.GetDoc(ctx, queryDoc)
}

func (k Keeper) GetDoc(ctx sdk.Context, queryDoc types.Document) types.Document {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(queryDoc.GetKeyDetailState())
	if len(bz) == 0 {
		return types.Document{}
	}

	ds := types.MustUnmarshalDocDetailState(k.cdc, bz)

	queryDoc.Version = ds.Version
	queryDoc.Data = ds.Data
	return queryDoc
}

func (k Keeper) IterateAllDocsOfAHolder(ctx sdk.Context, holderId string, cb func(doc types.Document) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	queryDoc := types.Document{Holder: holderId}
	it := sdk.KVStorePrefixIterator(store, queryDoc.GetKeyDetailOfHolder())

	defer it.Close()
	for ; it.Valid(); it.Next() {

		doc := types.MustMarshalFromDetailRawState(k.cdc, it.Key(), it.Value())
		if cb(doc) {
			break
		}
	}
}
func (k Keeper) IterateDocs(ctx sdk.Context, cb func(doc types.Document) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	it := sdk.KVStorePrefixIterator(store, types.DocDetailPrefix)

	defer it.Close()
	for ; it.Valid(); it.Next() {

		doc := types.MustMarshalFromDetailRawState(k.cdc, it.Key(), it.Value())
		if cb(doc) {
			break
		}
	}
}
