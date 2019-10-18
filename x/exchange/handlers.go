package exchange

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/exchange/messages"

	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
)

func NewHandler(k Keeper) sdkTypes.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdkTypes.Result {
		var ret sdk.Result
		switch msg := msg.(type) {
		case messages.MsgCreate:
			ret = handleMsgCreate(ctx, k, msg)
		case messages.MsgUpdate:
			ret = handleMsgUpdate(ctx, k, msg)
		case messages.MsgDelete:
			ret = handleMsgDelete(ctx, k, msg)
		case messages.MsgExchange:
			ret = handleMsgExchange(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
			return sdkTypes.NewResult(sdk.ErrUnknownRequest(errMsg).Result())
		}

		if !ret.IsOK() {
			return sdkTypes.NewResult(ret)
		}

		fee, denom := utils.GetMsgFee(msg)
		return sdkTypes.Result{
			Result:    ret,
			FeeDenom:  denom,
			FeeAmount: fee,
		}

	}
}

func handleMsgCreate(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgCreate,
) sdk.Result {

	exr, err := k.CreateExchangeRate(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// TODO: MsgFee is based on name of Msg. Currently, Asset and This module ( Exchagne) share the same set of names
	// Create, Delete, Update
	// fee, denom := utils.GetMsgFee(msg)
	event := sdk.NewEvent(
		EventTypeExchangeCreate,
		sdk.NewAttribute(AttributeFromDenom, msg.FromDenom),
		sdk.NewAttribute(AttributeToDenom, msg.ToDenom),
	)
	ctx.EventManager().EmitEvent(event)
	return sdk.Result{
		Log:    exr.String(),
		Events: ctx.EventManager().Events(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}

func handleMsgUpdate(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgUpdate,
) sdk.Result {

	exr, err := k.UpdateExchangeRate(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// TODO: MsgFee is based on name of Msg. Currently, Asset and This module ( Exchagne) share the same set of names
	// Create, Delete, Update
	// fee, denom := utils.GetMsgFee(msg)
	event := sdk.NewEvent(
		EventTypeExchangeUpdate,
		sdk.NewAttribute(AttributeFromDenom, msg.FromDenom),
		sdk.NewAttribute(AttributeToDenom, msg.ToDenom),
	)
	ctx.EventManager().EmitEvent(event)
	return sdk.Result{
		Log:    fmt.Sprintf("%s", exr),
		Events: ctx.EventManager().Events(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}

func handleMsgDelete(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgDelete,
) sdk.Result {

	exr, err := k.DeleteExchangeRate(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// TODO: MsgFee is based on name of Msg. Currently, Asset and This module ( Exchagne) share the same set of names
	// Create, Delete, Update
	// fee, denom := utils.GetMsgFee(msg)
	event := sdk.NewEvent(
		EventTypeExchangeDelete,
		sdk.NewAttribute(AttributeFromDenom, msg.FromDenom),
		sdk.NewAttribute(AttributeToDenom, msg.ToDenom),
	)
	ctx.EventManager().EmitEvent(event)
	return sdk.Result{
		Log:    fmt.Sprintf("%s", exr),
		Events: ctx.EventManager().Events(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}

func handleMsgExchange(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgExchange,
) sdk.Result {

	// The account sign this tx is the buying account
	signer := auth.GetSigner(ctx)

	// Get address
	address := signer.GetAddress()

	err := k.SellCoin(
		ctx,
		address,
		msg.Reserve,
		msg.FromDenom,
		msg.ToDenom,
		msg.Amount,
	)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}
	event := sdk.NewEvent(
		EventTypeExchange,
		sdk.NewAttribute(AttributeFromDenom, msg.FromDenom),
		sdk.NewAttribute(AttributeToDenom, msg.ToDenom),
		sdk.NewAttribute(AttributeAmount, msg.Amount.String()),
	)
	ctx.EventManager().EmitEvent(event)
	balanceAfter := k.bankKeeper.GetCoins(ctx, address)

	return sdk.Result{
		Log:    fmt.Sprintf("%s", balanceAfter.String()),
		Events: ctx.EventManager().Events(),
	}
}
