package document

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/document/client/cli"
	"github.com/sharering/shareledger/x/document/types"
)

func (s *E2ETestSuite) createNewDocument(id string) {
	args := []string{id,
		"TestProof" + id,
		"ExtraData",
		network.MakeByAccount(network.KeyDocIssuer)}
	resp, err := tests.RunCmdWithRetry(&s.Suite, cli.CmdCreateDocument(), s.network.Validators[0], args, 100)
	s.NoError(err)
	// code !=0 mean that this tx failed on CheckTx call (ante step) => the tx is not committed
	s.Require().Equal(uint32(0), resp.Code)
}

func (s *E2ETestSuite) revokeDocument(holder, proof string) {
	args := []string{holder,
		proof,
		network.MakeByAccount(network.KeyDocIssuer)}
	resp, err := tests.RunCmdWithRetry(&s.Suite, cli.CmdRevokeDocument(), s.network.Validators[0], args, 100)
	s.NoError(err)
	// code !=0 mean that this tx failed on CheckTx call (ante step) => the tx is not committed
	s.Require().Equal(uint32(0), resp.Code)
}

func (s *E2ETestSuite) TestCreateDocument() {
	testCases := tests.TestCasesTx{
		{
			Name: "create new document",
			Args: []string{
				secondId.Id,
				"TestProof2",
				"ExtraData2",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "create new document with unauthorize account",
			Args: []string{
				secondId.Id,
				"TestProof2",
				"ExtraData2",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCreateDocument(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCreateDocuments() {
	holderIds := thirdId.Id + "," + fourthId.Id
	testCases := tests.TestCasesTx{
		{
			Name: "create new documents",
			Args: []string{
				holderIds,
				"TestProof3,TestProof3",
				"ExtraData4,ExtraData4",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "create new documents with unauthorized account",
			Args: []string{
				holderIds,
				"TestProof3,TestProof4",
				"ExtraData3,ExtraData4",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCreateDocuments(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCmdRevokeDocument() {
	s.createNewDocument(sixthId.Id)
	testCases := tests.TestCasesTx{
		{
			Name: "revoke document",
			Args: []string{
				sixthId.Id,
				"TestProof" + sixthId.Id,
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "revoke document with empty input argument",
			Args: []string{
				"",
				"TestProof3",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdRevokeDocument(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCmdUpdateDocument() {
	s.createNewDocument(fifthId.Id)
	s.createNewDocument(eightId.Id)
	s.revokeDocument(eightId.Id, "TestProof"+eightId.Id)
	testCases := tests.TestCasesTx{
		{
			Name: "update document ok",
			Args: []string{
				fifthId.Id,
				"TestProof" + fifthId.Id,
				"UpdatedExtraData",
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
				"UpdatedExtraData",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr: true,
		},
		{
			Name: "update document with mismatch doc issuer",
			Args: []string{
				secondDoc.Holder,
				"testProof2",
				"UpdatedExtraData",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: types.ErrDocNotExisted.ABCICode(),
		},
		{
			Name: "update document that already revoked",
			Args: []string{
				eightId.Id,
				"TestProof" + eightId.Id,
				"UpdatedExtraData",
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: types.ErrDocRevoked.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdUpdateDocument(), s.network.Validators[0])
}
