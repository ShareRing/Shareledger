package keeper

import (
	"bytes"
	"container/list"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	abci "github.com/tendermint/abci/types"

	"github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

//cache validator -> Refactor with a LRU ?
type cachedValidator struct {
	val        posTypes.Validator
	marshalled string // marshalled amino bytes for the validator object (not operator address)
}

const MaxCacheLength = 500

// validatorCache-key: validator amino bytes

var validatorCache = make(map[string]cachedValidator, MaxCacheLength)
var validatorCacheList = list.New()

// get a single validator
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.Address) (validator posTypes.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(GetValidatorKey(addr))
	if value == nil {
		return validator, false
	}

	// If these amino encoded bytes are in the cache, return the cached validator
	strValue := string(value)
	if val, ok := validatorCache[strValue]; ok {
		valToReturn := val.val
		// Doesn't mutate the cache's value
		valToReturn.Owner = addr
		return valToReturn, true
	}

	// amino bytes weren't found in cache, so amino unmarshal and add it to the cache
	validator = posTypes.MustUnmarshalValidator(k.cdc, addr, value)
	cachedVal := cachedValidator{validator, strValue}
	validatorCache[strValue] = cachedValidator{validator, strValue}
	validatorCacheList.PushBack(cachedVal)

	// if the cache is too big, pop off the last element from it
	if validatorCacheList.Len() > MaxCacheLength {
		valToRemove := validatorCacheList.Remove(validatorCacheList.Front()).(cachedValidator)
		delete(validatorCache, valToRemove.marshalled)
	}

	return validator, true
}

// return a given amount of all the validators
func (k Keeper) GetValidators(ctx sdk.Context, maxRetrieve uint16) (validators []posTypes.Validator) {
	store := ctx.KVStore(k.storeKey)

	// maxRetrieve = 10
	validators = make([]posTypes.Validator, maxRetrieve)

	iterator := sdk.KVStorePrefixIterator(store, ValidatorsKey)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxRetrieve); iterator.Next() {
		addr := iterator.Key()[1:]
		validator := posTypes.MustUnmarshalValidator(k.cdc, addr, iterator.Value())
		validators[i] = validator
		i++
	}
	return validators[:i] // trim if the array length < maxRetrieve
}

func (k Keeper) mustGetValidator(ctx sdk.Context, addr sdk.Address) posTypes.Validator {
	validator, found := k.GetValidator(ctx, addr)
	if !found {
		panic(fmt.Sprintf("validator record not found for address: %X\n", addr))
	}
	return validator
}

// set the main record holding validator details
func (k Keeper) SetValidator(ctx sdk.Context, validator posTypes.Validator) {
	store := ctx.KVStore(k.storeKey)
	bz := posTypes.MustMarshalValidator(k.cdc, validator)
	store.Set(GetValidatorKey(validator.Owner), bz)
}

// Update the tokens of an existing validator, update the validators power index key
func (k Keeper) AddValidatorTokensAndShares(ctx sdk.Context, validator posTypes.Validator,
	tokensToAdd types.Dec) (valOut posTypes.Validator, addedShares types.Dec) {

	pool := k.GetPool(ctx)

	k.DeleteValidatorByPowerIndex(ctx, validator, pool)
	validator, pool, addedShares = validator.AddTokensFromDel(pool, tokensToAdd /*.RoundInt64()*/)

	// increment the intra-tx counter
	// in case of a conflict, the validator which least recently changed power takes precedence
	counter := k.GetIntraTxCounter(ctx)
	validator.BondIntraTxCounter = counter

	k.SetIntraTxCounter(ctx, counter+1)
	k.SetValidator(ctx, validator)
	k.SetPool(ctx, pool)
	k.SetValidatorByPowerIndex(ctx, validator, pool)

	return validator, addedShares
}

// remove the validator record and associated indexes
func (k Keeper) RemoveValidator(ctx sdk.Context, address sdk.Address) {

	validator, found := k.GetValidator(ctx, address)
	if !found {
		return
	}

	// delete the old validator record
	store := ctx.KVStore(k.storeKey)
	pool := k.GetPool(ctx)
	store.Delete(GetValidatorKey(address))
	store.Delete(GetValidatorByConsAddrKey(sdk.Address(validator.PubKey.Address())))
	store.Delete(GetValidatorsByPowerIndexKey(validator, pool))

}

// Update the tokens of an existing validator, update the validators power index key
func (k Keeper) RemoveValidatorTokensAndShares(ctx sdk.Context, validator posTypes.Validator,
	sharesToRemove types.Dec) (valOut posTypes.Validator, removedTokens types.Dec) {

	pool := k.GetPool(ctx)
	k.DeleteValidatorByPowerIndex(ctx, validator, pool)
	validator, pool, removedTokens = validator.RemoveDelShares(pool, sharesToRemove)
	k.SetValidator(ctx, validator)
	k.SetPool(ctx, pool)
	k.SetValidatorByPowerIndex(ctx, validator, pool)
	return validator, removedTokens
}

