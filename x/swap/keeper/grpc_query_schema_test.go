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
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/swap/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestFormatQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNFormat(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetFormatRequest
		response *types.QueryGetFormatResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetFormatRequest{
				Network: msgs[0].Network,
			},
			response: &types.QueryGetFormatResponse{Format: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetFormatRequest{
				Network: msgs[1].Network,
			},
			response: &types.QueryGetFormatResponse{Format: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetFormatRequest{
				Network: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Format(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestFormatQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNFormat(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllFormatRequest {
		return &types.QueryAllFormatRequest{
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
			resp, err := keeper.FormatAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Format), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Format),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FormatAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Format), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Format),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.FormatAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Format),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.FormatAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
