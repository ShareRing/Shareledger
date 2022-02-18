package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Voter(goCtx context.Context, req *types.QueryVoterRequest) (result *types.QueryVoterResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	result = &types.QueryVoterResponse{
		Voter: types.AccState{
			Status: string(types.StatusInactive),
		},
	}

	voter, found := k.GetAccStateByAddress(ctx, addr, types.AccStateKeyVoter)
	if !found {
		return
	}
	result = &types.QueryVoterResponse{
		Voter: voter,
	}

	return
}
