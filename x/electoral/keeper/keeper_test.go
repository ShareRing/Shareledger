package keeper_test

import (
	"sort"
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/sharering/shareledger/x/electoral/keeper"
	electoralutil "github.com/sharering/shareledger/x/electoral/testutil"
	"github.com/sharering/shareledger/x/electoral/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx             sdk.Context
	queryClient     types.QueryClient
	msgServer       types.MsgServer
	electoralKeeper *keeper.Keeper
	encCfg          moduletestutil.TestEncodingConfig
	gk              *electoralutil.MockGentlemintKeeper
}

var (
	shrAddress = "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak"

	invalidReq = "invalid request"
)

func TestKeeperTestSuite(t *testing.T) {
	params.SetAddressPrefixes()
	t.Log("Running TestKeeperTestSuite for electoral module...")
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	encCfg := moduletestutil.MakeTestEncodingConfig()

	ctrl := gomock.NewController(s.T())
	gk := electoralutil.NewMockGentlemintKeeper(ctrl)
	electoralKeeper := keeper.NewKeeper(encCfg.Codec, key, gk)

	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, electoralKeeper)

	s.ctx = ctx
	s.encCfg = encCfg
	s.electoralKeeper = electoralKeeper
	s.gk = gk
	s.queryClient = types.NewQueryClient(queryHelper)
	s.msgServer = keeper.NewMsgServerImpl(*electoralKeeper)
}

func (s *KeeperTestSuite) TestAccState() {
	a := createNAccState(s.electoralKeeper, s.ctx, 1)
	accState, err := s.queryClient.AccState(s.ctx, &types.QueryAccStateRequest{Key: a[0].Key})

	// success
	s.Require().NoError(err)

	res := types.QueryAccStateResponse{
		AccState: a[0],
	}
	s.Require().Equal(&res, accState)

	// fail
	accState, err = s.electoralKeeper.AccState(s.ctx, &types.QueryAccStateRequest{Key: "invalid"})
	s.Require().Error(err)
	s.Require().Nil(accState)
}

func (s *KeeperTestSuite) TestAccStates() {
	a := s.createNAccRole(1, types.AccStateKeyAccOp)
	accStates, err := s.electoralKeeper.AccStates(s.ctx, &types.QueryAccStatesRequest{})
	// success
	s.Require().NoError(err)

	res := &types.QueryAccStatesResponse{
		AccState: a,
		Pagination: &query.PageResponse{
			Total:   uint64(len(a)),
			NextKey: nil,
		},
	}
	s.Require().Equal(res, accStates)

	// invalid request
	var req *types.QueryAccStatesRequest
	accStates, err = s.electoralKeeper.AccStates(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Nil(accStates)
}

func (s *KeeperTestSuite) TestGetAllAccState() {
	a := createNAccState(s.electoralKeeper, s.ctx, 2)
	list := s.electoralKeeper.GetAllAccState(s.ctx)

	// success
	s.Require().Equal(a, list)
}

func (s *KeeperTestSuite) TestGetAccState() {
	addr, _ := sdk.AccAddressFromBech32(shrAddress)

	key := types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner)
	s.electoralKeeper.SetAccState(s.ctx, types.AccState{
		Address: addr.String(),
		Key:     string(key),
	})
	accState, found := s.electoralKeeper.GetAccState(s.ctx, key)

	// success
	s.Require().True(found)
	s.Require().Equal(accState.Address, addr.String())

	// fail
	accState, found = s.electoralKeeper.GetAccState(s.ctx, "invalid")
	s.Require().False(found)
	s.Require().Equal(accState, types.AccState{})
}

func (s *KeeperTestSuite) TestSetAccState() {
	var acc types.AccState

	addr, _ := sdk.AccAddressFromBech32(shrAddress)
	acc.Address = addr.String()
	acc.Key = string(types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner))
	s.electoralKeeper.SetAccState(s.ctx, acc)

	// success
	key := types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner)
	acc, f := s.electoralKeeper.GetAccState(s.ctx, key)
	s.Require().True(f)
	s.Require().Equal(acc.Address, addr.String())
}

