package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/asset/types"
)

func (k msgServer) UpdateAsset(goCtx context.Context, msg *types.MsgUpdateAsset) (*types.MsgUpdateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	oldAsset, found := k.GetAsset(ctx, msg.GetUUID())

	if !found {
		return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.GetUUID())
	}

	if oldAsset.GetCreator() != msg.GetCreator() {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only creator can update asset")
	}

	asset := types.NewAssetFromMsgUpdate(*msg)
	k.SetAsset(ctx, msg.GetUUID(), asset)

	event := sdk.NewEvent(
		types.EventTypeUpdateAsset,
		sdk.NewAttribute(types.AttributeMsgModule, "asset"),
		sdk.NewAttribute(types.AttributeMsgAction, "update"),
		sdk.NewAttribute(types.AttributeAssetCreator, msg.GetCreator()),
		sdk.NewAttribute(types.AttributeAssetUUID, msg.GetUUID()),
		sdk.NewAttribute(types.AttributeAssetHash, fmt.Sprintf("%X", msg.GetHash())),
		sdk.NewAttribute(types.AttributeAssetStatus, strconv.FormatBool(msg.GetStatus())),
		sdk.NewAttribute(types.AttributeAssetFee, strconv.Itoa(int(msg.GetRate()))),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgUpdateAssetResponse{}, nil
}
