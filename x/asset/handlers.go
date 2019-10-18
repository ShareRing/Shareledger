package asset

import (
	"fmt"
	"reflect"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"

	// "github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/asset/messages"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
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

	event := sdk.NewEvent(
		EventTypeCreateAsset,
		sdk.NewAttribute(AttributeMsgModule, "asset"),
		sdk.NewAttribute(AttributeMsgAction, "create"),
		sdk.NewAttribute(AttributeAssetCreator, msg.Creator.String()),
		sdk.NewAttribute(AttributeAssetUUID, msg.UUID),
		sdk.NewAttribute(AttributeAssetHash, fmt.Sprintf("%X", msg.Hash)),
		sdk.NewAttribute(AttributeAssetStatus, strconv.FormatBool(msg.Status)),
		sdk.NewAttribute(AttributeAssetFee, strconv.Itoa(int(msg.Fee))),
	)
	// fee, denom := utils.GetMsgFee(msg)
	ctx.EventManager().EmitEvent(event)
	return sdk.Result{
		Log:    fmt.Sprintf("%s", asset),
		Events: ctx.EventManager().Events(),
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
	event := sdk.NewEvent(
		EventTypeUpdateAsset,
		sdk.NewAttribute(AttributeMsgModule, "asset"),
		sdk.NewAttribute(AttributeMsgAction, "update"),
		sdk.NewAttribute(AttributeAssetCreator, msg.Creator.String()),
		sdk.NewAttribute(AttributeAssetUUID, msg.UUID),
		sdk.NewAttribute(AttributeAssetHash, fmt.Sprintf("%X", msg.Hash)),
		sdk.NewAttribute(AttributeAssetStatus, strconv.FormatBool(msg.Status)),
		sdk.NewAttribute(AttributeAssetFee, strconv.Itoa(int(msg.Fee))),
	)
	ctx.EventManager().EmitEvent(event)
	return sdk.Result{
		Log:    fmt.Sprintf("%s", asset),
		Events: ctx.EventManager().Events(),
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
	event := sdk.NewEvent(
		EventTypeUpdateAsset,
		sdk.NewAttribute(AttributeMsgModule, "asset"),
		sdk.NewAttribute(AttributeMsgAction, "delete"),
		sdk.NewAttribute(AttributeAssetUUID, msg.UUID),
	)
	ctx.EventManager().EmitEvent(event)
	return sdk.Result{
		Log:    fmt.Sprintf("%s", asset),
		Events: ctx.EventManager().Events(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}
