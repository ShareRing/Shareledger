package keeper

import (
	"context"
	"strconv"

	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) RequestOut(goCtx context.Context, msg *types.MsgRequestOut) (*types.MsgOutSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.GetAmount()), true)
	if err != nil {
		return nil, err
	}

	fee, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.GetFee()), true)
	if err != nil {
		return nil, err
	}

	sumCoin := amount.Add(fee...)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.GetSigners()[0], types.ModuleName, sumCoin); err != nil {
		return nil, err
	}

	var insertAmountCoin sdk.Coin
	var insertFeeCoin sdk.Coin
	insertFeeCoin = sdk.NewCoin(denom.Base, fee.AmountOf(denom.Base))
	insertAmountCoin = sdk.NewCoin(denom.Base, amount.AmountOf(denom.Base))

	req, err := k.AppendPendingRequest(ctx, types.Request{
		SrcAddr:     msg.Creator,
		DestAddr:    msg.DestAddress,
		SrcNetwork:  types.NetworkNameShareLedger,
		DestNetwork: msg.Network,
		Amount:      insertAmountCoin,
		Fee:         insertFeeCoin,
		Status:      types.SwapStatusPending,
	})

	if err == nil {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeSwapOut,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(types.EventTypeSwapAmount, amount.String()),
				sdk.NewAttribute(types.EventTypeSwapFee, fee.String()),
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
