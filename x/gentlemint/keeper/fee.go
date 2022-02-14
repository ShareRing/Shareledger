package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/constant"
	"github.com/sharering/shareledger/x/fee"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

func (k Keeper) GetPShrFeeByMsg(ctx sdk.Context, msg sdk.Msg) (sdk.Coin, error) {
	feeD := k.GetFeeByMsg(ctx, msg)
	usdRate := k.GetExchangeRateD(ctx)
	return denom.NormalizeCoins(sdk.NewDecCoins(feeD), &usdRate)
}

func (k Keeper) LoadFeeFundFromShrp(ctx sdk.Context, msg *types.MsgLoadFee) error {
	if msg.Shrp == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "shrp is nil")
	}
	rate := k.GetExchangeRateD(ctx)
	boughtPShr, err := denom.NormalizeCoins(sdk.NewDecCoins(*msg.Shrp), &rate)
	if err != nil {
		return err
	}
	if err := k.buyPShr(ctx, boughtPShr.Amount, msg.GetSigners()[0]); err != nil {
		return sdkerrors.Wrapf(err, "load fee %+v pshr with %+v", boughtPShr, msg.Shrp.Amount)
	}
	return nil
}

func (k Keeper) GetPShrFeeByActionKey(ctx sdk.Context, action string) (sdk.Coin, error) {
	feeD := k.GetFeeByAction(ctx, action)
	usdRate := k.GetExchangeRateD(ctx)
	return denom.NormalizeCoins(sdk.NewDecCoins(feeD), &usdRate)
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