func (s *KeeperTestSuite) TestRemoveAccState() {
	a := s.createNAccRole(1, types.AccStateKeyIdSigner)
	key := types.GenAccStateIndexKey(sdk.AccAddress(a[0].Address), types.AccStateKeyIdSigner)
	s.electoralKeeper.RemoveAccState(s.ctx, types.IndexKeyAccState(types.AccStateKeyIdSigner))

	// success
	accState, found := s.electoralKeeper.GetAccState(s.ctx, key)
	s.Require().False(found)
	s.Require().Equal(accState, types.AccState{})
}

func (s *KeeperTestSuite) TestActiveAccState() {
	addr, _ := sdk.AccAddressFromBech32(shrAddress)

	s.electoralKeeper.ActiveAccState(s.ctx, addr, types.AccStateKeyIdSigner)

	acc, f := s.electoralKeeper.GetAccState(s.ctx, types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner))
	s.Require().True(f)
	s.Require().Equal("active", acc.Status)
}

func (s *KeeperTestSuite) TestAccountOperator() {
	a := s.createNAccRole(1, types.AccStateKeyAccOp)
	req := &types.QueryAccountOperatorRequest{
		Address: a[0].Address,
	}

	acc, err := s.electoralKeeper.AccountOperator(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryAccountOperatorResponse{
		AccState: &types.AccState{
			Address: a[0].Address,
			Key:     a[0].Key,
		},
	}

	s.Require().Equal(res, acc)

	// invalid request
	var req2 *types.QueryAccountOperatorRequest
	acc, err = s.electoralKeeper.AccountOperator(s.ctx, req2)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) createNAccRole(n int, role types.AccStateKeyType) []types.AccState {
	items := make([]types.AccState, n)
	for i := range items {
		addr, _ := sdk.AccAddressFromBech32(sample.AccAddress())
		items[i].Key = string(types.GenAccStateIndexKey(addr, role))
		items[i].Address = addr.String()

		s.electoralKeeper.SetAccState(s.ctx, items[i])
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Key < items[j].Key
	})

	return items
}

func (s *KeeperTestSuite) TestAccountOperators() {
	a := s.createNAccRole(1, types.AccStateKeyAccOp)
	accOp, err := s.electoralKeeper.AccountOperators(s.ctx, &types.QueryAccountOperatorsRequest{})
	res := &types.QueryAccountOperatorsResponse{
		AccStates: []*types.AccState{
			{
				Address: a[0].Address,
				Key:     a[0].Key,
				Status:  a[0].Status,
			},
		},
	}
	s.Require().NoError(err)
	s.Require().Equal(accOp, res)
}

