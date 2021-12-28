package tests

import (
	"fmt"

	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/network"
)

type AssetIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewAssetIntegrationTestSuite(cfg network.Config) *AssetIntegrationTestSuite {
	return &AssetIntegrationTestSuite{cfg: cfg}
}

func (s *AssetIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
	s.T().Log("setting up integration test suite successfully")
}
func (s *AssetIntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
}

func (s *AssetIntegrationTestSuite) TestCreateAsset() {

	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "id1"
	assetHash2 := "cc6f58bd1ada876f0a4941ad579908easa726d6da"
	assetStatus := "true"
	SHRPFee := "2"

	s.Run("create_the_valid_asset_it_should_be_success", func() {

		stdOut, err := ExCmdCreateAsset(s.network.Validators[0].ClientCtx, assetID, assetHash, assetStatus, SHRPFee,
			network.SHRFee10(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount1),
		)
		s.Require().NoError(err, "create asset must success")
		txResponse := network.ParseStdOut(s.T(), stdOut.Bytes())
		s.Equalf(network.ShareLedgerSuccessCode, txResponse.Code, "some thing wrong %v", stdOut.String())
		_ = s.network.WaitForNextBlock()
		cmdQueryResponse, err := ExCmdGetAsset(s.network.Validators[0].ClientCtx, assetID,
			network.JSONFlag(),
		)
		asset := AssetJsonUnmarshal(s.T(), cmdQueryResponse.Bytes())
		s.T().Logf("the asset %v", asset)
		s.Assert().Equal(assetID, asset.UUID, "asset UUID is not equal")
		s.Assert().Equal(assetHash, string(asset.Hash), "asset hash is not equal")
		s.Assert().Equal(assetStatus, fmt.Sprintf("%t", asset.Status), "asset status is not equal")
		s.Assert().Equal(SHRPFee, fmt.Sprintf("%d", asset.Rate), "asset rate is not equal")
	})
	s.Run("create_duplicate_the_asset", func() {
		_ = s.network.WaitForNextBlock()
		out, err := ExCmdCreateAsset(s.network.Validators[0].ClientCtx, assetID, assetHash2, assetStatus, SHRPFee,
			network.SHRFee10(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount1),
		)
		s.Assert().NoError(err)

		_ = s.network.WaitForNextBlock()

		txnResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Assert().Equalf(network.ShareLedgerErrorCodeAssetAlreadyExisted, txnResponse.Code, "response after create asset %s", out)

		cmdQueryResponse, err := ExCmdGetAsset(s.network.Validators[0].ClientCtx, assetID)

		asset := AssetJsonUnmarshal(s.T(), cmdQueryResponse.Bytes())

		s.Assert().Equal(assetID, asset.UUID)
		s.Assert().NotEqual(assetHash2, asset.Hash)

	})
}

func (s *AssetIntegrationTestSuite) TestUpdateAsset() {

	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da_new"
	asset2 := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "id1"
	assetStatus := "false"
	SHRPFee := "3"

	_, err := ExCmdCreateAsset(s.network.Validators[0].ClientCtx, assetID, asset2, assetStatus, SHRPFee,
		network.SHRFee10(),
		network.JSONFlag(),
		network.MakeByAccount(network.KeyAccount1),
	)
	s.NoError(err)

	s.Run("update_the_asset_success", func() {
		out, err := ExCmdUpdateAsset(s.network.Validators[0].ClientCtx, assetID, assetHash, assetStatus, SHRPFee,
			network.SHRFee6(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount1),
		)
		s.NoError(err)
		txnResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerSuccessCode, txnResponse.Code)
		_ = s.network.WaitForNextBlock()
		assetByte, err := ExCmdGetAsset(s.network.Validators[0].ClientCtx, assetID)
		s.NoError(err)
		asset := AssetJsonUnmarshal(s.T(), assetByte.Bytes())
		s.Equal(assetHash, string(asset.Hash))
		s.Equal(assetID, asset.UUID)
		s.Equal(assetStatus, fmt.Sprintf("%t", asset.Status))
		s.Equal(SHRPFee, fmt.Sprintf("%d", asset.Rate))
	})
	s.Run("update_the_asset_not_found_asset_should_be_fail", func() {

		out, err := ExCmdUpdateAsset(s.network.Validators[0].ClientCtx, "assetID", assetHash, assetStatus, SHRPFee,
			network.SHRFee6(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount1),
		)
		s.NoError(err)
		txnResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerErrorCodeAssetNotExisted, txnResponse.Code)

	})

	s.Run("update_the_asset_but_updater_is_not_owner_of_asset", func() {
		out, err := ExCmdUpdateAsset(s.network.Validators[0].ClientCtx, assetID, "newhash", assetStatus, SHRPFee,
			network.SHRFee6(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount2))
		s.NoError(err)
		txnResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerErrorCodeUnauthorized, txnResponse.Code)
		_ = s.network.WaitForNextBlock()

		//validating to ensure the asset isn't changed
		assetByte, err := ExCmdGetAsset(s.network.Validators[0].ClientCtx, assetID)
		s.NoError(err)
		asset := AssetJsonUnmarshal(s.T(), assetByte.Bytes())
		s.NotEqual("newhash", string(asset.Hash))
	})

}

func (s *AssetIntegrationTestSuite) TestDeleteAsset() {

	assetID := "id1"

	s.Run("delete_asset_success", func() {

		out, err := ExCmdDeleteAsset(s.network.Validators[0].ClientCtx, assetID,
			network.SHRFee4(),
			network.JSONFlag(),
			network.MakeByAccount(network.KeyAccount1),
		)
		s.NoError(err)
		txnResponse := network.ParseStdOut(s.T(), out.Bytes())
		s.Equal(network.ShareLedgerSuccessCode, txnResponse.Code)
		_ = s.network.WaitForNextBlock()
	})

}
