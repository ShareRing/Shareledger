package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"strconv"
)

func (k msgServer) RequestIn(goCtx context.Context, msg *types.MsgRequestIn) (*types.MsgSwapInResponse, error) {
	msg.ValidateBasic()
	ctx := sdk.UnwrapSDKContext(goCtx)
	amount, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.GetAmount()), true)
	if err != nil {
		return nil, err
	}

	fee, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.GetFee()), true)
	if err != nil {
		return nil, err
	}
	var insertAmountCoin sdk.Coin
	var insertFeeCoin sdk.Coin
	insertFeeCoin = sdk.NewCoin(denom.Base, fee.AmountOf(denom.Base))
	insertAmountCoin = sdk.NewCoin(denom.Base, amount.AmountOf(denom.Base))
	req, err := k.AppendPendingRequest(ctx, types.Request{
		SrcAddr:     msg.SrcAddress,
		DestAddr:    msg.DestAddress,
		SrcNetwork:  msg.Network,
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      &insertAmountCoin,
		Fee:         &insertFeeCoin,
		Status:      types.SwapStatusPending,
		TxHashes:    msg.TxHashes,
		CreatedAt:   0,
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
			sdk.NewAttribute(types.EventTypeSwapSrcNetwork, msg.Network),
		),
	)

	return &types.MsgSwapInResponse{Id: req.Id}, nil
}
