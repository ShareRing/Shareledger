package keeper

import (
	"bytes"

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

	//checking if the validator hold valid token number:
	pool := k.GetPool(ctx)
	if !validator.IsAddingTokenValid(pool, bondAmt.Amount) {
		return types.ZeroDec(), posTypes.ErrBadPercentStake(k.Codespace())
	}
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

// get info for begin functions: MinTime and CreationHeight
func (k Keeper) getBeginInfo(ctx sdk.Context, params posTypes.Params, valSrcAddr sdk.Address) (
	minTime int64, height int64, completeNow bool) {

	validator, found := k.GetValidator(ctx, valSrcAddr)
	switch {
	case !found || validator.Status == types.Bonded:

		// the longest wait - just unbonding period from now
		//minTime = ctx.BlockHeader().Time.Add(params.UnbondingTime)
		minTime = ctx.BlockHeader().Time + int64(params.UnbondingTime)
		height = ctx.BlockHeader().Height
		return minTime, height, false

	case validator.IsUnbonded(ctx):
		return minTime, height, true

	case validator.Status == types.Unbonding:
		minTime = validator.UnbondingMinTime
		height = validator.UnbondingHeight
		return minTime, height, false

	default:
		panic("unknown validator status")
	}
}

// return a given amount of all the delegator unbonding-delegations
func (k Keeper) GetUnbondingDelegations(ctx sdk.Context, delegator sdk.Address,
	maxRetrieve uint16) (unbondingDelegations []posTypes.UnbondingDelegation) {

	unbondingDelegations = make([]posTypes.UnbondingDelegation, maxRetrieve)

	store := ctx.KVStore(k.storeKey)
	delegatorPrefixKey := GetUBDsKey(delegator)
	iterator := sdk.KVStorePrefixIterator(store, delegatorPrefixKey)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxRetrieve); iterator.Next() {
		unbondingDelegation := posTypes.MustUnmarshalUBD(k.cdc, iterator.Key(), iterator.Value())
		unbondingDelegations[i] = unbondingDelegation
		i++
	}
	return unbondingDelegations[:i] // trim if the array length < maxRetrieve
}

// return a unbonding delegation
func (k Keeper) GetUnbondingDelegation(ctx sdk.Context,
	delAddr sdk.Address, valAddr sdk.Address) (ubd posTypes.UnbondingDelegation, found bool) {

	store := ctx.KVStore(k.storeKey)
	key := GetUBDKey(delAddr, valAddr)
	value := store.Get(key)
	if value == nil {
		return ubd, false
	}

	ubd = posTypes.MustUnmarshalUBD(k.cdc, key, value)
	return ubd, true
}

// set the unbonding delegation and associated index
func (k Keeper) SetUnbondingDelegation(ctx sdk.Context, ubd posTypes.UnbondingDelegation) {
	store := ctx.KVStore(k.storeKey)
	bz := posTypes.MustMarshalUBD(k.cdc, ubd)
	key := GetUBDKey(ubd.DelegatorAddr, ubd.ValidatorAddr)
	store.Set(key, bz)
	store.Set(GetUBDByValIndexKey(ubd.DelegatorAddr, ubd.ValidatorAddr), []byte{}) // index, store empty bytes
}

// remove the unbonding delegation object and associated index
func (k Keeper) RemoveUnbondingDelegation(ctx sdk.Context, ubd posTypes.UnbondingDelegation) {
	store := ctx.KVStore(k.storeKey)
	key := GetUBDKey(ubd.DelegatorAddr, ubd.ValidatorAddr)
	store.Delete(key)
	store.Delete(GetUBDByValIndexKey(ubd.DelegatorAddr, ubd.ValidatorAddr))
}

// begin unbonding an unbonding record

func (k Keeper) BeginUnbonding(ctx sdk.Context,
	delAddr sdk.Address, valAddr sdk.Address, sharesAmount types.Dec) sdk.Error {

	// TODO quick fix, instead we should use an index, see https://github.com/cosmos/cosmos-sdk/issues/1402
	_, found := k.GetUnbondingDelegation(ctx, delAddr, valAddr)
	if found {
		return posTypes.ErrExistingUnbondingDelegation(k.Codespace())
	}

	returnAmount, err := k.unbond(ctx, delAddr, valAddr, sharesAmount)
	if err != nil {
		return err
	}

	// create the unbonding delegation
	params := k.GetParams(ctx)
	minTime, height, completeNow := k.getBeginInfo(ctx, params, valAddr)
	balance := types.NewCoin(params.BondDenom, returnAmount.RoundInt64())

	// no need to create the ubd object just complete now
	if completeNow {
		_, err := k.bankKeeper.AddCoins(ctx, delAddr, types.Coins{balance})
		if err != nil {
			return err
		}
		return nil
	}

	ubd := posTypes.UnbondingDelegation{
		DelegatorAddr:  delAddr,
		ValidatorAddr:  valAddr,
		CreationHeight: height,
		MinTime:        minTime,
		Balance:        balance,
		InitialBalance: balance,
	}
	k.SetUnbondingDelegation(ctx, ubd)
	return nil
}

