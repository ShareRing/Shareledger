package keeper

import (
	"context"
	"github.com/sharering/shareledger/x/constant"
	"github.com/sharering/shareledger/x/fee"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActionLevelFees(c context.Context, req *types.QueryActionLevelFeesRequest) (*types.QueryActionLevelFeesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	actionFees := k.GetAllActionLevelFee(ctx)
	actionsDefaultLevels := fee.GetListActionsWithDefaultLevel()

	actionLevelFees := make([]types.ActionLevelFee, 0, len(actionsDefaultLevels))
	for _, a := range actionFees {
		actionLevelFees = append(actionLevelFees, types.ActionLevelFee{
			Action:  a.Action,
			Level:   a.Level,
			Creator: a.Creator,
		})
		delete(actionsDefaultLevels, a.Action)
	}
	for a, l := range actionsDefaultLevels {
		actionLevelFees = append(actionLevelFees, types.ActionLevelFee{
			Action: a,
			Level:  l,
		})
	}
	sort.Slice(actionLevelFees, func(i, j int) bool {
		return actionLevelFees[i].Action < actionLevelFees[j].Action
	})

	return &types.QueryActionLevelFeesResponse{ActionLevelFee: actionLevelFees}, nil
}

func (k Keeper) ActionLevelFee(c context.Context, req *types.QueryActionLevelFeeRequest) (*types.QueryActionLevelFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	defaultLevel := string(constant.MinFee)

	val, found := k.GetActionLevelFee(
		ctx,
		req.Action,
	)
	if found {
		defaultLevel = val.Level
	}

	return &types.QueryActionLevelFeeResponse{
		Action: req.Action,
		Level:  defaultLevel,
		Fee:    k.GetFeeByLevel(ctx, defaultLevel).String(),
	}, nil
}
