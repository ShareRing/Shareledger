package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	netutilts "github.com/sharering/shareledger/testutil/network"
	types2 "github.com/sharering/shareledger/x/electoral/types"
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

func (s *ElectoralIntegrationTestSuite) setupTestMaterial() {
	s.T().Log("setting up user operator")

	out, _ := ExCmdEnrollAccountOperator(s.network.Validators[0].ClientCtx,
		[]string{netutilts.Accounts[netutilts.KeyOperator].String()},
		netutilts.SHRFee2,
		netutilts.MakeByAccount(netutilts.KeyAuthority),
		netutilts.SkipConfirmation,
		netutilts.BlockBroadcast,
	)
	s.Require().NoError(s.network.WaitForNextBlock())

	txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equal(netutilts.ShareLedgerSuccessCode, txResponse.Code, "%s", out.String())

	intData := struct {
		Operator  []string
		DocIssuer []string
		IdSigner  []string
	}{
		Operator: []string{
			"shareledger1tc9yej24s698vm5w0jvt7452mgxl3ck74nqxjg",
			"shareledger1cfyrlfknvufzqap3yqsnyk4hxtp3xrqh7jlysm",
		},
		DocIssuer: []string{
			"shareledger19w2g2pdcwpj5kutn6ve3r4ay5twptlejlcvkq8",
			"shareledger1pasya6jtvglu9q36nl3tqh5drqzeh0s98czrep",
		},
		IdSigner: []string{
			"shareledger1sx9dp9289h8lxvg0u9z4g77y93zplvq63stkd5",
			"shareledger1kcu0fdn5f07wq9534yqy3t78p3uuc5rawsshe8",
		},
	}

	if len(intData.IdSigner) != 0 {
		out, err := ExCmdEnrollIdSigner(s.network.Validators[0].ClientCtx,
			intData.IdSigner,
			netutilts.SHRFee2,
			netutilts.MakeByAccount(netutilts.KeyOperator),
			netutilts.SkipConfirmation,
			netutilts.BlockBroadcast,
		)
		if err != nil {
			s.Require().NoError(err, "init id signer fail")
		}
		s.Require().NoError(s.network.WaitForNextBlock())
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerSuccessCode, txResponse.Code, "%s", out.String())
	}
	if len(intData.Operator) != 0 {
		out, err := ExCmdEnrollAccountOperator(s.network.Validators[0].ClientCtx,
			intData.Operator,
			netutilts.SHRFee2,
			netutilts.MakeByAccount(netutilts.KeyAuthority),
			netutilts.SkipConfirmation,
			netutilts.BlockBroadcast,
		)
		if err != nil {
			s.Require().NoError(err, "init operator fail")
		}
		s.Require().NoError(s.network.WaitForNextBlock())
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerSuccessCode, txResponse.Code, "%s", out.String())
	}

	if len(intData.DocIssuer) != 0 {
		out, err := ExCmdEnrollDocIssuer(s.network.Validators[0].ClientCtx,
			intData.DocIssuer,
			netutilts.SHRFee2,
			netutilts.MakeByAccount(netutilts.KeyOperator),
			netutilts.SkipConfirmation,
			netutilts.BlockBroadcast,
		)
		if err != nil {
			s.Require().NoError(err, "init doc issuer fail")
		}
		s.Require().NoError(s.network.WaitForNextBlock())
		txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerSuccessCode, txResponse.Code, "%s", out.String())
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
	s.setupTestMaterial()
	s.T().Log("setting up integration test suite successfully")

}
func (s *ElectoralIntegrationTestSuite) TearDownSuite() {
	s.network.Cleanup()
	s.NoError(os.RemoveAll(s.dir), "fail to cleanup")
	s.T().Log("tearing down integration test suite")
}