// unbond the the delegation return
func (k Keeper) unbond(ctx sdk.Context, delAddr sdk.Address, valAddr sdk.Address,
	shares types.Dec) (amount types.Dec, err sdk.Error) {

	//k.OnDelegationSharesModified(ctx, delAddr, valAddr)

	// check if delegation has any shares in it unbond
	delegation, found := k.GetDelegation(ctx, delAddr, valAddr)
	if !found {
		err = posTypes.ErrNoDelegatorForAddress(k.Codespace())
		return
	}

	// retrieve the amount to remove
	if delegation.Shares.LT(shares) {
		err = posTypes.ErrNotEnoughDelegationShares(k.Codespace(), delegation.Shares.String())
		return
	}

	// get validator
	validator, found := k.GetValidator(ctx, valAddr)
	if !found {
		err = posTypes.ErrNoValidatorFound(k.Codespace())
		return
	}

	// subtract shares from delegator
	delegation.Shares = delegation.Shares.Sub(shares)

	// remove the delegation
	if delegation.Shares.IsZero() {

		// if the delegation is the operator of the validator then
		// trigger a jail validator
		if bytes.Equal(delegation.DelegatorAddr, validator.Owner) && !validator.Revoked {
			//	k.JailValidator(ctx, validator) -> try revoke ???
			validator = k.mustGetValidator(ctx, validator.Owner)
		}

		k.RemoveDelegation(ctx, delegation)
	} else {
		// Update height
		delegation.Height = ctx.BlockHeight()
		k.SetDelegation(ctx, delegation)
	}

	// remove the coins from the validator
	validator, amount = k.RemoveValidatorTokensAndShares(ctx, validator, shares)

	if validator.DelegatorShares.IsZero() && validator.Status != types.Bonded {
		// if bonded, we must remove in EndBlocker instead
		k.RemoveValidator(ctx, validator.Owner)
	}

	//k.OnDelegationSharesModified(ctx, delegation.DelegatorAddr, validator.OperatorAddr)
	return amount, nil
}

// complete unbonding an unbonding record
func (k Keeper) CompleteUnbonding(ctx sdk.Context, delAddr sdk.Address, valAddr sdk.Address) sdk.Error {

	ubd, found := k.GetUnbondingDelegation(ctx, delAddr, valAddr)
	if !found {
		return posTypes.ErrNoUnbondingDelegation(k.Codespace())
	}

	// ensure that enough time has passed
	ctxTime := ctx.BlockHeader().Time
	if ubd.MinTime > ctxTime {
		//return posTypes.ErrNotMature(k.Codespace(), "unbonding", "unit-time", ubd.MinTime, ctxTime)
	}

	_, err := k.bankKeeper.AddCoins(ctx, ubd.DelegatorAddr, types.Coins{ubd.Balance})
	if err != nil {
		return err
	}
	k.RemoveUnbondingDelegation(ctx, ubd)
	return nil
}

func (k Keeper) BeginRedelegation(ctx sdk.Context, delAddr sdk.Address,
	valSrcAddr, valDstAddr sdk.Address, sharesAmount types.Dec) sdk.Error {

	// check if this is a transitive redelegation
	if k.HasReceivingRedelegation(ctx, delAddr, valSrcAddr) {
		return posTypes.ErrTransitiveRedelegation(k.Codespace())
	}

	returnAmount, err := k.unbond(ctx, delAddr, valSrcAddr, sharesAmount)
	if err != nil {
		return err
	}

	params := k.GetParams(ctx)
	returnCoin := types.NewCoin(params.BondDenom, returnAmount.RoundInt64())
	dstValidator, found := k.GetValidator(ctx, valDstAddr)
	if !found {
		return posTypes.ErrBadRedelegationDst(k.Codespace())
	}
	sharesCreated, err := k.Delegate(ctx, delAddr, returnCoin, dstValidator, false)
	if err != nil {
		return err
	}

	// create the unbonding delegation
	minTime, height, completeNow := k.getBeginInfo(ctx, params, valSrcAddr)

	if completeNow { // no need to create the redelegation object
		return nil
	}

	red := posTypes.Redelegation{
		DelegatorAddr:    delAddr,
		ValidatorSrcAddr: valSrcAddr,
		ValidatorDstAddr: valDstAddr,
		CreationHeight:   height,
		MinTime:          minTime,
		SharesDst:        sharesCreated,
		SharesSrc:        sharesAmount,
		Balance:          returnCoin,
		InitialBalance:   returnCoin,
	}
	k.SetRedelegation(ctx, red)
	return nil
}