// perform all the store operations for when a validator status becomes bonded
func (k Keeper) bondValidator(ctx sdk.Context, validator posTypes.Validator) posTypes.Validator {

	pool := k.GetPool(ctx)

	k.DeleteValidatorByPowerIndex(ctx, validator, pool)

	validator.BondHeight = ctx.BlockHeight()

	// set the status
	validator, pool = validator.UpdateStatus(pool, types.Bonded)
	k.SetPool(ctx, pool)

	// save the now bonded validator record to the two referenced stores
	k.SetValidator(ctx, validator)
	k.SetValidatorByPowerIndex(ctx, validator, pool)

	// delete from queue if present
	//k.DeleteValidatorQueue(ctx, validator)

	// call the bond hook if present
	/*	if k.hooks != nil {
		k.hooks.OnValidatorBonded(ctx, validator.ConsAddress(), validator.OperatorAddr)
	}*/

	return validator
}

func (k Keeper) beginUnbondingValidator(ctx sdk.Context, validator posTypes.Validator) posTypes.Validator {

	pool := k.GetPool(ctx)
	params := k.GetParams(ctx)

	k.DeleteValidatorByPowerIndex(ctx, validator, pool)

	// sanity check
	if validator.Status != types.Bonded {
		panic(fmt.Sprintf("should not already be unbonded or unbonding, validator: %v\n", validator))
	}

	// set the status
	validator, pool = validator.UpdateStatus(pool, types.Unbonding)
	k.SetPool(ctx, pool)

	validator.UnbondingMinTime = ctx.BlockHeader().Time + int64(params.UnbondingTime)
	validator.UnbondingHeight = ctx.BlockHeader().Height

	// save the now unbonded validator record and power index
	k.SetValidator(ctx, validator)
	k.SetValidatorByPowerIndex(ctx, validator, pool)

	// Adds to unbonding validator queue
	//k.InsertValidatorQueue(ctx, validator)

	// call the unbond hook if present
	//if k.hooks != nil {
	//	k.hooks.OnValidatorBeginUnbonding(ctx, validator.Owner(), validator.OperatorAddr)
	//}

	return validator
}

// Validator state transitions

func (k Keeper) bondedToUnbonding(ctx sdk.Context, validator posTypes.Validator) posTypes.Validator {
	if validator.Status != types.Bonded {
		panic(fmt.Sprintf("bad state transition bondedToUnbonding, validator: %v\n", validator))
	}
	return k.beginUnbondingValidator(ctx, validator)
}

func (k Keeper) unbondingToBonded(ctx sdk.Context, validator posTypes.Validator) posTypes.Validator {
	if validator.Status != types.Unbonding {
		panic(fmt.Sprintf("bad state transition unbondingToBonded, validator: %v\n", validator))
	}
	return k.bondValidator(ctx, validator)
}

func (k Keeper) unbondedToBonded(ctx sdk.Context, validator posTypes.Validator) posTypes.Validator {
	if validator.Status != types.Unbonded {
		panic(fmt.Sprintf("bad state transition unbondedToBonded, validator: %v\n", validator))
	}
	return k.bondValidator(ctx, validator)
}

// switches a validator from unbonding state to unbonded state
func (k Keeper) unbondingToUnbonded(ctx sdk.Context, validator posTypes.Validator) posTypes.Validator {
	if validator.Status != types.Unbonding {
		panic(fmt.Sprintf("bad state transition unbondingToBonded, validator: %v\n", validator))
	}
	return k.completeUnbondingValidator(ctx, validator)
}

