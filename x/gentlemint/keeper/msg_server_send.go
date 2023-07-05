package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	dstAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}
	baseCoins, err := denom.NormalizeToBaseCoins(msg.Coins, false)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if err := k.bankKeeper.SendCoins(ctx, msg.GetSigners()[0], dstAddr, baseCoins); err != nil {
		return nil, err
	}

	return &types.MsgSendResponse{}, nil
}
