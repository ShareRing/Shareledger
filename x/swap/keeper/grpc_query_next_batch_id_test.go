package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestNextBatchId() {
	_ = s.createNBatch([]uint64{}, "", 3)

	resp, err := s.swapKeeper.NextBatchId(sdk.WrapSDKContext(s.ctx), &types.QueryNextBatchIdRequest{})
	s.Require().NoError(err)
	s.Require().Equal(3, int(resp.NextCount))

	s.swapKeeper.SetBatchCount(s.ctx, 2)

	resp, err = s.swapKeeper.NextBatchId(sdk.WrapSDKContext(s.ctx), &types.QueryNextBatchIdRequest{})
	s.Require().NoError(err)
	s.Require().Equal(2, int(resp.NextCount))
}

func (s *KeeperTestSuite) TestNextBatchId_invalidRequest() {
	_ = s.createNBatch([]uint64{}, "", 1)

	_, err := s.swapKeeper.NextBatchId(sdk.WrapSDKContext(s.ctx), nil)
	s.Require().Error(err)
}
