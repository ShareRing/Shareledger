package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) CreateLevelFee(goCtx context.Context, msg *types.MsgCreateLevelFee) (*types.MsgCreateLevelFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// If already have, will update current value
	var levelFee = types.LevelFee{
		Creator: msg.Creator,
		Level:   msg.Level,
		Fee:     msg.Fee,
	}

	k.SetLevelFee(
		ctx,
		levelFee,
	)
	return &types.MsgCreateLevelFeeResponse{}, nil
}

func (k msgServer) UpdateLevelFee(goCtx context.Context, msg *types.MsgUpdateLevelFee) (*types.MsgUpdateLevelFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	var levelFee = types.LevelFee{
		Creator: msg.Creator,
		Level:   msg.Level,
		Fee:     msg.Fee,
	}

	k.SetLevelFee(ctx, levelFee)

	return &types.MsgUpdateLevelFeeResponse{}, nil
}

func (k msgServer) DeleteLevelFee(goCtx context.Context, msg *types.MsgDeleteLevelFee) (*types.MsgDeleteLevelFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetLevelFee(
		ctx,
		msg.Level,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveLevelFee(
		ctx,
		msg.Level,
	)

	return &types.MsgDeleteLevelFeeResponse{}, nil
}
