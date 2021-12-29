package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/document/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DocumentOfHolderByIssuer(goCtx context.Context, req *types.QueryDocumentOfHolderByIssuerRequest) (*types.QueryDocumentOfHolderByIssuerResponse, error) {
	if req == nil || len(req.Holder) == 0 || len(req.Issuer) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	docs := make([]*types.Document, 0)

	cb := func(doc types.Document) bool {
		docs = append(docs, &doc)
		return false
	}

	k.IterateAllDocOfHolderByIssuer(ctx, req.Holder, req.Issuer, cb)

	return &types.QueryDocumentOfHolderByIssuerResponse{Documents: docs}, nil
}
