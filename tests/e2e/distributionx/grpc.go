package distributionx

import (
	"fmt"

	query "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/distributionx/types"
)

func (s *E2ETestSuite) TestGRPC() {
	val := s.network.Validators[0]
	buildUrl := func(surfix string) string {
		return fmt.Sprintf("%s/sharering/shareledger/distributionx/%s", val.APIAddress, surfix)
	}
	testCases := tests.TestCasesGrpc{
		{
			Name:      "gRPC get params",
			URL:       buildUrl("params"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryParamsResponse{},
			Expected: &types.QueryParamsResponse{
				Params: params,
			},
		},
		{
			Name:      "gRPC get all rewards",
			URL:       buildUrl("reward"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryAllRewardResponse{},
			Expected: &types.QueryAllRewardResponse{
				Reward: []types.Reward{reward1},
				Pagination: &query.PageResponse{
					Total: 1,
				},
			},
		},
		{
			Name:      "gRPC get reward by index",
			URL:       buildUrl("reward/" + reward1.Index),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryGetRewardResponse{},
			Expected: &types.QueryGetRewardResponse{
				Reward: reward1,
			},
		},
		{
			Name:      "gRPC get reward by index not found",
			URL:       buildUrl("reward/notexists"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get builder count",
			URL:       buildUrl("/builder_count"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryAllBuilderCountResponse{},
			Expected: &types.QueryAllBuilderCountResponse{
				BuilderCount: []types.BuilderCount{
					builderCount1,
				},
				Pagination: &query.PageResponse{
					Total: 1,
				},
			},
		},
		{
			Name:      "gRPC get builder count by index",
			URL:       buildUrl("/builder_count/ContractAddress1"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryGetBuilderCountResponse{},
			Expected: &types.QueryGetBuilderCountResponse{
				BuilderCount: builderCount1,
			},
		},
		{
			Name:      "gRPC get builder count by index not exists",
			URL:       buildUrl("/builder_count/notexists"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get builder list",
			URL:       buildUrl("/builder_list"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryAllBuilderListResponse{},
			Expected: &types.QueryAllBuilderListResponse{
				BuilderList: []types.BuilderList{builderList1, builderList2},
				Pagination:  &query.PageResponse{Total: 2},
			},
		},
		{
			Name:      "gRPC get builder list",
			URL:       buildUrl("/builder_list/1"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryGetBuilderListResponse{},
			Expected: &types.QueryGetBuilderListResponse{
				BuilderList: builderList1,
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}
