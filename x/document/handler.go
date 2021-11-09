package document

import (
	"fmt"

	"github.com/ShareRing/Shareledger/x/document/keeper"
	"github.com/ShareRing/Shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	// return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
	// 	ctx = ctx.WithEventManager(sdk.NewEventManager())

	// 	switch msg := msg.(type) {
	// 	// this line is used by starport scaffolding # 1
	// 	default:
	// 		errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
	// 		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	// 	}
	// }

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case *types.MsgCreateDoc:
			res, err := msgServer.CreateDocument(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateDocBatch:
			res, err := msgServer.CreateDocumentInBatch(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateDoc:
			res, err := msgServer.UpdateDocument(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRevokeDoc:
			res, err := msgServer.RevokeDocument(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized %s Msg type: %v", types.ModuleName, msg.String()))
		}
	}
}
