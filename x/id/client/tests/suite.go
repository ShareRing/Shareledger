package tests

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/sharering/shareledger/x/electoral/client/tests"
)

type IDIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewIDIntegrationTestSuite(cfg network.Config) *IDIntegrationTestSuite {
	return &IDIntegrationTestSuite{cfg: cfg}
}

func (s *IDIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for booking module")

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.T().Log("setting up document data....")

	out, _ := tests.ExCmdEnrollAccountOperator(
		s.network.Validators[0].ClientCtx,
		s.T(),
		[]string{s.network.Accounts[network.KeyOperator].String()},
		network.MakeByAccount(network.KeyAuthority),
		network.SkipConfirmation(),
		network.BlockBroadcast(),
		network.SHRFee2(),
	)
	_ = s.network.WaitForNextBlock()
	res := network.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode, res.Code, "init operator fail %v", res.String())

	out, _ = tests.ExCmdEnrollIdSigner(
		s.network.Validators[0].ClientCtx,
		s.T(),
		[]string{s.network.Accounts[network.KeyIDSigner].String()},
		network.MakeByAccount(network.KeyOperator),
		network.SkipConfirmation(),
		network.BlockBroadcast(),
		network.SHRFee2(),
	)
	_ = s.network.WaitForNextBlock()
	res = network.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(network.ShareLedgerSuccessCode, res.Code, "init operator fail %v", res.String())

	s.T().Log("setting up integration test suite successfully")
}
func (s *IDIntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
}

func (s *IDIntegrationTestSuite) TestCreateID() {

	_, _, addr := sample.RandomAddr(1)

	validatorCtx := s.network.Validators[0].ClientCtx
	s.Run("create_the_valid_id_should_be_success", func() {
		out := CmdExNewID(validatorCtx, s.T(), "Id1", addr[0], addr[0], "this is the new id",
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyIDSigner),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "create ID fail %s", out)

		out = CmdExGetID(validatorCtx, s.T(), "Id1",
			network.JSONFlag(),
		)

		idData := IDJsonUnmarshal(s.T(), out.Bytes())
		address, err := sdk.AccAddressFromBech32(addr[0])
		s.NoError(err)

		s.Equalf("Id1", idData.Id, "id not equal")
		s.Equalf("this is the new id", idData.GetData().GetExtraData(), "data not equal")
		s.Equal(address.String(), idData.ToBaseID().BackupAddress)
		s.Equal(address.String(), idData.ToBaseID().OwnerAddress)
		s.Equal(s.network.Accounts[network.KeyIDSigner].String(), idData.ToBaseID().IssuerAddress)
		out = CmdExGetIDByAddress(validatorCtx, s.T(), addr[0],
			network.JSONFlag(),
		)

		idData = IDJsonUnmarshal(s.T(), out.Bytes())
		s.Equalf("Id1", idData.Id, "id not equal")
		s.Equalf("this is the new id", idData.GetData().GetExtraData(), "data not equal")
		s.Equal(address.String(), idData.ToBaseID().BackupAddress)
		s.Equal(address.String(), idData.ToBaseID().OwnerAddress)
		s.Equal(s.network.Accounts[network.KeyIDSigner].String(), idData.ToBaseID().IssuerAddress)
	})

	s.Run("create_the_valid_id_but_caller_is_not_id_signer", func() {
		out := CmdExNewID(validatorCtx, s.T(), "Id12", addr[0], addr[0], "this is the new id",
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount2),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code, "create ID fail %s", out)

	})
}

