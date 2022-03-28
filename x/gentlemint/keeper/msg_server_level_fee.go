package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) SetLevelFee(goCtx context.Context, msg *types.MsgSetLevelFee) (*types.MsgSetLevelFeeResponse, error) {
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

	k.Keeper.SetLevelFee(
		ctx,
		levelFee,
	)
	return &types.MsgSetLevelFeeResponse{}, nil
}

func (k msgServer) DeleteLevelFee(goCtx context.Context, msg *types.MsgDeleteLevelFee) (*types.MsgDeleteLevelFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetLevelFee(
		ctx,
		msg.Level,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveLevelFee(
		ctx,
		msg.Level,
	)

	return &types.MsgDeleteLevelFeeResponse{}, nil
}
