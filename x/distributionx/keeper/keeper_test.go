package keeper_test

import (
	"testing"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/x/distributionx/keeper"
	dtestutil "github.com/sharering/shareledger/x/distributionx/testutil"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	queryClient types.QueryClient
	msgServer   types.MsgServer
	dKeeper     *keeper.Keeper
	accKeeper   *dtestutil.MockAccountKeeper
	bankKeeper  *dtestutil.MockBankKeeper
	wasmKeeper  *dtestutil.MockWasmKeeper

	encCfg moduletestutil.TestEncodingConfig
}

func TestKeeperTestSuite(t *testing.T) {
	params.SetAddressPrefixes()
	t.Log("Running TestKeeperTestSuite for distributionx module...")
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)
	encCfg := moduletestutil.MakeTestEncodingConfig()

	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	s.Require().NoError(stateStore.LoadLatestVersion())
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	ctx = ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	ctrl := gomock.NewController(s.T())
	s.accKeeper = dtestutil.NewMockAccountKeeper(ctrl)
	s.bankKeeper = dtestutil.NewMockBankKeeper(ctrl)
	s.wasmKeeper = dtestutil.NewMockWasmKeeper(ctrl)

	paramsSubspace := typesparams.NewSubspace(encCfg.Codec,
		types.Amino,
		key,
		memStoreKey,
		"DistributionxParams",
	)

	dKeeper := keeper.NewKeeper(encCfg.Codec, key, paramsSubspace, s.accKeeper, s.bankKeeper, s.wasmKeeper)

	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, dKeeper)

	s.encCfg = encCfg
	s.ctx = ctx
	s.msgServer = keeper.NewMsgServerImpl(*dKeeper)
	s.queryClient = types.NewQueryClient(queryHelper)
	s.dKeeper = dKeeper
}

func (s *KeeperTestSuite) TestLogger() {
	resp := s.dKeeper.Logger(s.ctx)
	s.Require().Empty(resp)
}
