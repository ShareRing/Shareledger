package keeper

import (
	"context"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) Out(goCtx context.Context, msg *types.MsgOut) (*types.MsgOutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sumCoins := sdk.NewDecCoins().Add(*msg.Amount).Add(*msg.Fee)

	baseCoins, err := denom.NormalizeToBaseCoins(sumCoins, true)
	if err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.GetSigners()[0], types.ModuleName, baseCoins); err != nil {
		return nil, err
	}

	req, err := k.AppendPendingRequest(ctx, types.Request{
		SrcAddr:     msg.Creator,
		DestAddr:    msg.DestAddr,
		SrcNetwork:  types.NetworkNameShareLedger,
		DestNetwork: msg.Network,
		Amount:      msg.Amount,
		Fee:         msg.Fee,
		Status:      types.SwapStatusPending,
	})

	if err == nil {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeSwapOut,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(types.EventTypeSwapAmount, req.Amount.String()),
				sdk.NewAttribute(types.EventTypeSwapFee, req.Fee.String()),
				sdk.NewAttribute(types.EventTypeSwapId, strconv.FormatUint(req.Id, 10)),
				sdk.NewAttribute(types.EventTypeSwapDestAddr, req.DestAddr),
				sdk.NewAttribute(types.EventTypeSwapSrcAddr, req.SrcAddr),
				sdk.NewAttribute(types.EventTypeSwapDestNetwork, req.DestNetwork),
				sdk.NewAttribute(types.EventTypeSwapSrcNetwork, req.SrcNetwork),
			),
		)
	}
	return &types.MsgOutResponse{
		Rid: req.Id,
	}, err
}
