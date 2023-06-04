package electoral

// func (s *E2ETestSuite) TestGRPCQueryAccState() {
// 	val := s.network.Validators[0]
// 	getURL := func(key string) string {
// 		return fmt.Sprintf("%s/shareledger/electoral/%s", val.APIAddress, key)
// 	}

// 	testCases := tests.TestCasesGrpc{
// 		{
// 			Name:      "gRPC get acc state by key",
// 			URL:       getURL(accDocIssuer.Key),
// 			ExpectErr: false,
// 			RespType:  &types.QueryAccStateResponse{},
// 			Expected: &types.QueryAccStateResponse{
// 				AccState: accDocIssuer,
// 			},
// 		},
// 		{
// 			Name:      "gRPC get acc state by key empty",
// 			URL:       getURL(""),
// 			ExpectErr: true,
// 		},
// 		{
// 			Name:      "gRPC get acc state by key empty not exists",
// 			URL:       getURL("noneExistsKey"),
// 			ExpectErr: true,
// 		},
// 	}

// 	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
// }

// func (s *E2ETestSuite) TestGRPCQueryAccStates() {
// 	val := s.network.Validators[0]
// 	baseURL := val.APIAddress
// 	testCases := tests.TestCasesGrpc{
// 		{
// 			Name:      "gRPC get acc states",
// 			URL:       fmt.Sprintf("%s/shareledger/electoral", baseURL),
// 			ExpectErr: false,
// 			RespType:  &types.Q
// 			Expected: &types.QueryAccStatesResponse{
// 				AccStates: []*types.AccState{accDocIssuer, accDocHolder},
// 			},
// 		},
// 	}

// 	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
// }
