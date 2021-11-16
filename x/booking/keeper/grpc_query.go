package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/booking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Booking gets booking by id
func (k Querier) Booking(ctx context.Context, req *types.QueryBookingRequest) (*types.QueryBookingRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	booking := k.GetBooking(sdkCtx, req.BookID)

	return &types.QueryBookingRequestResponse{Booking: &booking}, nil
}
