package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/id/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IdByAddress(goCtx context.Context, req *types.QueryIdByAddressRequest) (*types.QueryIdByAddressResponse, error) {
	address, err := sdk.AccAddressFromBech32(req.Address)

	if req == nil || len(req.Address) == 0 || err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	id, found := k.GetFullIDByAddress(ctx, address)
	if !found {
		return nil, status.Error(codes.NotFound, "id not found")
	}

	return &types.QueryIdByAddressResponse{Id: id}, nil
}
