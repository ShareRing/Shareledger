package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActionLevelFeeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActionLevelFee(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActionLevelFeeRequest
		response *types.QueryGetActionLevelFeeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActionLevelFeeRequest{
				Action: msgs[0].Action,
			},
			response: &types.QueryGetActionLevelFeeResponse{ActionLevelFee: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActionLevelFeeRequest{
				Action: msgs[1].Action,
			},
			response: &types.QueryGetActionLevelFeeResponse{ActionLevelFee: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActionLevelFeeRequest{
				Action: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ActionLevelFee(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestActionLevelFeeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActionLevelFee(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActionLevelFeeRequest {
		return &types.QueryAllActionLevelFeeRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActionLevelFeeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionLevelFee), step)
			require.Subset(t, msgs, resp.ActionLevelFee)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActionLevelFeeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionLevelFee), step)
			require.Subset(t, msgs, resp.ActionLevelFee)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActionLevelFeeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActionLevelFeeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
