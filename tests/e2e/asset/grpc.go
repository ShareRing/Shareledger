package asset

import (
	"fmt"

	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/asset/types"
)

func (s *E2ETestSuite) TestGRPCQueryAsset() {
	val := s.network.Validators[0]
	getURL := func(uuid string) string {
		return fmt.Sprintf("%s/shareledger/asset/%s", val.APIAddress, uuid)
	}
	testCases := tests.TestCasesGrpc{
		{
			Name:      "gRPC asset by uuid",
			URL:       getURL(asset1.UUID),
			ExpectErr: false,
			RespType:  &types.QueryAssetByUUIDResponse{},
			Expected: &types.QueryAssetByUUIDResponse{
				Asset: asset1,
			},
		},
		{
			Name:      "gRPC asset by uuid not exists",
			URL:       getURL("new_uuid"),
			ExpectErr: true,
		},
	}

	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}
