package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

type CheckFeeDecorator struct {
	gk GentlemintKeeper
}

func NewCheckFeeDecorator(gk GentlemintKeeper) CheckFeeDecorator {
	return CheckFeeDecorator{
		gk: gk,
	}
}

func (cfd CheckFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}
	msgs := tx.GetMsgs()
	requiredFees := sdk.NewCoins()
	for _, msg := range msgs {
		fee, err := cfd.gk.GetBaseFeeByMsg(ctx, msg)
		if err != nil {
			return ctx, err
		}
		requiredFees = requiredFees.Add(fee)
	}
	baseTXFee := feeTx.GetFee().AmountOf(denom.Base)
	baseRequiredFee := requiredFees.AmountOf(denom.Base)

	if baseRequiredFee.GT(baseTXFee) {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "got %v, required %v", baseTXFee, baseRequiredFee)
	}
	return next(ctx, tx, simulate)
}
