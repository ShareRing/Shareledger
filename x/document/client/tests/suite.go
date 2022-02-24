package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	netutilts "github.com/sharering/shareledger/testutil/network"
	"os"
	"strings"

	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/sample"
	documenttypes "github.com/sharering/shareledger/x/document/types"
	"github.com/sharering/shareledger/x/electoral/client/tests"
	idtest "github.com/sharering/shareledger/x/id/client/tests"
)

type DocumentIntegrationTestSuite struct {
	suite.Suite

	cfg         network.Config
	network     *network.Network
	dir         string
	userID1     string
	userID1Addr string
	userID2     string
	userID2Addr string
	userID3     string
	userID3Addr string
}

func NewDocumentIntegrationTestSuite(cf network.Config) *DocumentIntegrationTestSuite {
	return &DocumentIntegrationTestSuite{
		cfg: cf,
	}
}

func (s *DocumentIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for document module")

	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//override the keyring by our keyring information
	s.network.Validators[0].ClientCtx.Keyring = kb

	s.T().Log("setting up document data....")

	//Enroll ACCOUNT_OPERATOR
	out, _ := tests.ExCmdEnrollAccountOperator(
		s.network.Validators[0].ClientCtx,
		s.T(),
		[]string{netutilts.Accounts[netutilts.KeyOperator].String()},
		netutilts.MakeByAccount(netutilts.KeyAuthority),
		netutilts.SkipConfirmation(),
		netutilts.BlockBroadcast(),
		netutilts.SHRFee2(),
	)
	_ = s.network.WaitForNextBlock()
	res := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init operator fail %v", res.String())

	//Enroll DOC_ISSUER
	out, _ = tests.ExCmdEnrollDocIssuer(
		s.network.Validators[0].ClientCtx,
		s.T(),
		[]string{netutilts.Accounts[netutilts.KeyAccount1].String()},
		netutilts.SHRFee2(),
		netutilts.MakeByAccount(netutilts.KeyOperator),
		netutilts.SkipConfirmation(),
	)
	_ = s.network.WaitForNextBlock()
	res = netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init doc issuer fail %v", res.String())

	//Enroll ID_SIGNER
	out, _ = tests.ExCmdEnrollIdSigner(
		s.network.Validators[0].ClientCtx,
		s.T(),
		[]string{netutilts.Accounts[netutilts.KeyAccount1].String()},
		netutilts.SHRFee2(),
		netutilts.MakeByAccount(netutilts.KeyOperator),
		netutilts.SkipConfirmation(),
	)
	_ = s.network.WaitForNextBlock()
	res = netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init doc issuer fail %v", res.String())

	s.T().Log("create three ID")

	_, _, addrs := sample.RandomAddr(3)

	s.userID1 = "id-1"
	s.userID2 = "id-2"
	s.userID3 = "id-3"

	out = idtest.CmdExNewIDInBatch(s.network.Validators[0].ClientCtx, s.T(), fmt.Sprintf("%s,%s,%s", s.userID1, s.userID2, s.userID3),
		strings.Join(addrs, ","),
		strings.Join(addrs, ","),
		fmt.Sprintf("%s,%s,%s", "hello1", "hello2", "hello3"),
		netutilts.MakeByAccount(netutilts.KeyAccount1),
		netutilts.SkipConfirmation(),
		netutilts.SHRFee2())

	_ = s.network.WaitForNextBlock()
	res = netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init id fail %v", res.String())
	_ = s.network.WaitForNextBlock()

	s.T().Log("setting up integration test suite successfully")

}
func (s *DocumentIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "cleanup test case fails")
	s.T().Log("tearing down integration test suite")
}

