package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

// SetRequestedIn set a specific requestedIn in the store from its index
func (k Keeper) SetRequestedIn(ctx sdk.Context, destAddress sdk.Address, srcAddr string, txEventHashes []string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestedInKeyPrefix))

	for _, txHash := range txEventHashes {
		addressPair := types.RequestedIn{
			Slp3Address:  destAddress.String(),
			Erc20Address: srcAddr,
		}
		b := k.cdc.MustMarshal(&addressPair)
		store.Set(types.RequestedInKey(txHash), b)
	}

}

// GetRequestedIn returns a requestedIn from its index
func (k Keeper) GetRequestedIn(
	ctx sdk.Context,
	txHash string,

) (val types.RequestedIn, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestedInKeyPrefix))

	b := store.Get(types.RequestedInKey(
		txHash,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRequestedIn removes a requestedIn from the store
func (k Keeper) RemoveRequestedIn(
	ctx sdk.Context,
	txHash string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestedInKeyPrefix))
	store.Delete(types.RequestedInKey(
		txHash,
	))
}

// GetAllRequestedIn returns all requestedIn
func (k Keeper) GetAllRequestedIn(ctx sdk.Context) (list []types.RequestedIn) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestedInKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RequestedIn
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
