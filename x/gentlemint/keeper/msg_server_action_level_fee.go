package keeper

import (
	"context"
	"github.com/sharering/shareledger/x/fee"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) SetActionLevelFee(goCtx context.Context, msg *types.MsgSetActionLevelFee) (*types.MsgSetActionLevelFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if found := fee.HaveActionKey(msg.Action); !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "%s action was not found in fee table", msg.Action)
	}

	if _, found := k.GetLevelFee(ctx, msg.Level); !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "%s level was not found", msg.Level)
	}

	var actionLevelFee = types.ActionLevelFee{
		Creator: msg.Creator,
		Action:  msg.Action,
		Level:   msg.Level,
	}

	k.Keeper.SetActionLevelFee(
		ctx,
		actionLevelFee,
	)
	return &types.MsgSetActionLevelFeeResponse{}, nil
}

func (k msgServer) DeleteActionLevelFee(goCtx context.Context, msg *types.MsgDeleteActionLevelFee) (*types.MsgDeleteActionLevelFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetActionLevelFee(
		ctx,
		msg.Action,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveActionLevelFee(
		ctx,
		msg.Action,
	)

	return &types.MsgDeleteActionLevelFeeResponse{}, nil
}
