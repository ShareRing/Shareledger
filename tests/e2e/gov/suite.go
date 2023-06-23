//go:build e2e

package gov

import (
	"fmt"
	"path/filepath"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/sharering/shareledger/testutil/network"
	dtxtypes "github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/suite"
)

var (
	globalProposalId = 0
	votingPeriod     = time.Second * 60
	reward1          dtxtypes.Reward
	builderCount1    = dtxtypes.BuilderCount{
		Index: "ContractAddress1",
		Count: 1,
	}
	builderList1 = dtxtypes.BuilderList{
		Id:              1,
		ContractAddress: "ContractAddress1",
	}
	builderList2 = dtxtypes.BuilderList{
		Id:              2,
		ContractAddress: "ContractAddress2",
	}
	params         = dtxtypes.DefaultParams()
	devPoolAccount string
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// type cliFunction func() *cobra.Command

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	cfg.NumValidators = 1
	return &E2ETestSuite{cfg: cfg}
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for gov module")
	// the nodeDir, and moniker hard code at here in cosmos-sdk:
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:398
	// So just reuse it
	rootDir := s.T().TempDir()
	moniker := fmt.Sprintf("node%d", s.cfg.NumValidators-1)
	// TestingGenesis should use the same KeyringDir as validator KeyringDir
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:400
	nodeDir := filepath.Join(rootDir, moniker, "simcli")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg, nodeDir, moniker)
	s.Require().NotNil(kr)
	govGenesis := v1.DefaultGenesisState()
	govGenesis.Params.VotingPeriod = &votingPeriod
	govGenesis.Params.MinDeposit = sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(1000)))
	govGenesisBz, err := s.cfg.Codec.MarshalJSON(govGenesis)
	s.Require().NoError(err)
	s.cfg.GenesisState[types.ModuleName] = govGenesisBz
	// add 1000shr coin for validator account
	s.cfg.StakingTokens = sdk.NewInt(1000000000000)

	// init genesis state for distributionx module
	devPoolAccount = network.MustAddressFormKeyring(kr, network.KeyAccount3).String()
	params.BuilderWindows = 15
	params.TxThreshold = 3
	params.DevPoolAccount = devPoolAccount
	reward1 = dtxtypes.Reward{
		Index:  network.MustAddressFormKeyring(kr, network.KeyAccount1).String(),
		Amount: sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(1000))),
	}
	distriXGenesis := &dtxtypes.GenesisState{
		Params: params,
		RewardList: []dtxtypes.Reward{
			reward1,
		},
		BuilderCountList: []dtxtypes.BuilderCount{
			builderCount1,
		},
		BuilderListList: []dtxtypes.BuilderList{
			builderList1, builderList2,
		},
		BuilderListCount: 2,
	}
	s.cfg.GenesisState[dtxtypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(distriXGenesis)

	s.network = network.New(s.T(), rootDir, s.cfg)

	s.network.Validators[0].ClientCtx.Keyring = kr
	s.Require().NoError(s.network.WaitForNextBlock())
}
