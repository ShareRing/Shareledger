package asset

import (
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/asset/client/cli"
	"github.com/sharering/shareledger/x/asset/types"
)

func (s *E2ETestSuite) TestQueryAssetByUUID() {
	testCases := tests.TestCases{
		{
			Name: "query asset by UUID",
			Args: []string{
				asset1.UUID,
			},
			ExpectErr: false,
			RespType:  &types.QueryAssetByUUIDResponse{},
			Expected: &types.QueryAssetByUUIDResponse{
				Asset: asset1,
			},
		},
		{
			Name:      "query asset by UUID not pass uuid",
			Args:      []string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name: "query asset by UUID not exists",
			Args: []string{
				"uuidnotexists",
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdAssetByUUID(), s.network.Validators[0])
}
