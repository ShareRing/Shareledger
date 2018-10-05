package keeper

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	types "github.com/sharering/shareledger/types"
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

	delegation = posTypes.Delegation{} //posTypes.MustUnmarshalDelegation(k.cdc, key, value)
	return delegation, true
}

// Perform a delegation, set/update everything necessary within the store.
func (k Keeper) Delegate(ctx sdk.Context, delAddr sdk.Address, bondAmt types.Coin,
	validator posTypes.Validator, subtractAccount bool) (newShares types.Dec, err sdk.Error) {

	// Get or create the delegator delegation
	//Work-Around, refactor later
	delegation, found := posTypes.Delegation{}, false //k.GetDelegation(ctx, delAddr, validator.OperatorAddr)
	if !found {
		delegation = posTypes.Delegation{
			DelegatorAddr: delAddr,
			ValidatorAddr: validator.Owner,
			Shares:        types.ZeroDec(),
		}
	}

	if subtractAccount {
		// Account new shares, save
		_, _, err = k.bankKeeper.SubtractCoins(ctx, delegation.DelegatorAddr, types.Coins{bondAmt})
		if err != nil {
			return
		}
	}

	validator, newShares = k.AddValidatorTokensAndShares(ctx, validator, bondAmt.Amount)

	// Update delegation
	delegation.Shares = delegation.Shares.Add(newShares)
	delegation.Height = ctx.BlockHeight()
	//k.SetDelegation(ctx, delegation)
	return newShares, nil
}
