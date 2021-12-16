package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IdByAddress(goCtx context.Context, req *types.QueryIdByAddressRequest) (*types.QueryIdByAddressResponse, error) {
	if req == nil || len(req.Address) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	address := sdk.AccAddress(req.Address)
	id, found := k.GetFullIDByAddress(ctx, address)
	if !found {
		return nil, status.Error(codes.NotFound, "id not found")
	}

	return &types.QueryIdByAddressResponse{Id: id}, nil
}
