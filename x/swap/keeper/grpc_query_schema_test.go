package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/swap/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestFormatQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSchema(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QuerySchemaRequest
		response *types.QuerySchemaResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QuerySchemaRequest{
				Network: msgs[0].Network,
			},
			response: &types.QuerySchemaResponse{Schema: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QuerySchemaRequest{
				Network: msgs[1].Network,
			},
			response: &types.QuerySchemaResponse{Schema: msgs[1]},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Schema(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, tc.err, err)
				// require.True(t, strings.Contains(tc.err.Error(), err.Error()))
				return
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
	msgs := createNSchema(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QuerySchemasRequest {
		return &types.QuerySchemasRequest{
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
			resp, err := keeper.Schemas(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Schemas), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Schemas),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.Schemas(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Schemas), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Schemas),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.Schemas(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Schemas),
		)
	})
}
