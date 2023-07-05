package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/swap/types"
)

func (k Keeper) PastTxEvent(goCtx context.Context, req *types.QueryPastTxEventRequest) (*types.QueryPastTxEventResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	r, found := k.GetPastTxEvent(ctx, req.TxHash, req.LogIndex)
	if !found {
		return &types.QueryPastTxEventResponse{}, nil
	}
	return &types.QueryPastTxEventResponse{
		Event: &r,
	}, nil
}

func (k Keeper) PastTxEventsByTxHash(goCtx context.Context, req *types.QueryPastTxEventsByTxHashRequest) (*types.QueryPastTxEventsByTxHashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	r := k.GetPastTxEventsByTxHash(ctx, req.TxHash)

	return &types.QueryPastTxEventsByTxHashResponse{
		Events: r,
	}, nil
}

func (k Keeper) PastTxEvents(goCtx context.Context, req *types.QueryPastTxEventsRequest) (*types.QueryPastTxEventsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var events []*types.PastTxEvent
	store := ctx.KVStore(k.storeKey)
	pastTxEventStore := prefix.NewStore(store, []byte(types.PastTxEventsKeyPrefix))

	pageRes, err := query.Paginate(pastTxEventStore, req.Pagination, func(key []byte, value []byte) error {
		var event types.PastTxEvent

		if err := k.cdc.Unmarshal(value, &event); err != nil {
			return err
		}

		events = append(events, &event)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPastTxEventsResponse{
		Events:     events,
		Pagination: pageRes,
	}, nil
}
