package keeper_test

import (
    "strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/gentlemint/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestLevelFeeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLevelFee(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLevelFeeRequest
		response *types.QueryGetLevelFeeResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLevelFeeRequest{
			    Level: msgs[0].Level,
                
			},
			response: &types.QueryGetLevelFeeResponse{LevelFee: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetLevelFeeRequest{
			    Level: msgs[1].Level,
                
			},
			response: &types.QueryGetLevelFeeResponse{LevelFee: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetLevelFeeRequest{
			    Level:strconv.Itoa(100000),
                
			},
			err:     status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LevelFee(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestLevelFeeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLevelFee(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLevelFeeRequest {
		return &types.QueryAllLevelFeeRequest{
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
			resp, err := keeper.LevelFeeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LevelFee), step)
			require.Subset(t, msgs, resp.LevelFee)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LevelFeeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LevelFee), step)
			require.Subset(t, msgs, resp.LevelFee)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LevelFeeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LevelFeeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
