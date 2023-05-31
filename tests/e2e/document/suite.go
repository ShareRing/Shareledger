//go:build e2e

package document

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/document/client/cli"
	"github.com/sharering/shareledger/x/document/types"
	"github.com/stretchr/testify/suite"
)

var (
	firstDoc = types.Document{
		Holder:  "USER-1",
		Issuer:  "shareledger19l9teyc2znfv630sv9gzjc92xurzxcers75xud",
		Proof:   "testProof",
		Data:    "testData",
		Version: 0,
	}
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger document module")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg)
	genesisState := s.cfg.GenesisState
	var docGenesis types.GenesisState
	s.Require().NoError(s.cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &docGenesis))
	docGenesis.Documents = []*types.Document{&firstDoc}
	docGenesisBz, err := s.cfg.Codec.MarshalJSON(&docGenesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = docGenesisBz
	s.cfg.GenesisState = genesisState
	s.network = network.New(s.T(), s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestSuite) TestQueryDocumentByHolderId() {
	testCases := tests.TestCases{
		{
			Name: "query document by holder id",
			Args: []string{
				firstDoc.Holder,
				fmt.Sprintf("--%s=json", flags.FlagOutput),
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
				fmt.Sprintf("--%s=json", flags.FlagOutput),
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by holder id with InvalidArgument 3",
			Args: []string{
				"KFhnI60lCMlwL1gtu2nwZKyNsCzd42eXodt9hmRrBsf4Y1L4fNasGSRibI7geMzcX",
				fmt.Sprintf("--%s=json", flags.FlagOutput),
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
				fmt.Sprintf("--%s=json", flags.FlagOutput),
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
				fmt.Sprintf("--%s=json", flags.FlagOutput),
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by proof with InvalidArgument 3",
			Args: []string{
				"XYZ",
				fmt.Sprintf("--%s=json", flags.FlagOutput),
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by proof with InvalidArgument 4",
			Args: []string{
				fmt.Sprintf("--%s=abc", flags.FlagKeyringBackend),
			},
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
				fmt.Sprintf("--%s=json", flags.FlagOutput),
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
				fmt.Sprintf("--%s=json", flags.FlagOutput),
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by issuer with InvalidArgument 3",
			Args: []string{
				"XYZ",
				fmt.Sprintf("--%s=json", flags.FlagOutput),
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		}, {
			Name: "query document by issuer with InvalidArgument 4",
			Args: []string{
				fmt.Sprintf("--%s=abc", flags.FlagKeyringBackend),
			},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdDocumentOfHolderByIssuer(), s.network.Validators[0])
}
