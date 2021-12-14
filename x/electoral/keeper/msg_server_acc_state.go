package keeper

//
//import (
//	"context"
//
//	"github.com/ShareRing/Shareledger/x/electoral/types"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//)
//
//func (k msgServer) CreateAccState(goCtx context.Context, msg *types.MsgCreateAccState) (*types.MsgCreateAccStateResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//	// Check if the value already exists
//	_, isFound := k.GetAccState(
//		ctx,
//		msg.Key,
//	)
//	if isFound {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
//	}
//
//	var accState = types.AccState{
//		Creator: msg.Creator,
//		Key:     msg.Key,
//		Address: msg.Address,
//		Status:  msg.Status,
//	}
//
//	k.SetAccState(
//		ctx,
//		accState,
//	)
//	return &types.MsgCreateAccStateResponse{}, nil
//}
//
//func (k msgServer) UpdateAccState(goCtx context.Context, msg *types.MsgUpdateAccState) (*types.MsgUpdateAccStateResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//	// Check if the value exists
//	valFound, isFound := k.GetAccState(
//		ctx,
//		msg.Key,
//	)
//	if !isFound {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
//	}
//
//	// Checks if the the msg creator is the same as the current owner
//	if msg.Creator != valFound.Creator {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
//	}
//
//	var accState = types.AccState{
//		Creator: msg.Creator,
//		Key:     msg.Key,
//		Address: msg.Address,
//		Status:  msg.Status,
//	}
//
//	k.SetAccState(ctx, accState)
//
//	return &types.MsgUpdateAccStateResponse{}, nil
//}
//
//func (k msgServer) DeleteAccState(goCtx context.Context, msg *types.MsgDeleteAccState) (*types.MsgDeleteAccStateResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//	// Check if the value exists
//	valFound, isFound := k.GetAccState(
//		ctx,
//		msg.Key,
//	)
//	if !isFound {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
//	}
//
//	// Checks if the the msg creator is the same as the current owner
//	if msg.Creator != valFound.Creator {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
//	}
//
//	k.RemoveAccState(
//		ctx,
//		msg.Key,
//	)
//
//	return &types.MsgDeleteAccStateResponse{}, nil
//}
