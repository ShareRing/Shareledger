package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

// AllPastEventGenesis  return all past swap in event
func (k Keeper) AllPastEventGenesis(ctx sdk.Context) (val []types.PastTxEventGenesis) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastTxEventsKeyPrefix))

	iterator := store.Iterator(nil, nil)

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		txHash, logIndex, err := types.PastTxEventKeyReverser(key)
		if err != nil {
			panic(fmt.Errorf("reverse key fail %s", err))
		}
		var pastEvent = types.PastTxEvent{}
		value := iterator.Value()
		k.cdc.MustUnmarshal(value, &pastEvent)

		val = append(val, types.PastTxEventGenesis{
			SrcAddr:  pastEvent.SrcAddr,
			DestAddr: pastEvent.DestAddr,
			TxHash:   txHash,
			LogIndex: logIndex,
		})
	}
	return
}

// SetPastEventFromGenesis set the past event from genesis
func (k Keeper) SetPastEventFromGenesis(ctx sdk.Context, pastEventList []types.PastTxEventGenesis) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastTxEventsKeyPrefix))
	for _, pastEvent := range pastEventList {
		addressPair := types.PastTxEvent{
			SrcAddr:  pastEvent.SrcAddr,
			DestAddr: pastEvent.DestAddr,
		}
		b := k.cdc.MustMarshal(&addressPair)
		store.Set(types.PastTxEventKey(pastEvent.TxHash, pastEvent.LogIndex), b)
	}
}
