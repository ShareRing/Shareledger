package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetExchange(goCtx context.Context, msg *types.MsgSetExchange) (*types.MsgSetExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetExchangeResponse{}, nil
}