func (s *ElectoralIntegrationTestSuite) TestEnrollAccountOperator() {

	testSuite := []struct {
		d         string
		iAddress  []string
		txCreator string
		txFee     int
		oErr      error
		oRes      *types.TxResponse
		oAccState []types2.AccState
	}{
		{
			d:         "enroll_account_operator",
			iAddress:  []string{"shareledger1pasya6jtvglu9q36nl3tqh5drqzeh0s98czrep"},
			txCreator: netutilts.KeyAuthority,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oAccState: []types2.AccState{
				{
					Key:     "accopshareledger1pasya6jtvglu9q36nl3tqh5drqzeh0s98czrep",
					Address: "shareledger1pasya6jtvglu9q36nl3tqh5drqzeh0s98czrep",
					Status:  "active",
				},
			},
		},
		{
			d:         "enroll_account_operator_but_not_authority",
			iAddress:  []string{"shareledger18ncqa238qk0xxpgq5v9yz88p79r0mc3tt0yyvx"},
			txCreator: netutilts.KeyAccount1,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			stdOut, err := ExCmdEnrollAccountOperator(validatorCtx, tc.iAddress,
				netutilts.SHRFee(tc.txFee),
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator))
			if tc.oErr != nil {
				s.NotNilf(err, "case %s require error this step", tc.d)
			}

			txnResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
			s.Require().NoError(s.network.WaitForNextBlock())
			if tc.oRes != nil {
				s.Equal(tc.oRes.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.d, txnResponse.String()))
			}

			if len(tc.oAccState) != 0 {
				for _, a := range tc.oAccState {
					accOut, err := ExCmdQueryAccountOperator(validatorCtx, a.Address, netutilts.JSONFlag)
					s.Require().NoError(err, "fail to get account operator")
					acc := JsonAccountOperatorUnmarshal(s.T(), accOut.Bytes())
					s.Equal(a.GetKey(), acc.GetAccState().Key, "account operator key no equal")
					s.Equal(a.GetAddress(), acc.GetAccState().Address, "account operator address no equal")
					s.Equal(a.GetStatus(), acc.GetAccState().Status, "account operator status no equal")
				}
			}
		})
	}

}

func (s *ElectoralIntegrationTestSuite) TestRevokeAccountOperator() {

	testSuite := []struct {
		d         string
		iAddress  []string
		txCreator string
		txFee     int
		oErr      error
		oRes      *types.TxResponse
		oAccState []types2.AccState
	}{
		{
			d:         "revoke_account_operator",
			iAddress:  []string{"shareledger1tc9yej24s698vm5w0jvt7452mgxl3ck74nqxjg"},
			txCreator: netutilts.KeyAuthority,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oAccState: []types2.AccState{
				{
					Key:     "accopshareledger1tc9yej24s698vm5w0jvt7452mgxl3ck74nqxjg",
					Address: "shareledger1tc9yej24s698vm5w0jvt7452mgxl3ck74nqxjg",
					Status:  "inactive",
				},
			},
		},
		{
			d:         "enroll_account_operator_but_not_authority",
			iAddress:  []string{"shareledger1cfyrlfknvufzqap3yqsnyk4hxtp3xrqh7jlysm"},
			txCreator: netutilts.KeyAccount1,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			stdOut, err := ExCmdRevokeAccountOperator(validatorCtx, tc.iAddress,
				netutilts.SHRFee(tc.txFee),
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator))
			if tc.oErr != nil {
				s.NotNilf(err, "case %s require error this step", tc.d)
			}

			txnResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
			s.Require().NoError(s.network.WaitForNextBlock())
			if tc.oRes != nil {
				s.Equal(tc.oRes.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.d, txnResponse.String()))
			}

			if len(tc.oAccState) != 0 {
				for _, a := range tc.oAccState {
					accOut, err := ExCmdQueryAccountOperator(validatorCtx, a.Address, netutilts.JSONFlag)
					s.Require().NoErrorf(err, "fail to get account operator %s", a.Address)
					acc := JsonAccountOperatorUnmarshal(s.T(), accOut.Bytes())
					s.Equal(a.GetStatus(), acc.GetAccState().Status, "account operator status no equal")
				}
			}
		})
	}

}

