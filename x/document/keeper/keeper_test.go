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
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"

	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/x/document/keeper"
	documenttestutil "github.com/sharering/shareledger/x/document/testutil"
	"github.com/sharering/shareledger/x/document/types"
)

var (
	issuerAddress            = "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak"
	holderAddress            = "shareledger1j4ndn4qed0ukulc0a6a5fxe6yxmnm7wh59pwpn"
	data                     = "documentData"
	proof                    = "proof"
	defaultMsgCreateDocument = &types.MsgCreateDocument{
		Data:   data,
		Holder: holderAddress,
		Issuer: issuerAddress,
		Proof:  proof,
	}
)

type KeeperTestSuite struct {
	suite.Suite

	ctx            sdk.Context
	queryClient    types.QueryClient
	msgServer      types.MsgServer
	documentKeeper *keeper.Keeper
	idKeeper       *documenttestutil.MockIDKeeper

	encCfg moduletestutil.TestEncodingConfig
}

func TestKeeperTestSuite(t *testing.T) {
	params.SetAddressPrefixes()
	t.Log("Running TestKeeperTestSuite for document module...")
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig()

	ctrl := gomock.NewController(s.T())
	idKeeper := documenttestutil.NewMockIDKeeper(ctrl)
	documentKeeper := keeper.NewKeeper(encCfg.Codec, key, idKeeper)

	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, s.encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, documentKeeper)

	s.idKeeper = idKeeper
	s.ctx = ctx
	s.encCfg = encCfg
	s.documentKeeper = documentKeeper
	s.msgServer = keeper.NewMsgServerImpl(*s.documentKeeper)
	s.queryClient = types.NewQueryClient(queryHelper)
}

func (s *KeeperTestSuite) TestCreateDocument() {
	// valid
	s.createDefaultDocument()
	// existed
	s.idKeeper.EXPECT().GetFullIDByIDString(s.ctx, holderAddress).Return(nil, true)
	_, err := s.msgServer.CreateDocument(s.ctx, defaultMsgCreateDocument)
	s.Error(err)
}

func (s *KeeperTestSuite) TestCreateDocuments() {
	req := &types.MsgCreateDocuments{
		Data:   []string{data},
		Holder: []string{holderAddress},
		Issuer: issuerAddress,
		Proof:  []string{proof},
	}
	// not found id for holder
	s.idKeeper.EXPECT().GetFullIDByIDString(s.ctx, holderAddress).Return(nil, false)
	_, err := s.msgServer.CreateDocuments(s.ctx, req)
	s.ErrorContains(err, types.ErrHolderIDNotExisted.Error())
	// valid
	s.idKeeper.EXPECT().GetFullIDByIDString(s.ctx, holderAddress).Return(nil, true)
	_, err = s.msgServer.CreateDocuments(s.ctx, req)
	s.NoError(err)
	// existed
	s.idKeeper.EXPECT().GetFullIDByIDString(s.ctx, holderAddress).Return(nil, true)
	_, err = s.msgServer.CreateDocuments(s.ctx, req)
	s.Error(err)
}

func (s *KeeperTestSuite) TestRevokeDocument() {
	req := &types.MsgRevokeDocument{
		Holder: holderAddress,
		Issuer: issuerAddress,
		Proof:  proof,
	}
	// not found
	_, err := s.msgServer.RevokeDocument(s.ctx, req)
	s.ErrorContains(err, types.ErrDocNotExisted.Error())
	// valid
	s.createDefaultDocument()
	_, err = s.msgServer.RevokeDocument(s.ctx, req)
}

func (s *KeeperTestSuite) TestUpdateDocument() {
	req := types.NewMsgUpdateDocument("New Document Data", holderAddress, issuerAddress, proof)
	// not found
	_, err := s.msgServer.UpdateDocument(s.ctx, req)
	s.ErrorContains(err, types.ErrDocNotExisted.Error())
	// valid
	s.createDefaultDocument()
	_, err = s.msgServer.UpdateDocument(s.ctx, req)
	s.NoError(err)
}

func (s *KeeperTestSuite) TestDocumentByHolderId() {
	req := &types.QueryDocumentByHolderIdRequest{
		Id: holderAddress,
	}
	// invalid
	res, err := s.queryClient.DocumentByHolderId(s.ctx, nil)
	s.Nil(res)
	s.ErrorContains(err, codes.InvalidArgument.String())
	// empty resp
	res, err = s.queryClient.DocumentByHolderId(s.ctx, req)
	s.NoError(err)
	s.Require().Empty(res.Documents)
	// valid
	s.createDefaultDocument()
	res, err = s.queryClient.DocumentByHolderId(s.ctx, req)
	s.NoError(err)
	s.Equal(data, res.Documents[0].Data)
}

func (s *KeeperTestSuite) TestDocumentByProof() {
	req := &types.QueryDocumentByProofRequest{
		Proof: proof,
	}
	// invalid
	res, err := s.queryClient.DocumentByProof(s.ctx, nil)
	s.Nil(res)
	s.ErrorContains(err, codes.InvalidArgument.String())
	// not found
	res, err = s.queryClient.DocumentByProof(s.ctx, req)
	s.Nil(res)
	s.ErrorContains(err, codes.NotFound.String())
	// valid
	s.createDefaultDocument()
	res, err = s.queryClient.DocumentByProof(s.ctx, req)
	s.NoError(err)
	s.Equal(proof, res.Document.Proof)
}

func (s *KeeperTestSuite) TestDocumentOfHolderByIssuer() {
	req := &types.QueryDocumentOfHolderByIssuerRequest{
		Holder: holderAddress,
		Issuer: issuerAddress,
	}
	// invalid
	res, err := s.queryClient.DocumentOfHolderByIssuer(s.ctx, nil)
	s.Nil(res)
	s.ErrorContains(err, codes.InvalidArgument.String())
	// valid
	s.createDefaultDocument()
	res, err = s.queryClient.DocumentOfHolderByIssuer(s.ctx, req)
	s.NoError(err)
	s.Equal(issuerAddress, res.Documents[0].Issuer)
	s.Equal(holderAddress, res.Documents[0].Holder)
}

func (s *KeeperTestSuite) createDefaultDocument() {
	s.idKeeper.EXPECT().GetFullIDByIDString(s.ctx, holderAddress).Return(nil, true)
	_, err := s.msgServer.CreateDocument(s.ctx, defaultMsgCreateDocument)
	s.NoError(err)
}
