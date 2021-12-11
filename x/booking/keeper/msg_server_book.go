package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/booking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Book(goCtx context.Context, msg *types.MsgBook) (*types.MsgBookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBookResponse{}, nil
}
