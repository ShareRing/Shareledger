package distributionx

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	cfg.NumValidators = 1
	return &E2ETestSuite{cfg: cfg}
}

var (
	reward1       types.Reward
	builderCount1 = types.BuilderCount{
		Index: "ContractAddress1",
		Count: 1,
	}
	builderList1 = types.BuilderList{
		Id:              1,
		ContractAddress: "ContractAddress1",
	}
	builderList2 = types.BuilderList{
		Id:              2,
		ContractAddress: "ContractAddress2",
	}
	params         = types.DefaultParams()
	devPoolAccount string
)

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("settings up e2e test suite for distributionx module")

	// set devPoolAccount == KeyAccount3
	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg)
	devPoolAccount = network.MustAddressFormKeyring(kr, network.KeyAccount3).String()

	params.BuilderWindows = 15
	params.TxThreshold = 3
	params.DevPoolAccount = devPoolAccount
	reward1 = types.Reward{
		Index:  network.MustAddressFormKeyring(kr, network.KeyAccount1).String(),
		Amount: sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(1000))),
	}
	distriXGenesis := &types.GenesisState{
		Params: params,
		RewardList: []types.Reward{
			reward1,
		},
		BuilderCountList: []types.BuilderCount{
			builderCount1,
		},
		BuilderListList: []types.BuilderList{
			builderList1, builderList2,
		},
		BuilderListCount: 2,
	}
	s.cfg.GenesisState[types.ModuleName] = s.cfg.Codec.MustMarshalJSON(distriXGenesis)

	s.network = network.New(s.T(), s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
}
