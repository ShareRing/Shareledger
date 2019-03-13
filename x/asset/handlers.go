package asset

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	// "github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/asset/messages"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		constants.LOGGER.Info(
			"Msg for Asset Module",
			"type", reflect.TypeOf(msg),
			"msg", msg,
		)

		switch msg := msg.(type) {

		case messages.MsgCreate:
			return handleAssetCreation(ctx, k, msg)
		case messages.MsgUpdate:
			return handleAssetUpdate(ctx, k, msg)
		case messages.MsgDelete:
			return handleAssetDelete(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleAssetCreation(ctx sdk.Context, k Keeper, msg messages.MsgCreate) sdk.Result {

	asset, err := k.CreateAsset(ctx, msg)
	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:       fmt.Sprintf("%s", asset),
		Tags:      msg.Tags(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}

func handleAssetUpdate(ctx sdk.Context, k Keeper, msg messages.MsgUpdate) sdk.Result {

	asset, err := k.UpdateAsset(ctx, msg)
	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:       fmt.Sprintf("%s", asset),
		Tags:      msg.Tags(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}

func handleAssetDelete(ctx sdk.Context, k Keeper, msg messages.MsgDelete) sdk.Result {

	asset, err := k.DeleteAsset(ctx, msg)
	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:       fmt.Sprintf("%s", asset),
		Tags:      msg.Tags(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}