func (k Keeper) GetValidatorSetUpdates(ctx sdk.Context) []abci.Validator {
	/*var abciValidators []abci.Validator
	validators := k.GetValidators(ctx, 100)
	for _, val := range validators {
		abciValidators = append(abciValidators, val.ABCIValidator())
	}
	return abciValidators*/
	var updates []abci.Validator
	store := ctx.KVStore(k.storeKey)
	maxValidators := k.GetParams(ctx).MaxValidators
	totalPower := sdk.ZeroInt()

	// Retrieve the last validator set.
	// The persistent set is updated later in this function.
	// (see LastValidatorPowerKey).
	last := k.getLastValidatorsByAddr(ctx)

	// Iterate over validators, highest power to lowest.
	iterator := sdk.KVStoreReversePrefixIterator(store, ValidatorsByPowerIndexKey)
	count := 0
	for ; iterator.Valid() && count < int(maxValidators); iterator.Next() {

		// fetch the validator
		valAddr := sdk.Address(iterator.Value())
		validator := k.mustGetValidator(ctx, valAddr)

		if validator.Revoked {
			panic("should never retrieve a Revoked validator from the power store")
		}

		// if we get to a zero-power validator (which we don't bond),
		// there are no more possible bonded validators
		// note: we must check the ABCI power, since we round before sending to Tendermint
		if validator.Tokens.RoundInt64() == int64(0) {
			break
		}

		// apply the appropriate state change if necessary
		switch validator.Status {
		case types.Unbonded:
			validator = k.unbondedToBonded(ctx, validator)
		case types.Unbonding:
			validator = k.unbondingToBonded(ctx, validator)
		case types.Bonded:
			// no state change
		default:
			panic("unexpected validator status")
		}

		// fetch the old power bytes
		var valAddrBytes [types.ADDRESSLENGTH]byte
		copy(valAddrBytes[:], valAddr[:])
		oldPowerBytes, found := last[valAddrBytes]

		// calculate the new power bytes
		newPower := validator.BondedTokens().RoundInt64()
		newPowerBytes := k.cdc.MustMarshalBinary(sdk.NewInt(newPower))
		// update the validator set if power has changed
		if !found || !bytes.Equal(oldPowerBytes, newPowerBytes) {
			updates = append(updates, validator.ABCIValidator())

			// Assert that the validator had updated its ValidatorDistInfo.FeePoolWithdrawalHeight.
			// This hook is extremely useful, otherwise lazy accum bugs will be difficult to solve.
			/*	if k.hooks != nil {
				k.hooks.OnValidatorPowerDidChange(ctx, validator.ConsAddress(), valAddr)
			}*/

			// set validator power on lookup index.
			k.SetLastValidatorPower(ctx, valAddr, sdk.NewInt(newPower))
		}

		// validator still in the validator set, so delete from the copy
		delete(last, valAddrBytes)

		// keep count
		count++
		totalPower = totalPower.Add(sdk.NewInt(newPower))
	}

	// sort the no-longer-bonded validators
	noLongerBonded := k.sortNoLongerBonded(last)

	// iterate through the sorted no-longer-bonded validators
	for _, valAddrBytes := range noLongerBonded {

		// fetch the validator
		validator := k.mustGetValidator(ctx, sdk.Address(valAddrBytes))

		// bonded to unbonding
		k.bondedToUnbonding(ctx, validator)

		// delete from the bonded validator index
		k.DeleteLastValidatorPower(ctx, sdk.Address(valAddrBytes))

		// update the validator set
		updates = append(updates, validator.ABCIValidatorZero())
	}

	// set total power on lookup index if there are any updates
	if len(updates) > 0 {
		k.SetLastTotalPower(ctx, totalPower)
	}

	return updates

}

// map of operator addresses to serialized power
type validatorsByAddr map[[types.ADDRESSLENGTH]byte][]byte

// get the last validator set
func (k Keeper) getLastValidatorsByAddr(ctx sdk.Context) validatorsByAddr {
	last := make(validatorsByAddr)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, LastValidatorPowerKey)
	for ; iterator.Valid(); iterator.Next() {
		var valAddr [types.ADDRESSLENGTH]byte
		copy(valAddr[:], iterator.Key()[1:])
		powerBytes := iterator.Value()
		last[valAddr] = make([]byte, len(powerBytes))
		copy(last[valAddr][:], powerBytes[:])
	}
	return last
}

// validator index
func (k Keeper) SetValidatorByPowerIndex(ctx sdk.Context, validator posTypes.Validator, pool posTypes.Pool) {
	// jailed validators are not kept in the power index
	if validator.Revoked {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(GetValidatorsByPowerIndexKey(validator, pool), validator.Owner)
}

// validator index
func (k Keeper) DeleteValidatorByPowerIndex(ctx sdk.Context, validator posTypes.Validator, pool posTypes.Pool) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(GetValidatorsByPowerIndexKey(validator, pool))
}

// validator index
func (k Keeper) SetNewValidatorByPowerIndex(ctx sdk.Context, validator posTypes.Validator) {
	store := ctx.KVStore(k.storeKey)
	pool := k.GetPool(ctx)
	store.Set(GetValidatorsByPowerIndexKey(validator, pool), validator.Owner)
}

// perform all the store operations for when a validator status becomes unbonded
func (k Keeper) completeUnbondingValidator(ctx sdk.Context, validator posTypes.Validator) posTypes.Validator {
	pool := k.GetPool(ctx)
	validator, pool = validator.UpdateStatus(pool, types.Unbonded)
	k.SetPool(ctx, pool)
	k.SetValidator(ctx, validator)
	return validator
}
