package tests

import (
	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/network"
)

type ElectoralIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewElectoralIntegrationTestSuite(cf network.Config) *ElectoralIntegrationTestSuite {
	return &ElectoralIntegrationTestSuite{
		cfg: cf,
	}
}

func (s *ElectoralIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for document module")

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.T().Log("setting up user operator")
	out, _ := ExCmdEnrollAccountOperator(s.network.Validators[0].ClientCtx, s.T(),
		[]string{s.network.Accounts[network.KeyOperator].String()},
		network.SHRFee2(),
		network.MakeByAccount(network.KeyAuthority),
		network.SkipConfirmation(),
		network.BlockBroadcast(),
	)
	_ = s.network.WaitForNextBlock()
	txResponse := network.ParseStdOut(s.T(), out.Bytes())
	s.Equal(network.ShareLedgerSuccessCode, txResponse.Code, "%s", out.String())
	s.T().Log("setting up integration test suite successfully")

}
func (s *ElectoralIntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
}

func (s *ElectoralIntegrationTestSuite) TestAccountOperator() {
	validatorCtx := s.network.Validators[0].ClientCtx
	s.Run("enroll_account_operator", func() {
		out, err := ExCmdEnrollAccountOperator(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyAccount1].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyAuthority),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerSuccessCode, txResponse.Code)

		accOut, err := ExCmdQueryAccountOperator(validatorCtx, s.T(), s.network.Accounts[network.KeyAccount1].String(), network.JSONFlag())
		s.NoErrorf(err, "query account operator should not error %v", err)
		acc := JsonAccountOperatorUnmarshal(s.T(), accOut.Bytes())
		s.Equal("accop"+s.network.Accounts[network.KeyAccount1].String(), acc.GetAccState().Key, "account operator key no equal")
		s.Equal(s.network.Accounts[network.KeyAccount1].String(), acc.GetAccState().Address, "account operator address no equal")
		s.Equal("active", acc.GetAccState().Status, "account operator status no equal")
	})
	s.Run("enroll_account_operator_but_not_authority", func() {
		out, err := ExCmdEnrollAccountOperator(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyEmpty1].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyAccount3),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code)
	})

	s.Run("revoke_account_operator", func() {
		out, err := ExCmdRevokeAccountOperator(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyAccount1].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyAuthority),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerSuccessCode, txResponse.Code)
		out, err = ExCmdQueryAccountOperator(validatorCtx, s.T(), s.network.Accounts[network.KeyAccount1].String(), network.JSONFlag())
		s.NoErrorf(err, "error should nil %v", err)
		acc := JsonAccountOperatorUnmarshal(s.T(), out.Bytes())
		s.Equalf("inactive", acc.GetAccState().Status, "account operator status no equal")
	})

	s.Run("revoke_account_operator_but_not_authorizer", func() {
		out, err := ExCmdRevokeAccountOperator(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyOperator].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyAccount1),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code)
		out, err = ExCmdQueryAccountOperator(validatorCtx, s.T(), s.network.Accounts[network.KeyOperator].String(), network.JSONFlag())
		s.NoErrorf(err, "error should nil %v", err)
		acc := JsonAccountOperatorUnmarshal(s.T(), out.Bytes())
		s.Equalf("active", acc.GetAccState().Status, "account operator status no equal")
	})

}

