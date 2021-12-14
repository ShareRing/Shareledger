package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnrollIdSigner(goCtx context.Context, msg *types.MsgEnrollIdSigner) (*types.MsgEnrollIdSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEnrollIdSignerResponse{}, nil
}
