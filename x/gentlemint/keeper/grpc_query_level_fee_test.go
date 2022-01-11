package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestLevelFeeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLevelFee(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryLevelFeeRequest
		response *types.QueryLevelFeeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryLevelFeeRequest{
				Level: msgs[0].Level,
			},
			response: &types.QueryLevelFeeResponse{LevelFee: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryLevelFeeRequest{
				Level: msgs[1].Level,
			},
			response: &types.QueryLevelFeeResponse{LevelFee: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryLevelFeeRequest{
				Level: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
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

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryLevelFeesRequest {
		return &types.QueryLevelFeesRequest{}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LevelFees(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LevelFee), step)
			require.Subset(t, msgs, resp.LevelFee)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2

		resp, err := keeper.LevelFees(wctx, &types.QueryLevelFeesRequest{})
		require.NoError(t, err)
		require.LessOrEqual(t, len(resp.LevelFee), step)
		require.Subset(t, msgs, resp.LevelFee)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LevelFees(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
