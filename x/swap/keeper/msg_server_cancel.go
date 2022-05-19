package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	refund := sdk.NewDecCoins()
	for i := range requests {
		rq := requests[i]
		pendingStore.Delete(GetRequestIDBytes(rq.GetId()))
		refund = refund.Add(sdk.NewDecCoinFromCoin(*rq.GetAmount())).Add(sdk.NewDecCoinFromCoin(*rq.GetFee()))
	}

	add, err := sdk.AccAddressFromBech32(txCreator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%+v", err)
	}
	bc, err := denom.NormalizeToBaseCoins(refund, false)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "%v", err)
	}
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, add, bc); err != nil {
		return nil, err
	}

	rStatusChange := sdk.NewEvent(types.EventTypeRequestCancelStatus).
		AppendAttributes(
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		)

	for i := range requests {
		ctx.EventManager().EmitEvent(
			rStatusChange.AppendAttributes(sdk.NewAttribute(types.EventTypeSwapId, fmt.Sprintf("%d", requests[i].GetId()))),
		)
	}

	return &types.MsgCancelResponse{}, nil
}
