package keeper

import (
	"context"
	"strconv"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetExchange(goCtx context.Context, msg *types.MsgSetExchange) (*types.MsgSetExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	v, err := strconv.ParseFloat(msg.Rate, 64)
	if err != nil {
		return nil, err
	}
	k.SetExchangeRate(ctx, types.ExchangeRate{Rate: v})
	return &types.MsgSetExchangeResponse{}, nil
}
