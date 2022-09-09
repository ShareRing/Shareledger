package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) EnrollApprovers(goCtx context.Context, msg *types.MsgEnrollApprovers) (*types.MsgEnrollApproversResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.EventManager().EmitEvent(sdk.NewEvent("approve").AppendAttributes(sdk.NewAttribute("address", msg.String())))
	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		k.activeApprover(ctx, addr)
	}

	return &types.MsgEnrollApproversResponse{}, nil
}
