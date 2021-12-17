package tests

import (
	"github.com/ShareRing/Shareledger/testutil/network"
	documenttypes "github.com/ShareRing/Shareledger/x/document/types"
	"github.com/ShareRing/Shareledger/x/electoral/client/tests"
	"github.com/stretchr/testify/suite"
	"strings"
)

type DocumentIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network


}

func NewDocumentIntegrationTestSuite(cf network.Config) *DocumentIntegrationTestSuite {
	return &DocumentIntegrationTestSuite{
		cfg: cf,
	}
}

func (s *DocumentIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for booking module")

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.T().Log("setting up document data....")

	out :=tests.ExCmdEnrollAccountOperator(
		s.network.Validators[0].ClientCtx,
		s.T(),
		s.network.Accounts[network.KeyOperator].String(),
		network.MakeByAccount(network.KeyAuthority),
		network.SkipConfirmation(),
		network.BlockBroadcast(),
		network.SHRFee2(),
		)
	_=s.network.WaitForNextBlock()
	res :=network.ParseStdOut(s.T(),out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode,res.Code,"init operator fail %v",res.String())

	out =tests.ExCmdEnrollDocIssuer(
		s.network.Validators[0].ClientCtx,
		s.T(),
		s.network.Accounts[network.KeyAccount1].String(),
		network.SHRFee2(),
		network.MakeByAccount(network.KeyOperator),
		network.SkipConfirmation(),
		)
	_=s.network.WaitForNextBlock()
	res =network.ParseStdOut(s.T(),out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode,res.Code,"init doc issuer fail %v",res.String())

	s.T().Log("setting up integration test suite successfully")
}
func (s *DocumentIntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
}

func (s *DocumentIntegrationTestSuite)TestCreateDocument() {

	validationCtx := s.network.Validators[0].ClientCtx

	s.Run("create_document_success_by_using_authority_account", func() {

		holderId := "id-001"
		proof := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
		data := "extradata-001"


		out := CmdExCreateDocument(validationCtx,
			s.T(), holderId, proof, data,
			network.MakeByAccount(network.KeyAccount1),
			network.BlockBroadcast(),
			network.SHRFee2(),
			network.SkipConfirmation(),
		)

		txnResponse := network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode,txnResponse.Code,"create document fail %v",out.String())
		_=s.network.WaitForNextBlock()
		queryDocResponse := CmdExGetDocByProof(validationCtx,s.T(),proof,network.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(holderId,docResponse.GetHolder(),"holder ID isn't equal")
		s.Equalf(proof,docResponse.GetProof(),"proof isn't equal")
		s.Equalf(data,docResponse.GetData(),"data ID isn't equal")

	})

	s.Run("create_document_but_duplicated_the_tx_must_be_fail", func() {
		holderId := "id-003"
		proof := "c89efdaa54c0f20c7adf612882rr0950f5arfe637e0307cdcb4c672f298b8bc6"
		data := "extradata-003"


		out := CmdExCreateDocument(validationCtx,
			s.T(),holderId,proof,data ,
			network.MakeByAccount(network.KeyAccount1),
			network.BlockBroadcast(),
			network.SHRFee2(),
			network.SkipConfirmation(),
		)

		txnResponse := network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode,txnResponse.Code,"create document fail %v",out.String())
		_=s.network.WaitForNextBlock()
		queryDocResponse := CmdExGetDocByProof(validationCtx,s.T(),proof,network.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(holderId,docResponse.GetHolder(),"holder ID isn't equal")
		s.Equalf(proof,docResponse.GetProof(),"proof isn't equal")
		s.Equalf(data,docResponse.GetData(),"data ID isn't equal")

		out = CmdExCreateDocument(validationCtx,
			s.T(),holderId,proof,data ,
			network.MakeByAccount(network.KeyAccount1),
			network.BlockBroadcast(),
			network.SHRFee2(),
			network.SkipConfirmation(),
		)

		txnResponse = network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerDocumentDuplicated,txnResponse.Code,"create duplicate asset fail %v",out.String())
		_=s.network.WaitForNextBlock()
	})


}


