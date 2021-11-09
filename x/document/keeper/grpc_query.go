package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Gets document by proof
func (k Querier) DocumentByProof(ctx context.Context, req *types.QueryDocumentByProofRequest) (*types.QueryDocumentByProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	queryDoc := types.Document{Proof: req.Proof}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	doc := k.GetDocByProof(sdkCtx, queryDoc)

	return &types.QueryDocumentByProofResponse{Document: &doc}, nil
}

// Gets document by holder
func (k Querier) DocumentByHolderId(ctx context.Context, req *types.QueryDocumentByHolderIdRequest) (*types.QueryDocumentByHolderIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if len(req.Id) == 0 || len(req.Id) > types.MAX_LEN {
		return nil, status.Error(codes.InvalidArgument, "invalid request data")
	}

	docs := make([]*types.Document, 0)

	cb := func(doc types.Document) bool {
		docs = append(docs, &doc)
		return false
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.IterateAllDocsOfAHolder(sdkCtx, req.Id, cb)

	return &types.QueryDocumentByHolderIdResponse{Document: docs}, nil
}
