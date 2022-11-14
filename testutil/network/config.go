package network

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/sharering/shareledger/app"
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
	config.GenesisState = app.ModuleBasics.DefaultGenesis(encodingConfig.Codec)
	config.NumValidators = 2
	config.MinGasPrices = fmt.Sprintf("0.000006%s", denom.Base)

	return config
}
