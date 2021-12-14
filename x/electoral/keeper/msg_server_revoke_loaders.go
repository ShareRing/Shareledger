package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevokeLoaders(goCtx context.Context, msg *types.MsgRevokeLoaders) (*types.MsgRevokeLoadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevokeLoadersResponse{}, nil
}
