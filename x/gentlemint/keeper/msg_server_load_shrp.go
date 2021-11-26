package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LoadShrp(goCtx context.Context, msg *types.MsgLoadShrp) (*types.MsgLoadShrpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: double check this, since in the old logic, it does not cover this auth:  only enrolled shrp-loaders (approver) can submit this transaction

	shrp, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return
	}

	return &types.MsgLoadShrpResponse{}, nil
}
