package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/x/asset/keeper"
	"github.com/sharering/shareledger/x/asset/types"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	queryClient types.QueryClient
	msgServer   types.MsgServer
	assetKeeper keeper.Keeper
	encCfg      moduletestutil.TestEncodingConfig
}

var (
	creatorAddress        = "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak"
	UUID                  = "UUID"
	defaultMsgCreateAsset = &types.MsgCreateAsset{
		Creator: creatorAddress,
		Hash:    []byte{},
		UUID:    UUID,
		Status:  true,
		Rate:    1,
	}
)

func TestKeeperTestSuite(t *testing.T) {
	params.SetAddressPrefixes()
	t.Log("Running TestKeeperTestSuite for asset module...")
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig()
	s.assetKeeper = *keeper.NewKeeper(encCfg.Codec, key)
	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, s.encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, s.assetKeeper)

	s.encCfg = encCfg
	s.ctx = ctx
	s.msgServer = keeper.NewMsgServerImpl(s.assetKeeper)
	s.queryClient = types.NewQueryClient(queryHelper)
}

func (s *KeeperTestSuite) TestCreateAsset() {
	s.createDefaultAsset()
	// existed
	_, err := s.msgServer.CreateAsset(s.ctx, defaultMsgCreateAsset)
	s.ErrorContains(err, types.ErrAssetExist.Error())
}

func (s *KeeperTestSuite) TestDeleteAsset() {
	req := &types.MsgDeleteAsset{
		Owner: creatorAddress,
		UUID:  UUID,
	}
	// not found
	_, err := s.msgServer.DeleteAsset(s.ctx, req)
	s.ErrorContains(err, sdkerrors.ErrNotFound.Error())
	// valid
	s.createDefaultAsset()
	_, err = s.msgServer.DeleteAsset(s.ctx, req)
	s.NoError(err)
	// Unauthorized
	s.createDefaultAsset()
	req.Owner = "shareledger1j4ndn4qed0ukulc0a6a5fxe6yxmnm7wh59pwpn"
	_, err = s.msgServer.DeleteAsset(s.ctx, req)
	s.ErrorContains(err, sdkerrors.ErrUnauthorized.Error())
}

func (s *KeeperTestSuite) TestUpdateAsset() {
	req := &types.MsgUpdateAsset{
		Creator: creatorAddress,
		Hash:    []byte{},
		UUID:    UUID,
		Status:  false,
		Rate:    10000,
	}
	// not found
	_, err := s.msgServer.UpdateAsset(s.ctx, req)
	s.ErrorContains(err, types.ErrNameDoesNotExist.Error())
	// valid
	s.createDefaultAsset()
	_, err = s.msgServer.UpdateAsset(s.ctx, req)
}

func (s *KeeperTestSuite) TestAssetByUUID() {
	req := &types.QueryAssetByUUIDRequest{
		Uuid: UUID,
	}
	// not found
	_, err := s.queryClient.AssetByUUID(s.ctx, req)
	s.ErrorContains(err, codes.NotFound.String())
	// valid
	s.createDefaultAsset()
	res, err := s.queryClient.AssetByUUID(s.ctx, req)
	s.NoError(err)
	s.Equal(UUID, res.Asset.UUID)
}

func (s *KeeperTestSuite) createDefaultAsset() {
	_, err := s.msgServer.CreateAsset(s.ctx, defaultMsgCreateAsset)
	s.NoError(err)
}
