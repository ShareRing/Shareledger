package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/constant"
	"github.com/sharering/shareledger/x/fee"
)

// GetFeeByMsg return fee based on message
// return min fee if msg not found
func (k Keeper) GetFeeByMsg(ctx sdk.Context, msg sdk.Msg) sdk.DecCoin {
	return k.GetFeeByAction(ctx, fee.GetActionKey(msg))
}

// GetFeeByAction return fee based on action
// return min fee if not found
func (k Keeper) GetFeeByAction(ctx sdk.Context, action string) sdk.DecCoin {
	level := string(constant.MinFee)
	if len(action) == 0 {
		return k.GetFeeByLevel(ctx, level)
	}

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
