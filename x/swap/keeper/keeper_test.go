package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/x/swap/keeper"
	swaptestutil "github.com/sharering/shareledger/x/swap/testutil"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/stretchr/testify/suite"
)

var (
	tenSHR            = sdk.NewCoin("shr", sdk.NewInt(10))
	tenMilNSHR        = sdk.NewCoin("nshr", sdk.NewInt(10000000000))
	testAddr          = "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf"
	swapModuleAddress = []byte("shareledger1mfru9azs5nua2wxcd4sq64g5nt7nn4n85mcr0s")

	normalRequest = types.Request{
		SrcAddr:     testAddr,
		DestAddr:    "0x97b98d335c28f9ad9c123e344a78f00c84146431",
		SrcNetwork:  types.NetworkNameShareLedger,
		DestNetwork: "eth",
		Amount:      tenSHR,
		Fee:         tenSHR,
		Status:      types.SwapStatusPending,
		BatchId:     0,
	}

	normalRequestWithNSHR = types.Request{
		SrcAddr:     testAddr,
		DestAddr:    "0x97b98d335c28f9ad9c123e344a78f00c84146431",
		SrcNetwork:  types.NetworkNameShareLedger,
		DestNetwork: "eth",
		Amount:      tenMilNSHR,
		Fee:         tenMilNSHR,
		Status:      types.SwapStatusPending,
	}

	emptySrcRequest = types.Request{
		SrcAddr:     testAddr,
		DestAddr:    "0x97b98d335c28f9ad9c123e344a78f00c84146431",
		DestNetwork: "eth",
		Amount:      tenSHR,
		Fee:         tenSHR,
		Status:      types.SwapStatusPending,
	}
	request1 = types.Request{
		Id:          0,
		SrcAddr:     "",
		DestAddr:    testAddr,
		SrcNetwork:  "eth",
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      tenSHR,
		Fee:         tenSHR,
		Status:      types.SwapStatusPending,
		BatchId:     0,
		TxEvents: []*types.TxEvent{{
			TxHash:   "0xXXX",
			Sender:   "",
			LogIndex: 0,
		}},
	}

	request2 = types.Request{
		Id:          0,
		SrcAddr:     "",
		DestAddr:    testAddr,
		SrcNetwork:  "eth",
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      tenSHR,
		Fee:         tenSHR,
		Status:      types.SwapStatusPending,
		BatchId:     0,
		TxEvents: []*types.TxEvent{{
			TxHash:   "0xABC",
			Sender:   "",
			LogIndex: 0,
		}},
	}

	msgApproveIn = types.MsgApproveIn{
		Creator: "",
		Ids:     []uint64{0},
	}
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	queryClient types.QueryClient
	msgServer   types.MsgServer
	swapKeeper  *keeper.Keeper
	accKeeper   *swaptestutil.MockAccountKeeper
	bankKeeper  *swaptestutil.MockBankKeeper
	gmKeeper    *swaptestutil.MockGentlemintKeeper

	encCfg moduletestutil.TestEncodingConfig
}

func TestKeeperTestSuite(t *testing.T) {
	params.SetAddressPrefixes()
	t.Log("Running TestKeeperTestSuite for swap module...")
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig()
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	ctrl := gomock.NewController(s.T())
	s.accKeeper = swaptestutil.NewMockAccountKeeper(ctrl)
	s.bankKeeper = swaptestutil.NewMockBankKeeper(ctrl)
	s.gmKeeper = swaptestutil.NewMockGentlemintKeeper(ctrl)

	paramsSubspace := typesparams.NewSubspace(encCfg.Codec,
		types.Amino,
		key,
		memStoreKey,
		"SwapParams",
	)

	swapKeeper := keeper.NewKeeper(encCfg.Codec, key, paramsSubspace, s.bankKeeper, s.accKeeper)

	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, swapKeeper)

	s.encCfg = encCfg
	s.ctx = ctx
	s.msgServer = keeper.NewMsgServerImpl(*swapKeeper)
	s.queryClient = types.NewQueryClient(queryHelper)
	s.swapKeeper = swapKeeper
}

func (s *KeeperTestSuite) TestSetupMsgServer_ok() {
	s.swapKeeper.ImportRequest(s.ctx, []types.Request{normalRequest})
	stores := s.swapKeeper.GetStoreRequestMap(s.ctx)
	store, ok := stores[normalRequest.Status]
	s.Require().True(ok)
	found := store.Has(keeper.GetRequestIDBytes(normalRequest.Id))
	s.Require().True(found)
}

func (s *KeeperTestSuite) TestSetupMsgServer_nOk() {
	s.swapKeeper.ImportRequest(s.ctx, []types.Request{normalRequest})
	stores := s.swapKeeper.GetStoreRequestMap(s.ctx)
	store, ok := stores[normalRequest.Status]
	s.Require().True(ok)
	found := store.Has(keeper.GetRequestIDBytes(1))
	s.Require().True(!found)
}
