package network

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/sharering/shareledger/app"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/tendermint/spm/cosmoscmd"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"time"
)

func ShareLedgerTestingConfig() network.Config {
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	SettingAccountPrefix()
	config := network.DefaultConfig()

	config.Codec = encCfg.Marshaler
	config.TxConfig = encCfg.TxConfig
	config.LegacyAmino = encCfg.Amino
	config.InterfaceRegistry = encCfg.InterfaceRegistry
	config.AccountRetriever = authtypes.AccountRetriever{}
	config.AppConstructor = NewAppConstructor()
	config.GenesisState = app.ModuleBasics.DefaultGenesis(encCfg.Marshaler)
	config.TimeoutCommit = 2 * time.Second
	config.ChainID = "chain-" + tmrand.NewRand().Str(6)
	config.NumValidators = 2
	config.BondDenom = sdk.DefaultBondDenom
	config.MinGasPrices = fmt.Sprintf("0.000006%s", denom.Base)
	config.AccountTokens = sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction)
	config.StakingTokens = sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction)
	config.BondedTokens = sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)
	config.PruningStrategy = storetypes.PruningOptionNothing
	config.CleanupDir = true
	config.SigningAlgo = string(hd.Secp256k1Type)
	config.KeyringOptions = []keyring.Option{}

	return config
}

func SettingAccountPrefix() {
	config := sdk.GetConfig()

	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	config.Seal()

}
