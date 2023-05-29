package id

import (
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/id/client/cli"
	"github.com/sharering/shareledger/x/id/types"
)

var id1 = types.Id{
	Id: "Id1",
	Data: &types.BaseID{
		IssuerAddress: "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak",
		BackupAddress: "BackupAddress",
		OwnerAddress:  "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr",
		ExtraData:     "ExtraData",
	},
}

func (s *E2ETestSuite) TestGetByID() {
	testCases := tests.TestCases{
		{
			Name: "valid query id by id",
			Args: []string{
				id1.Id,
				network.JSONFlag,
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
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdIdById(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestGetByAddress() {
	testCases := tests.TestCases{{
		Name: "valid query id by address",
		Args: []string{
			id1.Data.OwnerAddress,
			network.JSONFlag,
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
		RespType:  nil,
		Expected:  nil,
	}}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdIdByAddress(), s.network.Validators[0])
}
