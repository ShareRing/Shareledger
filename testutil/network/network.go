package network

import (
	"bufio"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/testutil/simapp"
	electoraltypes "github.com/sharering/shareledger/x/electoral/types"
	"io/ioutil"
	"os"
	"testing"
)

const (
	Bech32MainPrefix = "shareledger"
)

// AppConstructor defines a function which accepts a network configuration and
// creates an ABCI Application to provide to Tendermint.

var (
	Accounts = map[string]sdk.Address{}
)

type (
	AppConstructor = func(val network.Validator) servertypes.Application

	AccountInfo struct {
		Key     string
		Balance sdk.Coins
	}
)

func CompileGenesis(t *testing.T, config *network.Config, genesisState map[string]json.RawMessage, au []authtypes.GenesisAccount, b []banktypes.Balance, elGen electoraltypes.GenesisState) map[string]json.RawMessage {
	var bankGenesis types.GenesisState
	var authGenesis authtypes.GenesisState

	err := config.Codec.UnmarshalJSON(genesisState[types.ModuleName], &bankGenesis)
	if err != nil {
		t.Errorf("fail to init test")
	}

	err = config.Codec.UnmarshalJSON(genesisState[authtypes.ModuleName], &authGenesis)
	if err != nil {
		t.Errorf("fail to init test")
	}

	accounts, err := authtypes.PackAccounts(au)
	if err != nil {
		t.Errorf("int fails")
	}

	authGenesis.Accounts = append(authGenesis.Accounts, accounts...)

	bankGenesis.Balances = b
	bankGenesisBz, err := config.Codec.MarshalJSON(&bankGenesis)
	if err != nil {
		t.Errorf("init test fail %v", err)
	}
	authGenesisBz, err := config.Codec.MarshalJSON(&authGenesis)
	if err != nil {
		t.Errorf("init test fail %v", err)
	}

	genElectoralBz := config.Codec.MustMarshalJSON(&elGen)
	if err != nil {
		t.Errorf("init test fail %v", err)
	}
	genesisState[types.ModuleName] = bankGenesisBz
	genesisState[authtypes.ModuleName] = authGenesisBz
	genesisState[electoraltypes.ModuleName] = genElectoralBz
	return genesisState
}

// GetTestingGenesis init the genesis state for testing in here
func GetTestingGenesis(t *testing.T, config *network.Config) (keyring.Keyring, string) {
	genesisState := config.GenesisState

	buf := bufio.NewReader(os.Stdin)
	baseDir, err := ioutil.TempDir(t.TempDir(), config.ChainID)
	if err != nil {
		t.Errorf("fail to create temp dir %v", err)
	}
	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, baseDir, buf, config.KeyringOptions...)
	accountBuilder := NewKeyringBuilder(t, kb)

	users := []AccountInfo{
		{Key: KeyAuthority, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyTreasurer, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyOperator, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyIDSigner, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyDocIssuer, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyMillionaire, Balance: OneMillionSHRSHRPCoins},
		{Key: KeyLoader, Balance: OneMillionSHRSHRPCoins},

		{Key: KeyEmpty1, Balance: ZeroSHRSHRP},
		{Key: KeyEmpty2, Balance: ZeroSHRSHRP},
		{Key: KeyEmpty3, Balance: ZeroSHRSHRP},
		{Key: KeyEmpty4, Balance: ZeroSHRSHRP},
		{Key: KeyEmpty5, Balance: ZeroSHRSHRP},

		{Key: KeyAccount1, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyAccount2, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyAccount3, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyAccount4, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyAccount5, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyAccount6, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyAccount7, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyAccount8, Balance: OneThousandSHROneHundredSHRPCoins},
	}

	for _, u := range users {
		accountBuilder.InitUser(u.Key, u.Balance)
	}

	newKeyringService, genAccounts, genBalances := accountBuilder.BuildGenesis()

	var genElectoral electoraltypes.GenesisState
	genElectoral = electoraltypes.GenesisState{
		Authority: &electoraltypes.Authority{
			Address: Accounts[KeyAuthority].String(),
		},
		Treasurer: &electoraltypes.Treasurer{
			Address: Accounts[KeyTreasurer].String(),
		},
	}

	genesisState = CompileGenesis(t, config, genesisState, genAccounts, genBalances, genElectoral)
	config.GenesisState = genesisState
	return newKeyringService, baseDir
}

// ShareLedgerChainConstructor returns a new shareLedger AppConstructor
func ShareLedgerChainConstructor() AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return simapp.New(val.Ctx.Config.RootDir, TestAppOptions{})
	}
}

type TestAppOptions struct{}

// Get implements TestAppOptions
func (ao TestAppOptions) Get(o string) interface{} {
	if o == app.FlagAppOptionSkipCheckVoter {
		return true
	}
	return nil
}
