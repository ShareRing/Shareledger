package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/document/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DocumentByHolderId(goCtx context.Context, req *types.QueryDocumentByHolderIdRequest) (*types.QueryDocumentByHolderIdResponse, error) {
	if req == nil || len(req.Id) == 0 || len(req.Id) > types.MAX_LEN {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	docs := make([]*types.Document, 0)

	cb := func(doc types.Document) bool {
		docs = append(docs, &doc)
		return false
	}

	k.IterateAllDocsOfAHolder(ctx, req.Id, cb)

	return &types.QueryDocumentByHolderIdResponse{Documents: docs}, nil
}
