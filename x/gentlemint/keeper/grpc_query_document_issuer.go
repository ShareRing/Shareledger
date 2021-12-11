package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DocumentIssuer(goCtx context.Context, req *types.QueryDocumentIssuerRequest) (*types.QueryDocumentIssuerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	key := types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner)
	v, f := k.GetAccState(ctx, key)
	if !f {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryDocumentIssuerResponse{
		AccState: &types.AccState{
			Address: v.Address,
			Key:     v.Key,
			Status:  v.Status,
		},
	}, nil
}