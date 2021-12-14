package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevokeDocIssuer(goCtx context.Context, msg *types.MsgRevokeDocIssuer) (*types.MsgRevokeDocIssuerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevokeDocIssuerResponse{}, nil
}