// complete unbonding an ongoing redelegation
func (k Keeper) CompleteRedelegation(ctx sdk.Context, delAddr sdk.Address,
	valSrcAddr, valDstAddr sdk.Address) sdk.Error {

	red, found := k.GetRedelegation(ctx, delAddr, valSrcAddr, valDstAddr)
	if !found {
		return posTypes.ErrNoRedelegation(k.Codespace())
	}

	// ensure that enough time has passed
	ctxTime := ctx.BlockHeader().Time
	if red.MinTime > ctxTime { //red.MinTime.After(ctxTime) {
		return posTypes.ErrNotMature(k.Codespace(), "redelegation", "unit-time", red.MinTime, ctxTime)
	}

	k.RemoveRedelegation(ctx, red)
	return nil
}

//_____________________________________________________________________________________

// return a given amount of all the delegator redelegations
func (k Keeper) GetRedelegations(ctx sdk.Context, delegator sdk.Address,
	maxRetrieve uint16) (redelegations []posTypes.Redelegation) {
	redelegations = make([]posTypes.Redelegation, maxRetrieve)

	store := ctx.KVStore(k.storeKey)
	delegatorPrefixKey := GetREDsKey(delegator)
	iterator := sdk.KVStorePrefixIterator(store, delegatorPrefixKey)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxRetrieve); iterator.Next() {
		redelegation := posTypes.MustUnmarshalRED(k.cdc, iterator.Key(), iterator.Value())
		redelegations[i] = redelegation
		i++
	}
	return redelegations[:i] // trim if the array length < maxRetrieve
}

// return a redelegation
func (k Keeper) GetRedelegation(ctx sdk.Context,
	delAddr sdk.Address, valSrcAddr, valDstAddr sdk.Address) (red posTypes.Redelegation, found bool) {

	store := ctx.KVStore(k.storeKey)
	key := GetREDKey(delAddr, valSrcAddr, valDstAddr)
	value := store.Get(key)
	if value == nil {
		return red, false
	}

	red = posTypes.MustUnmarshalRED(k.cdc, key, value)
	return red, true
}

// check if validator is receiving a redelegation
func (k Keeper) HasReceivingRedelegation(ctx sdk.Context,
	delAddr sdk.Address, valDstAddr sdk.Address) bool {

	store := ctx.KVStore(k.storeKey)
	prefix := GetREDsByDelToValDstIndexKey(delAddr, valDstAddr)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	found := false
	if iterator.Valid() {
		found = true
	}
	return found
}

// set a redelegation and associated index
func (k Keeper) SetRedelegation(ctx sdk.Context, red posTypes.Redelegation) {
	store := ctx.KVStore(k.storeKey)
	bz := posTypes.MustMarshalRED(k.cdc, red)
	key := GetREDKey(red.DelegatorAddr, red.ValidatorSrcAddr, red.ValidatorDstAddr)
	store.Set(key, bz)
	store.Set(GetREDByValSrcIndexKey(red.DelegatorAddr, red.ValidatorSrcAddr, red.ValidatorDstAddr), []byte{})
	store.Set(GetREDByValDstIndexKey(red.DelegatorAddr, red.ValidatorSrcAddr, red.ValidatorDstAddr), []byte{})
}

// remove a redelegation object and associated index
func (k Keeper) RemoveRedelegation(ctx sdk.Context, red posTypes.Redelegation) {
	store := ctx.KVStore(k.storeKey)
	redKey := GetREDKey(red.DelegatorAddr, red.ValidatorSrcAddr, red.ValidatorDstAddr)
	store.Delete(redKey)
	store.Delete(GetREDByValSrcIndexKey(red.DelegatorAddr, red.ValidatorSrcAddr, red.ValidatorDstAddr))
	store.Delete(GetREDByValDstIndexKey(red.DelegatorAddr, red.ValidatorSrcAddr, red.ValidatorDstAddr))
}
