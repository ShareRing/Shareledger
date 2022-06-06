package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) Reject(goCtx context.Context, msg *types.MsgReject) (*types.MsgRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Only support swap out
	reqs, err := k.Keeper.RejectSwap(ctx, msg)
	if err != nil {
		return nil, err
	}
	total := sdk.NewCoin(denom.Base, sdk.NewInt(0))
	reqIds := make([]string, 0, len(reqs))
	refunds := make(map[string]sdk.Coin)

	for _, r := range reqs {
		total = total.Add(r.GetAmount()).Add(r.Fee)
		reqIds = append(reqIds, fmt.Sprintf("%v", r.Id))
		re, found := refunds[r.SrcAddr]
		if !found {
			re = sdk.NewCoin(denom.Base, sdk.NewInt(0))
		}
		refunds[r.SrcAddr] = re.Add(r.Amount).Add(r.Fee)
	}

	for addr, refund := range refunds {
		refundAdd, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, refundAdd, sdk.NewCoins(refund)); err != nil {
			return nil, err
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeSwapReject,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.EventAttrApproverAction, types.SwapStatusRejected),
			sdk.NewAttribute(types.EventAttrRejectArr, msg.Creator),
			sdk.NewAttribute(types.EventAttrBatchTotal, total.String()),
			sdk.NewAttribute(types.EventAttrSwapId, strings.Join(reqIds, ",")),
		),
	)

	return &types.MsgRejectResponse{}, nil
}
