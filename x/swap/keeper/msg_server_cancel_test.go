package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestCancelRequests() {
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, gomock.Any(), gomock.Any()).Return(nil)
	msgRequest := types.MsgCancel{
		// Creator: normalRequest.SrcAddr,
		Ids: []uint64{0},
	}
	_ = s.createNRequest(normalRequest, 1)

	_, err := s.msgServer.Cancel(sdk.WrapSDKContext(s.ctx), &msgRequest)
	s.Require().Error(err)

	msgRequest.Creator = normalRequest.SrcAddr
	_, err = s.msgServer.Cancel(sdk.WrapSDKContext(s.ctx), &msgRequest)
	s.Require().NoError(err)
	_, found := s.swapKeeper.GetRequest(s.ctx, 0)
	s.Require().False(found)
}
