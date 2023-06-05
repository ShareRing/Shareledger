package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/x/gentlemint/keeper"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/stretchr/testify/suite"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	dtestutil "github.com/sharering/shareledger/x/gentlemint/testutil"
)

var (
	testAcc                        = "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf"
	testAcc1                       = "shareledger1mfru9azs5nua2wxcd4sq64g5nt7nn4n85mcr0s"
	tenSHR                         = sdk.NewCoin("shr", sdk.NewInt(10))
	oneThousandCent                = sdk.NewCoin("cent", sdk.NewInt(1000))
	tenMilNSHR                     = sdk.NewCoin("nshr", sdk.NewInt(10000000000))
	zeroNSHR                       = sdk.NewCoin("nshr", sdk.NewInt(0))
	tenDecCoin                     = sdk.NewDecCoin("shr", sdk.NewInt(10))
	fiveHundredThousandNSHRDecCoin = sdk.NewDecCoin("nshr", sdk.NewInt(500000))

	msgTest = types.NewMsgLoad(testAcc, testAcc, sdk.NewDecCoins(tenDecCoin))

	testActionFee = types.ActionLevelFee{
		Action:  "gentlemint_load",
		Level:   "low",
		Creator: "",
	}
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	queryClient types.QueryClient
	msgServer   types.MsgServer
	gKeeper     *keeper.Keeper
	accKeeper   *dtestutil.MockAccountKeeper
	bankKeeper  *dtestutil.MockBankKeeper

	encCfg moduletestutil.TestEncodingConfig
}

func TestKeeperTestSuite(t *testing.T) {
	params.SetAddressPrefixes()
	t.Log("Running TestKeeperTestSuite for gentlemint module...")
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)

	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	s.Require().NoError(stateStore.LoadLatestVersion())
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	ctx = ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	encCfg := moduletestutil.MakeTestEncodingConfig()

	ctrl := gomock.NewController(s.T())
	s.accKeeper = dtestutil.NewMockAccountKeeper(ctrl)
	s.bankKeeper = dtestutil.NewMockBankKeeper(ctrl)

	paramsSubspace := typesparams.NewSubspace(encCfg.Codec,
		types.Amino,
		key,
		memStoreKey,
		"DistributionxParams",
	)

	gKeeper := keeper.NewKeeper(encCfg.Codec, key, s.bankKeeper, s.accKeeper, paramsSubspace)

	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, gKeeper)

	s.encCfg = encCfg
	s.ctx = ctx
	s.msgServer = keeper.NewMsgServerImpl(*gKeeper)
	s.queryClient = types.NewQueryClient(queryHelper)
	s.gKeeper = gKeeper
}

func (s *KeeperTestSuite) TestLogger() {
	resp := s.gKeeper.Logger(s.ctx)
	s.Require().Empty(resp)
}

func (s *KeeperTestSuite) TestBaseMintPossible_false() {
	s.bankKeeper.EXPECT().GetSupply(gomock.Any(), gomock.Any()).Return(tenSHR)
	s.False(s.gKeeper.BaseMintPossible(s.ctx, types.MaxBaseSupply))
}

func (s *KeeperTestSuite) TestBaseMintPossible_true() {
	s.bankKeeper.EXPECT().GetSupply(gomock.Any(), gomock.Any()).Return(tenSHR)
	s.True(s.gKeeper.BaseMintPossible(s.ctx, math.NewInt(1000)))
}

func (s *KeeperTestSuite) TestLoadAllowanceLoader() {
	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.Require().NoError(err)
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), types.ModuleName, gomock.Any()).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, acc, gomock.Any()).Return(nil)
	s.Nil(s.gKeeper.LoadAllowanceLoader(s.ctx, sdk.MustAccAddressFromBech32(testAcc)))
}
