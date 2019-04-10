package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
)

type Keeper struct {
	storeKey sdk.StoreKey // key used to access the store from Context
}

func NewKeeper(key sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: key,
	}
}

//----------------------------------------------

func addressToKey(addr sdk.AccAddress) []byte {
	return append([]byte(constants.PREFIX_IDENTITY), addr.Bytes()...)
}

// Store ...
func (k Keeper) Store(ctx sdk.Context, address sdk.AccAddress, hash string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(addressToKey(address), []byte(hash))
}

// Get ...
func (k Keeper) Get(ctx sdk.Context, address sdk.AccAddress) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	hashBytes := store.Get(addressToKey(address))
	if hashBytes == nil {
		return "", false
	}
	return string(hashBytes), true
}

// Delete ...
func (k Keeper) Delete(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(addressToKey(address))
}