func (s *DocumentIntegrationTestSuite)TestCreateBatchDocument() {

	validationCtx := s.network.Validators[0].ClientCtx
	s.Run("creat_batch_document_success_by_using_authority_account", func() {
		holderId1 := "id-004"
		holderId2 := "id-005"
		holderId3 := "id-006"

		proof1 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
		proof2 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc7"
		proof3 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc8"

		data1:= "data-1"
		data2:= "data-2"
		data3:= "data-3"

		holderId := strings.Join([]string{holderId1,holderId2,holderId3},",")
		proof := strings.Join([]string{proof1,proof2,proof3},",")
		data := strings.Join([]string{data1,data2,data3},",")


		out := CmdExCreateDocumentInBatch(validationCtx,
			s.T(), holderId, proof, data,
			network.MakeByAccount(network.KeyAccount1),
			network.BlockBroadcast(),
			network.SHRFee2(),
			network.SkipConfirmation(),
		)

		txnResponse := network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode,txnResponse.Code,"create document fail %v",out.String())
		_=s.network.WaitForNextBlock()
		//doc-1
		queryDocResponse := CmdExGetDocByProof(validationCtx,s.T(),proof1,network.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(holderId1,docResponse.GetHolder(),"holder ID isn't equal")
		s.Equalf(proof1,docResponse.GetProof(),"proof isn't equal")
		s.Equalf(data1,docResponse.GetData(),"data ID isn't equal")

		//doc-2
		queryDocResponse = CmdExGetDocByProof(validationCtx,s.T(),proof2,network.JSONFlag())
		docResponse = queryDocResponse.GetDocument()
		s.Equalf(holderId2,docResponse.GetHolder(),"holder ID isn't equal")
		s.Equalf(proof2,docResponse.GetProof(),"proof isn't equal")
		s.Equalf(data2,docResponse.GetData(),"data ID isn't equal")


		//doc-2
		queryDocResponse = CmdExGetDocByProof(validationCtx,s.T(),proof3,network.JSONFlag())
		docResponse = queryDocResponse.GetDocument()
		s.Equalf(holderId3,docResponse.GetHolder(),"holder ID isn't equal")
		s.Equalf(proof3,docResponse.GetProof(),"proof isn't equal")
		s.Equalf(data3,docResponse.GetData(),"data ID isn't equal")


	})

	s.Run("create_document_in_batch_but_duplicated_the_tx_must_be_fail", func() {
		holderId1 := "id-004"
		holderId2 := "id-0011"
		holderId3 := "id-0012"

		proof1 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
		proof2 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8ba2"
		proof3 := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bgg"

		data1:= "data-1"
		data2:= "data-2"
		data3:= "data-3"

		holderId := strings.Join([]string{holderId1,holderId2,holderId3},",")
		proof := strings.Join([]string{proof1,proof2,proof3},",")
		data := strings.Join([]string{data1,data2,data3},",")


		out := CmdExCreateDocumentInBatch(validationCtx,
			s.T(), holderId, proof, data,
			network.MakeByAccount(network.KeyAccount1),
			network.BlockBroadcast(),
			network.SHRFee2(),
			network.SkipConfirmation(),
		)

		txnResponse := network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerDocumentDuplicated,txnResponse.Code,"create document fail %v",out.String())
		_=s.network.WaitForNextBlock()
	})


}


func (s *DocumentIntegrationTestSuite)TestUpdateDocument() {

	validationCtx := s.network.Validators[0].ClientCtx


	holderId1 := "id-100"
	holderId2 := "id-200"
	holderId3 := "id-300"

	proof1 := "proof-12232"
	proof2 := "proof-12233"
	proof3 := "proof-12234"

	data1:= "data-1"
	data2:= "data-2"
	data3:= "data-3"

	holderId := strings.Join([]string{holderId1,holderId2,holderId3},",")
	proof := strings.Join([]string{proof1,proof2,proof3},",")
	data := strings.Join([]string{data1,data2,data3},",")


	out := CmdExCreateDocumentInBatch(validationCtx,
		s.T(), holderId, proof, data,
		network.MakeByAccount(network.KeyAccount1),
		network.BlockBroadcast(),
		network.SHRFee2(),
		network.SkipConfirmation(),
	)

	txnResponse := network.ParseStdOut(s.T(),out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode,txnResponse.Code,"create document fail %v",out.String())
	_=s.network.WaitForNextBlock()


	queryDocResponse := CmdExGetDocByProof(validationCtx,s.T(),proof1,network.JSONFlag())
	docVersion := queryDocResponse.GetDocument().GetVersion()

	s.Run("update_document_success", func() {
		out =CmdExUpdateDocument(validationCtx, s.T(),holderId1,proof1, "new-dataaaaa",
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount1),
			network.SHRFee2(),
			network.BlockBroadcast(),
			network.JSONFlag(),
		)
		txnResponse = network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode,txnResponse.Code,"update document fail %v",out.String())
		_=s.network.WaitForNextBlock()

		//doc-2
		queryDocResponse := CmdExGetDocByProof(validationCtx,s.T(),proof1,network.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(holderId1,docResponse.GetHolder(),"holder ID isn't equal")
		s.Equalf(proof1,docResponse.GetProof(),"proof isn't equal")
		s.Equalf("new-dataaaaa",docResponse.GetData(),"data ID isn't equal")
		s.Equalf(docVersion+1,docResponse.GetVersion(),"doc version isn't equal")
	})

	s.Run("update_document_with_not_exist_hold_id_should_be_fail", func() {
		out =CmdExUpdateDocument(validationCtx, s.T(),holderId1+"hi there",proof1, "new-dataaaaa",
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount1),
			network.SHRFee2(),
			network.BlockBroadcast(),
			network.JSONFlag(),
		)
		txnResponse := network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerDocumentNotFound,txnResponse.Code,"update recheck status %v",out.String())
		_=s.network.WaitForNextBlock()
	})

	s.Run("update_document_with_not_exist_proof_should_be_fail", func() {
		out =CmdExUpdateDocument(validationCtx, s.T(),holderId1,proof1+"hi there", "new-dataaaaa",
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount1),
			network.SHRFee2(),
			network.BlockBroadcast(),
			network.JSONFlag(),
		)
		txnResponse := network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerDocumentNotFound,txnResponse.Code,"update recheck status %v",out.String())
		_=s.network.WaitForNextBlock()
	})

	s.Run("update_document_holderId_should_be_fail", func() {
		out =CmdExUpdateDocument(validationCtx, s.T(),holderId2,proof1, "new-dataaaaa_with_new_holder",
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount1),
			network.SHRFee2(),
			network.BlockBroadcast(),
			network.JSONFlag(),
		)
		txnResponse := network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerDocumentNotFound,txnResponse.Code,"update recheck status %v",out.String())
		_=s.network.WaitForNextBlock()
	})

	s.Run("update_document_by_another_doc_issuer", func() {
		out =CmdExUpdateDocument(validationCtx, s.T(),holderId1,proof1, "new-dataaaaa_with_new_holder",
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount2),
			network.SHRFee2(),
			network.BlockBroadcast(),
			network.JSONFlag(),
		)
		txnResponse := network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerErrorCodeUnauthorized,txnResponse.Code,"update recheck status %v",out.String())
		_=s.network.WaitForNextBlock()
	})
}

