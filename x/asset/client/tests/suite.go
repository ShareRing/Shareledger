package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	types2 "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/asset/types"
	"os"

	netutilts "github.com/sharering/shareledger/testutil/network"

	"github.com/stretchr/testify/suite"
)

type (
	AssetIntegrationTestSuite struct {
		suite.Suite

		cfg     *network.Config
		network *network.Network
		dir     string

		existedAsset []GenesisAsset
	}

	GenesisAsset struct {
		AssetHash    string
		AssetID      string
		AssetRate    string
		AssetStatus  string
		AssetCreator string
	}
)

func NewAssetIntegrationTestSuite(cfg *network.Config) *AssetIntegrationTestSuite {
	return &AssetIntegrationTestSuite{cfg: cfg}
}

func (s *AssetIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	kb, dir := netutilts.GetTestingGenesis(s.T(), s.cfg)
	s.dir = dir

	s.network = network.New(s.T(), *s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//override the keyring by our keyring information
	s.network.Validators[0].ClientCtx.Keyring = kb

	s.setupTestMaterial()

	s.T().Log("setting up integration test suite successfully")
}
func (s *AssetIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "cleanup test case fails")
	s.network.Cleanup()
	s.T().Log("tearing down integration test suite")
}

func (s *AssetIntegrationTestSuite) TestCreateAsset() {

	type testCase struct {
		iAssetID     string
		iAssetHash   string
		iAssetStatus string
		iShrPFee     string
		iTxFee       int
		iAcc         string
		oAsset       *types.Asset
		oResult      *types2.TxResponse
		oErr         error
		d            string
	}
	//write your test suite here
	testCases := []testCase{
		{
			iAssetID:     "id1",
			iAssetHash:   "cc6f58bd1ada876f0a4941ad579908eda726d6da",
			iAssetStatus: "true",
			iShrPFee:     "2",
			iTxFee:       6,
			iAcc:         netutilts.KeyAccount1,
			oResult:      &types2.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oAsset: &types.Asset{
				Hash:   []byte("cc6f58bd1ada876f0a4941ad579908eda726d6da"),
				UUID:   "id1",
				Status: true,
				Rate:   2,
			},
			oErr: nil,
			d:    "create the asset successfully and getting new asset for recheck 1",
		},
		{
			iAssetID:     "id2",
			iAssetHash:   "cc6f58bd1ada876f0a4941ad579908eda726d6sw",
			iAssetStatus: "true",
			iShrPFee:     "2",
			iTxFee:       6,
			iAcc:         netutilts.KeyAccount1,
			oResult:      &types2.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			oAsset: &types.Asset{
				Hash:   []byte("cc6f58bd1ada876f0a4941ad579908eda726d6sw"),
				UUID:   "id2",
				Status: true,
				Rate:   2,
			},
			oErr: nil,
			d:    "create the asset successfully and getting new asset for recheck 2",
		},
		{
			iAssetID:     "e_id_1",
			iAssetHash:   "0b44517f76eab863c3c4ab13c9774fadd62080ee97af663ec76ffeb671e9e064",
			iAssetStatus: "true",
			iShrPFee:     "2",
			iTxFee:       6,
			iAcc:         netutilts.KeyAccount1,
			oResult:      &types2.TxResponse{Code: types.ErrAssetExist.ABCICode()},
			oErr:         nil,
			d:            "create the duplicate asset",
		},
		{
			iAssetID:     "e_id_4",
			iAssetHash:   "0b44517f76eab863c3c4ab13c9774fadd62080ee97af663ec76ffeb671e9ess4",
			iAssetStatus: "true",
			iShrPFee:     "2",
			iTxFee:       6,
			iAcc:         netutilts.KeyEmpty1,
			oResult:      &types2.TxResponse{Code: sdkerrors.ErrInsufficientFunds.ABCICode()},
			oErr:         nil,
			d:            "create the asset but not enough shr to make txn",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.d, func() {
			out, err := ExCmdCreateAsset(s.network.Validators[0].ClientCtx, tc.iAssetID, tc.iAssetHash, tc.iAssetStatus, tc.iShrPFee,
				netutilts.SHRFee(tc.iTxFee),
				netutilts.JSONFlag,
				netutilts.MakeByAccount(tc.iAcc),
			)

			if tc.oErr != nil {
				s.NotNilf(err, tc.d)
			}
			if tc.oResult != nil {
				txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equal(tc.oResult.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.d, txnResponse.String()))
			}
			if tc.oAsset != nil {
				assetByte, err := ExCmdGetAsset(s.network.Validators[0].ClientCtx, tc.iAssetID)
				s.NoError(err)
				asset := AssetJsonUnmarshal(s.T(), assetByte.Bytes())
				s.Equal(string(tc.oAsset.Hash), string(asset.Hash))
				s.Equal(tc.oAsset.UUID, asset.UUID)
				s.Equal(tc.oAsset.Status, asset.Status, tc.d)
				s.Equal(tc.oAsset.Rate, asset.Rate, tc.d)
			}
		})

	}

}

