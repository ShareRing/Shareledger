package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/electoral/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccStates(c context.Context, req *types.QueryAccStatesRequest) (*types.QueryAccStatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accStates []types.AccState
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accStateStore := prefix.NewStore(store, types.KeyPrefix(types.AccStateKeyPrefix))

	pageRes, err := query.Paginate(accStateStore, req.Pagination, func(key []byte, value []byte) error {
		var accState types.AccState
		if err := k.cdc.Unmarshal(value, &accState); err != nil {
			return err
		}

		accStates = append(accStates, accState)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAccStatesResponse{AccState: accStates, Pagination: pageRes}, nil
}

func (k Keeper) AccState(c context.Context, req *types.QueryAccStateRequest) (*types.QueryAccStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAccState(
		ctx,
		types.IndexKeyAccState(req.Key),
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryAccStateResponse{AccState: val}, nil
}
