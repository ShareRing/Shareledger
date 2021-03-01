package gatecheck

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/electoral"
)

type CheckValDecorator struct {
	ek electoral.Keeper
}

func NewCheckValDecorator(ek electoral.Keeper) CheckValDecorator {
	return CheckValDecorator{
		ek: ek,
	}
}

func (cvd CheckValDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	msg := msgs[0]
	switch msg.Type() {
	case "create_validator":
		if !electoral.IsEnrolledVoter(ctx, msg.GetSigners()[0], cvd.ek) {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Signer is not enrolled voter")
		}
	default:
	}
	return next(ctx, tx, simulate)
}