func (s *IDIntegrationTestSuite) TestCreateIDInBatch() {

	_, _, addr := sample.RandomAddr(3)

	validatorCtx := s.network.Validators[0].ClientCtx
	s.Run("create_the_valid_id_should_be_success", func() {

		extras := []string{"extras-b-1", "extras-b-2", "extras-b-3"}
		ids := []string{"id-12", "id-13", "id-14"}
		out := CmdExNewIDInBatch(validatorCtx, s.T(), strings.Join(ids, ","), strings.Join(addr, ","), strings.Join(addr, ","), strings.Join(extras, ","),
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyIDSigner),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "create ID fail %s", out)

		for i := 0; i < 3; i++ {

			out = CmdExGetID(validatorCtx, s.T(), ids[i],
				network.JSONFlag(),
			)

			idData := IDJsonUnmarshal(s.T(), out.Bytes())
			address, err := sdk.AccAddressFromBech32(addr[i])
			s.NoError(err)
			s.Equalf(ids[i], idData.Id, "id not equal")
			s.Equalf(extras[i], idData.GetData().GetExtraData(), "data not equal")
			s.Equal(address.String(), idData.ToBaseID().BackupAddress)
			s.Equal(address.String(), idData.ToBaseID().OwnerAddress)
			s.Equal(s.network.Accounts[network.KeyIDSigner].String(), idData.ToBaseID().IssuerAddress)

			out = CmdExGetIDByAddress(validatorCtx, s.T(), addr[i],
				network.JSONFlag(),
			)

			idData = IDJsonUnmarshal(s.T(), out.Bytes())
			s.Equalf(ids[i], idData.Id, "id not equal")
			s.Equalf(extras[i], idData.GetData().GetExtraData(), "data not equal")
			s.Equal(address.String(), idData.ToBaseID().BackupAddress)
			s.Equal(address.String(), idData.ToBaseID().OwnerAddress)
			s.Equal(s.network.Accounts[network.KeyIDSigner].String(), idData.ToBaseID().IssuerAddress)
		}

	})

	s.Run("create_the_valid_id_but_caller_is_not_id_signer", func() {
		extras := []string{"extras-b-1", "extras-b-2", "extras-b-3"}
		ids := []string{"id-12", "id-13", "id-14"}
		out := CmdExNewIDInBatch(validatorCtx, s.T(), strings.Join(ids, ","), strings.Join(addr, ","), strings.Join(addr, ","), strings.Join(extras, ","),
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount1),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code, "create ID fail %s", out)

	})
}

func (s *IDIntegrationTestSuite) TestUpdateID() {
	s.Run("update_id_data_success", func() {
		_, _, addr := sample.RandomAddr(1)

		validatorCtx := s.network.Validators[0].ClientCtx
		out := CmdExNewID(validatorCtx, s.T(), "ID-forUpdate", addr[0], s.network.Accounts[network.KeyIDSigner].String(), "thisISBackup",
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyIDSigner),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "create ID fail %s", out)
		_ = s.network.WaitForNextBlock()
		out = CmdExUpdateID(validatorCtx, s.T(), "ID-forUpdate", "https://sharering.network/id/1",
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyIDSigner),
			network.BlockBroadcast(),
			network.JSONFlag(),
			network.SHRFee2())
		txResponse = network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "update ID fail %s", out)

		out = CmdExGetID(validatorCtx, s.T(), "ID-forUpdate",
			network.JSONFlag(),
		)
		idData := IDJsonUnmarshal(s.T(), out.Bytes())
		s.Equalf("https://sharering.network/id/1", idData.GetData().GetExtraData(), "data not equal")
	})
	s.Run("update_id_data_caller_is_not_id_signer", func() {
		_, _, addr := sample.RandomAddr(1)

		validatorCtx := s.network.Validators[0].ClientCtx
		out := CmdExNewID(validatorCtx, s.T(), "ID-forUpdate2", addr[0], addr[0], "thisISBackup2",
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyIDSigner),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "create ID fail %s", out)
		_ = s.network.WaitForNextBlock()
		out = CmdExUpdateID(validatorCtx, s.T(), "ID-forUpdate2", "https://sharering.network/id/1",
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyAccount3),
			network.BlockBroadcast(),
			network.JSONFlag(),
			network.SHRFee2())
		txResponse = network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerErrorCodeUnauthorized, txResponse.Code, "update ID fail %s", out)

		out = CmdExGetID(validatorCtx, s.T(), "ID-forUpdate2",
			network.JSONFlag(),
		)
		idData := IDJsonUnmarshal(s.T(), out.Bytes())
		s.Equalf("thisISBackup2", idData.GetData().GetExtraData(), "data was changed by unauthorized user")
	})
}

