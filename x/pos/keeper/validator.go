package keeper

import (
	"container/list"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

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

	validator = posTypes.MustUnmarshalValidator(k.cdc, addr, value)
	return validator, true
}

// return a given amount of all the validators
func (k Keeper) GetValidators(ctx sdk.Context, maxRetrieve uint16) (validators []posTypes.Validator) {
	store := ctx.KVStore(k.storeKey)
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
	//k.DeleteValidatorByPowerIndex(ctx, validator, pool)
	validator, pool, addedShares = validator.AddTokensFromDel(pool, tokensToAdd.RoundInt64())
	// increment the intra-tx counter
	// in case of a conflict, the validator which least recently changed power takes precedence
	counter := k.GetIntraTxCounter(ctx)
	validator.BondIntraTxCounter = counter
	k.SetIntraTxCounter(ctx, counter+1)
	k.SetValidator(ctx, validator)
	k.SetPool(ctx, pool)
	//k.SetValidatorByPowerIndex(ctx, validator, pool)
	return validator, addedShares
}
