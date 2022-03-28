package keeper_test

// TODO: update test cases
//import (
//	"github.com/sharering/shareledger/testutil/sample"
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
//func TestActionLevelFeeMsgServerCreate(t *testing.T) {
//	k, ctx := keepertest.GentlemintKeeper(t)
//	srv := keeper.NewMsgServerImpl(*k)
//	wctx := sdk.WrapSDKContext(ctx)
//	creator := sample.AccAddress()
//	for i := 0; i < 5; i++ {
//		expected := &types.MsgCreateActionLevelFee{Creator: creator,
//			Action: strconv.Itoa(i),
//		}
//		_, err := srv.CreateActionLevelFee(wctx, expected)
//		require.NoError(t, err, err.Error())
//		rst, found := k.GetActionLevelFee(ctx,
//			expected.Action,
//		)
//		require.True(t, found)
//		require.Equal(t, expected.Creator, rst.Creator)
//	}
//}
//
//func TestActionLevelFeeMsgServerUpdate(t *testing.T) {
//	creator := "A"
//
//	for _, tc := range []struct {
//		desc    string
//		request *types.MsgUpdateActionLevelFee
//		err     error
//	}{
//		{
//			desc: "Completed",
//			request: &types.MsgUpdateActionLevelFee{Creator: creator,
//				Action: strconv.Itoa(0),
//			},
//		},
//		{
//			desc: "Unauthorized",
//			request: &types.MsgUpdateActionLevelFee{Creator: "B",
//				Action: strconv.Itoa(0),
//			},
//			err: sdkerrors.ErrUnauthorized,
//		},
//		{
//			desc: "KeyNotFound",
//			request: &types.MsgUpdateActionLevelFee{Creator: creator,
//				Action: strconv.Itoa(100000),
//			},
//			err: sdkerrors.ErrKeyNotFound,
//		},
//	} {
//		t.Run(tc.desc, func(t *testing.T) {
//			k, ctx := keepertest.GentlemintKeeper(t)
//			srv := keeper.NewMsgServerImpl(*k)
//			wctx := sdk.WrapSDKContext(ctx)
//			expected := &types.MsgCreateActionLevelFee{Creator: creator,
//				Action: strconv.Itoa(0),
//			}
//			_, err := srv.CreateActionLevelFee(wctx, expected)
//			require.NoError(t, err)
//
//			_, err = srv.UpdateActionLevelFee(wctx, tc.request)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//				rst, found := k.GetActionLevelFee(ctx,
//					expected.Action,
//				)
//				require.True(t, found)
//				require.Equal(t, expected.Creator, rst.Creator)
//			}
//		})
//	}
//}
//
//func TestActionLevelFeeMsgServerDelete(t *testing.T) {
//	creator := "A"
//
//	for _, tc := range []struct {
//		desc    string
//		request *types.MsgDeleteActionLevelFee
//		err     error
//	}{
//		{
//			desc: "Completed",
//			request: &types.MsgDeleteActionLevelFee{Creator: creator,
//				Action: strconv.Itoa(0),
//			},
//		},
//		{
//			desc: "Unauthorized",
//			request: &types.MsgDeleteActionLevelFee{Creator: "B",
//				Action: strconv.Itoa(0),
//			},
//			err: sdkerrors.ErrUnauthorized,
//		},
//		{
//			desc: "KeyNotFound",
//			request: &types.MsgDeleteActionLevelFee{Creator: creator,
//				Action: strconv.Itoa(100000),
//			},
//			err: sdkerrors.ErrKeyNotFound,
//		},
//	} {
//		t.Run(tc.desc, func(t *testing.T) {
//			k, ctx := keepertest.GentlemintKeeper(t)
//			srv := keeper.NewMsgServerImpl(*k)
//			wctx := sdk.WrapSDKContext(ctx)
//
//			_, err := srv.CreateActionLevelFee(wctx, &types.MsgCreateActionLevelFee{Creator: creator,
//				Action: strconv.Itoa(0),
//			})
//			require.NoError(t, err)
//			_, err = srv.DeleteActionLevelFee(wctx, tc.request)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//				_, found := k.GetActionLevelFee(ctx,
//					tc.request.Action,
//				)
//				require.False(t, found)
//			}
//		})
//	}
//}
