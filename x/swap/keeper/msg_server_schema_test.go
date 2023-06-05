package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestFormatMsgServerCreate() {
	wctx := sdk.WrapSDKContext(s.ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateSchema{Creator: creator,
			Network: strconv.Itoa(i),
			In:      sdk.NewDecCoin(denom.Base, sdk.NewInt(100)),
			Out:     sdk.NewDecCoin(denom.Base, sdk.NewInt(100)),
		}
		_, err := s.msgServer.CreateSchema(wctx, expected)
		s.Require().NoError(err)
		rst, found := s.swapKeeper.GetSchema(s.ctx,
			expected.Network,
		)
		s.Require().True(found)
		s.Require().Equal(expected.Creator, rst.Creator)
	}
}

func (s *KeeperTestSuite) TestFormatMsgServerUpdate() {
	creator := "A"
	in := sdk.NewDecCoin(denom.Base, sdk.NewInt(100))
	out := sdk.NewDecCoin(denom.Base, sdk.NewInt(100))

	wctx := sdk.WrapSDKContext(s.ctx)
	expected := &types.MsgCreateSchema{Creator: creator,
		Network: strconv.Itoa(0),
		In:      in,
		Out:     out,
	}
	_, err := s.msgServer.CreateSchema(wctx, expected)
	s.Require().NoError(err)
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
		s.Run(tc.desc, func() {
			_, err = s.msgServer.UpdateSchema(wctx, tc.request)
			if tc.err != nil {
				s.Require().ErrorIs(err, tc.err)
			} else {
				s.Require().NoError(err)
				rst, found := s.swapKeeper.GetSchema(s.ctx,
					expected.Network,
				)
				s.Require().True(found)
				s.Require().Equal(expected.Creator, rst.Creator)
			}
		})
	}
}

func (s *KeeperTestSuite) TestFormatMsgServerDelete() {
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
		s.Run(tc.desc, func() {
			wctx := sdk.WrapSDKContext(s.ctx)

			_, err := s.msgServer.CreateSchema(wctx, &types.MsgCreateSchema{Creator: creator,
				Network: strconv.Itoa(0),
				In:      sdk.NewDecCoin(denom.Base, sdk.NewInt(100)),
				Out:     sdk.NewDecCoin(denom.Base, sdk.NewInt(100)),
			})
			s.Require().NoError(err)
			_, err = s.msgServer.DeleteSchema(wctx, tc.request)
			if tc.err != nil {
				s.Require().ErrorIs(err, tc.err)
			} else {
				s.Require().NoError(err)
				_, found := s.swapKeeper.GetSchema(s.ctx,
					tc.request.Network,
				)
				s.Require().False(found)
			}
		})
	}
}
