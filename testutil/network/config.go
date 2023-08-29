package network

import (
	"fmt"

	dbm "github.com/cometbft/cometbft-db"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	pruningtypes "github.com/cosmos/cosmos-sdk/store/pruning/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/testutil/sims"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/app/params"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func DefaultConfig() network.Config {
	encodingConfig := app.MakeTestEncodingConfig()
	config := network.DefaultConfig(func() network.TestFixture { return network.TestFixture{} })

	config.Codec = encodingConfig.Codec
	config.TxConfig = encodingConfig.TxConfig
	config.LegacyAmino = encodingConfig.Amino
	config.InterfaceRegistry = encodingConfig.InterfaceRegistry
	// config.AppConstructor = ShareLedgerChainConstructor()
	genesisState := app.ModuleBasics.DefaultGenesis(encodingConfig.Codec)
	config.NumValidators = 2
	config.MinGasPrices = fmt.Sprintf("1%s", denom.Base)
	config.BondDenom = denom.Base

	// change staking denom `stake` -> `nshr`
	stakingState := stakingtypes.GetGenesisStateFromAppState(config.Codec, genesisState)
	stakingState.Params.BondDenom = denom.Base
	genesisState[stakingtypes.ModuleName] = config.Codec.MustMarshalJSON(stakingState)
	config.GenesisState = genesisState
	var govGenesisState v1.GenesisState

	if genesisState[govtypes.ModuleName] != nil {
		encodingConfig.Codec.MustUnmarshalJSON(genesisState[govtypes.ModuleName], &govGenesisState)
	}

	config.AppConstructor = func(val network.ValidatorI) servertypes.Application {
		return app.New(val.GetCtx().Logger, dbm.NewMemDB(), nil, true,
			sims.AppOptionsMap{
				app.FlagAppOptionSkipCheckVoter: true,
				flags.FlagHome:                  app.DefaultNodeHome,
			},
			bam.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
			bam.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
			bam.SetChainID(val.GetCtx().Viper.GetString(flags.FlagChainID)),
		)
	}

	params.SetAddressPrefixes()
	return config
}

type EmptyAppOptions struct{}

func (EmptyAppOptions) Get(_ string) interface{} { return nil }
