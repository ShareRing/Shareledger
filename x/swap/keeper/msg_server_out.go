package keeper

import (
	"context"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) RequestOut(goCtx context.Context, msg *types.MsgRequestOut) (*types.MsgOutSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sumCoins := sdk.NewDecCoins().Add(*msg.Amount).Add(*msg.Fee)

	baseCoins, err := denom.NormalizeToBaseCoins(sumCoins, true)
	if err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.GetSigners()[0], types.ModuleName, baseCoins); err != nil {
		return nil, err
	}
	tn := time.Now().Unix()
	req, err := k.AppendPendingRequest(ctx, types.Request{
		SrcAddr:     msg.Creator,
		DestAddr:    msg.DestAddress,
		SrcNetwork:  types.NetworkNameShareLedger,
		DestNetwork: msg.Network,
		Amount:      msg.Amount,
		Fee:         msg.Fee,
		Status:      types.SwapStatusPending,
		CreatedAt:   uint64(tn),
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
	return &types.MsgOutSwapResponse{
		Id: req.Id,
	}, err
}
