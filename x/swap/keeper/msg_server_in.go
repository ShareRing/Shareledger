package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (k msgServer) RequestIn(goCtx context.Context, msg *types.MsgRequestIn) (*types.MsgRequestInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.GetSchema(ctx, msg.Network)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network %s, is not supported", msg.Network)
	}

	amount, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.GetAmount()), false)
	if err != nil {
		return nil, err
	}

	var insertAmountCoin sdk.Coin
	insertAmountCoin = sdk.NewCoin(denom.Base, amount.AmountOf(denom.Base))

	req, err := k.AppendPendingRequest(ctx, types.Request{
		DestAddr:    msg.DestAddress,
		SrcNetwork:  msg.Network,
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      insertAmountCoin,
		Fee:         *schema.Fee.In,
		Status:      types.SwapStatusPending,
		TxEvents:    msg.TxEvents,
		SrcAddr:     msg.SrcAddress,
	})
	if err != nil {
		return nil, err
	}

	hashEvent := make(map[string]string)

	for i, e := range msg.TxEvents {
		hashEvent[fmt.Sprintf("tx_%d", i)] = fmt.Sprintf("%s:%s:%d", e.TxHash, e.Sender, e.LogIndex)
	}
	ctx.EventManager().EmitEvent(
		types.NewCreateRequestsEvent(
			msg.GetCreator(),
			req.Id,
			insertAmountCoin,
			*schema.Fee.In,
			msg.SrcAddress,
			msg.Network,
			msg.DestAddress, types.NetworkNameShareLedger,
			hashEvent,
		),
	)

	return &types.MsgRequestInResponse{Id: req.Id}, nil
}
