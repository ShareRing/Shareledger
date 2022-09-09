package keeper

import (
	"context"

	denom "github.com/sharering/shareledger/x/utils/denom"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) RequestOut(goCtx context.Context, msg *types.MsgRequestOut) (*types.MsgRequestOutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.GetAmount()), true)
	if err != nil {
		return nil, err
	}

	schema, found := k.GetSchema(ctx, msg.Network)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network, %s, is not supported", msg.Network)
	}

	sumCoin := amount.Add(*schema.Fee.Out)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.GetSigners()[0], types.ModuleName, sumCoin); err != nil {
		return nil, err
	}

	var insertAmountCoin sdk.Coin
	insertAmountCoin = sdk.NewCoin(denom.Base, amount.AmountOf(denom.Base))

	req, err := k.AppendPendingRequest(ctx, types.Request{
		SrcAddr:     msg.SrcAddress,
		DestAddr:    msg.DestAddress,
		SrcNetwork:  types.NetworkNameShareLedger,
		DestNetwork: msg.Network,
		Amount:      insertAmountCoin,
		Fee:         *schema.Fee.Out,
		Status:      types.SwapStatusPending,
	})

	if err == nil {
		ctx.EventManager().EmitEvent(
			types.NewCreateRequestsEvent(msg.GetCreator(), req.Id, insertAmountCoin, *schema.Fee.Out, req.SrcAddr, req.SrcNetwork, req.DestAddr, req.DestNetwork, nil),
		)
	}
	return &types.MsgRequestOutResponse{
		Id: req.Id,
	}, err
}
