package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevokeAccountOperator(goCtx context.Context, msg *types.MsgRevokeAccountOperator) (*types.MsgRevokeAccountOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevokeAccountOperatorResponse{}, nil
}