func (s *ElectoralIntegrationTestSuite) TestEnrollDocIssuer() {

	testSuite := []struct {
		d         string
		iAddress  []string
		txCreator string
		txFee     int
		oErr      error
		oRes      *types.TxResponse
		oAccState []types2.AccState
	}{
		{
			d:         "enroll_doc_issuer",
			iAddress:  []string{"shareledger1syxg95th7whx23fn6vwaan34w4kal49azlawnq"},
			txCreator: netutilts.KeyOperator,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oAccState: []types2.AccState{
				{
					Key:     "docIssuershareledger1syxg95th7whx23fn6vwaan34w4kal49azlawnq",
					Address: "shareledger1syxg95th7whx23fn6vwaan34w4kal49azlawnq",
					Status:  "active",
				},
			},
		},
		{
			d:         "enroll_doc_issuer_but_not_authority",
			iAddress:  []string{"shareledger18ncqa238qk0xxpgq5v9yz88p79r0mc3tt0yyvx"},
			txCreator: netutilts.KeyAccount1,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			stdOut, err := ExCmdEnrollDocIssuer(validatorCtx, tc.iAddress,
				netutilts.SHRFee(tc.txFee),
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator))
			if tc.oErr != nil {
				s.NotNilf(err, "case %s require error this step", tc.d)
			}

			txnResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
			s.Require().NoError(s.network.WaitForNextBlock())
			if tc.oRes != nil {
				s.Equal(tc.oRes.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.d, txnResponse.String()))
			}

			if len(tc.oAccState) != 0 {
				for _, a := range tc.oAccState {
					accOut, err := ExCmdGetDocIssuer(validatorCtx, a.Address, netutilts.JSONFlag)
					s.Require().NoError(err, "fail to get doc issuer")
					acc := JsonIDSignerUnmarshal(s.T(), accOut.Bytes())
					s.Equal(a.GetKey(), acc.GetAccState().Key, "doc issuer key no equal")
					s.Equal(a.GetAddress(), acc.GetAccState().Address, "doc issuer address no equal")
					s.Equal(a.GetStatus(), acc.GetAccState().Status, "doc issuer status no equal")
				}
			}
		})
	}
}

func (s *ElectoralIntegrationTestSuite) TestRevokeDocIssuer() {

	testSuite := []struct {
		d         string
		iAddress  []string
		txCreator string
		txFee     int
		oErr      error
		oRes      *types.TxResponse
		oAccState []types2.AccState
	}{
		{
			d:         "reovke_doc_issuer_successfully",
			iAddress:  []string{"shareledger19w2g2pdcwpj5kutn6ve3r4ay5twptlejlcvkq8"},
			txCreator: netutilts.KeyOperator,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oAccState: []types2.AccState{
				{
					Key:     "docIssuershareledger19w2g2pdcwpj5kutn6ve3r4ay5twptlejlcvkq8",
					Address: "shareledger19w2g2pdcwpj5kutn6ve3r4ay5twptlejlcvkq8",
					Status:  "inactive",
				},
			},
		},
		{
			d:         "revoke_doc_issuer_but_not_authority",
			iAddress:  []string{"shareledger1pasya6jtvglu9q36nl3tqh5drqzeh0s98czrep"},
			txCreator: netutilts.KeyAccount1,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			stdOut, err := ExCmdRevokeDocIssuer(validatorCtx, tc.iAddress,
				netutilts.SHRFee(tc.txFee),
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator))
			if tc.oErr != nil {
				s.NotNilf(err, "case %s require error this step", tc.d)
			}

			txnResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
			s.Require().NoError(s.network.WaitForNextBlock())
			if tc.oRes != nil {
				s.Equal(tc.oRes.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.d, txnResponse.String()))
			}

			if len(tc.oAccState) != 0 {
				for _, a := range tc.oAccState {
					accOut, err := ExCmdGetDocIssuer(validatorCtx, a.Address, netutilts.JSONFlag)
					s.Require().NoError(err, "fail to get doc issuer")
					acc := JsonIDSignerUnmarshal(s.T(), accOut.Bytes())
					s.Equal(a.GetKey(), acc.GetAccState().Key, "doc issuer key no equal")
					s.Equal(a.GetAddress(), acc.GetAccState().Address, "doc issuer address no equal")
					s.Equal(a.GetStatus(), acc.GetAccState().Status, "doc issuer status no equal")
				}
			}
		})
	}
}

