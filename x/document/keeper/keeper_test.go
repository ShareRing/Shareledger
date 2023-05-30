package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/x/document/keeper"
	documenttestutil "github.com/sharering/shareledger/x/document/testutil"
	"github.com/sharering/shareledger/x/document/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	queryClient types.QueryClient
	msgServer   types.MsgServer
	encCfg      moduletestutil.TestEncodingConfig

	documentKeeper keeper.Keeper
	idKeeper       *documenttestutil.MockIDKeeper
}

func TestKeeperTestSuite(t *testing.T) {
	params.SetAddressPrefixes()
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	idKeeper := documenttestutil.NewMockIDKeeper(ctrl)

	s.idKeeper = idKeeper
}
