package keeper_test

// TODO: update test case
//import (
//	"strconv"
//	"testing"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//	"github.com/stretchr/testify/require"
//
//	keepertest "github.com/sharering/shareledger/testutil/keeper"
//	"github.com/sharering/shareledger/x/gentlemint/keeper"
//	"github.com/sharering/shareledger/x/gentlemint/types"
//)
//
//// Prevent strconv unused error
//var _ = strconv.IntSize
//
//func TestLevelFeeMsgServerCreate(t *testing.T) {
//	k, ctx := keepertest.GentlemintKeeper(t)
//	srv := keeper.NewMsgServerImpl(*k)
//	wctx := sdk.WrapSDKContext(ctx)
//	creator := "A"
//	for i := 0; i < 5; i++ {
//		expected := &types.MsgCreateLevelFee{Creator: creator,
//			Level: strconv.Itoa(i),
//		}
//		_, err := srv.CreateLevelFee(wctx, expected)
//		require.NoError(t, err)
//		rst, found := k.GetLevelFee(ctx,
//			expected.Level,
//		)
//		require.True(t, found)
//		require.Equal(t, expected.Creator, rst.Creator)
//	}
//}
//
//func TestLevelFeeMsgServerUpdate(t *testing.T) {
//	creator := "A"
//
//	for _, tc := range []struct {
//		desc    string
//		request *types.MsgUpdateLevelFee
//		err     error
//	}{
//		{
//			desc: "Completed",
//			request: &types.MsgUpdateLevelFee{Creator: creator,
//				Level: strconv.Itoa(0),
//			},
//		},
//		{
//			desc: "Unauthorized",
//			request: &types.MsgUpdateLevelFee{Creator: "B",
//				Level: strconv.Itoa(0),
//			},
//			err: sdkerrors.ErrUnauthorized,
//		},
//		{
//			desc: "KeyNotFound",
//			request: &types.MsgUpdateLevelFee{Creator: creator,
//				Level: strconv.Itoa(100000),
//			},
//			err: sdkerrors.ErrKeyNotFound,
//		},
//	} {
//		t.Run(tc.desc, func(t *testing.T) {
//			k, ctx := keepertest.GentlemintKeeper(t)
//			srv := keeper.NewMsgServerImpl(*k)
//			wctx := sdk.WrapSDKContext(ctx)
//			expected := &types.MsgCreateLevelFee{Creator: creator,
//				Level: strconv.Itoa(0),
//			}
//			_, err := srv.CreateLevelFee(wctx, expected)
//			require.NoError(t, err)
//
//			_, err = srv.UpdateLevelFee(wctx, tc.request)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//				rst, found := k.GetLevelFee(ctx,
//					expected.Level,
//				)
//				require.True(t, found)
//				require.Equal(t, expected.Creator, rst.Creator)
//			}
//		})
//	}
//}
//
//func TestLevelFeeMsgServerDelete(t *testing.T) {
//	creator := "A"
//
//	for _, tc := range []struct {
//		desc    string
//		request *types.MsgDeleteLevelFee
//		err     error
//	}{
//		{
//			desc: "Completed",
//			request: &types.MsgDeleteLevelFee{Creator: creator,
//				Level: strconv.Itoa(0),
//			},
//		},
//		{
//			desc: "Unauthorized",
//			request: &types.MsgDeleteLevelFee{Creator: "B",
//				Level: strconv.Itoa(0),
//			},
//			err: sdkerrors.ErrUnauthorized,
//		},
//		{
//			desc: "KeyNotFound",
//			request: &types.MsgDeleteLevelFee{Creator: creator,
//				Level: strconv.Itoa(100000),
//			},
//			err: sdkerrors.ErrKeyNotFound,
//		},
//	} {
//		t.Run(tc.desc, func(t *testing.T) {
//			k, ctx := keepertest.GentlemintKeeper(t)
//			srv := keeper.NewMsgServerImpl(*k)
//			wctx := sdk.WrapSDKContext(ctx)
//
//			_, err := srv.CreateLevelFee(wctx, &types.MsgCreateLevelFee{Creator: creator,
//				Level: strconv.Itoa(0),
//			})
//			require.NoError(t, err)
//			_, err = srv.DeleteLevelFee(wctx, tc.request)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//				_, found := k.GetLevelFee(ctx,
//					tc.request.Level,
//				)
//				require.False(t, found)
//			}
//		})
//	}
//}
