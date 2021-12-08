package keeper

import (
	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) isAccountOperator(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.isActive(ctx, address, types.AccStateKeyAccOp)
}

func (k Keeper) isSHRPLoader(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.isActive(ctx, address, types.AccStateKeyShrpLoaders)
}

func (k Keeper) isActive(ctx sdk.Context, address sdk.AccAddress, keyType types.AccStateKeyType) bool {
	key := types.GenAccStateIndexKey(address, types.AccStateKeyAccOp)
	r, found := k.GetAccState(ctx, key)
	return found && r.Status == string(types.StatusActive)
}

func (k Keeper) activeShrpLoader(ctx sdk.Context, addr sdk.AccAddress) {
	key := types.GenAccStateIndexKey(addr, types.AccStateKeyShrpLoaders)
	k.SetAccState(ctx, types.AccState{
		Key:     key,
		Address: addr.String(),
		Status:  string(types.StatusActive),
	})
}

func (k Keeper) activeIdSigner(ctx sdk.Context, addr sdk.AccAddress) {
	key := types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner)
	k.SetAccState(ctx, types.AccState{
		Key:     key,
		Address: addr.String(),
		Status:  string(types.StatusActive),
	})
}

func (k Keeper) activeDocIssuer(ctx sdk.Context, addr sdk.AccAddress) {
	key := types.GenAccStateIndexKey(addr, types.AccStateKeyDocIssuer)
	k.SetAccState(ctx, types.AccState{
		Key:     key,
		Address: addr.String(),
		Status:  string(types.StatusActive),
	})
}

func (k Keeper) activeAccOperator(ctx sdk.Context, addr sdk.AccAddress) {
	key := types.GenAccStateIndexKey(addr, types.AccStateKeyAccOp)
	k.SetAccState(ctx, types.AccState{
		Key:     key,
		Address: addr.String(),
		Status:  string(types.StatusActive),
	})
}

func (k Keeper) revokeAccOperator(ctx sdk.Context, addr sdk.AccAddress) (err error) {
	return k.revokeAccAccount(ctx, addr, types.AccStateKeyAccOp)
}
func (k Keeper) revokeShrpLoader(ctx sdk.Context, addr sdk.AccAddress) (err error) {
	return k.revokeAccAccount(ctx, addr, types.AccStateKeyShrpLoaders)
}

// revokeDocIssuer set addr doc issuer to inactive
// return err if there is passed addr not found
func (k Keeper) revokeDocIssuer(ctx sdk.Context, addr sdk.AccAddress) (err error) {
	return k.revokeAccAccount(ctx, addr, types.AccStateKeyDocIssuer)
}

// revokeIdSigner set addr signer to inactive
// return err if there is passed addr not found
func (k Keeper) revokeIdSigner(ctx sdk.Context, addr sdk.AccAddress) (err error) {
	return k.revokeAccAccount(ctx, addr, types.AccStateKeyIdSigner)
}

func (k Keeper) revokeAccAccount(ctx sdk.Context, addr sdk.AccAddress, keyType types.AccStateKeyType) error {
	key := types.GenAccStateIndexKey(addr, keyType)
	r, found := k.GetAccState(ctx, key)
	if !found {
		return sdkerrors.ErrNotFound
	}
	r.Status = string(types.StatusInactive)
	k.SetAccState(ctx, r)
	return nil
}

// SetAccState set a specific accState in the store from its index
func (k Keeper) SetAccState(ctx sdk.Context, accState types.AccState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccStateKeyPrefix))
	b := k.cdc.MustMarshal(&accState)
	store.Set(types.AccStateKey(
		accState.Key,
	), b)
}

// GetAccState returns a accState from its index
func (k Keeper) GetAccState(
	ctx sdk.Context,
	key string,

) (val types.AccState, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccStateKeyPrefix))

	b := store.Get(types.AccStateKey(
		key,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccState removes a accState from the store
func (k Keeper) RemoveAccState(
	ctx sdk.Context,
	key string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccStateKeyPrefix))
	store.Delete(types.AccStateKey(
		key,
	))
}

// GetAllAccState returns all accState
func (k Keeper) GetAllAccState(ctx sdk.Context) (list []types.AccState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccStateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	// iterator := sdk.KVStorePrefixIterator(store, []byte(types.AccStateKeyAccOp))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccState
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) IterateAccState(ctx sdk.Context, accTypeIndex types.AccStateKeyType) (list []types.AccState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccStateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte(accTypeIndex))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccState
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
