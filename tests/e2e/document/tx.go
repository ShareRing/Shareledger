package document

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/document/client/cli"
)

func (s *E2ETestSuite) TestCreateDocument() {
	testCases := tests.TestCasesTx{
		{
			Name: "create new document",
			Args: []string{
				"Holder-ID",
				"TestProof",
				"ExtraData",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "create new document with unauthorize account",
			Args: []string{
				"Holder-ID",
				"TestProof",
				"ExtraData",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCreateDocument(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCreateDocuments() {
	testCases := tests.TestCasesTx{
		{
			Name: "create new documents",
			Args: []string{
				"Holder-ID1,Holder-ID2,Holder-ID3",
				"TestProof1,TestProof2,TestProof3",
				"ExtraData1,ExtraData2,ExtraData3",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "create new documents",
			Args: []string{
				"Holder-ID1,Holder-ID2,Holder-ID3",
				"TestProof1,TestProof2,TestProof3",
				"ExtraData1,ExtraData2,ExtraData3",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCreateDocuments(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCmdRevokeDocument() {
	testCases := tests.TestCasesTx{
		{
			Name: "revoke document",
			Args: []string{
				"Holder-ID",
				"TestProof",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "revoke document with empty input argument",
			Args: []string{
				"",
				"TestProof",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    true,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdRevokeDocument(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCmdUpdateDocument() {
	testCases := tests.TestCasesTx{
		{
			Name: "update document ok",
			Args: []string{
				"Holder-ID",
				"TestProof",
				"ExtraData",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "update document with empty input argument",
			Args: []string{
				"",
				"TestProof",
				"ExtraData",
				network.SkipConfirmation,
				network.SyncBroadcast,
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    true,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdUpdateDocument(), s.network.Validators[0])
}
