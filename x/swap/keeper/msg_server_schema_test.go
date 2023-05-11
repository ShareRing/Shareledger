package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestFormatMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.SwapKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateSchema{Creator: creator,
			Network: strconv.Itoa(i),
			In:      sdk.NewDecCoin(denom.Base, sdk.NewInt(100)),
			Out:     sdk.NewDecCoin(denom.Base, sdk.NewInt(100)),
		}
		_, err := srv.CreateSchema(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetSchema(ctx,
			expected.Network,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestFormatMsgServerUpdate(t *testing.T) {
	creator := "A"
	in := sdk.NewDecCoin(denom.Base, sdk.NewInt(100))
	out := sdk.NewDecCoin(denom.Base, sdk.NewInt(100))
	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateSchema
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateSchema{Creator: creator,
				In:      &in,
				Out:     &out,
				Network: strconv.Itoa(0),
			},
		},

		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateSchema{Creator: creator,
				In:      &in,
				Out:     &out,
				Network: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SwapKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateSchema{Creator: creator,
				Network: strconv.Itoa(0),
				In:      in,
				Out:     out,
			}
			_, err := srv.CreateSchema(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateSchema(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetSchema(ctx,
					expected.Network,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestFormatMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteSchema
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteSchema{Creator: creator,
				Network: strconv.Itoa(0),
			},
		},

		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteSchema{Creator: creator,
				Network: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SwapKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateSchema(wctx, &types.MsgCreateSchema{Creator: creator,
				Network: strconv.Itoa(0),
				In:      sdk.NewDecCoin(denom.Base, sdk.NewInt(100)),
				Out:     sdk.NewDecCoin(denom.Base, sdk.NewInt(100)),
			})
			require.NoError(t, err)
			_, err = srv.DeleteSchema(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetSchema(ctx,
					tc.request.Network,
				)
				require.False(t, found)
			}
		})
	}
}
