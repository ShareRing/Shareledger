// ensure that user has enough shr token to pay fee
package gatecheck

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/gentlemint"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

type FeeEnsureDecorator struct {
	gk gentlemint.Keeper
}

func NewFeeEnsureDecorator(gk gentlemint.Keeper) FeeEnsureDecorator {
	return FeeEnsureDecorator{
		gk: gk,
	}
}

func (fsd FeeEnsureDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, _ := tx.(ante.FeeTx)
	fee := feeTx.GetFee()
	shrAmt := fee.AmountOf("shr")
	amtToBuy, notEnough := fsd.gk.NotEnoughShr(ctx, shrAmt, feeTx.FeePayer())
	if notEnough {
		err := fsd.gk.BuyShr(ctx, amtToBuy, feeTx.FeePayer())
		if err != nil {
			return ctx, err
		}
	}
	return next(ctx, tx, simulate)
}
