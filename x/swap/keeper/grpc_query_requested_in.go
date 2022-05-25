package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/swap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RequestedInAll(c context.Context, req *types.QueryAllRequestedInRequest) (*types.QueryAllRequestedInResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var requestedIns []types.RequestedIn
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	requestedInStore := prefix.NewStore(store, types.KeyPrefix(types.RequestedInKeyPrefix))

	pageRes, err := query.Paginate(requestedInStore, req.Pagination, func(key []byte, value []byte) error {
		var requestedIn types.RequestedIn
		if err := k.cdc.Unmarshal(value, &requestedIn); err != nil {
			return err
		}

		requestedIns = append(requestedIns, requestedIn)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRequestedInResponse{RequestedIn: requestedIns, Pagination: pageRes}, nil
}

func (k Keeper) RequestedIn(c context.Context, req *types.QueryGetRequestedInRequest) (*types.QueryGetRequestedInResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRequestedIn(
	    ctx,
	    req.Address,
        )
	if !found {
	    return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRequestedInResponse{RequestedIn: val}, nil
}