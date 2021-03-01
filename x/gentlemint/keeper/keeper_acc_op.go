package keeper

import (
	"fmt"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetAccOp(ctx sdk.Context, acc types.AccState) {
	store := ctx.KVStore(k.storeKey)
	store.Set(k.createAccOpKey(acc.Address), k.cdc.MustMarshalBinaryBare(acc.Status))
}

func (k Keeper) GetAccOp(ctx sdk.Context, accAddr sdk.AccAddress) types.AccState {
	store := ctx.KVStore((k.storeKey))
	bz := store.Get(k.createAccOpKey(accAddr))
	if len(bz) == 0 {
		return types.AccState{}
	}
	var status string
	k.cdc.MustUnmarshalBinaryBare(bz, &status)
	acc := types.NewAccState(accAddr, status)
	return acc
}

func (k Keeper) DeleteAccOp(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(k.createAccOpKey(addr))
}

func (k Keeper) createAccOpKey(addr sdk.AccAddress) []byte {
	key := fmt.Sprintf("%s%s", AccOpKey, addr)
	return []byte(key)
}

func (k Keeper) IterateAccOps(ctx sdk.Context, cb func(loader types.AccState) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(AccOpKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		key := string(iterator.Key())
		var status string
		sAddr, err := sdk.AccAddressFromBech32(key[len(AccOpKey):])

		if err != nil {
			panic(err)
		}
		err = k.cdc.UnmarshalBinaryBare(iterator.Value(), &status)

		acc := types.AccState{Address: sAddr, Status: status}
		if cb(acc) {
			break
		}
	}
}
