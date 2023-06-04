package distributionx

import (
	query "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/distributionx/client/cli"
	"github.com/sharering/shareledger/x/distributionx/types"
)

func (s *E2ETestSuite) TestQueryParams() {
	testCases := tests.TestCases{
		{
			Name:      "query params",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryParamsResponse{},
			Expected: &types.QueryParamsResponse{
				Params: params,
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdQueryParams(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryListReward() {
	testCases := tests.TestCases{
		// NOTE: why cli not return correct pagination value?
		{
			Name:      "query list reward",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryAllRewardResponse{},
			Expected: &types.QueryAllRewardResponse{
				Reward:     []types.Reward{reward1},
				Pagination: &query.PageResponse{},
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdListReward(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryShowReward() {
	testCases := tests.TestCases{
		{
			Name:      "query show reward",
			Args:      []string{reward1.Index},
			ExpectErr: false,
			RespType:  &types.QueryGetRewardResponse{},
			Expected: &types.QueryGetRewardResponse{
				Reward: reward1,
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdShowReward(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryListBuilderCount() {
	testCases := tests.TestCases{
		{
			Name:      "query list builder count",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryAllBuilderCountResponse{},
			Expected: &types.QueryAllBuilderCountResponse{
				BuilderCount: []types.BuilderCount{
					builderCount1,
				},
				Pagination: &query.PageResponse{},
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdListBuilderCount(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryShowBuilderCount() {
	testCases := tests.TestCases{
		{
			Name:      "query show builder count",
			Args:      []string{"ContractAddress1"},
			ExpectErr: false,
			RespType:  &types.QueryGetBuilderCountResponse{},
			Expected: &types.QueryGetBuilderCountResponse{
				BuilderCount: builderCount1,
			},
		},
		{
			Name:      "query show builder count invalid args",
			Args:      []string{},
			ExpectErr: true,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdShowBuilderCount(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryListBuilderList() {
	testCases := tests.TestCases{
		{
			Name:      "query list builder list",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryAllBuilderListResponse{},
			Expected: &types.QueryAllBuilderListResponse{
				BuilderList: []types.BuilderList{builderList1, builderList2},
				Pagination:  &query.PageResponse{},
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdListBuilderList(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryShowBuilderList() {
	testCases := tests.TestCases{
		{
			Name:      "query show builder list",
			Args:      []string{"1"},
			ExpectErr: false,
			RespType:  &types.QueryGetBuilderListResponse{},
			Expected: &types.QueryGetBuilderListResponse{
				BuilderList: builderList1,
			},
		},
		{
			Name:      "query show builder list invalid",
			Args:      []string{},
			ExpectErr: true,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdShowBuilderList(), s.network.Validators[0])
}
