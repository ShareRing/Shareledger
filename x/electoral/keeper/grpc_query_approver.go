package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/electoral/types"
)

func (k Keeper) Approver(goCtx context.Context, req *types.QueryApproverRequest) (*types.QueryApproverResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	key := types.GenAccStateIndexKey(addr, types.AccStateKeyApprover)
	m, f := k.GetAccState(ctx, key)
	if !f {
		return nil, status.Errorf(codes.InvalidArgument, "accState not found %s", addr)
	}

	return &types.QueryApproverResponse{AccState: m}, nil
}
