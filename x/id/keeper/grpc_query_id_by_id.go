package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/id/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IdById(goCtx context.Context, req *types.QueryIdByIdRequest) (*types.QueryIdByIdResponse, error) {
	if req == nil || len(req.Id) == 0 || len(req.Id) > types.MAX_ID_LEN {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	id, found := k.GetFullIDByIDString(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "id not found")
	}

	return &types.QueryIdByIdResponse{Id: id}, nil
}
