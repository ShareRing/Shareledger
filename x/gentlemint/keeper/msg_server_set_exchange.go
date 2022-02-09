package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) SetExchange(goCtx context.Context, msg *types.MsgSetExchange) (*types.MsgSetExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	k.SetExchangeRate(ctx, types.ExchangeRate{ShrpToShr: msg.Rate})
	return &types.MsgSetExchangeResponse{}, nil
}
