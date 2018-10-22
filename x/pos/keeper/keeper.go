package keeper

import (
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	bank "github.com/sharering/shareledger/x/bank"
	posTypes "github.com/sharering/shareledger/x/pos/type"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *wire.Codec
	bankKeeper bank.Keeper

	// codespace
	codespace sdk.CodespaceType
}

func NewKeeper(posKey sdk.StoreKey, bk bank.Keeper, cdc *wire.Codec) Keeper {
	keeper := Keeper{
		storeKey:   posKey,
		cdc:        cdc,
		bankKeeper: bk,
	}
	return keeper
}

//_________________________________________________________________________

// return the codespace
func (k Keeper) Codespace() sdk.CodespaceType {
	return k.codespace
}

//_________________________________________________________________________
// some generic reads/writes that don't need their own files

// load/save the global staking params
func (k Keeper) GetParams(ctx sdk.Context) (params posTypes.Params) {
	store := ctx.KVStore(k.storeKey)

	b := store.Get(ParamKey)
	if b == nil {
		panic("Stored params should not have been nil")
	}

	k.cdc.MustUnmarshalBinary(b, &params)
	return
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params posTypes.Params) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinary(params)
	store.Set(ParamKey, b)
}

//__________________________________________________________________________

// get the current in-block validator operation counter
func (k Keeper) InitIntraTxCounter(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(IntraTxCounterKey)
	if b == nil {
		k.SetIntraTxCounter(ctx, 0)
	}
}

// get the current in-block validator operation counter
func (k Keeper) GetIntraTxCounter(ctx sdk.Context) int16 {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(IntraTxCounterKey)
	var counter int16
	k.cdc.MustUnmarshalBinary(b, &counter)
	return counter
}

// set the current in-block validator operation counter
func (k Keeper) SetIntraTxCounter(ctx sdk.Context, counter int16) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinary(counter)
	store.Set(IntraTxCounterKey, bz)
}
