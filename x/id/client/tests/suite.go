package tests

import (
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	netutilts "github.com/sharering/shareledger/testutil/network"
	idtypes "github.com/sharering/shareledger/x/id/types"
	"github.com/stretchr/testify/suite"
	"os"
	"strings"

	"github.com/sharering/shareledger/x/electoral/client/tests"
)

type IDIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	dir     string
	network *network.Network
}

func NewIDIntegrationTestSuite(cfg network.Config) *IDIntegrationTestSuite {
	return &IDIntegrationTestSuite{cfg: cfg}
}
func (s *IDIntegrationTestSuite) setupTestMaterial() {

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

	initIDSig := []struct {
		accID    string
		id       string
		idOwner  string
		idBackup string
		idData   string
	}{
		{
			id:       "[owner_Acc1_Backup_Acc2]existed_replace_owner_id_1",
			idData:   "existed_replace_owner_id_data_1",
			idOwner:  netutilts.Accounts[netutilts.KeyAccount1].String(),
			idBackup: netutilts.Accounts[netutilts.KeyAccount2].String(),
			accID:    netutilts.KeyAccount1,
		},
		{
			accID: netutilts.KeyIDSigner,
		},
		{
			id:       "[owner_Acc3_Backup_Acc4]existed_replace_owner_id_2",
			idData:   "existed_replace_owner_id_data_2",
			idOwner:  netutilts.Accounts[netutilts.KeyAccount3].String(),
			idBackup: netutilts.Accounts[netutilts.KeyAccount4].String(),
			accID:    netutilts.KeyAccount3,
		},
		{
			id:       "[owner_Acc5_Backup_Acc6]existed_update_id_1",
			idData:   "existed_update_id_data_2",
			accID:    netutilts.KeyAccount5,
			idOwner:  netutilts.Accounts[netutilts.KeyAccount5].String(),
			idBackup: netutilts.Accounts[netutilts.KeyAccount6].String(),
		},
		{
			id:       "[owner_Acc6_Backup_Acc6]existed_update_id_1",
			idData:   "existed_update_id_data_1",
			accID:    netutilts.KeyAccount5,
			idOwner:  netutilts.Accounts[netutilts.KeyAccount6].String(),
			idBackup: netutilts.Accounts[netutilts.KeyAccount6].String(),
		},
		{
			id:       "[owner_Acc7_Backup_Acc8]existed_update_id_2",
			idData:   "existed_update_id_data_2",
			accID:    netutilts.KeyAccount7,
			idOwner:  netutilts.Accounts[netutilts.KeyAccount7].String(),
			idBackup: netutilts.Accounts[netutilts.KeyAccount8].String(),
		},
	}

	for _, id := range initIDSig {
		if id.accID != "" {
			out, _ = tests.ExCmdEnrollIdSigner(
				s.network.Validators[0].ClientCtx,
				[]string{netutilts.Accounts[id.accID].String()},
				netutilts.MakeByAccount(netutilts.KeyOperator),
				netutilts.SkipConfirmation,
				netutilts.BlockBroadcast,
				netutilts.SHRFee2,
			)
			s.Require().NoError(s.network.WaitForNextBlock())
			res = netutilts.ParseStdOut(s.T(), out.Bytes())
			s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init id signer fail %v", res.String())
		}
		if id.id != "" {
			out, _ = CmdExNewID(s.network.Validators[0].ClientCtx, id.id, id.idBackup, id.idOwner, id.idData,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(id.accID),
				netutilts.BlockBroadcast,
				netutilts.SHRFee2,
			)
			s.Require().NoError(s.network.WaitForNextBlock())
			res = netutilts.ParseStdOut(s.T(), out.Bytes())
			s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "init id fail %v", res.String())
		}

	}
}
func (s *IDIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for booking module")

	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//override the keyring by our keyring information
	s.network.Validators[0].ClientCtx.Keyring = kb

	s.T().Log("setting up id data....")
	s.setupTestMaterial()
	s.T().Log("setting up integration test suite successfully")
}
func (s *IDIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "tearing down fail")
	s.network.Cleanup()
	s.T().Log("tearing down integration test suite")
}

