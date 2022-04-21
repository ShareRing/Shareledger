package keeper

import (
	"context"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Funding(goCtx context.Context, req *types.QueryFundingRequest) (*types.QueryFundingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleAddr := k.ak.GetModuleAddress(types.ModuleName)
	if moduleAddr.Empty() {
		return nil, status.Error(codes.InvalidArgument, "can't get the module address")
	}
	coins := k.bankKeeper.SpendableCoins(ctx, moduleAddr)
	dCoin := denom.ToDisplayCoins(coins)

	r := &types.QueryFundingResponse{
		Availiable: make([]*sdk.DecCoin, 0, dCoin.Len()),
	}
	for _, c := range dCoin {
		r.Availiable = append(r.Availiable, &sdk.DecCoin{
			Denom:  c.Denom,
			Amount: c.Amount,
		})
	}

	return r, nil
}
