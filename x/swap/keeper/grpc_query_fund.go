package keeper

import (
	"context"

	denom "github.com/sharering/shareledger/x/utils/denom"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// func (k Keeper) Balance(goCtx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "invalid request")
// 	}

// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
// 	if moduleAddr.Empty() {
// 		return nil, status.Error(codes.InvalidArgument, "can't get the module address")
// 	}
// 	coins := k.bankKeeper.SpendableCoins(ctx, moduleAddr)

// 	var b sdk.Coin
// 	for _, c := range coins {
// 		if c.Denom == denom.Base {
// 			b = c
// 			break
// 		}
// 	}

// 	return &types.QueryBalanceResponse{Balance: &b}, nil
// }

func (k Keeper) Balance(goCtx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAddr.Empty() {
		return nil, status.Error(codes.InvalidArgument, "can't get the module address")
	}
	b := k.bankKeeper.GetBalance(ctx, moduleAddr, denom.Base)

	return &types.QueryBalanceResponse{Balance: &b}, nil
}
