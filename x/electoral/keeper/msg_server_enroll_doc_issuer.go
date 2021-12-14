package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnrollDocIssuer(goCtx context.Context, msg *types.MsgEnrollDocIssuer) (*types.MsgEnrollDocIssuerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEnrollDocIssuerResponse{}, nil
}
