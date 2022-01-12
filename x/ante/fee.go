package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gentleminttypes "github.com/sharering/shareledger/x/gentlemint/types"
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
		fee := cfd.gk.GetShrFeeByMsg(ctx, msg)
		requiredFees = requiredFees.Add(fee)
	}
	shrTXFee := feeTx.GetFee().AmountOf(gentleminttypes.DenomSHR)
	shrRequiredFee := requiredFees.AmountOf(gentleminttypes.DenomSHR)

	if shrRequiredFee.GT(shrTXFee) {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "got %s shr, required %s shr", shrTXFee, shrRequiredFee)
	}
	return next(ctx, tx, simulate)
}
