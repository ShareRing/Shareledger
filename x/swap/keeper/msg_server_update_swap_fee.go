package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/denom"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) UpdateSwapFee(goCtx context.Context, msg *types.MsgUpdateSwapFee) (*types.MsgUpdateSwapFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	v, found := k.GetSchema(ctx, msg.Network)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "the network %s schema is not found", msg.GetNetwork())
	}
	if msg.Out != nil && !msg.Out.IsZero() {
		outF, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(*msg.Out), sdk.NewDec(0), false)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid fee")
		}
		v.Fee.Out = &outF

	}
	if msg.In != nil && !msg.In.IsZero() {
		inF, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(*msg.In), sdk.NewDec(0), false)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid fee")
		}
		v.Fee.In = &inF
	}

	k.SetSchema(ctx, v)
	return &types.MsgUpdateSwapFeeResponse{}, nil
}
