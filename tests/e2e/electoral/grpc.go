package electoral

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/electoral/types"
)

func buildURL(s *E2ETestSuite, suffix string) string {
	return fmt.Sprintf("%s/shareledger/electoral/%s", s.network.Validators[0].APIAddress, suffix)
}

func buildShrURL(s *E2ETestSuite, suffix string) string {
	return fmt.Sprintf("%s/sharering/shareledger/electoral/%s", s.network.Validators[0].APIAddress, suffix)
}

func (s *E2ETestSuite) TestGRPCQuerySwapManager() {
	testCases := []tests.TestCaseGrpc{
		{
			Name:      "gRPC get swapmanagers",
			URL:       buildShrURL(s, "swap_managers"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QuerySwapManagersResponse{},
			Expected: &types.QuerySwapManagersResponse{
				SwapManagers: []*types.AccState{&accSwapManager},
			},
		},
		{
			Name: "gRPC get swapmanagers by address not exists",
			URL:  buildShrURL(s, "swap_managers/"+"notExistsAddress"),

			Headers:   map[string]string{},
			ExpectErr: true,
		},
	}

	tests.RunTestCasesGrpc(&s.Suite, testCases, s.network.Validators[0])
}

func (s *E2ETestSuite) TestGRPCQueryApprovers() {
	testCases := []tests.TestCaseGrpc{
		{
			Name:      "gRPC get approvers by address",
			URL:       buildShrURL(s, "approver/"+accApprover.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryApproverResponse{},
			Expected: &types.QueryApproverResponse{
				AccState: accApprover,
			},
		},
		{
			Name: "gRPC get approvers by address not exists",
			URL:  buildShrURL(s, "approver/"+"notExistsAddress"),

			Headers:   map[string]string{},
			ExpectErr: true,
		},
	}

	tests.RunTestCasesGrpc(&s.Suite, testCases, s.network.Validators[0])
}

func (s *E2ETestSuite) TestGRPCQueryRelayers() {
	testCases := []tests.TestCaseGrpc{
		{
			Name:      "gRPC get relayers",
			URL:       buildShrURL(s, "relayers"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryRelayersResponse{},
			Expected: &types.QueryRelayersResponse{
				Relayers: []*types.AccState{&accKeyRelayer},
			},
		},
		{
			Name:      "gRPC get relayer by address",
			URL:       buildShrURL(s, "relayer/"+accKeyRelayer.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryRelayerResponse{},
			Expected: &types.QueryRelayerResponse{
				AccState: accKeyRelayer,
			},
		},
		{
			Name: "gRPC get relayer by address not exists",
			URL:  buildShrURL(s, "relayer/"+"notExistsAddress"),

			Headers:   map[string]string{},
			ExpectErr: true,
		},
	}

	tests.RunTestCasesGrpc(&s.Suite, testCases, s.network.Validators[0])
}

func (s *E2ETestSuite) TestGRPCQueryLoaders() {
	testCases := []tests.TestCaseGrpc{
		{
			Name:      "gRPC get loaders",
			URL:       buildURL(s, "loaders"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryLoadersResponse{},
			Expected: &types.QueryLoadersResponse{
				Loaders: []*types.AccState{&accKeyShrpLoaders},
			},
		},
		{
			Name:      "gRPC get loader by address",
			URL:       buildURL(s, "loaders/"+accKeyShrpLoaders.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryLoaderResponse{},
			Expected: &types.QueryLoaderResponse{
				AccState: &accKeyShrpLoaders,
			},
		},
		{
			Name: "gRPC get loader by address not exists",
			URL:  buildURL(s, "loaders/"+"notExistsAddress"),

			Headers:   map[string]string{},
			ExpectErr: true,
		},
	}

	tests.RunTestCasesGrpc(&s.Suite, testCases, s.network.Validators[0])
}

func (s *E2ETestSuite) TestGRPCQueryVoter() {
	testCases := []tests.TestCaseGrpc{
		{
			Name:      "gRPC get voter",
			URL:       buildURL(s, "voters/"+accVoter.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryVoterResponse{},
			Expected: &types.QueryVoterResponse{
				Voter: accVoter,
			},
		},
		{
			Name: "gRPC get voter by address not exists",
			URL:  buildURL(s, "voters/"+"notExistsAddress"),

			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get voters",
			URL:       buildURL(s, "voters"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryVotersResponse{},
			Expected: &types.QueryVotersResponse{
				Voters: []*types.AccState{&accVoter},
			},
		},
	}

	tests.RunTestCasesGrpc(&s.Suite, testCases, s.network.Validators[0])
}

func (s *E2ETestSuite) TestGRPCQueryAccountOperator() {
	testCases := []tests.TestCaseGrpc{
		{
			Name:      "gRPC get account operator",
			URL:       buildURL(s, "accountOperators/"+accOp.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryAccountOperatorResponse{},
			Expected: &types.QueryAccountOperatorResponse{
				AccState: &accOp,
			},
		},
		{
			Name: "gRPC get account operator by address not exists",
			URL:  buildURL(s, "accountOperators/"+"notExistsAddress"),

			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get account operators",
			URL:       buildURL(s, "accountOperators"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryAccountOperatorsResponse{},
			Expected: &types.QueryAccountOperatorsResponse{
				AccStates: []*types.AccState{&accOp},
			},
		},
	}

	tests.RunTestCasesGrpc(&s.Suite, testCases, s.network.Validators[0])
}

func (s *E2ETestSuite) TestGRPCQueryAccState() {
	testCases := []tests.TestCaseGrpc{
		{
			Name:      "gRPC get account state",
			URL:       buildURL(s, "accStates/"+accOp.Key),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryAccStateResponse{},
			Expected: &types.QueryAccStateResponse{
				AccState: accOp,
			},
		},
		{
			Name: "gRPC get account state by address not exists",
			URL:  buildURL(s, "accStates/"+"notExistsAddress"),

			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get account states",
			URL:       buildURL(s, "accStates"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryAccStatesResponse{},
			Expected: &types.QueryAccStatesResponse{
				AccState: []types.AccState{
					accOp,
					accApprover,
					accDocIssuer,
					accIDSigner,
					accKeyRelayer,
					accKeyShrpLoaders,
					accSwapManager,
					accVoter,
				},
				Pagination: &query.PageResponse{
					Total: 8,
				},
			},
		},
	}

	tests.RunTestCasesGrpc(&s.Suite, testCases, s.network.Validators[0])
}

func (s *E2ETestSuite) TestGRPCQueryDocIssuer() {
	testCases := []tests.TestCaseGrpc{
		{
			Name:      "gRPC get doc issuer",
			URL:       buildURL(s, "documentIssuers/"+accDocIssuer.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentIssuerResponse{},
			Expected: &types.QueryDocumentIssuerResponse{
				AccState: &accDocIssuer,
			},
		},
		{
			Name: "gRPC get doc issuer by address not exists",
			URL:  buildURL(s, "documentIssuers/"+"notExistsAddress"),

			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get doc issuers",
			URL:       buildURL(s, "documentIssuers"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentIssuersResponse{},
			Expected: &types.QueryDocumentIssuersResponse{
				AccStates: []*types.AccState{&accDocIssuer},
			},
		},
	}

	tests.RunTestCasesGrpc(&s.Suite, testCases, s.network.Validators[0])
}
