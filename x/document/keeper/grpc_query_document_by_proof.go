package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/document/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DocumentByProof(goCtx context.Context, req *types.QueryDocumentByProofRequest) (*types.QueryDocumentByProofResponse, error) {
	if req == nil || len(req.Proof) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	queryDoc := types.Document{Proof: req.Proof}
	doc, found := k.GetDocByProof(ctx, queryDoc)
	if !found {
		return nil, status.Error(codes.NotFound, "document not found")
	}

	return &types.QueryDocumentByProofResponse{Document: &doc}, nil
}
