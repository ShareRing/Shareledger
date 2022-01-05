package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/asset/types"
)

func (k msgServer) CreateAsset(goCtx context.Context, msg *types.MsgCreateAsset) (*types.MsgCreateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_, found := k.GetAsset(ctx, msg.GetUUID())
	if found {
		return nil, sdkerrors.Wrap(types.ErrAssetExist, msg.GetUUID())
	}

	asset := types.NewAssetFromMsgCreate(*msg)
	k.SetAsset(ctx, msg.GetUUID(), asset)

	event := sdk.NewEvent(
		types.EventTypeCreateAsset,
		sdk.NewAttribute(types.AttributeMsgModule, "asset"),
		sdk.NewAttribute(types.AttributeMsgAction, "create"),
		sdk.NewAttribute(types.AttributeAssetCreator, msg.GetCreator()),
		sdk.NewAttribute(types.AttributeAssetUUID, msg.GetUUID()),
		sdk.NewAttribute(types.AttributeAssetHash, fmt.Sprintf("%X", msg.GetHash())),
		sdk.NewAttribute(types.AttributeAssetStatus, strconv.FormatBool(msg.GetStatus())),
		sdk.NewAttribute(types.AttributeAssetFee, strconv.Itoa(int(msg.GetRate()))),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgCreateAssetResponse{}, nil
}
