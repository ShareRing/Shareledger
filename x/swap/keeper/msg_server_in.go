package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"strconv"
)

func (k msgServer) RequestIn(goCtx context.Context, msg *types.MsgRequestIn) (*types.MsgSwapInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reqHistory, found := k.GetRequestedIn(ctx, msg.DestAddress)
	if found {
		for _, hash := range msg.TxHashes {
			if _, processed := reqHistory.TxHashes[hash]; processed {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "tx hash was already processed")
			}
		}
	}

	fee, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.Fee), true)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}
	baseFee := sdk.NewCoin(denom.Base, fee.AmountOf(denom.Base))

	schema, found := k.GetSchema(ctx, msg.Network)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network, %s, is not supported", msg.Network)
	}
	if schema.Fee != nil && schema.Fee.In != nil && baseFee.IsLT(*schema.Fee.In) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "required fee for swap in is expected %s, but got %s", baseFee.String(), baseFee.String())
	}

	amount, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.GetAmount()), true)
	if err != nil {
		return nil, err
	}

	var insertAmountCoin sdk.Coin
	insertAmountCoin = sdk.NewCoin(denom.Base, amount.AmountOf(denom.Base))

	slpAddress, err := sdk.AccAddressFromBech32(msg.DestAddress)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	k.SetRequestedIn(ctx, slpAddress, msg.TxHashes)
	req, err := k.AppendPendingRequest(ctx, types.Request{
		DestAddr:    msg.DestAddress,
		SrcNetwork:  msg.Network,
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      &insertAmountCoin,
		Fee:         &baseFee,
		Status:      types.SwapStatusPending,
		TxHashes:    msg.TxHashes,
	})
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSwapIn,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.EventTypeSwapAmount, msg.Amount.String()),
			sdk.NewAttribute(types.EventTypeSwapFee, msg.Fee.String()),
			sdk.NewAttribute(types.EventTypeSwapId, strconv.FormatUint(req.Id, 10)),
			sdk.NewAttribute(types.EventTypeSwapDestAddr, msg.DestAddress),
			sdk.NewAttribute(types.EventTypeSwapDestNetwork, types.NetworkNameShareLedger),
			sdk.NewAttribute(types.EventTypeSwapSrcNetwork, msg.Network),
		),
	)

	return &types.MsgSwapInResponse{Id: req.Id}, nil
}
