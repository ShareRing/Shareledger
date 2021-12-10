package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Delete(goCtx context.Context, msg *types.MsgDelete) (*types.MsgDeleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteResponse{}, nil
}
