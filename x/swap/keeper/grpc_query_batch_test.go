package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/swap/types"
)

func TestBatchQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBatch(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryBatchRequest
		response *types.QueryBatchResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetBatchRequest{Id: msgs[0].Id},
			response: &types.QueryGetBatchResponse{Batch: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetBatchRequest{Id: msgs[1].Id},
			response: &types.QueryGetBatchResponse{Batch: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetBatchRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Batch(wctx, tc.request)
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
