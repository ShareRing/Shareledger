package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	baseCoins, err := denom.NormalizeToBaseCoins(msg.Coins, false)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	if err := k.burnCoins(ctx, msg.GetSigners()[0], baseCoins); err != nil {
		return nil, sdkerrors.Wrapf(err, "burns %v coins from %v", baseCoins, msg.Creator)
	}
	return &types.MsgBurnResponse{}, nil
}