func (s *IDIntegrationTestSuite) TestCreateID() {

	validatorCtx := s.network.Validators[0].ClientCtx
	type (
		TestCase struct {
			d           string
			iId         string
			iBackupAddr string
			iOwnerAddr  string
			iExData     string
			txFee       int
			txCreator   string
			oErr        error
			oRes        *sdk.TxResponse
			oId         *idtypes.Id
		}
	)

	testSuite := []TestCase{
		{
			d:           "create_the_valid_id_should_be_success",
			iId:         "ID_1",
			iOwnerAddr:  "shareledger1pclpkwn6vc6lmrmv7407cd7em4cypdekc6kvn4",
			iBackupAddr: "shareledger1pclpkwn6vc6lmrmv7407cd7em4cypdekc6kvn4",
			iExData:     "the ex_data",
			txCreator:   netutilts.KeyIDSigner,
			txFee:       2,
			oErr:        nil,
			oRes:        &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oId: &idtypes.Id{
				Id: "ID_1",
				Data: &idtypes.BaseID{
					IssuerAddress: netutilts.Accounts[netutilts.KeyIDSigner].String(),
					BackupAddress: "shareledger1pclpkwn6vc6lmrmv7407cd7em4cypdekc6kvn4",
					OwnerAddress:  "shareledger1pclpkwn6vc6lmrmv7407cd7em4cypdekc6kvn4",
					ExtraData:     "the ex_data",
				},
			},
		},
		{
			d:           "create_the_valid_id_but_caller_is_not_id_signer",
			iId:         "ID_1",
			iOwnerAddr:  "shareledger1pclpkwn6vc6lmrmv7407cd7em4cypdekc6kvn4",
			iBackupAddr: "shareledger1pclpkwn6vc6lmrmv7407cd7em4cypdekc6kvn4",
			iExData:     "the ex_data",
			txCreator:   netutilts.KeyAccount2,
			txFee:       2,
			oErr:        nil,
			oRes:        &sdk.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}

	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			out, err := CmdExNewID(validatorCtx, tc.iId, tc.iBackupAddr, tc.iOwnerAddr, tc.iExData,
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txFee),
			)
			if tc.oErr != nil {
				s.NotNilf(err, "this case need return error")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "create ID fail %s", out)
			}
			if tc.oId != nil {
				out = CmdExGetID(validatorCtx, s.T(), tc.iId,
					netutilts.JSONFlag,
				)
				idData := IDJsonUnmarshal(s.T(), out.Bytes())

				s.Equalf(tc.oId.Id, idData.Id, "id not equal")
				s.Equalf(tc.oId.GetData().GetExtraData(), idData.GetData().GetExtraData(), "data not equal")
				s.Equal(tc.oId.GetData().BackupAddress, idData.ToBaseID().BackupAddress)
				s.Equal(tc.oId.GetData().OwnerAddress, idData.ToBaseID().OwnerAddress)
				s.Equal(netutilts.Accounts[netutilts.KeyIDSigner].String(), idData.ToBaseID().IssuerAddress)
			}
		})
	}

}

