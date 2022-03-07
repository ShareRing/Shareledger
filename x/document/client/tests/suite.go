package tests

import (
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/types"
	netutilts "github.com/sharering/shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
	"os"
	"strings"

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
func (s *DocumentIntegrationTestSuite) setupTestingMaterial() {

	initIDSigAndDocIssuer := []struct {
		id        string
		accountID string
		idData    string
		docProof  string
		docData   string
	}{
		{
			id:        "id_1",
			accountID: netutilts.KeyAccount1,
			idData:    "extra_1",
		},
		{
			id:        "id_2",
			accountID: netutilts.KeyAccount2,
			idData:    "extra_2",
		},
		{
			id:        "id_3",
			accountID: netutilts.KeyAccount3,
			idData:    "extra_3",
		},
		{
			id:        "id_4",
			accountID: netutilts.KeyAccount4,
			idData:    "extra_4",
			docData:   "doc_data",
			docProof:  "proofff1",
		},
		{
			id:        "id_5",
			accountID: netutilts.KeyAccount5,
			idData:    "extra_5_update",
			docData:   "doc_for_update_1",
			docProof:  "proof_doc_for_update_1",
		},
		{
			id:        "id_6",
			accountID: netutilts.KeyAccount6,
			idData:    "extra_6_update",
			docData:   "doc_for_update_2",
			docProof:  "proof_doc_for_update_2",
		},
		{
			id:        "id_7",
			accountID: netutilts.KeyAccount7,
			idData:    "extra_6_update",
			docData:   "doc_for_update_3",
			docProof:  "proof_doc_for_update_3",
		},
	}

	//Enroll ACCOUNT_OPERATOR
	out, _ := tests.ExCmdEnrollAccountOperator(
		s.network.Validators[0].ClientCtx,
		[]string{netutilts.Accounts[netutilts.KeyOperator].String()},
		netutilts.MakeByAccount(netutilts.KeyAuthority),
		netutilts.SkipConfirmation,
		netutilts.BlockBroadcast,
		netutilts.SHRFee2,
	)
	s.Require().NoError(s.network.WaitForNextBlock())
	res := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init operator fail %v", res.String())

	for _, iz := range initIDSigAndDocIssuer {
		//Enroll ID_SIGNER
		out, _ = tests.ExCmdEnrollIdSigner(
			s.network.Validators[0].ClientCtx,
			[]string{netutilts.Accounts[iz.accountID].String()},
			netutilts.SHRFee2,
			netutilts.MakeByAccount(netutilts.KeyOperator),
			netutilts.SkipConfirmation,
		)
		s.Require().NoError(s.network.WaitForNextBlock())
		res = netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init doc issuer fail %v", res.String())

		out, _ = tests.ExCmdEnrollDocIssuer(
			s.network.Validators[0].ClientCtx,
			[]string{netutilts.Accounts[iz.accountID].String()},
			netutilts.SHRFee2,
			netutilts.MakeByAccount(netutilts.KeyOperator),
			netutilts.SkipConfirmation,
		)
		s.Require().NoError(s.network.WaitForNextBlock())
		res = netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init doc issuer fail %v", res.String())

		out, _ = idtest.CmdExNewID(s.network.Validators[0].ClientCtx,
			iz.id,
			sample.AccAddress(),
			netutilts.Accounts[iz.accountID].String(),
			iz.idData,
			netutilts.MakeByAccount(iz.accountID),
			netutilts.SkipConfirmation,
			netutilts.BlockBroadcast,
			netutilts.SHRFee2)

		s.Require().NoError(s.network.WaitForNextBlock())
		res = netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init id fail %v for id %s", res.String(), iz.id)

		if strings.TrimSpace(iz.docProof) != "" && strings.TrimSpace(iz.docData) != "" {
			out, _ := CmdExCreateDocument(s.network.Validators[0].ClientCtx,
				iz.id, iz.docProof, iz.docData,
				netutilts.MakeByAccount(iz.accountID),
				netutilts.BlockBroadcast,
				netutilts.SHRFee2,
				netutilts.SkipConfirmation,
			)
			s.Require().NoError(s.network.WaitForNextBlock())
			res = netutilts.ParseStdOut(s.T(), out.Bytes())
			s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init doc %s fail %v", iz.docProof, res.String())
		}

	}

}
func (s *DocumentIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for document module")

	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//override the keyring by our keyring informatirevoke_the_document_successfullyon
	s.network.Validators[0].ClientCtx.Keyring = kb

	s.setupTestingMaterial()

	s.T().Log("setting up integration test suite successfully")

}
func (s *DocumentIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "cleanup test case fails")
	s.network.Cleanup()
	s.T().Log("tearing down integration test suite")
}

