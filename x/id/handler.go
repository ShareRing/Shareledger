package id

import (
	"fmt"

	"github.com/ShareRing/Shareledger/x/id/keeper"
	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateId:
			res, err := msgServer.CreateId(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateId:
			res, err := msgServer.UpdateId(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateIdBatch:
			res, err := msgServer.CreateIdInBatch(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgReplaceIdOwner:
			res, err := msgServer.ReplaceIdOwner(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			fmt.Printf("Unrecognized Id Msg type: %v", msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized identity Msg type: %v", msg))
		}
	}
	// return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
	// 	ctx = ctx.WithEventManager(sdk.NewEventManager())

	// 	switch msg := msg.(type) {
	// 	// this line is used by starport scaffolding # 1
	// 	default:
	// 		errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
	// 		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	// 	}
	// }
}