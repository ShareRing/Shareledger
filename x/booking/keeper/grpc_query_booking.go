package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/booking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Booking(goCtx context.Context, req *types.QueryBookingRequest) (*types.QueryBookingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	booking, found := k.GetBooking(ctx, req.BookID)
	if !found {
		return nil, status.Error(codes.NotFound, "Booking not found")
	}

	return &types.QueryBookingResponse{Booking: &booking}, nil
}