func (s *KeeperTestSuite) TestApprover() {
	a := s.createNAccRole(1, types.AccStateKeyApprover)
	req := &types.QueryApproverRequest{
		Address: a[0].Address,
	}

	acc, err := s.electoralKeeper.Approver(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryApproverResponse{
		AccState: types.AccState{
			Address: a[0].Address,
			Key:     a[0].Key,
			Status:  a[0].Status,
		},
	}

	s.Require().Equal(res, acc)

	// request nil - invalid request
	req = nil
	acc, err = s.electoralKeeper.Approver(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)

	// notfound
	req = &types.QueryApproverRequest{
		Address: sample.AccAddress(),
	}
	acc, err = s.electoralKeeper.Approver(s.ctx, req)
	s.Require().Error(err, "not found")
	s.Require().Nil(acc)

	// invalid address
	req = &types.QueryApproverRequest{
		Address: "invalid address",
	}
	acc, err = s.electoralKeeper.Approver(s.ctx, req)
	s.Require().Error(err, "invalid address")
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestApprovers() {
	a := s.createNAccRole(1, types.AccStateKeyApprover)
	req := &types.QueryApproversRequest{}

	acc, err := s.electoralKeeper.Approvers(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryApproversResponse{
		Approvers: []*types.AccState{
			{
				Address: a[0].Address,
				Key:     a[0].Key,
				Status:  a[0].Status,
			},
		},
	}

	s.Require().Equal(res, acc)
	// request nil - invalid request
	req = nil
	acc, err = s.electoralKeeper.Approvers(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestDocumentIssuers() {
	a := s.createNAccRole(1, types.AccStateKeyDocIssuer)
	req := &types.QueryDocumentIssuersRequest{}

	acc, err := s.electoralKeeper.DocumentIssuers(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryDocumentIssuersResponse{
		AccStates: []*types.AccState{
			{
				Address: a[0].Address,
				Key:     a[0].Key,
				Status:  a[0].Status,
			},
		},
	}

	s.Require().Equal(res, acc)
	// request nil - invalid request
	req = nil
	acc, err = s.electoralKeeper.DocumentIssuers(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestDocumentIssuer() {
	a := s.createNAccRole(1, types.AccStateKeyDocIssuer)

	req := &types.QueryDocumentIssuerRequest{
		Address: a[0].Address,
	}
	acc, err := s.electoralKeeper.DocumentIssuer(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryDocumentIssuerResponse{
		AccState: &types.AccState{
			Address: a[0].Address,
			Key:     a[0].Key,
			Status:  a[0].Status,
		},
	}
	s.Require().Equal(res, acc)
	// request nil - invalid request
	req = nil
	acc, err = s.electoralKeeper.DocumentIssuer(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestSetAuthority() {
	a := s.createNAccRole(1, types.AccStateKeyIdSigner)

	s.electoralKeeper.SetAuthority(s.ctx, types.Authority{Address: a[0].Address})

	auth, f := s.electoralKeeper.GetAuthority(s.ctx)
	s.Require().True(f)
	s.Require().Equal(auth.Address, a[0].Address)
}

func (s *KeeperTestSuite) TestIDSigners() {
	a := s.createNAccRole(1, types.AccStateKeyIdSigner)

	acc, err := s.electoralKeeper.IdSigners(s.ctx, &types.QueryIdSignersRequest{})
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryIdSignersResponse{
		AccStates: []*types.AccState{
			{
				Address: a[0].Address,
				Key:     a[0].Key,
				Status:  a[0].Status,
			},
		},
	}

	s.Require().Equal(res, acc)

	// request nil - invalid request
	var req *types.QueryIdSignersRequest
	acc, err = s.electoralKeeper.IdSigners(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestIdSigner() {
	a := s.createNAccRole(1, types.AccStateKeyIdSigner)

	req := &types.QueryIdSignerRequest{
		Address: a[0].Address,
	}
	acc, err := s.electoralKeeper.IdSigner(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryIdSignerResponse{
		AccState: &types.AccState{
			Address: a[0].Address,
			Key:     a[0].Key,
			Status:  a[0].Status,
		},
	}

	s.Require().Equal(res, acc)
	// request nil - invalid request
	req = nil
	acc, err = s.electoralKeeper.IdSigner(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestLoader() {
	a := s.createNAccRole(1, types.AccStateKeyShrpLoaders)

	acc, err := s.electoralKeeper.Loader(s.ctx, &types.QueryLoaderRequest{Address: a[0].Address})
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - invalid request
	var req *types.QueryLoaderRequest
	acc, err = s.electoralKeeper.Loader(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestLoaders() {
	a := s.createNAccRole(1, types.AccStateKeyShrpLoaders)

	acc, err := s.electoralKeeper.Loaders(s.ctx, &types.QueryLoadersRequest{})
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryLoadersResponse{
		Loaders: []*types.AccState{
			{
				Address: a[0].Address,
				Key:     a[0].Key,
				Status:  a[0].Status,
			},
		},
	}

	s.Require().Equal(res, acc)

	// request nil - invalid request
	var req *types.QueryLoadersRequest
	acc, err = s.electoralKeeper.Loaders(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestRelayer() {
	a := s.createNAccRole(1, types.AccStateKeyRelayer)

	acc, err := s.electoralKeeper.Relayer(s.ctx, &types.QueryRelayerRequest{Address: a[0].Address})
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - invalid request
	var req *types.QueryRelayerRequest
	acc, err = s.electoralKeeper.Relayer(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestRelayers() {
	a := s.createNAccRole(1, types.AccStateKeyRelayer)

	acc, err := s.electoralKeeper.Relayers(s.ctx, &types.QueryRelayersRequest{})
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryRelayersResponse{
		Relayers: []*types.AccState{
			{
				Address: a[0].Address,
				Key:     a[0].Key,
				Status:  a[0].Status,
			},
		},
	}

	s.Require().Equal(res, acc)

	// request nil - invalid request
	var req *types.QueryRelayersRequest
	acc, err = s.electoralKeeper.Relayers(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestSwapManager() {
	a := s.createNAccRole(1, types.AccStateKeySwapManager)

	acc, err := s.electoralKeeper.SwapManager(s.ctx, &types.QuerySwapManagerRequest{Address: a[0].Address})
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - invalid request
	var req *types.QuerySwapManagerRequest
	acc, err = s.electoralKeeper.SwapManager(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestSwapManagers() {
	a := s.createNAccRole(1, types.AccStateKeySwapManager)

	acc, err := s.electoralKeeper.SwapManagers(s.ctx, &types.QuerySwapManagersRequest{})
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QuerySwapManagersResponse{
		SwapManagers: []*types.AccState{
			{
				Address: a[0].Address,
				Key:     a[0].Key,
				Status:  a[0].Status,
			},
		},
	}

	s.Require().Equal(res, acc)

	// request nil - invalid request
	var req *types.QuerySwapManagersRequest
	acc, err = s.electoralKeeper.SwapManagers(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestVoter() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)

	acc, err := s.electoralKeeper.Voter(s.ctx, &types.QueryVoterRequest{Address: a[0].Address})
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - invalid request
	var req *types.QueryVoterRequest
	acc, err = s.electoralKeeper.Voter(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestVoters() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)

	acc, err := s.electoralKeeper.Voters(s.ctx, &types.QueryVotersRequest{})
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	res := &types.QueryVotersResponse{
		Voters: []*types.AccState{
			{
				Address: a[0].Address,
				Key:     a[0].Key,
				Status:  a[0].Status,
			},
		},
	}

	s.Require().Equal(res, acc)

	// request nil - invalid request
	var req *types.QueryVotersRequest
	acc, err = s.electoralKeeper.Voters(s.ctx, req)
	s.Require().Error(err, invalidReq)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestEnrollAccountOperators() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)
	req := &types.MsgEnrollAccountOperators{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}

	acc, err := s.msgServer.EnrollAccountOperators(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgEnrollAccountOperators
	acc, err = s.msgServer.EnrollAccountOperators(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestEnrollApprovers() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)
	req := &types.MsgEnrollApprovers{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}

	acc, err := s.msgServer.EnrollApprovers(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)
}

func (s *KeeperTestSuite) TestEnrollDocIssuers() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)
	req := &types.MsgEnrollDocIssuers{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}

	acc, err := s.msgServer.EnrollDocIssuers(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgEnrollDocIssuers
	acc, err = s.msgServer.EnrollDocIssuers(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)

}

func (s *KeeperTestSuite) TestEnrollIdSigners() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)
	req := &types.MsgEnrollIdSigners{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}

	acc, err := s.msgServer.EnrollIdSigners(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgEnrollIdSigners
	acc, err = s.msgServer.EnrollIdSigners(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

// func (s *KeeperTestSuite) TestEnrollLoaders() {
// 	a := s.createNAccRole(1, types.AccStateKeyVoter)
// 	req := &types.MsgEnrollLoaders{
// 		Creator: a[0].Address,
// 		Addresses: []string{
// 			a[0].Address,
// 		},
// 	}
// 	s.gk.EXPECT().LoadAllowanceLoader(s.ctx, a[0].Address).Return(nil)
// 	acc, err := s.msgServer.EnrollLoaders(s.ctx, req)
// 	s.Require().NoError(err)
// 	s.Require().NotNil(acc)

// 	// request nil - validate fail
// 	var reqNil types.MsgEnrollLoaders
// 	acc, err = s.msgServer.EnrollLoaders(s.ctx, &reqNil)
// 	s.Require().Error(err)
// 	s.Require().Nil(acc)
// }

func (s *KeeperTestSuite) TestEnrollSwapManagers() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)
	req := &types.MsgEnrollSwapManagers{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}

	acc, err := s.msgServer.EnrollSwapManagers(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	reqFail := &types.MsgEnrollSwapManagers{
		Creator: a[0].Address,
		Addresses: []string{
			"invalid address",
		},
	}
	acc, err = s.msgServer.EnrollSwapManagers(s.ctx, reqFail)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestEnrollRelayers() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)
	req := &types.MsgEnrollRelayers{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}

	acc, err := s.msgServer.EnrollRelayers(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

}

func (s *KeeperTestSuite) TestEnrollVoter() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)
	req := &types.MsgEnrollVoter{
		Creator: a[0].Address,
		Address: a[0].Address,
	}

	acc, err := s.msgServer.EnrollVoter(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgEnrollVoter
	acc, err = s.msgServer.EnrollVoter(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestRevokeAccountOperators() {
	a := s.createNAccRole(1, types.AccStateKeyAccOp)
	req := &types.MsgRevokeAccountOperators{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}

	acc, err := s.msgServer.RevokeAccountOperators(s.ctx, req)
	s.Require().NoError(err)
	s.Require().Equal(acc, &types.MsgRevokeAccountOperatorsResponse{})

	// // request nil - validate fail
	var reqNil types.MsgRevokeAccountOperators
	_, err = s.msgServer.RevokeAccountOperators(s.ctx, &reqNil)
	s.Require().Error(err)
}

func (s *KeeperTestSuite) TestRevokeApprovers() {
	a := s.createNAccRole(1, types.AccStateKeyApprover)
	req := &types.MsgRevokeApprovers{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}

	acc, err := s.msgServer.RevokeApprovers(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgRevokeApprovers
	acc, err = s.msgServer.RevokeApprovers(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestRevokeDocIssuers() {
	a := s.createNAccRole(1, types.AccStateKeyDocIssuer)
	req := &types.MsgRevokeDocIssuers{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}

	acc, err := s.msgServer.RevokeDocIssuers(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgRevokeDocIssuers
	acc, err = s.msgServer.RevokeDocIssuers(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestRevokeIdSigners() {
	a := s.createNAccRole(1, types.AccStateKeyIdSigner)
	req := &types.MsgRevokeIdSigners{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}
	acc, err := s.msgServer.RevokeIdSigners(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgRevokeIdSigners
	acc, err = s.msgServer.RevokeIdSigners(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestRevokeLoaders() {
	a := s.createNAccRole(1, types.AccStateKeyShrpLoaders)
	req := &types.MsgRevokeLoaders{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}
	acc, err := s.msgServer.RevokeLoaders(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgRevokeLoaders
	acc, err = s.msgServer.RevokeLoaders(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestRevokeRelayers() {
	a := s.createNAccRole(1, types.AccStateKeyRelayer)
	req := &types.MsgRevokeRelayers{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}
	acc, err := s.msgServer.RevokeRelayers(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgRevokeRelayers
	acc, err = s.msgServer.RevokeRelayers(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestRevokeSwapManagers() {
	a := s.createNAccRole(1, types.AccStateKeySwapManager)
	req := &types.MsgRevokeSwapManagers{
		Creator: a[0].Address,
		Addresses: []string{
			a[0].Address,
		},
	}
	acc, err := s.msgServer.RevokeSwapManagers(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgRevokeSwapManagers
	acc, err = s.msgServer.RevokeSwapManagers(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}

func (s *KeeperTestSuite) TestRevokeVoter() {
	a := s.createNAccRole(1, types.AccStateKeyVoter)
	req := &types.MsgRevokeVoter{
		Creator: a[0].Address,
		Address: a[0].Address,
	}
	acc, err := s.msgServer.RevokeVoter(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(acc)

	// request nil - validate fail
	var reqNil types.MsgRevokeVoter
	acc, err = s.msgServer.RevokeVoter(s.ctx, &reqNil)
	s.Require().Error(err)
	s.Require().Nil(acc)
}
