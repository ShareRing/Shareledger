package network

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/sharering/shareledger/app"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/tendermint/spm/cosmoscmd"
)

func ShareLedgerTestingConfig() network.Config {
	shareRingEncCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	cosmoscmd.SetPrefixes(Bech32MainPrefix)

	config := network.DefaultConfig()

	config.Codec = shareRingEncCfg.Marshaler
	config.TxConfig = shareRingEncCfg.TxConfig
	config.LegacyAmino = shareRingEncCfg.Amino
	config.InterfaceRegistry = shareRingEncCfg.InterfaceRegistry
	config.AppConstructor = ShareLedgerChainConstructor()
	config.GenesisState = app.ModuleBasics.DefaultGenesis(shareRingEncCfg.Marshaler)
	config.NumValidators = 2
	config.MinGasPrices = fmt.Sprintf("0.000006%s", denom.Base)

	return config
}
