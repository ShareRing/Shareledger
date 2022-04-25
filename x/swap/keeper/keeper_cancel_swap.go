package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k Keeper) CancelSwap(ctx sdk.Context, msg *types.MsgCancel) error {
	pendingStore := k.GetStoreRequestMap(ctx)[types.SwapStatusPending]
	requests := k.GetRequestsByIdsFromStore(ctx, pendingStore, msg.GetIds())
	if len(msg.GetIds()) != len(requests) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "transactions don't have same status or not found, with required current status, %s", types.SwapStatusPending)
	}
	txCreator := msg.GetCreator()
	if txCreator == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator address is empty %s", txCreator)
	}
	for i := range requests {
		reqSrcAddr := requests[i].GetSrcAddr()
		if reqSrcAddr != txCreator {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "your address %s isn't owner of swap request id=%s", txCreator, reqSrcAddr)
		}
	}
	for i := range requests {
		pendingStore.Delete(GetRequestIDBytes(requests[i].GetId()))
	}
	return nil
}
