package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/constant"
)

// GetFeeByAction return fee based on action
// return min fee if not found
func (k Keeper) GetFeeByAction(ctx sdk.Context, action string) sdk.DecCoin {
	level := string(constant.MinFee)
	m, found := k.GetActionLevelFee(ctx, action)
	if found {
		level = m.Level
	}
	return k.GetFeeByLevel(ctx, level)
}

// GetFeeByLevel get fee by level
// return min fee if not found
func (k Keeper) GetFeeByLevel(ctx sdk.Context, level string) sdk.DecCoin {
	levelCost, found := k.GetLevelFee(ctx, level)
	if !found {
		return constant.DefaultFeeLevel[constant.MinFee]
	}
	dec, err := sdk.ParseDecCoin(levelCost.Fee)
	if err != nil {
		k.Logger(ctx).Error(err.Error(), "level", levelCost.Level, "cost", levelCost.Fee)
		dec = constant.DefaultFeeLevel[constant.MinFee]
	}
	return dec

}
