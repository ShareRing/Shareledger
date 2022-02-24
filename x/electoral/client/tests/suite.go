package tests

import (
	"github.com/cosmos/cosmos-sdk/testutil/network"
	netutilts "github.com/sharering/shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
	"os"
)

type ElectoralIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
	dir     string
}

func NewElectoralIntegrationTestSuite(cf network.Config) *ElectoralIntegrationTestSuite {
	return &ElectoralIntegrationTestSuite{
		cfg: cf,
	}
}

func (s *ElectoralIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for document module")
	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//override the keyring by our keyring information
	s.network.Validators[0].ClientCtx.Keyring = kb

	s.T().Log("setting up user operator")
	out, _ := ExCmdEnrollAccountOperator(s.network.Validators[0].ClientCtx, s.T(),
		[]string{netutilts.Accounts[netutilts.KeyOperator].String()},
		netutilts.SHRFee2(),
		netutilts.MakeByAccount(netutilts.KeyAuthority),
		netutilts.SkipConfirmation(),
		netutilts.BlockBroadcast(),
	)
	_ = s.network.WaitForNextBlock()
	txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equal(netutilts.ShareLedgerSuccessCode, txResponse.Code, "%s", out.String())
	s.T().Log("setting up integration test suite successfully")

}
func (s *ElectoralIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "fail to cleanup")
	s.T().Log("tearing down integration test suite")
}

func (s *ElectoralIntegrationTestSuite) TestAccountOperator() {
	validatorCtx := s.network.Validators[0].ClientCtx
	s.Run("enroll_account_operator", func() {
		out, err := ExCmdEnrollAccountOperator(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyAccount1].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyAuthority),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerSuccessCode, txResponse.Code)

		accOut, err := ExCmdQueryAccountOperator(validatorCtx, s.T(), netutilts.Accounts[netutilts.KeyAccount1].String(), netutilts.JSONFlag())
		s.NoErrorf(err, "query account operator should not error %v", err)
		acc := JsonAccountOperatorUnmarshal(s.T(), accOut.Bytes())
		s.Equal("accop"+netutilts.Accounts[netutilts.KeyAccount1].String(), acc.GetAccState().Key, "account operator key no equal")
		s.Equal(netutilts.Accounts[netutilts.KeyAccount1].String(), acc.GetAccState().Address, "account operator address no equal")
		s.Equal("active", acc.GetAccState().Status, "account operator status no equal")
	})
	s.Run("enroll_account_operator_but_not_authority", func() {
		out, err := ExCmdEnrollAccountOperator(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyEmpty1].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyAccount3),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerErrorCodeUnauthorized, txResponse.Code)
	})

	s.Run("revoke_account_operator", func() {
		out, err := ExCmdRevokeAccountOperator(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyAccount1].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyAuthority),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerSuccessCode, txResponse.Code)
		out, err = ExCmdQueryAccountOperator(validatorCtx, s.T(), netutilts.Accounts[netutilts.KeyAccount1].String(), netutilts.JSONFlag())
		s.NoErrorf(err, "error should nil %v", err)
		acc := JsonAccountOperatorUnmarshal(s.T(), out.Bytes())
		s.Equalf("inactive", acc.GetAccState().Status, "account operator status no equal")
	})

	s.Run("revoke_account_operator_but_not_authorizer", func() {
		out, err := ExCmdRevokeAccountOperator(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyOperator].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyAccount1),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerErrorCodeUnauthorized, txResponse.Code)
		out, err = ExCmdQueryAccountOperator(validatorCtx, s.T(), netutilts.Accounts[netutilts.KeyOperator].String(), netutilts.JSONFlag())
		s.NoErrorf(err, "error should nil %v", err)
		acc := JsonAccountOperatorUnmarshal(s.T(), out.Bytes())
		s.Equalf("active", acc.GetAccState().Status, "account operator status no equal")
	})

}

