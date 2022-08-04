package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RequestedIns(goCtx context.Context, req *types.QueryRequestedInsRequest) (*types.QueryRequestedInsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	r, found := k.GetRequestedIn(ctx, req.TxHash, req.LogEventIndex)
	if !found {
		return &types.QueryRequestedInsResponse{}, nil
	}
	return &types.QueryRequestedInsResponse{
		RequestedIn: &r,
	}, nil
}
