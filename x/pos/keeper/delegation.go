package keeper

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// return a specific delegation
func (k Keeper) GetDelegation(ctx sdk.Context,
	delAddr sdk.Address, valAddr sdk.Address) (
	delegation posTypes.Delegation, found bool) {

	store := ctx.KVStore(k.storeKey)
	key := GetDelegationKey(delAddr, valAddr)
	value := store.Get(key)
	if value == nil {
		return delegation, false
	}

	delegation = posTypes.MustUnmarshalDelegation(k.cdc, key, value)
	return delegation, true
}

// return all delegations  during POS withdrawlReward
func (k Keeper) GetAllDelegations(ctx sdk.Context) (delegations []posTypes.Delegation) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, DelegationKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		delegation := posTypes.MustUnmarshalDelegation(k.cdc, iterator.Key(), iterator.Value())
		delegations = append(delegations, delegation)
	}
	return delegations
}

// set the delegation
func (k Keeper) SetDelegation(ctx sdk.Context, delegation posTypes.Delegation) {
	store := ctx.KVStore(k.storeKey)
	b := posTypes.MustMarshalDelegation(k.cdc, delegation)
	store.Set(GetDelegationKey(delegation.DelegatorAddr, delegation.ValidatorAddr), b)
}

// remove a delegation from store
func (k Keeper) RemoveDelegation(ctx sdk.Context, delegation posTypes.Delegation) {
	//	k.OnDelegationRemoved(ctx, delegation.DelegatorAddr, delegation.ValidatorAddr) //hookig
	store := ctx.KVStore(k.storeKey)
	store.Delete(GetDelegationKey(delegation.DelegatorAddr, delegation.ValidatorAddr))
}

// Perform a delegation, set/update everything necessary within the store.
func (k Keeper) Delegate(ctx sdk.Context, delAddr sdk.Address, bondAmt types.Coin,
	validator posTypes.Validator, subtractAccount bool) (newShares types.Dec, err sdk.Error) {

	// Get or create the delegator delegation
	delegation, found := k.GetDelegation(ctx, delAddr, validator.Owner)
	if !found {
		delegation = posTypes.Delegation{
			DelegatorAddr: delAddr,
			ValidatorAddr: validator.Owner,
			Shares:        types.ZeroDec(),
		}
	}

	if subtractAccount {
		// Account new shares, save
		_, err = k.bankKeeper.SubtractCoins(ctx, delegation.DelegatorAddr, types.Coins{bondAmt})
		if err != nil {
			return
		}
	}

	validator, newShares = k.AddValidatorTokensAndShares(ctx, validator, bondAmt.Amount)

	// Update delegation
	delegation.Shares = delegation.Shares.Add(newShares)
	delegation.Height = ctx.BlockHeight()
	k.SetDelegation(ctx, delegation)
	return newShares, nil
}
