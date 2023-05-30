package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"

	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/x/id/keeper"
	"github.com/sharering/shareledger/x/id/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context

	queryClient types.QueryClient
	msgServer   types.MsgServer
	idKeeper    keeper.Keeper

	encCfg moduletestutil.TestEncodingConfig
}

var (
	idID            = "ID123"
	issuerAddress   = "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak"
	ownerAddress    = "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr"
	newOwnerAddress = "shareledger16k53l0v5fr93u2kpe7ek2lc3tm37k4954zalvv"
	createIdReq     = &types.MsgCreateId{
		Id:            idID,
		IssuerAddress: issuerAddress,
		BackupAddress: "",
		ExtraData:     "",
		OwnerAddress:  ownerAddress,
	}
)

func TestKeeperTestSuite(t *testing.T) {
	params.SetAddressPrefixes()
	t.Log("Running TestKeeperTestSuite for id module...")
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig()
	s.idKeeper = *keeper.NewKeeper(encCfg.Codec, key)

	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, s.encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, s.idKeeper)

	s.encCfg = encCfg
	s.ctx = ctx
	s.msgServer = keeper.NewMsgServerImpl(s.idKeeper)
	s.queryClient = types.NewQueryClient(queryHelper)
}

func (s *KeeperTestSuite) TestGetIDByID() {
	_, err := s.msgServer.CreateId(s.ctx, createIdReq)
	s.NoError(err)
	resp, err := s.queryClient.IdById(s.ctx, &types.QueryIdByIdRequest{
		Id: idID,
	})
	s.NoError(err)
	s.Equal(idID, resp.Id.Id)
	s.Equal(ownerAddress, resp.Id.Data.OwnerAddress)
}

func (s *KeeperTestSuite) TestGetByAddress() {
	_, err := s.msgServer.CreateId(s.ctx, createIdReq)
	s.NoError(err)
	resp, err := s.queryClient.IdByAddress(s.ctx, &types.QueryIdByAddressRequest{
		Address: ownerAddress,
	})
	s.NoError(err)
	s.Equal(idID, resp.Id.Id)
	s.Equal(ownerAddress, resp.Id.Data.OwnerAddress)
}

func (s *KeeperTestSuite) TestCreateID() {
	// valid
	_, err := s.msgServer.CreateId(s.ctx, createIdReq)
	s.NoError(err)
	// existed
	_, err = s.msgServer.CreateId(s.ctx, createIdReq)
	s.Error(err)
}

func (s *KeeperTestSuite) TestCreateIDs() {
	_, err := s.msgServer.CreateIds(s.ctx, &types.MsgCreateIds{
		IssuerAddress: issuerAddress,
		BackupAddress: []string{"backup"},
		ExtraData:     []string{"extra1"},
		Id:            []string{idID},
		OwnerAddress:  []string{ownerAddress},
	})
	s.NoError(err)
	_, err = s.msgServer.CreateIds(s.ctx, &types.MsgCreateIds{
		IssuerAddress: issuerAddress,
		BackupAddress: []string{"backup"},
		ExtraData:     []string{"extra1"},
		Id:            []string{idID},
		OwnerAddress:  []string{ownerAddress},
	})
	s.Error(err)
}

func (s *KeeperTestSuite) TestReplaceIdOwner() {
	_, err := s.msgServer.CreateId(s.ctx, createIdReq)
	s.NoError(err)
	_, err = s.msgServer.ReplaceIdOwner(s.ctx, &types.MsgReplaceIdOwner{
		BackupAddress: "",
		Id:            idID,
		OwnerAddress:  newOwnerAddress,
	})
	s.NoError(err)
}

func (s *KeeperTestSuite) TestUpdateID() {
	_, err := s.msgServer.CreateId(s.ctx, createIdReq)
	s.NoError(err)
	_, err = s.msgServer.UpdateId(s.ctx, &types.MsgUpdateId{
		IssuerAddress: ownerAddress,
		Id:            idID,
		ExtraData:     "NewExtraData",
	})
	s.NoError(err)
}
