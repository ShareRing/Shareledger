package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// AssetByUUID gets asset data by uuid
func (k Querier) AssetByUUID(ctx context.Context, req *types.QueryAssetByUUIDRequest) (*types.QueryAssetByUUIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	asset := k.GetAsset(sdkCtx, req.Uuid)

	return &types.QueryAssetByUUIDResponse{Asset: &asset}, nil
}
