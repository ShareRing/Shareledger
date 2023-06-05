package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestParamsQuery() {
	params := types.DefaultParams()
	s.swapKeeper.SetParams(s.ctx, params)

	response, err := s.swapKeeper.Params(sdk.WrapSDKContext(s.ctx), &types.QueryParamsRequest{})
	s.Require().NoError(err)
	s.Require().Equal(&types.QueryParamsResponse{Params: params}, response)
}

func (s *KeeperTestSuite) TestParamsQuery_invalidRequest() {
	params := types.DefaultParams()
	s.swapKeeper.SetParams(s.ctx, params)

	_, err := s.swapKeeper.Params(sdk.WrapSDKContext(s.ctx), nil)
	s.Require().Error(err)
}
