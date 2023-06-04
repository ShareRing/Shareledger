package swap

import (
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/id/types"
	"github.com/sharering/shareledger/x/swap/client/cli"
)

func (s *E2ETestSuite) TestQueryRequest() {

	testCases := tests.TestCases{
		{
			Name:      "query the valid request",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryIdByIdResponse{},
			Expected:  &types.QueryIdByIdResponse{},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdRequests(), s.network.Validators[0])
}