func (s *DocumentIntegrationTestSuite) TestCreateDocument() {

	validationCtx := s.network.Validators[0].ClientCtx

	s.Run("create_document_success_by_using_authority_account", func() {

		proof := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8sa6"
		data := "extradata-001"

		out := CmdExCreateDocument(validationCtx,
			s.T(), s.userID1, proof, data,
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.BlockBroadcast(),
			netutilts.SHRFee2(),
			netutilts.SkipConfirmation(),
		)

		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "create document fail %v", out.String())
		_ = s.network.WaitForNextBlock()
		queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), proof, netutilts.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(s.userID1, docResponse.GetHolder(), "holder ID isn't equal")
		s.Equalf(proof, docResponse.GetProof(), "proof isn't equal")
		s.Equalf(data, docResponse.GetData(), "data ID isn't equal")

	})

	s.Run("create_document_but_duplicated_the_tx_must_be_fail", func() {
		holderId := "id-3"
		proof := "c89efdaa54c0f20c7adf612882rr0950f5arfe637e0307cdcb4c672f298b8bc6"
		data := "extradata-003"

		out := CmdExCreateDocument(validationCtx,
			s.T(), holderId, proof, data,
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.BlockBroadcast(),
			netutilts.SHRFee2(),
			netutilts.SkipConfirmation(),
		)

		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "create document fail %v", out.String())
		_ = s.network.WaitForNextBlock()
		queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), proof, netutilts.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(holderId, docResponse.GetHolder(), "holder ID isn't equal")
		s.Equalf(proof, docResponse.GetProof(), "proof isn't equal")
		s.Equalf(data, docResponse.GetData(), "data ID isn't equal")

		out = CmdExCreateDocument(validationCtx,
			s.T(), holderId, proof, data,
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.BlockBroadcast(),
			netutilts.SHRFee2(),
			netutilts.SkipConfirmation(),
		)

		txnResponse = netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerErrorCodeDocumentAlreadyExisted, txnResponse.Code, "create duplicate asset fail %v", out.String())
		_ = s.network.WaitForNextBlock()
	})

}

func (s *DocumentIntegrationTestSuite) TestCreateBatchDocument() {

	validationCtx := s.network.Validators[0].ClientCtx
	s.Run("creat_batch_document_success_by_using_authority_account", func() {
		holderId1 := "id-1"
		holderId2 := "id-2"
		holderId3 := "id-3"

		proof1 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
		proof2 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc7"
		proof3 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc8"

		data1 := "data-1"
		data2 := "data-2"
		data3 := "data-3"

		holderId := strings.Join([]string{holderId1, holderId2, holderId3}, ",")
		proof := strings.Join([]string{proof1, proof2, proof3}, ",")
		data := strings.Join([]string{data1, data2, data3}, ",")

		out := CmdExCreateDocumentInBatch(validationCtx,
			s.T(), holderId, proof, data,
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.BlockBroadcast(),
			netutilts.SHRFee2(),
			netutilts.SkipConfirmation(),
		)

		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "create document fail %v", out.String())
		_ = s.network.WaitForNextBlock()
		//doc-1
		queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), proof1, netutilts.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(holderId1, docResponse.GetHolder(), "holder ID isn't equal")
		s.Equalf(proof1, docResponse.GetProof(), "proof isn't equal")
		s.Equalf(data1, docResponse.GetData(), "data ID isn't equal")

		//doc-2
		queryDocResponse = CmdExGetDocByProof(validationCtx, s.T(), proof2, netutilts.JSONFlag())
		docResponse = queryDocResponse.GetDocument()
		s.Equalf(holderId2, docResponse.GetHolder(), "holder ID isn't equal")
		s.Equalf(proof2, docResponse.GetProof(), "proof isn't equal")
		s.Equalf(data2, docResponse.GetData(), "data ID isn't equal")

		//doc-2
		queryDocResponse = CmdExGetDocByProof(validationCtx, s.T(), proof3, netutilts.JSONFlag())
		docResponse = queryDocResponse.GetDocument()
		s.Equalf(holderId3, docResponse.GetHolder(), "holder ID isn't equal")
		s.Equalf(proof3, docResponse.GetProof(), "proof isn't equal")
		s.Equalf(data3, docResponse.GetData(), "data ID isn't equal")

	})

	s.Run("create_document_in_batch_but_duplicated_the_tx_must_be_fail", func() {
		holderId1 := "id-1"
		holderId2 := "id-2"
		holderId3 := "id-2"

		proof1 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
		proof2 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
		proof3 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bgg"

		data1 := "data-1"
		data2 := "data-2"
		data3 := "data-3"

		holderId := strings.Join([]string{holderId1, holderId2, holderId3}, ",")
		proof := strings.Join([]string{proof1, proof2, proof3}, ",")
		data := strings.Join([]string{data1, data2, data3}, ",")

		out := CmdExCreateDocumentInBatch(validationCtx,
			s.T(), holderId, proof, data,
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.BlockBroadcast(),
			netutilts.SHRFee2(),
			netutilts.SkipConfirmation(),
		)

		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.NotEqual(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "create document fail %v", out.String())
		_ = s.network.WaitForNextBlock()
	})

}

