package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestCancelBatches() {
	ctx := sdk.WrapSDKContext(s.ctx)
	reqs := s.createNRequest(normalRequest, 3)

	for i := 0; i < 3; i++ {
		req := types.NewMsgApprove("", "", []uint64{reqs[i].GetId()})
		_, err := s.msgServer.ApproveOut(ctx, req)
		s.Require().NoError(err)
	}
	msgReq := types.MsgCancelBatches{
		Creator: "",
		Ids:     []uint64{0, 1, 2},
	}

	_, err := s.msgServer.CancelBatches(ctx, &msgReq)
	s.Require().NoError(err)
	batches := s.swapKeeper.GetAllBatch(s.ctx)
	s.Require().Equal(0, len(batches))
}
