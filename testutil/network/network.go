package network

import (
	"bufio"
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simapp2 "github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/testutil/simapp"
	electoraltypes "github.com/sharering/shareledger/x/electoral/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"
	dbm "github.com/tendermint/tm-db"
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tmrand "github.com/tendermint/tendermint/libs/rand"
)

const (
	PrefixValidator = "val"

	PrefixConsensus = "cons"

	PrefixPublic = "pub"

	PrefixOperator = "oper"

	Bech32MainPrefix = "shareledger"

	Bech32PrefixAccAddr = Bech32MainPrefix

	Bech32PrefixAccPub = Bech32MainPrefix + PrefixPublic

	Bech32PrefixValAddr = Bech32MainPrefix + PrefixValidator + PrefixOperator

	Bech32PrefixValPub = Bech32MainPrefix + PrefixValidator + PrefixOperator + PrefixPublic

	Bech32PrefixConsAddr = Bech32MainPrefix + PrefixValidator + PrefixConsensus

	Bech32PrefixConsPub = Bech32MainPrefix + PrefixValidator + PrefixConsensus + PrefixPublic
)

// package-wide network lock to only allow one test network at a time
var lock = new(sync.Mutex)

// AppConstructor defines a function which accepts a network configuration and
// creates an ABCI Application to provide to Tendermint.

var (
	Accounts = map[string]sdk.Address{}
)

type (
	// Network defines a local in-process testing network using SimApp. It can be
	// configured to start any number of validators, each with its own RPC and API
	// clients. Typically, this test network would be used in client and integration
	// testing where user input is expected.
	//
	// Note, due to Tendermint constraints in regards to RPC functionality, there
	// may only be one test network running at a time. Thus, any caller must be
	// sure to Cleanup after testing is finished in order to allow other tests
	// to create networks. In addition, only the first validator will have a valid
	// RPC and API server/client.
	Network struct {
		T          *testing.T
		BaseDir    string
		Validators []*network.Validator

		Config network.Config

		Accounts map[string]sdk.Address
	}

	AppConstructor = func(val network.Validator) servertypes.Application
)

// GetTestingGenesis init the genesis state for testing in here
func GetTestingGenesis(t *testing.T, config *network.Config) (keyring.Keyring, string) {
	genesisState := config.GenesisState
	var bankGenesis types.GenesisState
	var authGenesis authtypes.GenesisState
	var genAccounts []authtypes.GenesisAccount
	var genElectoral electoraltypes.GenesisState
	err := config.Codec.UnmarshalJSON(genesisState[types.ModuleName], &bankGenesis)
	if err != nil {
		t.Errorf("fail to init test")
	}

	err = config.Codec.UnmarshalJSON(genesisState[authtypes.ModuleName], &authGenesis)
	if err != nil {
		t.Errorf("fail to init test")
	}
	buf := bufio.NewReader(os.Stdin)

	baseDir, err := ioutil.TempDir(t.TempDir(), config.ChainID)

	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, baseDir, buf, config.KeyringOptions...)
	var genBalances []banktypes.Balance

	info, _, err := kb.NewMnemonic(KeyAuthority, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyAuthority] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   defaultCoins,
	})

	info, _, err = kb.NewMnemonic(KeyTreasurer, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyTreasurer] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   defaultCoins,
	})

	info, _, err = kb.NewMnemonic(KeyOperator, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyOperator] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   defaultCoins,
	})

	info, _, err = kb.NewMnemonic(KeyIDSigner, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyIDSigner] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   defaultCoins,
	})

	info, _, err = kb.NewMnemonic(KeyDocIssuer, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyDocIssuer] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   defaultCoins,
	})

	info, _, err = kb.NewMnemonic(KeyMillionaire, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyMillionaire] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   becauseImRich,
	})

	info, _, err = kb.NewMnemonic(KeyLoader, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyLoader] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   becauseImRich,
	})

	info, _, err = kb.NewMnemonic(KeyEmpty1, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyEmpty1] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   poorMen,
	})

	info, _, err = kb.NewMnemonic(KeyEmpty2, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyEmpty2] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   poorMen,
	})
	info, _, err = kb.NewMnemonic(KeyEmpty3, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyEmpty3] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   poorMen,
	})
	info, _, err = kb.NewMnemonic(KeyEmpty4, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyEmpty4] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   poorMen,
	})

	info, _, err = kb.NewMnemonic(KeyEmpty5, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyEmpty5] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   poorMen,
	})

	info, _, err = kb.NewMnemonic(KeyAccount1, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyAccount1] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   defaultCoins,
	})

	info, _, err = kb.NewMnemonic(KeyAccount2, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyAccount2] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   defaultCoins,
	})
	info, _, err = kb.NewMnemonic(KeyAccount3, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyAccount3] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   defaultCoins,
	})

	info, _, err = kb.NewMnemonic(KeyAccount4, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(t, err, "init fail")
	Accounts[KeyAccount4] = info.GetAddress()
	genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	genBalances = append(genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   defaultCoins,
	})

	for i, _ := range genAccounts {
		err = genAccounts[i].SetAccountNumber(uint64(i + 1))
		require.NoError(t, err, "init fail")
	}

	genElectoral = electoraltypes.GenesisState{
		Authority: &electoraltypes.Authority{
			Address: Accounts[KeyAuthority].String(),
		},
		Treasurer: &electoraltypes.Treasurer{
			Address: Accounts[KeyTreasurer].String(),
		},
	}

	accounts, err := authtypes.PackAccounts(genAccounts)
	if err != nil {
		t.Errorf("int fails")
	}
	authGenesis.Accounts = append(authGenesis.Accounts, accounts...)

	bankGenesis.Balances = genBalances
	bankGenesisBz, err := config.Codec.MarshalJSON(&bankGenesis)
	if err != nil {
		t.Errorf("init test fail %v", err)
	}
	authGenesisBz, err := config.Codec.MarshalJSON(&authGenesis)
	if err != nil {
		t.Errorf("init test fail %v", err)
	}

	genElectoralBz := config.Codec.MustMarshalJSON(&genElectoral)
	if err != nil {
		t.Errorf("init test fail %v", err)
	}
	genesisState[types.ModuleName] = bankGenesisBz
	genesisState[authtypes.ModuleName] = authGenesisBz
	genesisState[electoraltypes.ModuleName] = genElectoralBz
	config.GenesisState = genesisState
	return kb, baseDir
}

