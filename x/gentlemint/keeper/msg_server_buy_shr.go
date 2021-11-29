package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BuyShr(goCtx context.Context, msg *types.MsgBuyShr) (*types.MsgBuyShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBuyShrResponse{}, nil
}

func (k msgServer) buyShr(ctx sdk.Context, amount sdk.Int, buyer sdk.AccAddress) error {
	panic("implement me")
}
