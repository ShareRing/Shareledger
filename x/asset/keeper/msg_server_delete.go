package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Delete(goCtx context.Context, msg *types.MsgDelete) (*types.MsgDeleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	oldAsset, found := k.GetAsset(ctx, msg.GetUUID())

	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "Asset not found")
	}

	if oldAsset.GetCreator() != msg.GetOwner() {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only owner can update asset")
	}

	k.DeleteAsset(ctx, msg.GetUUID())

	event := sdk.NewEvent(
		types.EventTypeDeleteAsset,
		sdk.NewAttribute(types.AttributeMsgModule, "asset"),
		sdk.NewAttribute(types.AttributeMsgAction, "delete"),
		sdk.NewAttribute(types.AttributeAssetUUID, msg.UUID),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgDeleteResponse{}, nil
}