func (s *AssetIntegrationTestSuite) TestUpdateAsset() {

	type TestCase struct {
		iAssetID     string
		iAssetHash   string
		iAssetStatus string
		iShrPFee     string
		iTxFee       int
		iAcc         string
		oAsset       *types.Asset
		oResult      *types2.TxResponse
		oErr         error
		d            string
	}
	//write your test suite here
	testCase := []TestCase{
		{
			iAssetID:     "e_id_1",
			iAssetHash:   "0b44517f76eab863c3c4ab13c9774fadd62080ee97af663ec76ffeb671e9e064",
			iAssetStatus: "false",
			iShrPFee:     "4",
			iTxFee:       6,
			iAcc:         netutilts.KeyAccount1,
			oAsset: &types.Asset{
				Hash:   []byte("0b44517f76eab863c3c4ab13c9774fadd62080ee97af663ec76ffeb671e9e064"),
				UUID:   "e_id_1",
				Status: false,
				Rate:   4,
			},
			oErr:    nil,
			oResult: &types2.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			d:       "update exited asset by update shrp fee and status",
		},
		{
			iAssetID:     "not_found",
			iAssetHash:   "not_found_0b44517f76eab863c3c4ab13c9774fadd62080ee97af663ec76ffeb671e9e064",
			iAssetStatus: "false",
			iShrPFee:     "4",
			iTxFee:       6,
			iAcc:         netutilts.KeyAccount1,
			oErr:         nil,
			oResult:      &types2.TxResponse{Code: types.ErrNameDoesNotExist.ABCICode()},
			d:            "update the not found asset should return an error code",
		},
		{
			iAssetID:     "e_id_2",
			iAssetHash:   "1a70dac5cb91142a8b5eda54dbeed46a42d1d7af83788c05fa43485acba09fa8",
			iAssetStatus: "false",
			iShrPFee:     "4",
			iTxFee:       6,
			iAcc:         netutilts.KeyAccount3,
			oErr:         nil,
			oResult:      &types2.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
			d:            "update exited asset but the who make transaction wasn't owner of this asset",
		},
	}
	for _, tc := range testCase {
		s.Run(tc.d, func() {
			out, err := ExCmdUpdateAsset(s.network.Validators[0].ClientCtx,
				tc.iAssetID, tc.iAssetHash, tc.iAssetStatus, tc.iShrPFee,
				netutilts.SHRFee(tc.iTxFee),
				netutilts.JSONFlag,
				netutilts.MakeByAccount(tc.iAcc),
			)
			// expect test case co error
			if tc.oErr != nil {
				s.NotNilf(err, tc.d)
			}
			if tc.oResult != nil {
				txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equal(tc.oResult.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.d, txnResponse.String()))
			}
			if tc.oAsset != nil {
				assetByte, err := ExCmdGetAsset(s.network.Validators[0].ClientCtx, tc.iAssetID)
				s.NoError(err)
				asset := AssetJsonUnmarshal(s.T(), assetByte.Bytes())
				s.Equal(string(tc.oAsset.Hash), string(asset.Hash))
				s.Equal(tc.oAsset.UUID, asset.UUID)
				s.Equal(tc.oAsset.Status, asset.Status, tc.d)
				s.Equal(tc.oAsset.Rate, asset.Rate, tc.d)
			}
		})
	}

}

