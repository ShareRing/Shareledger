package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateId(goCtx context.Context, msg *types.MsgUpdateId) (*types.MsgUpdateIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateIdResponse{}, nil
}