// NewAppConstructor returns a new simapp AppConstructor
func NewAppConstructor(encodingCfg cosmoscmd.EncodingConfig) AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return app.New(
			val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			encodingCfg,
			simapp2.EmptyAppOptions{},
			baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices))
	}
}

//DefaultConfig returns a sane default configuration suitable for nearly all
//testing requirements.
func DefaultConfig() network.Config {
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	return network.Config{
		Codec:             encCfg.Marshaler,
		TxConfig:          encCfg.TxConfig,
		LegacyAmino:       encCfg.Amino,
		InterfaceRegistry: encCfg.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor:    NewAppConstructor(encCfg),
		GenesisState:      simapp.ModuleBasics.DefaultGenesis(encCfg.Marshaler),
		TimeoutCommit:     2 * time.Second,
		ChainID:           "chain-" + tmrand.NewRand().Str(6),
		NumValidators:     4,
		BondDenom:         sdk.DefaultBondDenom,
		MinGasPrices:      fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		AccountTokens:     sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:     sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:      sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy:   storetypes.PruningOptionNothing,
		CleanupDir:        true,
		SigningAlgo:       string(hd.Secp256k1Type),
		KeyringOptions:    []keyring.Option{},
	}
}

func SettingAccountPrefix() {
	config := sdk.GetConfig()

	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	config.Seal()

}

func ShareLedgerTestingConfig() network.Config {
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	SettingAccountPrefix()
	return network.Config{
		Codec:             encCfg.Marshaler,
		TxConfig:          encCfg.TxConfig,
		LegacyAmino:       encCfg.Amino,
		InterfaceRegistry: encCfg.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor:    NewAppConstructor(encCfg),
		GenesisState:      app.ModuleBasics.DefaultGenesis(encCfg.Marshaler),
		TimeoutCommit:     2 * time.Second,
		ChainID:           "chain-" + tmrand.NewRand().Str(6),
		NumValidators:     2,
		BondDenom:         sdk.DefaultBondDenom,
		MinGasPrices:      fmt.Sprintf("0.000006%s", denom.Base),
		AccountTokens:     sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:     sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:      sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy:   storetypes.PruningOptionNothing,
		CleanupDir:        true,
		SigningAlgo:       string(hd.Secp256k1Type),
		KeyringOptions:    []keyring.Option{},
	}
}
