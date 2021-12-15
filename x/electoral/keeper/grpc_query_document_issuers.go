package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DocumentIssuers(goCtx context.Context, req *types.QueryDocumentIssuersRequest) (*types.QueryDocumentIssuersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	list := k.IterateAccState(ctx, types.AccStateKeyDocIssuer)
	res := make([]*types.AccState, 0, len(list))
	for _, i := range list {
		res = append(res, &types.AccState{
			Key:     i.Key,
			Address: i.Address,
			Status:  i.Status,
		})
	}

	return &types.QueryDocumentIssuersResponse{
		AccStates: res,
	}, nil
}
