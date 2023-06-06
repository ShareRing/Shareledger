package electoral

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/electoral/types"

	"github.com/sharering/shareledger/tests"
)

func (s *E2ETestSuite) TestGRPC() {
	val := s.network.Validators[0]
	buildURL := func(suffix string) string {
		return fmt.Sprintf("%s/shareledger/electoral/%s", val.APIAddress, suffix)
	}
	buildURLShareRing := func(suffix string) string {
		return fmt.Sprintf("%s/sharering/shareledger/electoral/%s", val.APIAddress, suffix)
	}
	testCase := tests.TestCasesGrpc{
		{
			Name:      "gRPC get all accstate",
			URL:       buildURL("accStates"),
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
		{
			Name:      "gRPC get accstate by key",
			URL:       buildURL("accStates/" + accApprover.Key),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryAccStateResponse{},
			Expected: &types.QueryAccStateResponse{
				AccState: accApprover,
			},
		},
		{
			Name:      "gRPC get accstate by key not exists",
			URL:       buildURL("accStates/" + "notExistsKey"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get accstate by key empty",
			URL:       buildURL("accStates/" + ""),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get operator by Address",
			URL:       buildURL("accountOperators/" + accOp.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryAccountOperatorResponse{},
			Expected: &types.QueryAccountOperatorResponse{
				AccState: &accOp,
			},
		},
		{
			Name:      "gRPC get operator by address not exists",
			URL:       buildURL("accountOperators/" + "notExistsAddress"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get operator by address empty",
			URL:       buildURL("accountOperators/" + ""),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get approver by address",
			URL:       buildURLShareRing("approver/" + accApprover.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryApproverResponse{},
			Expected: &types.QueryApproverResponse{
				AccState: accApprover,
			},
		},
		{
			Name:      "gRPC get approver by address not exists",
			URL:       buildURLShareRing("approver/" + "notExistsAddress"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get approver by address empty",
			URL:       buildURLShareRing("approver/" + ""),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get doc issuer by address",
			URL:       buildURL("documentIssuers/" + accDocIssuer.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentIssuerResponse{},
			Expected: &types.QueryDocumentIssuerResponse{
				AccState: &accDocIssuer,
			},
		},
		{
			Name:      "gRPC get doc issuer by address not exists",
			URL:       buildURL("documentIssuers/" + "notExistsAddress"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get doc issuer by address empty",
			URL:       buildURL("documentIssuers/" + ""),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get id signer by address",
			URL:       buildURL("idSigners/" + accIDSigner.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryIdSignerResponse{},
			Expected: &types.QueryIdSignerResponse{
				AccState: &accIDSigner,
			},
		},
		{
			Name:      "gRPC get id signer by address not exists",
			URL:       buildURL("idSigners/" + "notExistsAddress"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get id signer by address empty",
			URL:       buildURL("idSigners/" + ""),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get key relayer by address",
			URL:       buildURLShareRing("relayer/" + accKeyRelayer.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryRelayerResponse{},
			Expected: &types.QueryRelayerResponse{
				AccState: accKeyRelayer,
			},
		},
		{
			Name:      "gRPC get key relayer by address not exists",
			URL:       buildURLShareRing("relayer/" + "notExistsAddress"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get key relayer by address empty",
			URL:       buildURLShareRing("relayer/" + ""),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get voter by address",
			URL:       buildURL("voters/" + accVoter.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryVoterResponse{},
			Expected: &types.QueryVoterResponse{
				Voter: accVoter,
			},
		},
		{
			Name:      "gRPC get voter by address not exists",
			URL:       buildURL("voters/" + "notExistsAddress"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get voter by address empty",
			URL:       buildURL("voters/" + ""),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get voters",
			URL:       buildURL("voters"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryVotersResponse{},
			Expected: &types.QueryVotersResponse{
				Voters: []*types.AccState{&accVoter},
			},
		},
		{
			Name:      "gRPC get approvers",
			URL:       buildURLShareRing("approves"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryApproversResponse{},
			Expected: &types.QueryApproversResponse{
				Approvers: []*types.AccState{&accApprover},
			},
		},
		{
			Name:      "gRPC get doc issuers",
			URL:       buildURL("documentIssuers"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentIssuersResponse{},
			Expected: &types.QueryDocumentIssuersResponse{
				AccStates: []*types.AccState{&accDocIssuer},
			},
		},
		{
			Name:      "gRPC get id signers",
			URL:       buildURL("idSigners"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryIdSignersResponse{},
			Expected: &types.QueryIdSignersResponse{
				AccStates: []*types.AccState{&accIDSigner},
			},
		},
		{
			Name:      "gRPC get key relayers",
			URL:       buildURLShareRing("relayers"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryRelayersResponse{},
			Expected: &types.QueryRelayersResponse{
				Relayers: []*types.AccState{&accKeyRelayer},
			},
		},
		{
			Name:      "gRPC get relayer by address",
			URL:       buildURLShareRing("relayer/" + accKeyRelayer.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryRelayerResponse{},
			Expected: &types.QueryRelayerResponse{
				AccState: accKeyRelayer,
			},
		},
		{
			Name:      "gRPC get loader by address",
			URL:       buildURL("loaders/" + accKeyShrpLoaders.Address),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryLoaderResponse{},
			Expected: &types.QueryLoaderResponse{
				AccState: &accKeyShrpLoaders,
			},
		},
		{
			Name:      "gRPC get loader by address not exists",
			URL:       buildURL("loaders/" + "notExistsAddress"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get loader by address empty",
			URL:       buildURL("loaders/" + ""),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC get loaders",
			URL:       buildURL("loaders"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryLoadersResponse{},
			Expected: &types.QueryLoadersResponse{
				Loaders: []*types.AccState{&accKeyShrpLoaders},
			},
		},
		{
			Name:      "gRPC get swapmanagers",
			URL:       buildURLShareRing("swap_managers"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QuerySwapManagersResponse{},
			Expected: &types.QuerySwapManagersResponse{
				SwapManagers: []*types.AccState{&accSwapManager},
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCase, val)
}
