package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnrollLoaders(goCtx context.Context, msg *types.MsgEnrollLoaders) (*types.MsgEnrollLoadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEnrollLoadersResponse{}, nil
}
