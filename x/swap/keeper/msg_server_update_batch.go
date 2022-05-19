package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) UpdateBatch(goCtx context.Context, msg *types.MsgUpdateBatch) (*types.MsgUpdateBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	batch, found := k.GetBatch(ctx, msg.GetBatchId())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "batch id=%s not found", msg.GetBatchId())
	}

	//var wantedStatus = msg.GetStatus()
	//var requireStatus string
	//switch wantedStatus {
	//case types.BatchStatusDone, types.BatchStatusFail:
	//	requireStatus = types.BatchStatusProcessing
	//case types.BatchStatusProcessing:
	//	requireStatus = types.BatchStatusPending
	//
	//default:
	//	return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the %s status is not supported", wantedStatus)
	//}
	//
	//if batch.GetStatus() != requireStatus {
	//	return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "updating batch status [%s] => [%s] is not allowed", batch.GetStatus(), msg.GetStatus())
	//}

	batch.Status = msg.GetStatus()
	batch.Network = msg.GetNetwork()
	k.SetBatch(ctx, batch)

	return &types.MsgUpdateBatchResponse{}, nil
}
