package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) CreateSignSchema(goCtx context.Context, msg *types.MsgCreateSignSchema) (*types.MsgCreateSignSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetFormat(
		ctx,
		msg.Network,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var format = types.SignSchema{
		Creator: msg.Creator,
		Network: msg.Network,
		Schema:  msg.Schema,
	}

	k.SetFormat(
		ctx,
		format,
	)
	return &types.MsgCreateSignSchemaResponse{}, nil
}

func (k msgServer) UpdateSignSchema(goCtx context.Context, msg *types.MsgUpdateSignSchema) (*types.MsgUpdateSignSchemaResponse, error) {
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

	var format = types.SignSchema{
		Creator: msg.Creator,
		Network: msg.Network,
		Schema:  msg.Schema,
	}

	k.SetFormat(ctx, format)

	return &types.MsgUpdateSignSchemaResponse{}, nil
}

func (k msgServer) DeleteSignSchema(goCtx context.Context, msg *types.MsgDeleteSignSchema) (*types.MsgDeleteSignSchemaResponse, error) {
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

	return &types.MsgDeleteSignSchemaResponse{}, nil
}
