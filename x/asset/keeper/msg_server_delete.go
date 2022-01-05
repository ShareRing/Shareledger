package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/asset/types"
)

func (k msgServer) DeleteAsset(goCtx context.Context, msg *types.MsgDeleteAsset) (*types.MsgDeleteAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	oldAsset, found := k.GetAsset(ctx, msg.GetUUID())

	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "Asset not found")
	}

	if oldAsset.GetCreator() != msg.GetOwner() {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only owner can update asset")
	}

	k.Keeper.DeleteAsset(ctx, msg.GetUUID())

	event := sdk.NewEvent(
		types.EventTypeDeleteAsset,
		sdk.NewAttribute(types.AttributeMsgModule, "asset"),
		sdk.NewAttribute(types.AttributeMsgAction, "delete"),
		sdk.NewAttribute(types.AttributeAssetUUID, msg.UUID),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgDeleteAssetResponse{}, nil
}
