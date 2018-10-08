package keeper

// import (
// 	"container/list"

// 	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
// 	posTypes "github.com/sharering/shareledger/x/pos/type"
// )

// // Cache the amino decoding of validators, as it can be the case that repeated slashing calls
// // cause many calls to GetValidator, which were shown to throttle the state machine in our
// // simulation. Note this is quite biased though, as the simulator does more slashes than a
// // live chain should, however we require the slashing to be fast as noone pays gas for it.
// type cachedValidator struct {
// 	val        posTypes.Validator
// 	marshalled string // marshalled amino bytes for the validator object (not operator address)
// }

// const MaxCacheLength = 500

// // validatorCache-key: validator amino bytes
// var validatorCache = make(map[string]cachedValidator, MaxCacheLength)
// var validatorCacheList = list.New() //TODO: refactor to avoid using cachedList ->LRU ?

// // get a single validator
// func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.Address) (validator posTypes.Validator, found bool) {
// 	store := ctx.KVStore(k.storeKey)
// 	value := store.Get(GetValidatorKey(addr))
// 	if value == nil {
// 		return validator, false
// 	}

// 	// If these amino encoded bytes are in the cache, return the cached validator
// 	strValue := string(value)
// 	if val, ok := validatorCache[strValue]; ok {
// 		valToReturn := val.val
// 		// Doesn't mutate the cache's value
// 		valToReturn.Owner = addr
// 		return valToReturn, true
// 	}

// 	// amino bytes weren't found in cache, so amino unmarshal and add it to the cache
// 	validator, _ = posTypes.UnmarshalValidator(k.cdc, addr, value)
// 	cachedVal := cachedValidator{validator, strValue}
// 	validatorCache[strValue] = cachedValidator{validator, strValue}
// 	validatorCacheList.PushBack(cachedVal)

// 	// if the cache is too big, pop off the last element from it
// 	if validatorCacheList.Len() > MaxCacheLength {
// 		valToRemove := validatorCacheList.Remove(validatorCacheList.Front()).(cachedValidator)
// 		delete(validatorCache, valToRemove.marshalled)
// 	}

// 	validator, _ = posTypes.UnmarshalValidator(k.cdc, addr, value)
// 	return validator, true
// }

// // set the main record holding validator details
// func (k Keeper) SetValidator(ctx sdk.Context, validator posTypes.Validator) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := posTypes.MustMarshalValidator(k.cdc, validator)
// 	store.Set(GetValidatorKey(validator.Owner), bz)
// }

// // Update the tokens of an existing validator, update the validators power index key
// func (k Keeper) AddValidatorTokensAndShares(ctx sdk.Context, validator types.Validator,
// 	tokensToAdd sdk.Int) (valOut types.Validator, addedShares sdk.Dec) {

// 	//pool := k.GetPool(ctx)
// 	//k.DeleteValidatorByPowerIndex(ctx, validator, pool)
// 	validator, pool, addedShares = validator.AddTokensFromDel(pool, tokensToAdd)
// 	// increment the intra-tx counter
// 	// in case of a conflict, the validator which least recently changed power takes precedence
// 	//counter := k.GetIntraTxCounter(ctx)
// 	//validator.BondIntraTxCounter = counter
// 	//k.SetIntraTxCounter(ctx, counter+1)
// 	k.SetValidator(ctx, validator)
// 	//k.SetPool(ctx, pool)
// 	//k.SetValidatorByPowerIndex(ctx, validator, pool)
// 	return validator, addedShares
// }
