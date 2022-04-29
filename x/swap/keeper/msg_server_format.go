package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) CreateFormat(goCtx context.Context, msg *types.MsgCreateFormat) (*types.MsgCreateFormatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetFormat(
		ctx,
		msg.Network,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var format = types.Format{
		Creator:    msg.Creator,
		Network:    msg.Network,
		DataFormat: msg.DataTypeFormat,
	}

	k.SetFormat(
		ctx,
		format,
	)
	return &types.MsgCreateFormatResponse{}, nil
}

func (k msgServer) UpdateFormat(goCtx context.Context, msg *types.MsgUpdateFormat) (*types.MsgUpdateFormatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetFormat(
		ctx,
		msg.Network,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var format = types.Format{
		Creator:    msg.Creator,
		Network:    msg.Network,
		DataFormat: msg.DataTypeFormat,
	}

	k.SetFormat(ctx, format)

	return &types.MsgUpdateFormatResponse{}, nil
}

func (k msgServer) DeleteFormat(goCtx context.Context, msg *types.MsgDeleteFormat) (*types.MsgDeleteFormatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetFormat(
		ctx,
		msg.Network,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveFormat(
		ctx,
		msg.Network,
	)

	return &types.MsgDeleteFormatResponse{}, nil
}
