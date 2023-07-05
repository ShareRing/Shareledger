package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) Cancel(goCtx context.Context, msg *types.MsgCancel) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pendingStore := k.GetStoreRequestMap(ctx)[types.SwapStatusPending]
	requests := k.GetRequestsByIdsFromStore(ctx, pendingStore, msg.GetIds())
	if len(msg.GetIds()) != len(requests) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "transactions don't have same status or not found, with required current status, %s", types.SwapStatusPending)
	}
	txCreator := msg.GetCreator()

	for i := range requests {
		reqSrcAddr := requests[i].GetSrcAddr()
		if reqSrcAddr != txCreator {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "your address %s isn't owner of swap request id=%s", txCreator, reqSrcAddr)
		}
	}

	refund := sdk.NewCoins()
	reqIds := make([]string, 0, len(requests))
	for i := range requests {
		rq := requests[i]
		pendingStore.Delete(GetRequestIDBytes(rq.GetId()))
		refund = refund.Add(rq.GetAmount()).Add(rq.GetFee())
		reqIds = append(reqIds, fmt.Sprintf("%v", rq.Id))
	}

	addr, err := sdk.AccAddressFromBech32(txCreator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%+v", err)
	}

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "%v", err)
	}
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, refund); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewCancelRequestEvent(msg.Creator, reqIds),
	)
	ctx.EventManager().EmitEvent(
		types.NewChangeRequestStatusesEvent(reqIds, types.SwapStatusCancelled))

	return &types.MsgCancelResponse{}, nil
}
