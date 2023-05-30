package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/app/params"
	assetTypes "github.com/sharering/shareledger/x/asset/types"
	"github.com/sharering/shareledger/x/booking/keeper"
	bookingtestutl "github.com/sharering/shareledger/x/booking/testutil"
	"github.com/sharering/shareledger/x/booking/types"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx           sdk.Context
	queryClient   types.QueryClient
	msgServer     types.MsgServer
	bookingKeeper *keeper.Keeper
	assetKeeper   *bookingtestutl.MockAssetKeeper
	bankKeeper    *bookingtestutl.MockBankKeeper

	encCfg moduletestutil.TestEncodingConfig
}

var (
	bookerAddress           = "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak"
	UUID                    = "UUID"
	defaultMsgCreateBooking = &types.MsgCreateBooking{
		Booker:   bookerAddress,
		UUID:     UUID,
		Duration: 1,
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

	ctrl := gomock.NewController(s.T())
	s.assetKeeper = bookingtestutl.NewMockAssetKeeper(ctrl)
	s.bankKeeper = bookingtestutl.NewMockBankKeeper(ctrl)
	bookingKeeper := keeper.NewKeeper(encCfg.Codec, key, s.assetKeeper, s.bankKeeper)

	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, bookingKeeper)

	s.encCfg = encCfg
	s.ctx = ctx
	s.msgServer = keeper.NewMsgServerImpl(*bookingKeeper)
	s.queryClient = types.NewQueryClient(queryHelper)
	s.bookingKeeper = bookingKeeper
}

func (s *KeeperTestSuite) TestCreateBooking() {
	// asset not found
	s.assetKeeper.EXPECT().GetAsset(s.ctx, UUID).Return(assetTypes.Asset{}, false)
	_, err := s.msgServer.CreateBooking(s.ctx, defaultMsgCreateBooking)
	s.ErrorContains(err, types.ErrAssetDoesNotExist.Error())
	// valid
	s.createDefaultBooking()
}

func (s *KeeperTestSuite) TestBooking() {
	// not found
	_, err := s.queryClient.Booking(s.ctx, nil)
	s.ErrorContains(err, codes.NotFound.String())
	s.createDefaultBooking()
	bookID, err := keeper.GenBookID(defaultMsgCreateBooking)
	s.NoError(err)
	_, err = s.queryClient.Booking(s.ctx, &types.QueryBookingRequest{
		BookID: bookID,
	})
	s.NoError(err)
}

func (s *KeeperTestSuite) TestCompleteBooking() {
	// not found
	_, err := s.msgServer.CompleteBooking(s.ctx, nil)
	s.ErrorContains(err, types.ErrBookingDoesNotExist.Error())
	// valid
	s.createDefaultBooking()
	bookID, err := keeper.GenBookID(defaultMsgCreateBooking)
	s.NoError(err)
	s.assetKeeper.EXPECT().GetAsset(s.ctx, UUID).Return(assetTypes.Asset{
		Creator: "shareledger18pf3zdwqjntd9wkvfcjvmdc7hua6c0q2eck5h5",
		UUID:    UUID,
		Status:  false,
		Rate:    1,
	}, true)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, gomock.Any(), gomock.Any()).Return(nil)
	s.assetKeeper.EXPECT().SetAssetStatus(s.ctx, UUID, true)
	_, err = s.msgServer.CompleteBooking(s.ctx, &types.MsgCompleteBooking{
		Booker: bookerAddress,
		BookID: bookID,
	})
	s.NoError(err)
}

func (s *KeeperTestSuite) createDefaultBooking() {
	s.assetKeeper.EXPECT().GetAsset(s.ctx, UUID).Return(assetTypes.Asset{UUID: UUID, Status: true, Rate: 1}, true)
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(s.ctx, gomock.Any(), types.ModuleName, gomock.Any()).Return(nil)
	s.assetKeeper.EXPECT().SetAssetStatus(s.ctx, UUID, false)
	_, err := s.msgServer.CreateBooking(s.ctx, defaultMsgCreateBooking)
	s.NoError(err)
}
