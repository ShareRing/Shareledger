package keeper

import (
	"bytes"
	"container/list"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/bank"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

var validatorDistCache = make(map[string]*list.Element, MaxCacheLength)
var validatorDistCacheList = list.New()

func (k Keeper) GetValidatorDistInfo(
	ctx sdk.Context, addr sdk.Address,
) (
	vdi posTypes.ValidatorDistInfo, found bool,
) {

	// If exist in cache, return cached value
	if vdiElem, ok := validatorDistCache[string(addr)]; ok {
		// Update frequency in validatorCacheList
		validatorCacheList.MoveToBack(vdiElem)
		// Update frequency in validatorCacheList

		return vdiElem.Value.(posTypes.ValidatorDistInfo), true
	}

	// It doesn't exist in cache, retrieve from store
	store := ctx.KVStore(k.storeKey)
	value := store.Get(GetValidatorDistKey(addr))

	// If such validator doesn't exist
	if value == nil {
		return vdi, false
	}

	vdi = posTypes.MustUnmarshalValidatorDist(k.cdc, value)

	vdiElem := validatorDistCacheList.PushBack(vdi)
	validatorDistCache[string(addr)] = vdiElem

	// If exceeding MaxCacheLength, delete the LRU item, meaning in the front of the queue
	if validatorDistCacheList.Len() > MaxCacheLength {
		vdiToRemove := validatorDistCacheList.Remove(validatorDistCacheList.Front()).(*list.Element)
		delete(validatorDistCache, string(vdiToRemove.Value.(posTypes.ValidatorDistInfo).ValidatorAddr))
	}

	return vdi, true
}

func (k Keeper) SetValidatorDistInfo(
	ctx sdk.Context, vdi posTypes.ValidatorDistInfo,
) {
	store := ctx.KVStore(k.storeKey)
	bz := posTypes.MustMarshalValidatorDist(k.cdc, vdi)

	// Update cache if exists in cache
	if vdiElem, ok := validatorDistCache[string(vdi.ValidatorAddr)]; ok {
		validatorDistCacheList.Remove(vdiElem)
		vdiElem = validatorDistCacheList.PushBack(vdi)
		validatorDistCache[string(vdi.ValidatorAddr)] = vdiElem
	} else {
		vdiElem = validatorCacheList.PushBack(vdi)
		validatorDistCache[string(vdi.ValidatorAddr)] = vdiElem
	}

	store.Set(GetValidatorDistKey(vdi.ValidatorAddr), bz)
}

// UpdateBlockReward is called everytime this validator is selected as forger
func (k Keeper) UpdateBlockReward(
	ctx sdk.Context,
	validatorAddr sdk.Address,
	commissionRate types.Dec,
	rewardPerBlock types.Coin,
) (
	posTypes.ValidatorDistInfo, sdk.Error,
) {

	vdi, found := k.GetValidatorDistInfo(ctx, validatorAddr)

	if !found {
		return vdi, sdk.ErrInternal(fmt.Sprintf(constants.POS_VALIDATOR_DIST_NOT_FOUND, validatorAddr))
	}

	// commission for being a validation
	// fmt.Printf("RewardPerBlock: %v, CommissionRate: %v\n", rewardPerBlock, commissionRate)

	commissionCoin := rewardPerBlock.Mul(commissionRate)
	// fmt.Printf("Commission: %s\n", commissionCoin.String())

	// rewardAccum both delegators and validators itself
	// fmt.Printf("Old RewardAccum: %s\n", vdi.RewardAccum)
	vdi.RewardAccum = vdi.RewardAccum.Plus(rewardPerBlock.Minus(commissionCoin))
	// fmt.Printf("New RewardAccum: %s\n", vdi.RewardAccum.String())

	// Update new commission
	vdi.Commission = vdi.Commission.Plus(commissionCoin)

	// Save ValidatorDistInfo
	k.SetValidatorDistInfo(ctx, vdi)

	return vdi, nil
}

// UpdateDelAccum - Update Delegation Accum of a certain delegator is called everytime reward delegation settlement
func (k Keeper) UpdateDelAccum(
	ctx sdk.Context,
	validatorAddr sdk.Address,
	currentHeight int64,
) (
	posTypes.ValidatorDistInfo, sdk.Error,
) {

	vdi, found := k.GetValidatorDistInfo(ctx, validatorAddr)

	if !found {
		return vdi, sdk.ErrInternal(fmt.Sprintf(constants.POS_VALIDATOR_DIST_NOT_FOUND, validatorAddr))
	}

	// Get all delegations
	allDelegations := k.GetAllDelegations(ctx)

	// totalRewardAccum is RewardAccum of this vaidator
	totalRewardAccum := vdi.RewardAccum

	totalShare := types.ZeroDec()

	// Distribute RewardAccum to Delegator of this validator
	for _, delegation := range allDelegations {

		// TODO time consuming, need a way to retrieve all delegators of a validator
		if bytes.Equal(delegation.ValidatorAddr, vdi.ValidatorAddr) {

			// Update current Height, total Reward Accum
			delegation = delegation.UpdateDelAccum(currentHeight, totalRewardAccum)

			totalShare = totalShare.Add(delegation.Shares)

			// Save to store
			k.SetDelegation(ctx, delegation)
		}
	}

	// Distribute RewardAccum to Validator
	validatorShare := types.NewDec(100).Sub(totalShare)

	vdi.ValidatorReward = vdi.ValidatorReward.Plus(totalRewardAccum.Mul(validatorShare))

	// Done with distribution, RewardAccum is set to Zero
	vdi.RewardAccum = types.NewZeroPOSCoin()

	return vdi, nil
}

func (k Keeper) WithdrawDelReward(
	ctx sdk.Context,
	bkeeper bank.Keeper,
	currentHeight int64,
	validatorAddr sdk.Address,
	delegatorAddr sdk.Address,
) (posTypes.ValidatorDistInfo, types.Coin, sdk.Error) {

	vdi, found := k.GetValidatorDistInfo(ctx, validatorAddr)
	if !found {
		return vdi,
			types.NewZeroPOSCoin(),
			sdk.ErrInternal(fmt.Sprintf(constants.POS_VALIDATOR_DIST_NOT_FOUND, validatorAddr))
	}

	//  Update all new reward
	vdi, err := k.UpdateDelAccum(ctx, vdi.ValidatorAddr, currentHeight)
	if err != nil {
		return vdi, types.NewZeroPOSCoin(), err
	}

	// get delegator
	delegation, found := k.GetDelegation(
		ctx,
		delegatorAddr,
		vdi.ValidatorAddr,
	)

	if !found {
		return vdi,
			types.NewZeroPOSCoin(),
			sdk.ErrInternal(fmt.Sprintf(constants.POS_DELEGATION_NOT_FOUND, delegatorAddr))
	}

	rewardCoin := delegation.RewardAccum

	// Reset this delegation RewardAccum to Zero
	delegation.RewardAccum = types.NewZeroPOSCoin()

	// update balance of delegator

	_, err = bkeeper.AddCoin(
		ctx,
		delegatorAddr,
		rewardCoin,
	)

	if err != nil {
		return vdi,
			types.NewZeroPOSCoin(),
			sdk.ErrInternal(fmt.Sprintf(constants.POS_WITHDRAWAL_ERROR, err.Error()))
	}

	return vdi, rewardCoin, nil

}
