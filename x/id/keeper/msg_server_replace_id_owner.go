package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ReplaceIdOwner(goCtx context.Context, msg *types.MsgReplaceIdOwner) (*types.MsgReplaceIdOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgReplaceIdOwnerResponse{}, nil
}
