package id

import (
	"fmt"

	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/id/types"
)

func (s *E2ETestSuite) Test() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress
	testCases := tests.TestCasesGrpc{
		{
			Name:      "gRPC id by id valid",
			URL:       fmt.Sprintf("%s/shareledger/id/id/%s", baseURL, id1.Id),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryIdByIdResponse{},
			Expected: &types.QueryIdByIdResponse{
				Id: &id1,
			},
		},
		{
			Name:      "gRPC id by id not pass id",
			URL:       fmt.Sprintf("%s/shareledger/id/id", baseURL),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC id by id not exists id",
			URL:       fmt.Sprintf("%s/shareledger/id/id/%s", baseURL, "invalid_id"),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}
