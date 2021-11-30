package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BurnShr(goCtx context.Context, msg *types.MsgBurnShr) (*types.MsgBurnShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if !k.isTreasurer(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not Treasurer")
	}

	shrCoins, err := types.ParseShrCoinsStr(msg.Amount)
	if err != nil {
		return nil, err
	}
	if err := k.burnCoins(ctx, msg.GetSigners()[0], shrCoins); err != nil {
		return nil, sdkerrors.Wrapf(err, "burns %v coins from %v", shrCoins, msg.Creator)
	}

	return &types.MsgBurnShrResponse{
		Log: fmt.Sprintf("Successfully burn %d shr", msg.Amount),
	}, nil
}