func (s *DocumentIntegrationTestSuite)TestRevokeDocument() {

	validationCtx := s.network.Validators[0].ClientCtx


	holderIdForRevoke1 := "rv-101"
	proofRevoke1 := "rv-proof-12232"
	dataRevoke1:= "rv-data-1"

	holderIdForRevoke2 := "rv-102"
	proofRevoke2 := "rv-proof-12233"
	dataRevoke2:= "rv-data-2"


	out := CmdExCreateDocument(validationCtx,
		s.T(), holderIdForRevoke1, proofRevoke1, dataRevoke1,
		network.MakeByAccount(network.KeyAccount1),
		network.BlockBroadcast(),
		network.SHRFee2(),
		network.SkipConfirmation(),
	)

	txnResponse := network.ParseStdOut(s.T(),out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode,txnResponse.Code,"create document fail %v",out.String())
	_=s.network.WaitForNextBlock()


	out = CmdExCreateDocument(validationCtx,
		s.T(), holderIdForRevoke2, proofRevoke2, dataRevoke2,
		network.MakeByAccount(network.KeyAccount1),
		network.BlockBroadcast(),
		network.SHRFee2(),
		network.SkipConfirmation(),
	)

	txnResponse = network.ParseStdOut(s.T(),out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode,txnResponse.Code,"create document fail %v",out.String())
	_=s.network.WaitForNextBlock()

	s.Run("revoke_document_success", func() {
		out =CmdExRevokeDocument(validationCtx, s.T(),holderIdForRevoke1,proofRevoke1,
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount1),
			network.SHRFee2(),
			network.BlockBroadcast(),
			network.JSONFlag(),
		)
		txnResponse = network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode,txnResponse.Code,"revoke document fail %v",out.String())
		_=s.network.WaitForNextBlock()

		//doc-2
		queryDocResponse := CmdExGetDocByProof(validationCtx,s.T(),proofRevoke1,network.JSONFlag())
		docResponse := queryDocResponse.GetDocument()
		s.Equalf(int32(documenttypes.DocRevokeFlag),docResponse.GetVersion(),"doc version isn't equal")
	})

	s.Run("revoke_document_not_exist_should_be_fail", func() {
		out =CmdExRevokeDocument(validationCtx, s.T(),holderIdForRevoke1,proofRevoke1+"hi there",
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount1),
			network.SHRFee2(),
			network.BlockBroadcast(),
			network.JSONFlag(),
		)
		txnResponse := network.ParseStdOut(s.T(),out.Bytes())
		s.Equalf(network.ShareLedgerDocumentNotFound,txnResponse.Code,"update recheck status %v",out.String())
		_=s.network.WaitForNextBlock()
	})

}