func (s *DocumentIntegrationTestSuite) TestCreateDocument() {
	type (
		TestCase struct {
			d          string
			iHolderID  string
			iDocProof  string
			iDocData   string
			txnCreator string
			txnFee     int
			oErr       error
			oRes       *types.TxResponse
			oDoc       *documenttypes.Document
		}
	)

	testSuite := []TestCase{
		{
			d:          "create_document_success_by_using_authority_account",
			iHolderID:  "id_1",
			iDocProof:  "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8sa6",
			iDocData:   "extradata-001",
			txnCreator: netutilts.KeyAccount1,
			txnFee:     2,
			oErr:       nil,
			oRes:       &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oDoc: &documenttypes.Document{
				Holder:  "id_1",
				Issuer:  netutilts.Accounts[netutilts.KeyAccount1].String(),
				Proof:   "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8sa6",
				Data:    "extradata-001",
				Version: 0,
			},
		},
		{
			d:          "create_document_but_duplicated_the_tx_must_be_fail",
			iHolderID:  "id_4",
			iDocProof:  "proofff1",
			iDocData:   "extra_4",
			txnCreator: netutilts.KeyAccount4,
			txnFee:     2,
			oErr:       nil,
			oRes:       &types.TxResponse{Code: documenttypes.ErrDocExisted.ABCICode()},
			oDoc:       nil,
		},
	}

	validationCtx := s.network.Validators[0].ClientCtx

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			out, err := CmdExCreateDocument(validationCtx,
				tc.iHolderID, tc.iDocProof, tc.iDocData,
				netutilts.MakeByAccount(tc.txnCreator),
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txnFee),
				netutilts.SkipConfirmation,
			)
			if tc.oErr != nil {
				s.NotNilf(err, "this case need got error")
			}
			if tc.oRes != nil {
				txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txnResponse.Code, "create document fail %s", txnResponse.String())
			}
			if tc.oDoc != nil {
				queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), tc.iDocProof, netutilts.JSONFlag)
				docResponse := queryDocResponse.GetDocument()
				s.Equalf(tc.oDoc.Holder, docResponse.GetHolder(), "holder ID isn't equal")
				s.Equalf(tc.oDoc.Proof, docResponse.GetProof(), "proof isn't equal")
				s.Equalf(tc.oDoc.Data, docResponse.GetData(), "data ID isn't equal")
			}

		})
	}

}

func (s *DocumentIntegrationTestSuite) TestCreateBatchDocument() {
	validationCtx := s.network.Validators[0].ClientCtx
	type (
		TestCase struct {
			d          string
			iHolderIDs []string
			iDocProofs []string
			iDocDatas  []string
			txnCreator string
			txnFee     int
			oErr       error
			oRes       *types.TxResponse
			oDocs      []documenttypes.Document
		}
	)

	testSuite := []TestCase{
		{
			d: "create_documents_success_by_using_authority_account",
			iHolderIDs: []string{
				"id_1",
				"id_2",
				"id_3",
			},
			iDocProofs: []string{
				"c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6",
				"c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc7",
				"c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc8",
			},
			iDocDatas: []string{
				"extradata_batch_001",
				"extradata_batch_002",
				"extradata_batch_003"},
			txnCreator: netutilts.KeyAccount3,
			txnFee:     2,
			oErr:       nil,
			oRes:       &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oDocs: []documenttypes.Document{
				{
					Holder:  "id_1",
					Issuer:  netutilts.Accounts[netutilts.KeyAccount3].String(),
					Proof:   "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6",
					Data:    "extradata_batch_001",
					Version: 0,
				},
				{
					Holder:  "id_2",
					Issuer:  netutilts.Accounts[netutilts.KeyAccount3].String(),
					Proof:   "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc7",
					Data:    "extradata_batch_002",
					Version: 0,
				},
				{
					Holder:  "id_3",
					Issuer:  netutilts.Accounts[netutilts.KeyAccount3].String(),
					Proof:   "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc8",
					Data:    "extradata_batch_003",
					Version: 0,
				},
			},
		},
	}
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			out, err := CmdExCreateDocumentInBatch(validationCtx,
				strings.Join(tc.iHolderIDs, ","),
				strings.Join(tc.iDocProofs, ","),
				strings.Join(tc.iDocDatas, ","),
				netutilts.MakeByAccount(tc.txnCreator),
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txnFee),
				netutilts.SkipConfirmation,
			)
			if tc.oErr != nil {
				s.NotNilf(err, "this case need got error")
			}
			if tc.oRes != nil {
				txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txnResponse.Code, "create documents fail %s", txnResponse.String())
			}
			if len(tc.oDocs) != 0 {
				for _, d := range tc.oDocs {
					queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), d.Proof, netutilts.JSONFlag)
					docResponse := queryDocResponse.GetDocument()
					s.Equalf(d.Holder, docResponse.GetHolder(), "holder ID isn't equal")
					s.Equalf(d.Proof, docResponse.GetProof(), "proof isn't equal")
					s.Equalf(d.Data, docResponse.GetData(), "data ID isn't equal")
				}

			}

		})
	}
}

