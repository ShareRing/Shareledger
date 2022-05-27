package keeper

import (
	"context"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Balance(goCtx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAddr.Empty() {
		return nil, status.Error(codes.InvalidArgument, "can't get the module address")
	}
	coins := k.bankKeeper.SpendableCoins(ctx, moduleAddr)
	dCoin := denom.ToDisplayCoins(coins)

	balance := sdk.DecCoin{}
	for _, c := range dCoin {
		if c.Denom == denom.Shr {
			balance = c
			break
		}
	}

	return &types.QueryBalanceResponse{Balance: &balance}, nil
}
