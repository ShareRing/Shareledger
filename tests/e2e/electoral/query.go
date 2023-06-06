package electoral

import (
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/electoral/client/cli"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (s *E2ETestSuite) TestQueryAccountOperator() {
	testCases := tests.TestCases{
		{
			Name:      "query account operator by address",
			Args:      []string{accOp.Address},
			ExpectErr: false,
			RespType:  &types.QueryAccountOperatorResponse{},
			Expected: &types.QueryAccountOperatorResponse{
				AccState: &accOp,
			},
		},
		{
			Name:      "query account operator by address empty",
			Args:      []string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "query account operator by address not exists",
			Args:      []string{"address2"},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdAccountOperator(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryAccountOperators() {
	testCases := tests.TestCases{
		{
			Name:      "query account operators",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryAccountOperatorsResponse{},
			Expected: &types.QueryAccountOperatorsResponse{
				AccStates: []*types.AccState{&accOp},
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdAccountOperators(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryApprover() {
	testCases := tests.TestCases{
		{
			Name:      "query approver by address",
			Args:      []string{accApprover.Address},
			ExpectErr: false,
			RespType:  &types.QueryApproverResponse{},
			Expected: &types.QueryApproverResponse{
				AccState: accApprover,
			},
		},
		{
			Name:      "query approver by address empty",
			Args:      []string{},
			RespType:  &types.QueryApproverResponse{},
			ExpectErr: true,
		},
		{
			Name: "query approver by address not exists",
			Args: []string{
				"notExistsAddress",
			},
			RespType:  &types.QueryApproverResponse{},
			ExpectErr: true,
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdApprover(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryApprovers() {
	testCases := tests.TestCases{
		{
			Name:      "query approvers",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryApproversResponse{},
			Expected: &types.QueryApproversResponse{
				Approvers: []*types.AccState{&accApprover},
			},
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdApproves(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryVoter() {
	testCases := tests.TestCases{
		{
			Name:      "query voter by address",
			Args:      []string{accVoter.Address},
			ExpectErr: false,
			RespType:  &types.QueryVoterResponse{},
			Expected: &types.QueryVoterResponse{
				Voter: accVoter,
			},
		},
		{
			Name:      "query voter by address empty",
			Args:      []string{},
			RespType:  &types.QueryVoterResponse{},
			ExpectErr: true,
		},
		{
			Name: "query voter by address not exists",
			Args: []string{
				"notExistsAddress",
			},
			RespType:  &types.QueryVoterResponse{},
			ExpectErr: true,
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdGetVoter(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryVoters() {
	testCases := tests.TestCases{
		{
			Name:      "query voters",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryVotersResponse{},
			Expected: &types.QueryVotersResponse{
				Voters: []*types.AccState{&accVoter},
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdGetVoters(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryLoader() {
	testCases := tests.TestCases{
		{
			Name:      "query loader by address",
			Args:      []string{accKeyShrpLoaders.Address},
			ExpectErr: false,
			RespType:  &types.QueryLoaderResponse{},
			Expected: &types.QueryLoaderResponse{
				AccState: &accKeyShrpLoaders,
			},
		},
		{
			Name:      "query loader by address empty",
			Args:      []string{},
			ExpectErr: true,
			RespType:  &types.QueryLoaderResponse{},
		},
		{
			Name: "query loader by address not exists",
			Args: []string{
				"notExistsAddress",
			},
			RespType:  &types.QueryLoaderResponse{},
			ExpectErr: true,
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdGetLoader(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryLoaders() {
	testCases := tests.TestCases{
		{
			Name:      "query loaders",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryLoadersResponse{},
			Expected: &types.QueryLoadersResponse{
				Loaders: []*types.AccState{&accKeyShrpLoaders},
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdGetLoaders(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryLoadersFromFile() {
	path, err := testutil.GetAbsolutePath("/electoral_loaders.json")
	s.Require().NoError(err)

	testCases := tests.TestCases{
		{
			Name:      "query loaders from file",
			Args:      []string{path},
			ExpectErr: false,
			RespType:  &types.QueryLoaderResponse{},
			Expected: &types.QueryLoaderResponse{
				AccState: &accKeyShrpLoaders,
			},
		},
		{
			Name:      "query loaders from file by address empty",
			Args:      []string{},
			ExpectErr: true,
			RespType:  &types.QueryLoaderResponse{},
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdGetLoadersFromFile(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryIDSigner() {
	testCases := tests.TestCases{
		{
			Name:      "query id signer by address",
			Args:      []string{accIDSigner.Address},
			ExpectErr: false,
			RespType:  &types.QueryIdSignerResponse{},
			Expected: &types.QueryIdSignerResponse{
				AccState: &accIDSigner,
			},
		},
		{
			Name:      "query id signer by address empty",
			Args:      []string{},
			ExpectErr: true,
			RespType:  &types.QueryIdSignerResponse{},
		},
		{
			Name: "query id signer by address not exists",
			Args: []string{
				"notExistsAddress",
			},
			RespType:  &types.QueryIdSignerResponse{},
			ExpectErr: true,
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdIdSigner(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryIDSigners() {
	testCases := tests.TestCases{
		{
			Name:      "query id signers",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryIdSignersResponse{},
			Expected: &types.QueryIdSignersResponse{
				AccStates: []*types.AccState{&accIDSigner},
			},
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdIdSigners(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryDocumentIssuer() {
	testCases := tests.TestCases{
		{
			Name:      "query document issuer by address",
			Args:      []string{accDocIssuer.Address},
			ExpectErr: false,
			RespType:  &types.QueryDocumentIssuerResponse{},
			Expected: &types.QueryDocumentIssuerResponse{
				AccState: &accDocIssuer,
			},
		},
		{
			Name:      "query document issuer by address empty",
			Args:      []string{},
			ExpectErr: true,
			RespType:  &types.QueryDocumentIssuerResponse{},
		},
		{
			Name: "query document issuer by address not exists",
			Args: []string{
				"notExistsAddress",
			},
			RespType:  &types.QueryDocumentIssuerResponse{},
			ExpectErr: true,
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdDocumentIssuer(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryDocumentIssuers() {
	testCases := tests.TestCases{
		{
			Name:      "query document issuers",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentIssuersResponse{},
			Expected: &types.QueryDocumentIssuersResponse{
				AccStates: []*types.AccState{&accDocIssuer},
			},
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdDocumentIssuers(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQuerySwapManager() {
	testCases := tests.TestCases{
		{
			Name:      "query swap manager by address",
			Args:      []string{accSwapManager.Address},
			ExpectErr: false,
			RespType:  &types.QuerySwapManagerResponse{},
			Expected: &types.QuerySwapManagerResponse{
				AccState: accSwapManager,
			},
		},
		{
			Name:      "query swap manager by address empty",
			Args:      []string{},
			ExpectErr: true,
			RespType:  &types.QuerySwapManagerResponse{},
		},
		{
			Name: "query swap manager by address not exists",
			Args: []string{
				"notExistsAddress",
			},
			RespType:  &types.QuerySwapManagerResponse{},
			ExpectErr: true,
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdSwapManager(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQuerySwapManagers() {
	testCases := tests.TestCases{
		{
			Name:      "query swap managers",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QuerySwapManagersResponse{},
			Expected: &types.QuerySwapManagersResponse{
				SwapManagers: []*types.AccState{&accSwapManager},
			},
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdSwapManagers(), s.network.Validators[0])
}
