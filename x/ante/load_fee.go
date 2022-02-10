package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gentleminttypes "github.com/sharering/shareledger/x/gentlemint/types"
)

type LoadFeeDecorator struct {
	gk GentlemintKeeper
}

func NewLoadFeeDecorator(gk GentlemintKeeper) LoadFeeDecorator {
	return LoadFeeDecorator{
		gk: gk,
	}
}

func (cfd LoadFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	for _, m := range msgs {
		switch loadFeeMessage := m.(type) {
		case *gentleminttypes.MsgLoadFee:
			if err := cfd.gk.LoadFeeFundFromShrp(ctx, loadFeeMessage); err != nil {
				return ctx, sdkerrors.Wrapf(err, "load fee from %v", loadFeeMessage.Shrp)
			}
		}
	}
	return next(ctx, tx, simulate)
}
