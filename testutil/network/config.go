package network

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/app/params"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func DefaultConfig() network.Config {
	encodingConfig := app.MakeTestEncodingConfig()
	config := network.DefaultConfig()

	config.Codec = encodingConfig.Codec
	config.TxConfig = encodingConfig.TxConfig
	config.LegacyAmino = encodingConfig.Amino
	config.InterfaceRegistry = encodingConfig.InterfaceRegistry
	config.AppConstructor = ShareLedgerChainConstructor()
	genesisState := app.ModuleBasics.DefaultGenesis(encodingConfig.Codec)
	config.NumValidators = 2
	config.MinGasPrices = fmt.Sprintf("0.000006%s", denom.Base)
	config.BondDenom = denom.Base

	// change staking denom `stake` -> `nshr`
	stakingState := stakingtypes.GetGenesisStateFromAppState(config.Codec, genesisState)
	stakingState.Params.BondDenom = denom.Base
	genesisState[stakingtypes.ModuleName] = config.Codec.MustMarshalJSON(stakingState)
	config.GenesisState = genesisState

	params.SetAddressPrefixes()

	return config
}
