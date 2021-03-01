package asset

import (
	"fmt"
	"strconv"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgCreate:
			return handleCreateAsset(ctx, keeper, msg)
		case MsgUpdate:
			return handleUpdateAsset(ctx, keeper, msg)
		case MsgDelete:
			return handleDeleteAsset(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized asset Msg type: %v", msg.Type()))
		}
	}
}

func handleCreateAsset(ctx sdk.Context, keeper Keeper, msg MsgCreate) (*sdk.Result, error) {
	oldAsset := keeper.GetAsset(ctx, msg.UUID)
	if !oldAsset.Creator.Empty() {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Asset already exists")
	}
	asset := types.NewAssetFromMsgCreate(msg)
	keeper.SetAsset(ctx, msg.UUID, asset)
	log, err := asset.GetString()
	if err != nil {
		return nil, err
	}
	event := sdk.NewEvent(
		EventTypeCreateAsset,
		sdk.NewAttribute(AttributeMsgModule, "asset"),
		sdk.NewAttribute(AttributeMsgAction, "create"),
		sdk.NewAttribute(AttributeAssetCreator, msg.Creator.String()),
		sdk.NewAttribute(AttributeAssetUUID, msg.UUID),
		sdk.NewAttribute(AttributeAssetHash, fmt.Sprintf("%X", msg.Hash)),
		sdk.NewAttribute(AttributeAssetStatus, strconv.FormatBool(msg.Status)),
		sdk.NewAttribute(AttributeAssetFee, strconv.Itoa(int(msg.Rate))),
	)
	ctx.EventManager().EmitEvent(event)
	return &sdk.Result{
		Log:    log,
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleUpdateAsset(ctx sdk.Context, keeper Keeper, msg MsgUpdate) (*sdk.Result, error) {
	oldAsset := keeper.GetAsset(ctx, msg.UUID)
	if !oldAsset.Creator.Equals(msg.Creator) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only creator can update asset")
	}
	asset := types.NewAssetFromMsgUpdate(msg)
	keeper.SetAsset(ctx, msg.UUID, asset)
	log, err := asset.GetString()
	if err != nil {
		return nil, err
	}
	event := sdk.NewEvent(
		EventTypeUpdateAsset,
		sdk.NewAttribute(AttributeMsgModule, "asset"),
		sdk.NewAttribute(AttributeMsgAction, "update"),
		sdk.NewAttribute(AttributeAssetCreator, msg.Creator.String()),
		sdk.NewAttribute(AttributeAssetUUID, msg.UUID),
		sdk.NewAttribute(AttributeAssetHash, fmt.Sprintf("%X", msg.Hash)),
		sdk.NewAttribute(AttributeAssetStatus, strconv.FormatBool(msg.Status)),
		sdk.NewAttribute(AttributeAssetFee, strconv.Itoa(int(msg.Rate))),
	)

	ctx.EventManager().EmitEvent(event)
	return &sdk.Result{
		Log:    log,
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleDeleteAsset(ctx sdk.Context, keeper Keeper, msg MsgDelete) (*sdk.Result, error) {
	oldAsset := keeper.GetAsset(ctx, msg.UUID)
	if !oldAsset.Creator.Equals(msg.Owner) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only Owner can update asset")
	}
	log, err := oldAsset.GetString()
	if err != nil {
		return nil, err
	}
	event := sdk.NewEvent(
		EventTypeDeleteAsset,
		sdk.NewAttribute(AttributeMsgModule, "asset"),
		sdk.NewAttribute(AttributeMsgAction, "delete"),
		sdk.NewAttribute(AttributeAssetUUID, msg.UUID),
	)
	ctx.EventManager().EmitEvent(event)
	keeper.DeleteAsset(ctx, msg.UUID)
	return &sdk.Result{
		Log:    log,
		Events: ctx.EventManager().Events(),
	}, nil
}
