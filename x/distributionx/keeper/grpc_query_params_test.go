package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/types"
)

func (s *KeeperTestSuite) TestParamsQuery() {
	wctx := sdk.WrapSDKContext(s.ctx)
	params := types.DefaultParams()
	s.dKeeper.SetParams(s.ctx, params)

	// nil request
	response, err := s.dKeeper.Params(wctx, nil)
	s.Require().NotNil(err)

	// normal request
	response, err = s.dKeeper.Params(wctx, &types.QueryParamsRequest{})
	s.Require().NoError(err)
	s.Require().Equal(&types.QueryParamsResponse{Params: params}, response)
}
