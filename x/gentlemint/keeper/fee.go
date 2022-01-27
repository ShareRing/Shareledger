package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/constant"
	"github.com/sharering/shareledger/x/fee"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k Keeper) GetShrFeeByMsg(ctx sdk.Context, msg sdk.Msg) sdk.Coin {
	feeD := k.GetFeeByMsg(ctx, msg)
	return k.convertShrCoin(ctx, feeD)
}

func (k Keeper) LoadFeeFromShrp(ctx sdk.Context, msg *types.MsgLoadFee) error {
	shrp := types.ShrpDecToCoins(msg.Shrp.Amount)
	boughtShr := types.CoinsToShr(shrp, k.GetExchangeRateD(ctx))
	if err := k.buyShr(ctx, boughtShr.Amount, msg.GetSigners()[0]); err != nil {
		return sdkerrors.Wrapf(err, "load fee %+v shr with %+v", boughtShr, shrp)
	}
	return nil
}

func (k Keeper) GetShrFeeByActionKey(ctx sdk.Context, action string) sdk.Coin {
	feeD := k.GetFeeByAction(ctx, action)
	return k.convertShrCoin(ctx, feeD)
}

func (k Keeper) convertShrCoin(ctx sdk.Context, amt sdk.DecCoin) sdk.Coin {
	amount := amt.Amount.TruncateInt()
	if amt.Denom != types.DenomSHR {
		return types.CoinsToShr(types.ShrpDecToCoins(amt.Amount), k.GetExchangeRateD(ctx))
	}
	return sdk.NewCoin(types.DenomSHR, amount)
}

// GetFeeByMsg return fee based on message
// return min fee if msg not found
func (k Keeper) GetFeeByMsg(ctx sdk.Context, msg sdk.Msg) sdk.DecCoin {
	return k.GetFeeByAction(ctx, fee.GetActionKey(msg))
}

// GetFeeByAction return fee based on action
// return min fee if not found
func (k Keeper) GetFeeByAction(ctx sdk.Context, action string) sdk.DecCoin {
	level := string(constant.MinFee)
	if fee.IsSpecialActionKey(action) {
		level = string(constant.NoFee)
	}
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