func (s *IDIntegrationTestSuite) TestReplaceOwner() {
	s.Run("replace_id_owner_data_success", func() {
		_, _, addr := sample.RandomAddr(1)

		validatorCtx := s.network.Validators[0].ClientCtx
		out := CmdExNewID(validatorCtx, s.T(), "ID-forReplace", s.network.Accounts[network.KeyAccount3].String(), addr[0], "thisISID",
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyIDSigner),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "create ID fail %s", out)
		_ = s.network.WaitForNextBlock()

		out = CmdExReplaceIdOwner(validatorCtx, s.T(), "ID-forReplace", s.network.Accounts[network.KeyAccount3].String(),
			network.MakeByAccount(network.KeyAccount3),
			network.SkipConfirmation(), network.JSONFlag(), network.SHRFee2(), network.BlockBroadcast())

		out = CmdExGetID(validatorCtx, s.T(), "ID-forReplace",
			network.JSONFlag(),
		)
		idData := IDJsonUnmarshal(s.T(), out.Bytes())
		s.Equalf(s.network.Accounts[network.KeyAccount3].String(), idData.GetData().GetOwnerAddress(), "owner data didn't change")
	})
	s.Run("replace_id_owner_data_caller_is_not_backup", func() {
		_, _, addr := sample.RandomAddr(1)

		validatorCtx := s.network.Validators[0].ClientCtx
		out := CmdExNewID(validatorCtx, s.T(), "ID-forReplace2", s.network.Accounts[network.KeyAccount3].String(), addr[0], "thisISID",
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyIDSigner),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "create ID fail %s", out)
		_ = s.network.WaitForNextBlock()

		out = CmdExReplaceIdOwner(validatorCtx, s.T(), "ID-forReplace2", s.network.Accounts[network.KeyAccount3].String(),
			network.MakeByAccount(network.KeyAccount1),
			network.SkipConfirmation(), network.JSONFlag(), network.SHRFee2(), network.BlockBroadcast())

		out = CmdExGetID(validatorCtx, s.T(), "ID-forReplace2",
			network.JSONFlag(),
		)
		idData := IDJsonUnmarshal(s.T(), out.Bytes())
		address, err := sdk.AccAddressFromBech32(addr[0])
		s.NoError(err)
		s.Equalf(address.String(), idData.GetData().GetOwnerAddress(), "owner data not equal")
	})

	s.Run("replace_id_owner_but_backup_already_hold_another_id", func() {
		_, _, addr := sample.RandomAddr(1)

		validatorCtx := s.network.Validators[0].ClientCtx
		out := CmdExNewID(validatorCtx, s.T(), "ID-forReplace4", addr[0], s.network.Accounts[network.KeyAccount4].String(), "thisISID3",
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyIDSigner),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "create ID fail %s", out)
		_ = s.network.WaitForNextBlock()

		out = CmdExNewID(validatorCtx, s.T(), "ID-forReplace3", s.network.Accounts[network.KeyAccount4].String(), addr[0], "thisISID3",
			network.JSONFlag(),
			network.SkipConfirmation(),
			network.MakeByAccount(network.KeyIDSigner),
			network.BlockBroadcast(),
			network.SHRFee2(),
		)
		txResponse = network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "create ID fail %s", out)
		_ = s.network.WaitForNextBlock()

		out = CmdExReplaceIdOwner(validatorCtx, s.T(), "ID-forReplace3", s.network.Accounts[network.KeyAccount4].String(),
			network.MakeByAccount(network.KeyAccount4),
			network.SkipConfirmation(), network.JSONFlag(), network.SHRFee2(), network.BlockBroadcast())

		txResponse = network.ParseStdOut(s.T(), out.Bytes())
		s.Equalf(network.ShareLedgerErrorCodeIDAddressOwnerID, txResponse.Code, "check replace response fail %s", out)

		_ = s.network.WaitForNextBlock()
		out = CmdExGetID(validatorCtx, s.T(), "ID-forReplace3",
			network.JSONFlag(),
		)
		idData := IDJsonUnmarshal(s.T(), out.Bytes())
		address, err := sdk.AccAddressFromBech32(addr[0])
		s.NoError(err)
		s.Equalf(address.String(), idData.GetData().GetOwnerAddress(), "owner data not equal")
	})
}
