package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Balances(goCtx context.Context, req *types.QueryBalancesRequest) (*types.QueryBalancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.GetAddress())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %v", err)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	currentBalances := k.bankKeeper.GetAllBalances(ctx, addr)
	displayCoins := denom.ToDisplayCoins(currentBalances)
	r := &types.QueryBalancesResponse{
		Coins: make([]*sdk.DecCoin, 0, displayCoins.Len()),
	}
	for _, c := range displayCoins {
		r.Coins = append(r.Coins, &sdk.DecCoin{
			Denom:  c.Denom,
			Amount: c.Amount,
		})
	}

	return r, nil
}
