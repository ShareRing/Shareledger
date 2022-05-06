package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) RevokeRelayers(goCtx context.Context, msg *types.MsgRevokeRelayers) (*types.MsgRevokeRelayersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevokeRelayersResponse{}, nil
}