func (s *IDIntegrationTestSuite) TestCreateIDInBatch() {

	type (
		TestCase struct {
			d            string
			iIds         []string
			iBackupAddrs []string
			iOwnerAddrs  []string
			iExDatas     []string
			txFee        int
			txCreator    string
			oErr         error
			oRes         *sdk.TxResponse
			oId          []idtypes.Id
		}
	)

	testSuite := []TestCase{
		{
			d:    "create_the_valid_id_should_be_success",
			iIds: []string{"batch_ID_1", "batch_ID_2", "batch_ID_3"},
			iOwnerAddrs: []string{
				"shareledger1ghrpxfgfy0kdnas8lsr9wjq3q0hg0m3cs3n8n8",
				"shareledger17papd8h9glkvx0ff0lexn9u42689y63ffrtxs2",
				"shareledger1hq7wjjgeymvs3q4vmkvac3dghfsjwvjvf8jdaw"},
			iBackupAddrs: []string{
				"shareledger1ghrpxfgfy0kdnas8lsr9wjq3q0hg0m3cs3n8n8",
				"shareledger17papd8h9glkvx0ff0lexn9u42689y63ffrtxs2",
				"shareledger1hq7wjjgeymvs3q4vmkvac3dghfsjwvjvf8jdaw"},
			iExDatas:  []string{"ex_data_1", "ex_data_2", "ex_data_3"},
			txCreator: netutilts.KeyIDSigner,
			txFee:     2,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oId: []idtypes.Id{
				{
					Id: "batch_ID_1",
					Data: &idtypes.BaseID{
						IssuerAddress: netutilts.Accounts[netutilts.KeyIDSigner].String(),
						BackupAddress: "shareledger1ghrpxfgfy0kdnas8lsr9wjq3q0hg0m3cs3n8n8",
						OwnerAddress:  "shareledger1ghrpxfgfy0kdnas8lsr9wjq3q0hg0m3cs3n8n8",
						ExtraData:     "ex_data_1",
					},
				},
				{
					Id: "batch_ID_2",
					Data: &idtypes.BaseID{
						IssuerAddress: netutilts.Accounts[netutilts.KeyIDSigner].String(),
						BackupAddress: "shareledger17papd8h9glkvx0ff0lexn9u42689y63ffrtxs2",
						OwnerAddress:  "shareledger17papd8h9glkvx0ff0lexn9u42689y63ffrtxs2",
						ExtraData:     "ex_data_2",
					},
				},
				{
					Id: "batch_ID_3",
					Data: &idtypes.BaseID{
						IssuerAddress: netutilts.Accounts[netutilts.KeyIDSigner].String(),
						BackupAddress: "shareledger1hq7wjjgeymvs3q4vmkvac3dghfsjwvjvf8jdaw",
						OwnerAddress:  "shareledger1hq7wjjgeymvs3q4vmkvac3dghfsjwvjvf8jdaw",
						ExtraData:     "ex_data_3",
					},
				},
			},
		},
		{
			d:            "create_id_but_creator_is_not_authorizer",
			iIds:         []string{"batch_ID_4"},
			iOwnerAddrs:  []string{"shareledger1pclpkwn6vc6lmrmv7407cd7em4cypdekc6kvn4"},
			iBackupAddrs: []string{"shareledger1pclpkwn6vc6lmrmv7407cd7em4cypdekc6kvn4"},
			iExDatas:     []string{"the ex_data"},
			txCreator:    netutilts.KeyAccount2,
			txFee:        2,
			oErr:         nil,
			oRes:         &sdk.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			out, err := CmdExNewIDInBatch(validatorCtx,
				strings.Join(tc.iIds, ","),
				strings.Join(tc.iBackupAddrs, ","),
				strings.Join(tc.iOwnerAddrs, ","),
				strings.Join(tc.iExDatas, ","),
				netutilts.JSONFlag,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.BlockBroadcast,
				netutilts.SHRFee(tc.txFee),
			)
			if tc.oErr != nil {
				s.NotNilf(err, "this case need return error")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "create ID fail %s", out)
			}
			if len(tc.oId) != 0 {
				for _, i := range tc.oId {
					out = CmdExGetID(validatorCtx, s.T(), i.GetId(),
						netutilts.JSONFlag,
					)
					idData := IDJsonUnmarshal(s.T(), out.Bytes())

					s.Equalf(i.Id, idData.Id, "id not equal")
					s.Equalf(i.GetData().GetExtraData(), idData.GetData().GetExtraData(), "data not equal")
					s.Equal(i.GetData().BackupAddress, idData.ToBaseID().BackupAddress)
					s.Equal(i.GetData().OwnerAddress, idData.ToBaseID().OwnerAddress)
					s.Equal(netutilts.Accounts[netutilts.KeyIDSigner].String(), i.ToBaseID().IssuerAddress)
				}

			}
		})
	}

}

