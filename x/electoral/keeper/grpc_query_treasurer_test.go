package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ShareRing/Shareledger/testutil/keeper"
	"github.com/ShareRing/Shareledger/x/electoral/types"
)

func TestTreasurerQuery(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTreasurer(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTreasurerRequest
		response *types.QueryGetTreasurerResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTreasurerRequest{},
			response: &types.QueryGetTreasurerResponse{Treasurer: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Treasurer(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
