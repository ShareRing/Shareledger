package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
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
			return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "your address %s isn't owner of swap request id=%s", txCreator, reqSrcAddr)
		}
	}

	refunds := make(map[string]sdk.DecCoins)
	for i := range requests {
		rq := requests[i]
		pendingStore.Delete(GetRequestIDBytes(rq.GetId()))
		total, found := refunds[rq.GetSrcAddr()]
		if !found {
			total = sdk.NewDecCoins()
		}
		total = total.Add(*rq.GetAmount()).Add(*rq.GetFee())
		refunds[rq.GetSrcAddr()] = total
	}

	for addr, coin := range refunds {
		add, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%+v", err)
		}
		bc, err := denom.NormalizeToBaseCoins(coin, false)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrLogic, "%v", err)
		}
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, add, bc); err != nil {
			return err
		}
	}

	return nil
}
