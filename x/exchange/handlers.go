package exchange

import (
	"fmt"
	"reflect"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/exchange/messages"
	etypes "github.com/sharering/shareledger/x/exchange/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case messages.MsgCreate:
			return handleMsgCreate(ctx, k, msg)
		case messages.MsgRetrieve:
			return handleMsgRetrieve(ctx, k, msg)
		case messages.MsgUpdate:
			return handleMsgUpdate(ctx, k, msg)
		case messages.MsgDelete:
			return handleMsgDelete(ctx, k, msg)
		case messages.MsgExchange:
			return handleMsgExchange(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
			return sdk.ErrUnknownRequest(errMsg).Result()
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
	fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:       fmt.Sprintf("%s", exr),
		Tags:      msg.Tags(),
		FeeAmount: fee,
		FeeDenom:  denom,
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
	fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:       fmt.Sprintf("%s", exr),
		Tags:      msg.Tags(),
		FeeAmount: fee,
		FeeDenom:  denom,
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
	fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:       fmt.Sprintf("%s", exr),
		Tags:      msg.Tags(),
		FeeAmount: fee,
		FeeDenom:  denom,
	}
}

func handleMsgRetrieve(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgRetrieve,
) sdk.Result {

	exr, err := k.RetrieveExchangeRate(ctx, msg.FromDenom, msg.ToDenom)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	return sdk.Result{
		Log:  fmt.Sprintf("%s", exr),
		Tags: msg.Tags(),
	}
}

func handleMsgExchange(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgExchange,
) sdk.Result {

	exr, err := k.RetrieveExchangeRate(ctx, msg.FromDenom, msg.ToDenom)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// The account sign this tx is the buying account
	signer := auth.GetSigner(ctx)

	// Get address
	address := signer.GetAddress()

	// Get balance
	fromAcc := k.bankKeeper.GetCoins(ctx, address)

	// Already check validity with the message
	reserve := etypes.NewReserve(msg.Reserve)

	reserveAcc := reserve.GetCoins(ctx, k.bankKeeper)

	sellingCoin := types.NewCoinFromDec(msg.FromDenom, msg.Amount)

	buyingCoin := exr.Convert(sellingCoin)

	if fromAcc.LT(sellingCoin) || reserveAcc.LT(buyingCoin) {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_INSUFFICIENT_BALANCE,
			fromAcc.String(),
			sellingCoin.String(),
			reserveAcc.String(),
			buyingCoin.String())).Result()
	}

	// Transfer selling currencies from FromAcc to ReserveAcc
	newFromAcc := fromAcc.Minus(sellingCoin)
	newReserveAcc := reserveAcc.Plus(sellingCoin)

	// Transfer Buying currencies from ReserveAcc to FromAcc
	newReserveAcc = newReserveAcc.Minus(buyingCoin)
	newFromAcc = newFromAcc.Plus(buyingCoin)

	// Save to store

	sdkErr := k.bankKeeper.SetCoins(ctx, address, newFromAcc)

	if sdkErr != nil {
		return sdkErr.Result()
	}

	sdkErr = reserve.SetCoins(ctx, k.bankKeeper, newReserveAcc)

	if sdkErr != nil {
		return sdkErr.Result()
	}

	return sdk.Result{
		Log:  fmt.Sprintf("%s", newFromAcc.String()),
		Tags: msg.Tags(),
	}
}
