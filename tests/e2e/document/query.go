package document

import (
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/document/client/cli"
	"github.com/sharering/shareledger/x/document/types"
)

func (s *E2ETestSuite) TestQueryDocumentByHolderId() {
	testCases := tests.TestCases{
		{
			Name: "query document by holder id",
			Args: []string{
				firstDoc.Holder,
			},
			ExpectErr: false,
			RespType:  &types.QueryDocumentByHolderIdResponse{},
			Expected: &types.QueryDocumentByHolderIdResponse{
				Documents: []*types.Document{&firstDoc},
			},
		}, {
			Name:      "query document by holder id with InvalidArgument 1",
			Args:      []string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by holder id with InvalidArgument 2",
			Args: []string{
				"",
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by holder id with InvalidArgument 3",
			Args: []string{
				"KFhnI60lCMlwL1gtu2nwZKyNsCzd42eXodt9hmRrBsf4Y1L4fNasGSRibI7geMzcX",
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdDocumentByHolderId(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryDocumentByProof() {
	testCases := tests.TestCases{
		{
			Name: "query document by proof",
			Args: []string{
				firstDoc.Proof,
			},
			ExpectErr: false,
			RespType:  &types.QueryDocumentByProofResponse{},
			Expected: &types.QueryDocumentByProofResponse{
				Document: &firstDoc,
			},
		}, {
			Name:      "query document by proof with InvalidArgument 1",
			Args:      []string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by proof with InvalidArgument 2",
			Args: []string{
				"",
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by proof with InvalidArgument 3",
			Args: []string{
				"XYZ",
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name:      "query document by proof with InvalidArgument 4",
			Args:      []string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdDocumentByProof(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryDocumentOfHolderByIssuer() {
	testCases := tests.TestCases{
		{
			Name: "query document by issuer",
			Args: []string{
				firstDoc.Holder,
				firstDoc.Issuer,
			},
			ExpectErr: false,
			RespType:  &types.QueryDocumentOfHolderByIssuerResponse{},
			Expected: &types.QueryDocumentOfHolderByIssuerResponse{
				Documents: []*types.Document{&firstDoc},
			},
		}, {
			Name:      "query document by issuer with InvalidArgument 1",
			Args:      []string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by issuer with InvalidArgument 2",
			Args: []string{
				"",
				"",
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by issuer with InvalidArgument 3",
			Args: []string{
				"XYZ",
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name:      "query document by issuer with InvalidArgument 4",
			Args:      []string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdDocumentOfHolderByIssuer(), s.network.Validators[0])
}
