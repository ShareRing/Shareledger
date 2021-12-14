package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/ShareRing/Shareledger/testutil/keeper"
	"github.com/ShareRing/Shareledger/x/electoral/keeper"
	"github.com/ShareRing/Shareledger/x/electoral/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAccStateMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ElectoralKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateAccState{Creator: creator,
			Key: strconv.Itoa(i),
		}
		_, err := srv.CreateAccState(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetAccState(ctx,
			expected.Key,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestAccStateMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateAccState
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateAccState{Creator: creator,
				Key: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateAccState{Creator: "B",
				Key: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateAccState{Creator: creator,
				Key: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ElectoralKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateAccState{Creator: creator,
				Key: strconv.Itoa(0),
			}
			_, err := srv.CreateAccState(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateAccState(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetAccState(ctx,
					expected.Key,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestAccStateMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteAccState
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteAccState{Creator: creator,
				Key: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteAccState{Creator: "B",
				Key: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteAccState{Creator: creator,
				Key: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ElectoralKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateAccState(wctx, &types.MsgCreateAccState{Creator: creator,
				Key: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteAccState(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetAccState(ctx,
					tc.request.Key,
				)
				require.False(t, found)
			}
		})
	}
}
