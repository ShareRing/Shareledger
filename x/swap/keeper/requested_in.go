package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/swap/types"
)

// SetPastTxEvent set a specific requestedIn in the store from its index
func (k Keeper) SetPastTxEvent(ctx sdk.Context, destAddr sdk.Address, srcAddr string, txEvents []*types.TxEvent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastTxEventsKeyPrefix))

	for _, txEvent := range txEvents {
		addressPair := types.PastTxEvent{
			SrcAddr:  srcAddr,
			DestAddr: destAddr.String(),
		}
		b := k.cdc.MustMarshal(&addressPair)
		store.Set(types.PastTxEventKey(txEvent.TxHash, txEvent.LogIndex), b)
	}

}

// GetPastTxEvent returns a requestedIn from its index
func (k Keeper) GetPastTxEvent(
	ctx sdk.Context,
	txHash string,
	logIndex uint64,

) (val types.PastTxEvent, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastTxEventsKeyPrefix))

	b := store.Get(types.PastTxEventKey(
		txHash,
		logIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemovePastTxEvent removes a requestedIn from the store
func (k Keeper) RemovePastTxEvent(
	ctx sdk.Context,
	txHash string,
	logIndex uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastTxEventsKeyPrefix))
	store.Delete(types.PastTxEventKey(
		txHash,
		logIndex,
	))
}

// // GetPastTxEvents returns all requestedIn
// func (k Keeper) GetPastTxEvents(ctx sdk.Context) (list []*types.PastTxEvent) {
// 	store := ctx.KVStore(k.storeKey)
// 	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PastTxEventsKeyPrefix))

// 	defer iterator.Close()

// 	for ; iterator.Valid(); iterator.Next() {
// 		var val types.PastTxEvent
// 		k.cdc.MustUnmarshal(iterator.Value(), &val)
// 		list = append(list, &val)
// 	}

// 	return
// }

// GetPastTxEventsByTxHash
func (k Keeper) GetPastTxEventsByTxHash(ctx sdk.Context, txHash string) (events []*types.PastTxEvent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastTxEventsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, types.PastTxEventByTxHashKey(txHash))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PastTxEvent
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		events = append(events, &val)
	}

	return
}
