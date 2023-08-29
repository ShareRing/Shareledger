package distributionx

import (
	"fmt"
	"path/filepath"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
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

func (s *E2ETestSuite) NewE2ENetWork(nw *network.Network) {
	s.network = nw
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
	// the nodeDir, and moniker hard code at here in cosmos-sdk:
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:398
	// So just reuse it here
	rootDir := s.T().TempDir()
	moniker := fmt.Sprintf("node%d", s.cfg.NumValidators-1)
	// TestingGenesis should use the same KeyringDir as validator KeyringDir
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:400
	nodeDir := filepath.Join(rootDir, moniker, "simcli")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg, nodeDir, moniker)
	s.Require().NotNil(kr)
	devPoolAccount = network.MustAddressFormKeyring(kr, network.KeyAccount3).String()

	reward := sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(1000)))
	params.BuilderWindows = 25
	params.TxThreshold = 3
	params.DevPoolAccount = devPoolAccount
	reward1 = types.Reward{
		Index:  network.MustAddressFormKeyring(kr, network.KeyAccount1).String(),
		Amount: reward,
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

	// set distributionx module balance
	authtypes.NewModuleAddress(types.ModuleName)

	var bankGenesis banktypes.GenesisState
	s.Require().NoError(s.cfg.Codec.UnmarshalJSON(s.cfg.GenesisState[banktypes.ModuleName], &bankGenesis))

	bankGenesis.Balances = append(bankGenesis.Balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(types.ModuleName).String(),
		Coins:   reward,
	})
	bankGenesis.Supply.Add(reward...)

	s.cfg.GenesisState[types.ModuleName] = s.cfg.Codec.MustMarshalJSON(distriXGenesis)
	s.cfg.GenesisState[banktypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&bankGenesis)

	s.network = network.New(s.T(), rootDir, s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
}