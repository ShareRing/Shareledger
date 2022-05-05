package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) In(goCtx context.Context, msg *types.MsgIn) (*types.MsgInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	req, err := k.AppendPendingRequest(ctx, types.Request{
		SrcAddr:     msg.SrcAddress,
		DestAddr:    msg.DestAddress,
		SrcNetwork:  msg.SrcNetwork,
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      msg.Amount,
		Fee:         msg.Fee,
		Status:      types.SwapStatusPending,
	})
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSwapOut,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.EventTypeSwapAmount, msg.Amount.String()),
			sdk.NewAttribute(types.EventTypeSwapFee, msg.Fee.String()),
			sdk.NewAttribute(types.EventTypeSwapId, strconv.FormatUint(req.Id, 10)),
			sdk.NewAttribute(types.EventTypeSwapDestAddr, msg.DestAddress),
			sdk.NewAttribute(types.EventTypeSwapSrcAddr, msg.SrcAddress),
			sdk.NewAttribute(types.EventTypeSwapDestNetwork, types.NetworkNameShareLedger),
			sdk.NewAttribute(types.EventTypeSwapSrcNetwork, msg.SrcNetwork),
		),
	)

	return &types.MsgInResponse{RId: req.Id}, nil
}