func (s *IDIntegrationTestSuite) TestUpdateID() {
	testSuite := []struct {
		d         string
		iID       string
		iData     string
		txFee     int
		txCreator string
		oErr      error
		oRes      *sdk.TxResponse
		oId       *idtypes.Id
	}{
		{
			d:         "update_id_data_success",
			iID:       "[owner_Acc5_Backup_Acc6]existed_update_id_1",
			iData:     "update_to_new_1",
			txFee:     2,
			txCreator: netutilts.KeyAccount5,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oId: &idtypes.Id{
				Id: "[owner_Acc5_Backup_Acc6]existed_update_id_1",
				Data: &idtypes.BaseID{
					ExtraData: "update_to_new_1",
				},
			},
		}, {
			d:         "update_id_data_caller_is_not_id_signer",
			iID:       "[owner_Acc7_Backup_Acc8]existed_update_id_2",
			iData:     "update_to_new_2",
			txFee:     2,
			txCreator: netutilts.KeyAccount2,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
			oId: &idtypes.Id{
				Id: "[owner_Acc7_Backup_Acc8]existed_update_id_2",
				Data: &idtypes.BaseID{
					ExtraData: "existed_update_id_data_2",
				},
			},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			out, err := CmdExUpdateID(validatorCtx, tc.iID, tc.iData,
				netutilts.SkipConfirmation,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.BlockBroadcast,
				netutilts.JSONFlag,
				netutilts.SHRFee(tc.txFee))
			if tc.oErr != nil {
				s.NotNilf(err, "error is required in this case")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "update ID fail %s", out)
			}
			if tc.oId != nil {
				out = CmdExGetID(validatorCtx, s.T(), tc.iID,
					netutilts.JSONFlag,
				)
				idData := IDJsonUnmarshal(s.T(), out.Bytes())
				s.Equalf(tc.oId.GetData().GetExtraData(), idData.GetData().GetExtraData(), "data not equal")
			}
		})
	}

}

func (s *IDIntegrationTestSuite) TestReplaceOwner() {

	testSuite := []struct {
		d         string
		iID       string
		iNewAddr  string
		txFee     int
		txCreator string
		oErr      error
		oRes      *sdk.TxResponse
		oId       *idtypes.Id
	}{
		{
			d:         "replace_id_owner_data_success",
			iID:       "[owner_Acc1_Backup_Acc2]existed_replace_owner_id_1",
			iNewAddr:  netutilts.Accounts[netutilts.KeyAccount2].String(),
			txFee:     2,
			txCreator: netutilts.KeyAccount2,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oId: &idtypes.Id{
				Id: "[owner_Acc1_Backup_Acc2]existed_replace_owner_id_1",
				Data: &idtypes.BaseID{
					OwnerAddress: netutilts.Accounts[netutilts.KeyAccount2].String(),
				},
			},
		},
		{
			d:         "replace_id_owner_data_caller_is_not_backup",
			iID:       "[owner_Acc3_Backup_Acc4]existed_replace_owner_id_2",
			iNewAddr:  netutilts.Accounts[netutilts.KeyAccount4].String(),
			txFee:     2,
			txCreator: netutilts.KeyAccount5,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
			oId: &idtypes.Id{
				Id: "[owner_Acc3_Backup_Acc4]existed_replace_owner_id_2",
				Data: &idtypes.BaseID{
					OwnerAddress: netutilts.Accounts[netutilts.KeyAccount3].String(),
				},
			},
		},
		{
			d:         "replace_id_owner_but_backup_already_hold_another_id",
			iID:       "[owner_Acc5_Backup_Acc6]existed_update_id_1",
			iNewAddr:  netutilts.Accounts[netutilts.KeyAccount6].String(),
			txFee:     2,
			txCreator: netutilts.KeyAccount6,
			oErr:      nil,
			oRes:      &sdk.TxResponse{Code: idtypes.ErrOwnerHasID.ABCICode()},
			oId: &idtypes.Id{
				Id: "[owner_Acc5_Backup_Acc6]existed_update_id_1",
				Data: &idtypes.BaseID{
					OwnerAddress: netutilts.Accounts[netutilts.KeyAccount5].String(),
				},
			},
		},
	}
	validatorCtx := s.network.Validators[0].ClientCtx
	for _, tc := range testSuite {
		s.Run(tc.d, func() {
			out, err := CmdExReplaceIdOwner(validatorCtx, tc.iID, tc.iNewAddr,
				netutilts.MakeByAccount(tc.txCreator),
				netutilts.SkipConfirmation, netutilts.JSONFlag, netutilts.SHRFee(tc.txFee), netutilts.BlockBroadcast)
			if tc.oErr != nil {
				s.NotNilf(err, "require error in this case")
			}
			if tc.oRes != nil {
				txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equalf(tc.oRes.Code, txResponse.Code, "update ID fail %s", out)
			}
			if tc.oId != nil {
				out = CmdExGetID(validatorCtx, s.T(), tc.iID,
					netutilts.JSONFlag,
				)
				idData := IDJsonUnmarshal(s.T(), out.Bytes())
				s.Equalf(tc.oId.GetData().GetOwnerAddress(), idData.GetData().GetOwnerAddress(), "owner not equal")
			}
		})
	}

}