func (s *DocumentIntegrationTestSuite) TestUpdateDocument() {

	validationCtx := s.network.Validators[0].ClientCtx

	holderId1 := "id-100"
	holderId2 := "id-200"
	holderId3 := "id-300"

	_, _, addr := sample.RandomAddr(3)

	out := idtest.CmdExNewIDInBatch(s.network.Validators[0].ClientCtx, s.T(), fmt.Sprintf("%s,%s,%s", holderId1, holderId2, holderId3),
		strings.Join(addr, ","),
		strings.Join(addr, ","),
		fmt.Sprintf("%s,%s,%s", "hello1", "hello2", "hello3"),
		netutilts.MakeByAccount(netutilts.KeyAccount1),
		netutilts.SkipConfirmation(),
		netutilts.SHRFee2())

	_ = s.network.WaitForNextBlock()
	res := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init id fail %v", res.String())
	_ = s.network.WaitForNextBlock()

	proof1 := "proof-12232"
	proof2 := "proof-12233"
	proof3 := "proof-12234"

	data1 := "data-1"
	data2 := "data-2"
	data3 := "data-3"

	holderId := strings.Join([]string{holderId1, holderId2, holderId3}, ",")
	proof := strings.Join([]string{proof1, proof2, proof3}, ",")
	data := strings.Join([]string{data1, data2, data3}, ",")

	out = CmdExCreateDocumentInBatch(validationCtx,
		s.T(), holderId, proof, data,
		netutilts.MakeByAccount(netutilts.KeyAccount1),
		netutilts.BlockBroadcast(),
		netutilts.SHRFee2(),
		netutilts.SkipConfirmation(),
	)

	txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "create document fail %v", out.String())
	_ = s.network.WaitForNextBlock()

	queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), proof1, netutilts.JSONFlag())
	docVersion := queryDocResponse.GetDocument().GetVersion()

	s.Run("update_document_success", func() {
		out = CmdExUpdateDocument(validationCtx, s.T(), holderId1, proof1, "new-dataaaaa",
			netutilts.SkipConfirmation(),
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.SHRFee2(),
			netutilts.BlockBroadcast(),
			netutilts.JSONFlag(),
		)
		txnResponse = netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "update document fail %v", out.String())
		_ = s.network.WaitForNextBlock()

		//doc-2
		queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), proof1, netutilts.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(holderId1, docResponse.GetHolder(), "holder ID isn't equal")
		s.Equalf(proof1, docResponse.GetProof(), "proof isn't equal")
		s.Equalf("new-dataaaaa", docResponse.GetData(), "data ID isn't equal")
		s.Equalf(docVersion+1, docResponse.GetVersion(), "doc version isn't equal")
	})

	s.Run("update_document_with_not_exist_hold_id_should_be_fail", func() {
		out = CmdExUpdateDocument(validationCtx, s.T(), holderId1+"hi there", proof1, "new-dataaaaa",
			netutilts.SkipConfirmation(),
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.SHRFee2(),
			netutilts.BlockBroadcast(),
			netutilts.JSONFlag(),
		)
		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerDocumentNotFound, txnResponse.Code, "update recheck status %v", out.String())
		_ = s.network.WaitForNextBlock()
	})

	s.Run("update_document_with_not_exist_proof_should_be_fail", func() {
		out = CmdExUpdateDocument(validationCtx, s.T(), holderId1, proof1+"hi there", "new-dataaaaa",
			netutilts.SkipConfirmation(),
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.SHRFee2(),
			netutilts.BlockBroadcast(),
			netutilts.JSONFlag(),
		)
		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerDocumentNotFound, txnResponse.Code, "update recheck status %v", out.String())
		_ = s.network.WaitForNextBlock()
	})

	s.Run("update_document_holderId_should_be_fail", func() {
		out = CmdExUpdateDocument(validationCtx, s.T(), holderId2, proof1, "new-dataaaaa_with_new_holder",
			netutilts.SkipConfirmation(),
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.SHRFee2(),
			netutilts.BlockBroadcast(),
			netutilts.JSONFlag(),
		)
		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerDocumentNotFound, txnResponse.Code, "update recheck status %v", out.String())
		_ = s.network.WaitForNextBlock()
	})

	s.Run("update_document_by_another_doc_issuer", func() {
		out = CmdExUpdateDocument(validationCtx, s.T(), holderId1, proof1, "new-dataaaaa_with_new_holder",
			netutilts.SkipConfirmation(),
			netutilts.MakeByAccount(netutilts.KeyAccount2),
			netutilts.SHRFee2(),
			netutilts.BlockBroadcast(),
			netutilts.JSONFlag(),
		)
		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerErrorCodeUnauthorized, txnResponse.Code, "update recheck status %v", out.String())
		_ = s.network.WaitForNextBlock()
	})
}

