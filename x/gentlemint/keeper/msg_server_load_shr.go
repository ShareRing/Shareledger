package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LoadShr(goCtx context.Context, msg *types.MsgLoadShr) (*types.MsgLoadShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Check Is Authority

	return &types.MsgLoadShrResponse{}, nil
}
