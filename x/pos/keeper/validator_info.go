package keeper

import (
	"bytes"
	"container/list"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// Cache for each check and deliver call
var validatorDistCacheDeliver = make(map[string]*list.Element, MaxCacheLength)
var validatorDistCacheListDeliver = list.New()
var validatorDistCacheCheck = make(map[string]*list.Element, MaxCacheLength)
var validatorDistCacheListCheck = list.New()

func (k Keeper) GetValidatorDistInfo(
	ctx sdk.Context, addr sdk.Address,
) (
	vdi posTypes.ValidatorDistInfo, found bool,
) {
	var validatorDistCache *(map[string]*list.Element)
	var validatorDistCacheList *(*list.List)

	// CHECK and DELIVER should have different cache
	// Note: Simulate alter DELIVER cache
	if ctx.IsCheckTx() {
		validatorDistCache = &validatorDistCacheCheck
		validatorDistCacheList = &validatorDistCacheListCheck
	} else {
		validatorDistCache = &validatorDistCacheDeliver
		validatorDistCacheList = &validatorDistCacheListDeliver
	}

	// If exist in cache, return cached value
	if vdiElem, ok := (*validatorDistCache)[string(addr)]; ok {
		// Update frequency in validatorCacheList
		(*validatorCacheList).MoveToBack(vdiElem)

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

	vdiElem := (*validatorDistCacheList).PushBack(vdi)
	(*validatorDistCache)[string(addr)] = vdiElem

	// If exceeding MaxCacheLength, delete the LRU item, meaning in the front of the queue
	if (*validatorDistCacheList).Len() > MaxCacheLength {
		vdiToRemove := (*validatorDistCacheList).Remove((*validatorDistCacheList).Front()).(*list.Element)
		delete((*validatorDistCache), string(vdiToRemove.Value.(posTypes.ValidatorDistInfo).ValidatorAddr))
	}

	return vdi, true
}

func (k Keeper) SetValidatorDistInfo(
	ctx sdk.Context, vdi posTypes.ValidatorDistInfo,
) {
	store := ctx.KVStore(k.storeKey)
	bz := posTypes.MustMarshalValidatorDist(k.cdc, vdi)

	// If this is for Deliver, update deliver
	if !ctx.IsCheckTx() {
		// Remove from cache if exists in cache
		vdiElem, ok := validatorDistCacheDeliver[string(vdi.ValidatorAddr)]
		if ok {
			(*validatorDistCacheListDeliver).Remove(vdiElem)
		}
		vdiElem = validatorDistCacheListDeliver.PushBack(vdi)
		validatorDistCacheDeliver[string(vdi.ValidatorAddr)] = vdiElem
	}

	// Check cache is always updated, even in deliverCtx
	// Remove from cache if exists in cache
	vdiElem, ok := validatorDistCacheCheck[string(vdi.ValidatorAddr)]
	if ok {
		(*validatorDistCacheListCheck).Remove(vdiElem)
	}
	vdiElem = validatorDistCacheListCheck.PushBack(vdi)
	validatorDistCacheCheck[string(vdi.ValidatorAddr)] = vdiElem

	store.Set(GetValidatorDistKey(vdi.ValidatorAddr), bz)

	// validatorDistCacheCheck = validatorDistCacheDeliver
	// validatorDistCacheListCheck = validatorDistCacheListDeliver
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
	fmt.Printf("CHECK TX?? %v\n", ctx.IsCheckTx())
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

	vdi, vdiFound := k.GetValidatorDistInfo(ctx, validatorAddr)

	validator, valFound := k.GetValidator(ctx, validatorAddr)

	if !vdiFound || !valFound {
		return vdi, sdk.ErrInternal(fmt.Sprintf(constants.POS_VALIDATOR_DIST_NOT_FOUND, validatorAddr))
	}

	// Get all delegations
	allDelegations := k.GetAllDelegations(ctx)

	// totalRewardAccum is RewardAccum of this vaidator
	totalRewardAccum := vdi.RewardAccum

	totalShare := types.ZeroDec()

	// Distribute RewardAccum to Delegator of this validator
	for _, delegation := range allDelegations {
		fmt.Printf("Delegation: %X\n", delegation.DelegatorAddr)
		// TODO time consuming, need a way to retrieve all delegators of a validator
		if bytes.Equal(delegation.ValidatorAddr, vdi.ValidatorAddr) {
			fmt.Printf("Validator: %X\n", delegation.ValidatorAddr)
			// Update current Height, total Reward Accum
			delegation = delegation.UpdateDelAccum(currentHeight, totalRewardAccum, validator.DelegatorShares)

			totalShare = totalShare.Add(delegation.Shares)
			fmt.Printf("TotalShare: %v\n", totalShare)

			// Save to store
			k.SetDelegation(ctx, delegation)
		}
	}

	// Distribute RewardAccum to Validator
	// totalShare in coins
	validatorShare := types.OneDec().Sub(totalShare.Quo(validator.DelegatorShares))
	fmt.Printf("totalShares(%v)/DelegatorShare(%v)=validatorShare(%v)\n", totalShare, validator.DelegatorShares, validatorShare)

	fmt.Printf("ValidatorReward Before: %v\n", vdi.ValidatorReward)
	vdi.ValidatorReward = vdi.ValidatorReward.Plus(totalRewardAccum.Mul(validatorShare))
	fmt.Printf("ValidatorReward After: %v\n", vdi.ValidatorReward)

	// Done with distribution, RewardAccum is set to Zero
	vdi.RewardAccum = types.NewZeroPOSCoin()

	return vdi, nil
}

func (k Keeper) WithdrawDelReward(
	ctx sdk.Context,
	validatorAddr sdk.Address,
	delegatorAddr sdk.Address,
) (posTypes.ValidatorDistInfo, types.Coin, sdk.Error) {

	currentHeight := ctx.BlockHeight()

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
	fmt.Printf("Withdrawed RewardCoin: %v\n", rewardCoin)

	// Reset this delegation RewardAccum to Zero
	delegation.RewardAccum = types.NewZeroPOSCoin()
	// Save new delegation
	dtxt, _ := delegation.HumanReadableString()
	fmt.Printf("Set delegation: %s\n", dtxt)
	k.SetDelegation(ctx, delegation)

	// if this withdraw is from validator
	if bytes.Equal(validatorAddr[:], delegatorAddr[:]) {
		fmt.Printf("Withdrawal is from Validator\n")
		rewardCoin = rewardCoin.Plus(vdi.RewardAccum)
		vdi.RewardAccum = types.NewZeroPOSCoin()
		txt := vdi.HumanReadableString()
		fmt.Printf("Set validator: %s\n", txt)
		k.SetValidatorDistInfo(ctx, vdi)
	}

	fmt.Printf("Reward Coin: %v\n", rewardCoin)

	coins := k.bankKeeper.GetCoins(ctx, delegatorAddr)
	fmt.Printf("Before update balance: %v\n", coins)

	// update balance of delegator
	after, err := k.bankKeeper.AddCoin(
		ctx,
		delegatorAddr,
		rewardCoin,
	)
	if err != nil {
		return vdi,
			types.NewZeroPOSCoin(),
			sdk.ErrInternal(fmt.Sprintf(constants.POS_WITHDRAWAL_ERROR, err.Error()))
	}

	err = k.bankKeeper.SetCoins(ctx, validatorAddr, after)
	fmt.Printf("After update balance %v\n", after)

	if err != nil {
		return vdi,
			types.NewZeroPOSCoin(),
			sdk.ErrInternal(fmt.Sprintf(constants.POS_WITHDRAWAL_ERROR, err.Error()))
	}

	return vdi, rewardCoin, nil

}