func (s *DocumentIntegrationTestSuite) TestUpdateDocument() {

	validationCtx := s.network.Validators[0].ClientCtx

	type (
		TestCase struct {
			d          string
			iHolderID  string
			iDocProof  string
			iDocData   string
			txnCreator string
			txnFee     int
			oErr       error
			oRes       *types.TxResponse
			oDoc       *documenttypes.Document
		}
	)

	testSuite := []TestCase{
		{
			d:          "update the document success",
			iHolderID:  "id_5",
			iDocProof:  "proof_doc_for_update_1",
			iDocData:   "doc_for_update_3_newwww-001",
			txnCreator: netutilts.KeyAccount5,
			txnFee:     2,
			oErr:       nil,
			oRes:       &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oDoc: &documenttypes.Document{
				Holder:  "id_5",
				Issuer:  netutilts.Accounts[netutilts.KeyAccount5].String(),
				Proof:   "proof_doc_for_update_1",
				Data:    "doc_for_update_3_newwww-001",
				Version: 1,
			},
		},
		{
			d:          "update_document_with_not_exist_hold_id_should_be_fail",
			iHolderID:  "id_5",
			iDocProof:  "proof_doc_for_update_1+hithere",
			iDocData:   "doc_for_update_3_newwww",
			txnCreator: netutilts.KeyAccount5,
			txnFee:     2,
			oErr:       nil,
			oRes:       &types.TxResponse{Code: documenttypes.ErrDocNotExisted.ABCICode()},
		},
		{
			d:          "update the document but by not authorize",
			iHolderID:  "id_7",
			iDocProof:  "proof_doc_for_update_3",
			iDocData:   "doc_for_update_3",
			txnCreator: netutilts.KeyAccount6,
			txnFee:     2,
			oErr:       nil,
			oRes:       &types.TxResponse{Code: documenttypes.ErrDocNotExisted.ABCICode()},
			oDoc: &documenttypes.Document{
				Holder:  "id_7",
				Issuer:  netutilts.Accounts[netutilts.KeyAccount5].String(),
				Proof:   "proof_doc_for_update_3",
				Data:    "doc_for_update_3",
				Version: 0,
			},
		},
	}

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			out, err := CmdExUpdateDocument(validationCtx,
				tc.iHolderID, tc.iDocProof, tc.iDocData,
				netutilts.MakeByAccount(tc.txnCreator),
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txnFee),
				netutilts.SkipConfirmation,
			)
			if tc.oErr != nil {
				s.NotNilf(err, "this case need got error")
			}
			if tc.oRes != nil {
				txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txnResponse.Code, "update document fail %s", txnResponse.String())
			}
			if tc.oDoc != nil {
				queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), tc.iDocProof, netutilts.JSONFlag)
				docResponse := queryDocResponse.GetDocument()
				s.Equalf(tc.oDoc.Holder, docResponse.GetHolder(), "holder ID isn't equal")
				s.Equalf(tc.oDoc.Proof, docResponse.GetProof(), "proof isn't equal")
				s.Equalf(tc.oDoc.Data, docResponse.GetData(), "data ID isn't equal")
			}

		})
	}

}

func (s *DocumentIntegrationTestSuite) TestRevokeDocument() {
	validationCtx := s.network.Validators[0].ClientCtx

	type (
		TestCase struct {
			d          string
			iHolderID  string
			iDocProof  string
			txnCreator string
			txnFee     int
			oErr       error
			oRes       *types.TxResponse
			oDoc       *documenttypes.Document
		}
	)

	testSuite := []TestCase{
		{
			d:          "revoke the document successfully",
			iHolderID:  "id_1",
			iDocProof:  "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8sa6",
			txnCreator: netutilts.KeyAccount1,
			txnFee:     2,
			oErr:       nil,
			oRes:       &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oDoc: &documenttypes.Document{
				Version: documenttypes.DocRevokeFlag,
			},
		},
		{
			d:          "revoke the document fail",
			iHolderID:  "id_33424",
			iDocProof:  "c89efdaa54c0f20c7adf612882df0950f5a958+not found",
			txnCreator: netutilts.KeyAccount1,
			txnFee:     2,
			oErr:       nil,
			oRes:       &types.TxResponse{Code: documenttypes.ErrDocNotExisted.ABCICode()},
		},
	}

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			out, err := CmdExRevokeDocument(validationCtx,
				tc.iHolderID, tc.iDocProof,
				netutilts.MakeByAccount(tc.txnCreator),
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txnFee),
				netutilts.SkipConfirmation,
			)
			if tc.oErr != nil {
				s.NotNilf(err, "this case need got error")
			}
			if tc.oRes != nil {
				txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txnResponse.Code, "revoke document fail %s", txnResponse.String())
			}
			if tc.oDoc != nil {
				queryDocResponse := CmdExGetDocByProof(validationCtx, s.T(), tc.iDocProof, netutilts.JSONFlag)
				docResponse := queryDocResponse.GetDocument()
				s.Equalf(tc.oDoc.Version, docResponse.GetVersion(), "version isn't equal")
			}

		})
	}

}
