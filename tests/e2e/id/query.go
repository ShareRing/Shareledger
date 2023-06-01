package id

import (
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/id/client/cli"
	"github.com/sharering/shareledger/x/id/types"
)

func (s *E2ETestSuite) TestGetByID() {
	testCases := tests.TestCases{
		{
			Name: "valid query id by id",
			Args: []string{
				id1.Id,
			},
			ExpectErr: false,
			RespType:  &types.QueryIdByIdResponse{},
			Expected: &types.QueryIdByIdResponse{
				Id: &id1,
			},
		}, {
			Name:      "query id by id not pass Id",
			Args:      []string{},
			ExpectErr: true,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdIdById(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestGetByAddress() {
	testCases := tests.TestCases{{
		Name: "valid query id by address",
		Args: []string{
			id1.Data.OwnerAddress,
		},
		ExpectErr: false,
		RespType:  &types.QueryIdByIdResponse{},
		Expected: &types.QueryIdByIdResponse{
			Id: &id1,
		},
	}, {
		Name:      "query id by address not pass address",
		Args:      []string{},
		ExpectErr: true,
	}}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdIdByAddress(), s.network.Validators[0])
}
