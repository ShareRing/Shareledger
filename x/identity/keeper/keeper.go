package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint"
	"github.com/sharering/shareledger/x/identity/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
	gmKeeper gentlemint.Keeper
}

const (
	IdSignerPrefix = "IdSigner"
	IdPrefix       = "Id"
)

// NewKeeper creates new instances of the gentlemint Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, gmKeeper gentlemint.Keeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		gmKeeper: gmKeeper,
	}
}

func (k Keeper) GetIdSigner(ctx sdk.Context, address string) types.IdSigner {
	if !k.IsIdSignerPresent(ctx, address) {
		return types.NewIdSigner()
	}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(address))
	var signer types.IdSigner
	k.cdc.MustUnmarshalBinaryBare(bz, &signer)
	return signer
}

func (k Keeper) GetId(ctx sdk.Context, address string) string {
	if !k.IsIdPresent(ctx, address) {
		return ""
	}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(address))
	return string(bz)
}

func (k Keeper) SetIdSigner(ctx sdk.Context, address string, signer types.IdSigner) {
	if signer.Status == "" {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(address), k.cdc.MustMarshalBinaryBare(signer))
}

func (k Keeper) SetId(ctx sdk.Context, address string, hash string) {
	if hash == "" {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(address), []byte(hash))
}

func (k Keeper) GetIdSignerStatus(ctx sdk.Context, address string) string {
	signer := k.GetIdSigner(ctx, address)
	return signer.Status
}

func (k Keeper) SetIdSignerStatus(ctx sdk.Context, address string, status string) {
	signer := k.GetIdSigner(ctx, address)
	signer.Status = status
	k.SetIdSigner(ctx, address, signer)
}

func (k Keeper) DeleteIdSigner(ctx sdk.Context, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(address))
}

func (k Keeper) DeleteId(ctx sdk.Context, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(address))
}

func (k Keeper) GetIdSignerIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(IdSignerPrefix))
}

func (k Keeper) IterateIdSigners(ctx sdk.Context, cb func(signerKey string, signer types.IdSigner) bool) {
	iterator := k.GetIdSignerIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var signer types.IdSigner
		signerKey := string(iterator.Key())
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &signer)
		if cb(signerKey, signer) {
			break
		}
	}
}

func (k Keeper) GetIdIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(IdPrefix))
}

func (k Keeper) IterateIds(ctx sdk.Context, cb func(idKey, hash string) bool) {
	iterator := k.GetIdIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		idKey := string(iterator.Key())
		hash := string(iterator.Value())
		if cb(idKey, hash) {
			break
		}
	}
}

func (k Keeper) IsIdSignerPresent(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(address))
}

func (k Keeper) IsIdPresent(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(address))
}

func (k Keeper) LoadCoins(ctx sdk.Context, toAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.gmKeeper.LoadCoins(ctx, toAddr, amt)
}

func (k Keeper) IsAuthority(ctx sdk.Context, addr sdk.AccAddress) bool {
	authAddr := k.gmKeeper.GetAuthorityAccount(ctx)

	return authAddr == addr.String()
}
