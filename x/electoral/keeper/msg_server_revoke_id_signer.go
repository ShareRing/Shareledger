package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevokeIdSigner(goCtx context.Context, msg *types.MsgRevokeIdSigner) (*types.MsgRevokeIdSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevokeIdSignerResponse{}, nil
}
