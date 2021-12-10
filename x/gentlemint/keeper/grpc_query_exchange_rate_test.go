package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ShareRing/Shareledger/testutil/keeper"
	"github.com/ShareRing/Shareledger/x/gentlemint/types"
)

func TestExchangeRateQuery(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestExchangeRate(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetExchangeRateRequest
		response *types.QueryGetExchangeRateResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetExchangeRateRequest{},
			response: &types.QueryGetExchangeRateResponse{Rate: item.Rate},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ExchangeRate(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