func (s *ElectoralIntegrationTestSuite) TestEnrollIDSigner() {

	testSuite := []struct {
		d         string
		iAddress  []string
		txCreator string
		txFee     int
		oErr      error
		oRes      *types.TxResponse
		oAccState []types2.AccState
	}{
		{
			d:         "enroll_id_signer",
			iAddress:  []string{"shareledger1dtd4krk8e69y9pze43qkprsqqpsaxljuus3maf"},
			txCreator: netutilts.KeyOperator,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oAccState: []types2.AccState{
				{
					Key:     "idsignershareledger1dtd4krk8e69y9pze43qkprsqqpsaxljuus3maf",
					Address: "shareledger1dtd4krk8e69y9pze43qkprsqqpsaxljuus3maf",
					Status:  "active",
				},
			},
		},
		{
			d:         "enroll_id_signer_but_not_authority",
			iAddress:  []string{"shareledger18ncqa238qk0xxpgq5v9yz88p79r0mc3tt0yyvx"},
			txCreator: netutilts.KeyAccount1,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			stdOut, err := ExCmdEnrollIdSigner(validatorCtx, tc.iAddress,
				netutilts.SHRFee(tc.txFee),
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator))
			if tc.oErr != nil {
				s.NotNilf(err, "case %s require error this step", tc.d)
			}

			txnResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
			s.Require().NoError(s.network.WaitForNextBlock())
			if tc.oRes != nil {
				s.Equal(tc.oRes.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.d, txnResponse.String()))
			}

			if len(tc.oAccState) != 0 {
				for _, a := range tc.oAccState {
					accOut, err := ExCmdGetIdSigner(validatorCtx, a.Address, netutilts.JSONFlag)
					s.Require().NoError(err, "fail to get doc issuer")
					acc := JsonIDSignerUnmarshal(s.T(), accOut.Bytes())
					s.Equal(a.GetKey(), acc.GetAccState().Key, "doc issuer key no equal")
					s.Equal(a.GetAddress(), acc.GetAccState().Address, "doc issuer address no equal")
					s.Equal(a.GetStatus(), acc.GetAccState().Status, "doc issuer status no equal")
				}
			}
		})
	}
}

func (s *ElectoralIntegrationTestSuite) TestRevokeIDSigner() {

	testSuite := []struct {
		d         string
		iAddress  []string
		txCreator string
		txFee     int
		oErr      error
		oRes      *types.TxResponse
		oAccState []types2.AccState
	}{
		{
			d:         "revoke_id_signer",
			iAddress:  []string{"shareledger1sx9dp9289h8lxvg0u9z4g77y93zplvq63stkd5"},
			txCreator: netutilts.KeyOperator,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oAccState: []types2.AccState{
				{
					Key:     "idsignershareledger1sx9dp9289h8lxvg0u9z4g77y93zplvq63stkd5",
					Address: "shareledger1sx9dp9289h8lxvg0u9z4g77y93zplvq63stkd5",
					Status:  "inactive",
				},
			},
		},
		{
			d:         "revoke_id_signer_but_not_authority",
			iAddress:  []string{"shareledger1kcu0fdn5f07wq9534yqy3t78p3uuc5rawsshe8"},
			txCreator: netutilts.KeyAccount1,
			txFee:     2,
			oErr:      nil,
			oRes:      &types.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			stdOut, err := ExCmdRevokeIdSigner(validatorCtx, tc.iAddress,
				netutilts.SHRFee(tc.txFee),
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator))
			if tc.oErr != nil {
				s.NotNilf(err, "case %s require error this step", tc.d)
			}

			txnResponse := netutilts.ParseStdOut(s.T(), stdOut.Bytes())
			s.Require().NoError(s.network.WaitForNextBlock())
			if tc.oRes != nil {
				s.Equal(tc.oRes.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.d, txnResponse.String()))
			}

			if len(tc.oAccState) != 0 {
				for _, a := range tc.oAccState {
					accOut, err := ExCmdGetIdSigner(validatorCtx, a.Address, netutilts.JSONFlag)
					s.Require().NoError(err, "fail to get doc issuer")
					acc := JsonIDSignerUnmarshal(s.T(), accOut.Bytes())
					s.Equal(a.GetKey(), acc.GetAccState().Key, "doc issuer key no equal")
					s.Equal(a.GetAddress(), acc.GetAccState().Address, "doc issuer address no equal")
					s.Equal(a.GetStatus(), acc.GetAccState().Status, "doc issuer status no equal")
				}
			}
		})
	}
}
