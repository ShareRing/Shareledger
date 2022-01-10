package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) CreateActionLevelFee(goCtx context.Context, msg *types.MsgCreateActionLevelFee) (*types.MsgCreateActionLevelFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if found := k.actionsTable.HaveAction(msg.Action); !found {
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

	k.SetActionLevelFee(
		ctx,
		actionLevelFee,
	)
	return &types.MsgCreateActionLevelFeeResponse{}, nil
}

func (k msgServer) UpdateActionLevelFee(goCtx context.Context, msg *types.MsgUpdateActionLevelFee) (*types.MsgUpdateActionLevelFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetActionLevelFee(
		ctx,
		msg.Action,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the  msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var actionLevelFee = types.ActionLevelFee{
		Creator: msg.Creator,
		Action:  msg.Action,
		Level:   msg.Level,
	}

	k.SetActionLevelFee(ctx, actionLevelFee)

	return &types.MsgUpdateActionLevelFeeResponse{}, nil
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
