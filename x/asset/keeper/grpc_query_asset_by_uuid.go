package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AssetByUUID(goCtx context.Context, req *types.QueryAssetByUUIDRequest) (*types.QueryAssetByUUIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	asset, found := k.GetAsset(ctx, req.Uuid)
	if !found {
		return nil, status.Error(codes.NotFound, "Asset not found")
	}

	return &types.QueryAssetByUUIDResponse{Asset: &asset}, nil
}
