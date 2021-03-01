package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k Keeper) SetDocIssuer(ctx sdk.Context, acc types.AccState) {
	store := ctx.KVStore(k.storeKey)
	store.Set(k.createDocIssuerKey(acc.Address), k.cdc.MustMarshalBinaryBare(acc.Status))
}

func (k Keeper) GetDocIssuer(ctx sdk.Context, accAddr sdk.AccAddress) types.AccState {
	store := ctx.KVStore((k.storeKey))
	bz := store.Get(k.createDocIssuerKey(accAddr))
	if len(bz) == 0 {
		return types.AccState{}
	}
	var status string
	k.cdc.MustUnmarshalBinaryBare(bz, &status)
	acc := types.NewAccState(accAddr, status)
	return acc
}

func (k Keeper) DeleteDocIssuer(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(k.createDocIssuerKey(addr))
}

func (k Keeper) DetactivateDocIssuer(ctx sdk.Context, addr sdk.AccAddress) {
	acc := k.GetDocIssuer(ctx, addr)
	acc.Status = types.Inactive

	k.SetDocIssuer(ctx, acc)
}

func (k Keeper) createDocIssuerKey(addr sdk.AccAddress) []byte {
	key := fmt.Sprintf("%s%s", DocIssuerKey, addr)
	return []byte(key)
}

func (k Keeper) IterateDocIssuers(ctx sdk.Context, cb func(loader types.AccState) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(DocIssuerKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		key := string(iterator.Key())
		var status string
		sAddr, err := sdk.AccAddressFromBech32(key[len(DocIssuerKey):])

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
