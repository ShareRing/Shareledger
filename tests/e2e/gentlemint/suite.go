//go:build e2e

package gentlemint

import (
	"fmt"
	"path/filepath"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/testutil/network"
	gmtypes "github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/suite"
)

var (
	defaultLevelFees = []gmtypes.LevelFee{
		{
			Level:   "low",
			Fee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0.02")),
			Creator: "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr",
		},
		{
			Level:   "medium",
			Fee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0.03")),
			Creator: "shareledger18pf3zdwqjntd9wkvfcjvmdc7hua6c0q2eck5h5",
		},
		{
			Level:   "high",
			Fee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0.05")),
			Creator: "shareledger1qn7whj8v3gjf3a2nncydu88kt325xtd3t3gc95",
		},
	}

	levelMin = gmtypes.LevelFee{
		Level:   "min",
		Fee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0.01")),
		Creator: "shareledger1qn7whj8v3gjf3a2nncydu88kt325xtd3t3gc95",
	}

	levelZero = gmtypes.LevelFee{
		Level:   "zero",
		Fee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0")),
		Creator: "shareledger1qn7whj8v3gjf3a2nncydu88kt325xtd3t3gc95",
	}

	defaultActionLevelFees = []gmtypes.ActionLevelFee{
		{
			Action: "gentlemint_load",
			Level:  "low",
		},
		{
			Action: "swap_approve-out",
			Level:  "medium",
		},
		{
			Action: "swap_deposit",
			Level:  "high",
		},
		{
			// this item used for delete, ignore check it when query
			Action: "new_for_delete",
			Level:  "low",
		},
	}

	defaultRate           = "400"
	rate, _               = strconv.Atoi(defaultRate)
	exchangeRate          = gmtypes.ExchangeRate{Rate: defaultRate}
	params                = gmtypes.Params{MinimumGasPrices: sdk.NewDecCoins(sdk.NewDecCoin("nshr", sdk.NewInt(1000)))}
	feeLevelLow           = defaultLevelFees[0].Fee.String()
	lowConvertedFee, _    = denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(defaultLevelFees[0].Fee), sdk.NewDec(int64(rate)), true)
	mediumConvertedFee, _ = denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(defaultLevelFees[1].Fee), sdk.NewDec(int64(rate)), true)
	highConvertedFee, _   = denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(defaultLevelFees[2].Fee), sdk.NewDec(int64(rate)), true)
	minConvertedFee, _    = denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(levelMin.Fee), sdk.NewDec(int64(rate)), true)
	zeroConvertedFee, _   = denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(levelZero.Fee), sdk.NewDec(int64(rate)), true)
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
	gmGenesis := gmtypes.GenesisState{
		ExchangeRate:       &exchangeRate,
		LevelFeeList:       defaultLevelFees,
		ActionLevelFeeList: defaultActionLevelFees,
		Params:             params,
	}

	gmGenesisBz, err := s.cfg.Codec.MarshalJSON(&gmGenesis)
	s.Require().NoError(err)
	s.cfg.GenesisState[gmtypes.ModuleName] = gmGenesisBz

	s.network = network.New(s.T(), rootDir, s.cfg)

	s.network.Validators[0].ClientCtx.Keyring = kr
	s.Require().NoError(s.network.WaitForNextBlock())
}