func (s *DocumentIntegrationTestSuite) TestRevokeDocument() {
	_, _, addrs := sample.RandomAddr(2)
	validationCtx := s.network.Validators[0].ClientCtx

	holderIdForRevoke1 := "rv-101"
	proofRevoke1 := "rv-proof-12232"
	dataRevoke1 := "rv-data-1"

	holderIdForRevoke2 := "rv-102"
	proofRevoke2 := "rv-proof-12233"
	dataRevoke2 := "rv-data-2"

	out := idtest.CmdExNewIDInBatch(s.network.Validators[0].ClientCtx, s.T(), fmt.Sprintf("%s,%s", holderIdForRevoke1, holderIdForRevoke2),
		strings.Join(addrs, ","),
		strings.Join(addrs, ","),
		fmt.Sprintf("%s,%s", "hello1", "hello2"),
		netutilts.MakeByAccount(netutilts.KeyAccount1),
		netutilts.SkipConfirmation(),
		netutilts.SHRFee2())

	_ = s.network.WaitForNextBlock()
	res := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init id fail %v", res.String())

	out = CmdExCreateDocument(validationCtx,
		s.T(), holderIdForRevoke1, proofRevoke1, dataRevoke1,
		netutilts.MakeByAccount(netutilts.KeyAccount1),
		netutilts.BlockBroadcast(),
		netutilts.SHRFee2(),
		netutilts.SkipConfirmation(),
	)

	txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "create document fail %v", out.String())
	_ = s.network.WaitForNextBlock()

	out = CmdExCreateDocument(validationCtx,
		s.T(), holderIdForRevoke2, proofRevoke2, dataRevoke2,
		netutilts.MakeByAccount(netutilts.KeyAccount1),
		netutilts.BlockBroadcast(),
		netutilts.SHRFee2(),
		netutilts.SkipConfirmation(),
	)

	txnResponse = netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "create document fail %v", out.String())
	_ = s.network.WaitForNextBlock()

	s.Run("revoke_document_success", func() {
		out = CmdExRevokeDocument(validationCtx, s.T(), holderIdForRevoke1, proofRevoke1,
			netutilts.SkipConfirmation(),
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.SHRFee2(),
			netutilts.BlockBroadcast(),
			netutilts.JSONFlag(),
		)
		txnResponse = netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txnResponse.Code, "revoke document fail %v", out.String())
		_ = s.network.WaitForNextBlock()

		//doc-2
		queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), proofRevoke1, netutilts.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(int32(documenttypes.DocRevokeFlag), docResponse.GetVersion(), "doc version isn't equal")
	})

	s.Run("revoke_document_not_exist_should_be_fail", func() {
		out = CmdExRevokeDocument(validationCtx, s.T(), holderIdForRevoke1, proofRevoke1+"hi there",
			netutilts.SkipConfirmation(),
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.SHRFee2(),
			netutilts.BlockBroadcast(),
			netutilts.JSONFlag(),
		)
		txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerDocumentNotFound, txnResponse.Code, "update recheck status %v", out.String())
		_ = s.network.WaitForNextBlock()
	})

}
