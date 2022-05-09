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
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestFormatMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.SwapKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFormat{Creator: creator,
			Network: strconv.Itoa(i),
		}
		_, err := srv.CreateFormat(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetFormat(ctx,
			expected.Network,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestFormatMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateFormat
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateFormat{Creator: creator,
				Network: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateFormat{Creator: "B",
				Network: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateFormat{Creator: creator,
				Network: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SwapKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateFormat{Creator: creator,
				Network: strconv.Itoa(0),
			}
			_, err := srv.CreateFormat(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateFormat(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetFormat(ctx,
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
		request *types.MsgDeleteFormat
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteFormat{Creator: creator,
				Network: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteFormat{Creator: "B",
				Network: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteFormat{Creator: creator,
				Network: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SwapKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateFormat(wctx, &types.MsgCreateFormat{Creator: creator,
				Network: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteFormat(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetFormat(ctx,
					tc.request.Network,
				)
				require.False(t, found)
			}
		})
	}
}
