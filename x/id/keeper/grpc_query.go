package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// IdByAddress gets id data by address
func (k Querier) IdByAddress(ctx context.Context, req *types.QueryIdByAddressRequest) (*types.QueryIdByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	id := k.GetIdByAddress(sdkCtx, address)

	return &types.QueryIdByAddressResponse{Id: id}, nil
}

// IdById gets id data by id
func (k Querier) IdById(ctx context.Context, req *types.QueryIdByIdRequest) (*types.QueryIdByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if len(req.Id) == 0 || len(req.Id) > types.MAX_ID_LEN {
		return nil, status.Error(codes.InvalidArgument, "empty request data")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	id := k.GetIDByIdString(sdkCtx, req.Id)

	return &types.QueryIdByAddressResponse{Id: id}, nil
}
