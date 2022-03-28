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
	"github.com/sharering/shareledger/x/electoral/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAccStateQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAccState(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryAccStateRequest
		response *types.QueryAccStateResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryAccStateRequest{
				Key: msgs[0].Key,
			},
			response: &types.QueryAccStateResponse{AccState: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryAccStateRequest{
				Key: msgs[1].Key,
			},
			response: &types.QueryAccStateResponse{AccState: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryAccStateRequest{
				Key: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.AccState(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestAccStateQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAccState(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAccStatesRequest {
		return &types.QueryAccStatesRequest{
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
			resp, err := keeper.AccStates(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AccState), step)
			require.Subset(t, msgs, resp.AccState)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AccStates(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AccState), step)
			require.Subset(t, msgs, resp.AccState)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AccStates(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AccStates(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
