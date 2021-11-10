package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ShareRing/Shareledger/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Create(goCtx context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oldAsset := k.GetAsset(ctx, msg.UUID)
	if len(oldAsset.Creator) > 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Asset already exists")
	}
	asset := types.NewAssetFromMsgCreate(*msg)
	k.SetAsset(ctx, msg.UUID, asset)
	// log, err := asset.GetString()
	// if err != nil {
	// 	return nil, err
	// }
	event := sdk.NewEvent(
		types.EventTypeCreateAsset,
		sdk.NewAttribute(types.AttributeMsgModule, "asset"),
		sdk.NewAttribute(types.AttributeMsgAction, "create"),
		sdk.NewAttribute(types.AttributeAssetCreator, msg.Creator),
		sdk.NewAttribute(types.AttributeAssetUUID, msg.UUID),
		sdk.NewAttribute(types.AttributeAssetHash, fmt.Sprintf("%X", msg.Hash)),
		sdk.NewAttribute(types.AttributeAssetStatus, strconv.FormatBool(msg.Status)),
		sdk.NewAttribute(types.AttributeAssetFee, strconv.Itoa(int(msg.Rate))),
	)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgCreateResponse{}, nil
}

func (k msgServer) Update(goCtx context.Context, msg *types.MsgUpdate) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oldAsset := k.GetAsset(ctx, msg.UUID)
	if oldAsset.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only creator can update asset")
	}
	asset := types.NewAssetFromMsgUpdate(*msg)
	k.SetAsset(ctx, msg.UUID, asset)
	// log, err := asset.GetString()
	// if err != nil {
	// 	return nil, err
	// }

	event := sdk.NewEvent(
		types.EventTypeUpdateAsset,
		sdk.NewAttribute(types.AttributeMsgModule, "asset"),
		sdk.NewAttribute(types.AttributeMsgAction, "update"),
		sdk.NewAttribute(types.AttributeAssetCreator, msg.Creator),
		sdk.NewAttribute(types.AttributeAssetUUID, msg.UUID),
		sdk.NewAttribute(types.AttributeAssetHash, fmt.Sprintf("%X", msg.Hash)),
		sdk.NewAttribute(types.AttributeAssetStatus, strconv.FormatBool(msg.Status)),
		sdk.NewAttribute(types.AttributeAssetFee, strconv.Itoa(int(msg.Rate))),
	)

	ctx.EventManager().EmitEvent(event)
	return &types.MsgUpdateResponse{}, nil
}

func (k msgServer) Delete(goCtx context.Context, msg *types.MsgDelete) (*types.MsgDeleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oldAsset := k.GetAsset(ctx, msg.UUID)
	if oldAsset.Creator != msg.Owner {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only Owner can update asset")
	}
	// log, err := oldAsset.GetString()
	// if err != nil {
	// 	return nil, err
	// }
	event := sdk.NewEvent(
		types.EventTypeDeleteAsset,
		sdk.NewAttribute(types.AttributeMsgModule, "asset"),
		sdk.NewAttribute(types.AttributeMsgAction, "delete"),
		sdk.NewAttribute(types.AttributeAssetUUID, msg.UUID),
	)
	ctx.EventManager().EmitEvent(event)
	k.DeleteAsset(ctx, msg.UUID)
	ctx.EventManager().EmitEvent(event)
	return &types.MsgDeleteResponse{}, nil
}
