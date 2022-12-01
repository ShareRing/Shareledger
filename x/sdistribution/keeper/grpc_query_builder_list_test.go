package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/sdistribution/types"
)

func TestBuilderListQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SdistributionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBuilderList(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBuilderListRequest
		response *types.QueryGetBuilderListResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetBuilderListRequest{Id: msgs[0].Id},
			response: &types.QueryGetBuilderListResponse{BuilderList: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetBuilderListRequest{Id: msgs[1].Id},
			response: &types.QueryGetBuilderListResponse{BuilderList: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetBuilderListRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.BuilderList(wctx, tc.request)
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

func TestBuilderListQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SdistributionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBuilderList(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllBuilderListRequest {
		return &types.QueryAllBuilderListRequest{
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
			resp, err := keeper.BuilderListAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BuilderList), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.BuilderList),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.BuilderListAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BuilderList), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.BuilderList),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.BuilderListAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.BuilderList),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.BuilderListAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
