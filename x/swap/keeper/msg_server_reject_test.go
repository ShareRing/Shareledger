package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestReject() {
	reqs := s.createNRequest(normalRequestWithNSHR, 1)
	acc, err := sdk.AccAddressFromBech32(testAddr)
	s.Require().NoError(err)
	expectedRefund := reqs[0].Amount.Add(reqs[0].Fee)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, acc, sdk.Coins{expectedRefund}).Return(nil)

	for _, tc := range []struct {
		desc        string
		request     *types.MsgReject
		expectedErr bool
	}{
		{
			desc: "ok",
			request: &types.MsgReject{
				Creator: reqs[0].SrcAddr,
				Ids:     []uint64{reqs[0].Id},
			},
			expectedErr: false,
		},
		{
			desc: "InvalidAccAddress",
			request: &types.MsgReject{
				Creator: "test1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
				Ids:     []uint64{reqs[0].Id},
			},
			expectedErr: true,
		},
	} {
		s.Run(tc.desc, func() {
			_, err := s.msgServer.Reject(sdk.WrapSDKContext(s.ctx), tc.request)
			if tc.expectedErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}
