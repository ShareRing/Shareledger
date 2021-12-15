package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateId(goCtx context.Context, msg *types.MsgCreateId) (*types.MsgCreateIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateIdResponse{}, nil
}