func (s *AssetIntegrationTestSuite) TestDeleteAsset() {

	type TestCase struct {
		iAssetID string
		iTxFee   int
		iAcc     string
		oResult  *types2.TxResponse
		oErr     error
		d        string
	}
	//write your test suite here
	testCases := []TestCase{
		{
			iAssetID: "e_id_3",
			iTxFee:   4,
			iAcc:     netutilts.KeyAccount1,
			oErr:     nil,
			oResult:  &types2.TxResponse{Code: netutilts.ShareLedgerSuccessCode},
			d:        "delete asset successfully",
		},
		{
			iAssetID: "not_found",
			iTxFee:   4,
			iAcc:     netutilts.KeyAccount1,
			oErr:     nil,
			oResult:  &types2.TxResponse{Code: sdkerrors.ErrNotFound.ABCICode()},
			d:        "delete not found asset",
		},
		{
			iAssetID: "e_id_2",
			iTxFee:   4,
			iAcc:     netutilts.KeyAccount3,
			oErr:     nil,
			oResult:  &types2.TxResponse{Code: sdkerrors.ErrUnauthorized.ABCICode()},
			d:        "delete asset but don't have the right to do this",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.d, func() {
			out, err := ExCmdDeleteAsset(s.network.Validators[0].ClientCtx, tc.iAssetID,
				netutilts.SHRFee(tc.iTxFee),
				netutilts.JSONFlag,
				netutilts.MakeByAccount(tc.iAcc),
			)
			// expect test case co error
			if tc.oErr != nil {
				s.NotNilf(err, tc.d)
			}
			if tc.oResult != nil {
				txnResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
				s.Equal(tc.oResult.Code, txnResponse.Code, fmt.Sprintf("the case %s: raw log :%s", tc.d, txnResponse.String()))
			}
		})
	}
}

//setupTestMaterial prepare the test data
func (s *AssetIntegrationTestSuite) setupTestMaterial() {

	var assetList = []GenesisAsset{
		{
			AssetCreator: netutilts.KeyAccount1,
			AssetHash:    "0b44517f76eab863c3c4ab13c9774fadd62080ee97af663ec76ffeb671e9e064",
			AssetID:      "e_id_1",
			AssetStatus:  "true",
			AssetRate:    "3",
		},
		{
			AssetCreator: netutilts.KeyAccount1,
			AssetHash:    "1a70dac5cb91142a8b5eda54dbeed46a42d1d7af83788c05fa43485acba09fa8",
			AssetID:      "e_id_2",
			AssetStatus:  "true",
			AssetRate:    "3",
		},
		{
			AssetCreator: netutilts.KeyAccount1,
			AssetHash:    "70cd57b00247b2eb2dfdcf2b30fa73fd79447c23fa6aa512209fab5b6dca4fd2",
			AssetID:      "e_id_3",
			AssetStatus:  "true",
			AssetRate:    "3",
		},
	}
	for _, ex := range assetList {
		out, err := ExCmdCreateAsset(s.network.Validators[0].ClientCtx,
			ex.AssetID, ex.AssetHash, ex.AssetStatus, ex.AssetRate,
			netutilts.SHRFee(6),
			netutilts.JSONFlag,
			netutilts.MakeByAccount(ex.AssetCreator),
		)
		s.NoError(err, "starting create asset fail")
		txOut := netutilts.ParseStdOut(s.T(), out.Bytes())
		s.Equal(netutilts.ShareLedgerSuccessCode, txOut.Code)
	}
	s.existedAsset = assetList
}