func (s *ElectoralIntegrationTestSuite) TestDocIssuer() {
	validatorCtx := s.network.Validators[0].ClientCtx

	s.Run("enroll_doc_issuer", func() {
		out, err := ExCmdEnrollDocIssuer(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyDocIssuer].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyOperator),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll account operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())

		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "the txn response %v", txResponse.String())

		accOut, err := ExCmdGetDocIssuer(validatorCtx, s.T(), s.network.Accounts[network.KeyDocIssuer].String(), network.JSONFlag())
		s.NoErrorf(err, "query doc issuer should not error %v", accOut.String())
		acc := JsonDocIssuerUnmarshal(s.T(), accOut.Bytes())
		s.Equal("docIssuer"+s.network.Accounts[network.KeyDocIssuer].String(), acc.GetAccState().Key, "doc issuer key no equal")
		s.Equal(s.network.Accounts[network.KeyDocIssuer].String(), acc.GetAccState().Address, "doc issuer address no equal")
		s.Equal("active", acc.GetAccState().Status, "doc issuer status no equal")
	})
	s.Run("enroll_doc_issuer_but_not_authority", func() {
		out, err := ExCmdEnrollDocIssuer(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyEmpty1].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyAccount3),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll doc issuer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code)
	})

	s.Run("revoke_doc_issuer_not_authorizer", func() {
		out, err := ExCmdRevokeDocIssuer(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyDocIssuer].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyAccount4),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "revoke doc issuer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code)

	})

	s.Run("revoke_doc_issuer", func() {
		out, err := ExCmdRevokeDocIssuer(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyDocIssuer].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyOperator),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "revoke_doc_issuer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerSuccessCode, txResponse.Code)
		out, err = ExCmdGetDocIssuer(validatorCtx, s.T(), s.network.Accounts[network.KeyDocIssuer].String(), network.JSONFlag())
		s.NoErrorf(err, "error should nil %v", err)
		acc := JsonDocIssuerUnmarshal(s.T(), out.Bytes())
		s.Equalf("inactive", acc.GetAccState().Status, "account operator status no equal")
	})
}

func (s *ElectoralIntegrationTestSuite) TestIdSigner() {
	validatorCtx := s.network.Validators[0].ClientCtx
	accOut, err := ExCmdQueryAccountOperator(validatorCtx, s.T(), s.network.Accounts[network.KeyOperator].String(), network.JSONFlag())
	s.NoErrorf(err, "query account operator should not error %v", err)
	acc := JsonAccountOperatorUnmarshal(s.T(), accOut.Bytes())
	s.T().Logf("account operator %s", acc.String())
	s.Run("enroll_id_signer", func() {
		out, err := ExCmdEnrollIdSigner(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyIDSigner].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyOperator),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll id signer operator should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())

		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "the txn response %v", txResponse.String())

		accOut, err := ExCmdGetIdSigner(validatorCtx, s.T(), s.network.Accounts[network.KeyIDSigner].String(), network.JSONFlag())
		s.NoErrorf(err, "query id signer should not error %v", accOut.String())
		acc := JsonIDSignerUnmarshal(s.T(), accOut.Bytes())
		s.Equal("idsigner"+s.network.Accounts[network.KeyIDSigner].String(), acc.GetAccState().Key, "id signer key no equal")
		s.Equal(s.network.Accounts[network.KeyIDSigner].String(), acc.GetAccState().Address, "id signer address no equal")
		s.Equal("active", acc.GetAccState().Status, "id signer status no equal")
	})
	s.Run("enroll_id_singer_but_not_authority", func() {
		out, err := ExCmdEnrollIdSigner(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyEmpty1].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyAccount3),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "enroll id signer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code)
	})

	s.Run("revoke_id_signer_not_authorizer", func() {
		out, err := ExCmdRevokeIdSigner(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyIDSigner].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyAccount4),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "revoke doc issuer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code)

	})

	s.Run("revoke_id_signer", func() {
		out, err := ExCmdRevokeIdSigner(validatorCtx, s.T(),
			[]string{s.network.Accounts[network.KeyIDSigner].String()},
			network.SHRFee2(),
			network.MakeByAccount(network.KeyOperator),
			network.SkipConfirmation(),
			network.BlockBroadcast(),
		)
		s.NoErrorf(err, "revoke id signer should not error %v", err)
		_ = s.network.WaitForNextBlock()
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "the response %v", out.String())
		out, err = ExCmdGetIdSigner(validatorCtx, s.T(), s.network.Accounts[network.KeyIDSigner].String(), network.JSONFlag())
		s.NoErrorf(err, "error should nil %v", err)
		acc := JsonIDSignerUnmarshal(s.T(), out.Bytes())
		s.Equalf("inactive", acc.GetAccState().Status, "id signer status no equal")
	})
}
