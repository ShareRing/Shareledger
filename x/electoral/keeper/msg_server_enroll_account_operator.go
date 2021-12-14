package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnrollAccountOperator(goCtx context.Context, msg *types.MsgEnrollAccountOperator) (*types.MsgEnrollAccountOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEnrollAccountOperatorResponse{}, nil
}