func (s *ElectoralIntegrationTestSuite) TestDocIssuer() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("enroll_doc_issuer", func() {
		out, err := ExCmdEnrollDocIssuer(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyDocIssuer].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyOperator),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())

		s.Equalf(netutilts.ShareLedgerSuccessCode, txResponse.Code, "the txn response %v", txResponse.String())

		accOut, err := ExCmdGetDocIssuer(validatorCtx, s.T(), netutilts.Accounts[netutilts.KeyDocIssuer].String(), netutilts.JSONFlag())
		s.NoErrorf(err, "query doc issuer should not error %v", accOut.String())
		acc := JsonDocIssuerUnmarshal(s.T(), accOut.Bytes())
		s.Equal("docIssuer"+netutilts.Accounts[netutilts.KeyDocIssuer].String(), acc.GetAccState().Key, "doc issuer key no equal")
		s.Equal(netutilts.Accounts[netutilts.KeyDocIssuer].String(), acc.GetAccState().Address, "doc issuer address no equal")
		s.Equal("active", acc.GetAccState().Status, "doc issuer status no equal")
	})
	s.Run("enroll_doc_issuer_but_not_authority", func() {
		out, err := ExCmdEnrollDocIssuer(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyEmpty1].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyAccount3),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll doc issuer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerErrorCodeUnauthorized, txResponse.Code)
	})

	s.Run("revoke_doc_issuer_not_authorizer", func() {
		out, err := ExCmdRevokeDocIssuer(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyDocIssuer].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyAccount4),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "revoke doc issuer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerErrorCodeUnauthorized, txResponse.Code)

	})

	s.Run("revoke_doc_issuer", func() {
		out, err := ExCmdRevokeDocIssuer(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyDocIssuer].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyOperator),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "revoke_doc_issuer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerSuccessCode, txResponse.Code)
		out, err = ExCmdGetDocIssuer(validatorCtx, s.T(), netutilts.Accounts[netutilts.KeyDocIssuer].String(), netutilts.JSONFlag())
		s.NoErrorf(err, "error should nil %v", err)
		acc := JsonDocIssuerUnmarshal(s.T(), out.Bytes())
		s.Equalf("inactive", acc.GetAccState().Status, "account operator status no equal")
	})
}

func (s *ElectoralIntegrationTestSuite) TestIdSigner() {
	validatorCtx := s.network.Validators[0].ClientCtx
	accOut, err := ExCmdQueryAccountOperator(validatorCtx, s.T(), netutilts.Accounts[netutilts.KeyOperator].String(), netutilts.JSONFlag())
	s.NoErrorf(err, "query account operator should not error %v", err)
	acc := JsonAccountOperatorUnmarshal(s.T(), accOut.Bytes())
	s.T().Logf("account operator %s", acc.String())
	s.Run("enroll_id_signer", func() {
		out, err := ExCmdEnrollIdSigner(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyIDSigner].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyOperator),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll id signer operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())

		s.Equalf(netutilts.ShareLedgerSuccessCode, txResponse.Code, "the txn response %v", txResponse.String())

		accOut, err := ExCmdGetIdSigner(validatorCtx, s.T(), netutilts.Accounts[netutilts.KeyIDSigner].String(), netutilts.JSONFlag())
		s.NoErrorf(err, "query id signer should not error %v", accOut.String())
		acc := JsonIDSignerUnmarshal(s.T(), accOut.Bytes())
		s.Equal("idsigner"+netutilts.Accounts[netutilts.KeyIDSigner].String(), acc.GetAccState().Key, "id signer key no equal")
		s.Equal(netutilts.Accounts[netutilts.KeyIDSigner].String(), acc.GetAccState().Address, "id signer address no equal")
		s.Equal("active", acc.GetAccState().Status, "id signer status no equal")
	})
	s.Run("enroll_id_singer_but_not_authority", func() {
		out, err := ExCmdEnrollIdSigner(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyEmpty1].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyAccount3),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll id signer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerErrorCodeUnauthorized, txResponse.Code)
	})

	s.Run("revoke_id_signer_not_authorizer", func() {
		out, err := ExCmdRevokeIdSigner(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyIDSigner].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyAccount4),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "revoke doc issuer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerErrorCodeUnauthorized, txResponse.Code)

	})

	s.Run("revoke_id_signer", func() {
		out, err := ExCmdRevokeIdSigner(validatorCtx, s.T(),
			[]string{netutilts.Accounts[netutilts.KeyIDSigner].String()},
			netutilts.SHRFee2(),
			netutilts.MakeByAccount(netutilts.KeyOperator),
			netutilts.SkipConfirmation(),
			netutilts.BlockBroadcast(),
		)
		s.NoErrorf(err, "revoke id signer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(netutilts.ShareLedgerSuccessCode, txResponse.Code, "the response %v", out.String())
		out, err = ExCmdGetIdSigner(validatorCtx, s.T(), netutilts.Accounts[netutilts.KeyIDSigner].String(), netutilts.JSONFlag())
		s.NoErrorf(err, "error should nil %v", err)
		acc := JsonIDSignerUnmarshal(s.T(), out.Bytes())
		s.Equalf("inactive", acc.GetAccState().Status, "id signer status no equal")
	})
}
