package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DocumentByHolderId(goCtx context.Context, req *types.QueryDocumentByHolderIdRequest) (*types.QueryDocumentByHolderIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryDocumentByHolderIdResponse{}, nil
}