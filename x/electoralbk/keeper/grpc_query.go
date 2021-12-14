package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoralbk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) Voter(ctx context.Context, req *types.QueryVoterRequest) (*types.QueryVoterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	voterID := types.VoterPrefix + req.Address
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	voter := k.GetVoter(sdkCtx, voterID)

	return &types.QueryVoterResponse{Voter: &voter}, nil
}
