package keeper

import (
	"context"
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
		SrcAddr:     msg.SrcAddress,
		DestAddr:    msg.DestAddress,
		SrcNetwork:  types.NetworkNameShareLedger,
		DestNetwork: msg.Network,
		Amount:      insertAmountCoin,
		Fee:         insertFeeCoin,
		Status:      types.SwapStatusPending,
	})

	if err == nil {
		ctx.EventManager().EmitEvent(
			types.NewCreateRequestsEvent(msg.GetCreator(), req.Id, amount, fee, req.SrcAddr, req.SrcNetwork, req.DestAddr, req.DestNetwork, nil),
		)
	}
	return &types.MsgOutSwapResponse{
		Id: req.Id,
	}, err
}
