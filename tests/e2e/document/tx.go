package document

import (
	"time"

	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/document/client/cli"
)

var waitForSync = 2

func (s *E2ETestSuite) TestCreateDocument() {
	// need to wait for previous transaction finish and sync before try a next one
	// to avoid the err: "account sequence mismatch, expected <num1>, got <num0>, incorrect account sequence"
	time.Sleep(time.Duration(waitForSync) * time.Second)
	testCases := tests.TestCasesTx{
		{
			Name: "create new document",
			Args: []string{
				"Holder-ID",
				"TestProof",
				"ExtraData",
				network.SkipConfirmation,
				network.JSONFlag,
				network.SyncBroadcast,
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: 0,
		},
		{
			Name: "create new document with unauthorize account",
			Args: []string{
				"Holder-ID",
				"TestProof",
				"ExtraData",
				network.SkipConfirmation,
				network.JSONFlag,
				network.SyncBroadcast,
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: 0x4,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCreateDocument(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCreateDocuments() {
	// need to wait for previous transaction finish and sync before try a next one
	// to avoid the err: "account sequence mismatch, expected <num1>, got <num0>, incorrect account sequence"
	time.Sleep(time.Duration(waitForSync) * time.Second)
	testCases := tests.TestCasesTx{
		{
			Name: "create new documents",
			Args: []string{
				"Holder-ID1,Holder-ID2,Holder-ID3",
				"TestProof1,TestProof2,TestProof3",
				"ExtraData1,ExtraData2,ExtraData3",
				network.SkipConfirmation,
				network.JSONFlag,
				network.SyncBroadcast,
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: 0,
		},
		{
			Name: "create new documents",
			Args: []string{
				"Holder-ID1,Holder-ID2,Holder-ID3",
				"TestProof1,TestProof2,TestProof3",
				"ExtraData1,ExtraData2,ExtraData3",
				network.SkipConfirmation,
				network.JSONFlag,
				network.SyncBroadcast,
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: 0x4,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCreateDocuments(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCmdRevokeDocument() {
	// need to wait for previous transaction finish and sync before try a next one
	// to avoid the err: "account sequence mismatch, expected <num1>, got <num0>, incorrect account sequence"
	time.Sleep(time.Duration(waitForSync) * time.Second)
	testCases := tests.TestCasesTx{
		{
			Name: "revoke document",
			Args: []string{
				"Holder-ID",
				"TestProof",
				network.SkipConfirmation,
				network.JSONFlag,
				network.SyncBroadcast,
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: 0,
		},
		{
			Name: "revoke document with empty input argument",
			Args: []string{
				"",
				"TestProof",
				network.SkipConfirmation,
				network.JSONFlag,
				network.SyncBroadcast,
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    true,
			ExpectedCode: 0,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdRevokeDocument(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCmdUpdateDocument() {
	// need to wait for previous transaction finish and sync before try a next one
	// to avoid the err: "account sequence mismatch, expected <num1>, got <num0>, incorrect account sequence"
	time.Sleep(time.Duration(waitForSync) * time.Second)
	testCases := tests.TestCasesTx{
		{
			Name: "update document ok",
			Args: []string{
				"Holder-ID",
				"TestProof",
				"ExtraData",
				network.SkipConfirmation,
				network.JSONFlag,
				network.SyncBroadcast,
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    false,
			ExpectedCode: 0,
		},
		{
			Name: "update document with empty input argument",
			Args: []string{
				"",
				"TestProof",
				"ExtraData",
				network.SkipConfirmation,
				network.JSONFlag,
				network.SyncBroadcast,
				network.MakeByAccount(network.KeyDocIssuer),
			},
			ExpectErr:    true,
			ExpectedCode: 0,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdUpdateDocument(), s.network.Validators[0])
}
