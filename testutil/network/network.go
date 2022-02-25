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
	"github.com/sharering/shareledger/testutil/simapp"
	electoraltypes "github.com/sharering/shareledger/x/electoral/types"
	"io/ioutil"
	"os"
	"sync"
	"testing"
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
	}

	AppConstructor = func(val network.Validator) servertypes.Application
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

	accountBuilder.InitUser(KeyAuthority, defaultCoins)
	accountBuilder.InitUser(KeyTreasurer, defaultCoins)
	accountBuilder.InitUser(KeyOperator, defaultCoins)
	accountBuilder.InitUser(KeyIDSigner, defaultCoins)
	accountBuilder.InitUser(KeyDocIssuer, defaultCoins)
	accountBuilder.InitUser(KeyMillionaire, becauseImRich)
	accountBuilder.InitUser(KeyLoader, becauseImRich)
	accountBuilder.InitUser(KeyEmpty1, poorMen)
	accountBuilder.InitUser(KeyEmpty2, poorMen)
	accountBuilder.InitUser(KeyEmpty3, poorMen)
	accountBuilder.InitUser(KeyEmpty4, poorMen)
	accountBuilder.InitUser(KeyEmpty5, poorMen)
	accountBuilder.InitUser(KeyAccount1, defaultCoins)
	accountBuilder.InitUser(KeyAccount2, defaultCoins)
	accountBuilder.InitUser(KeyAccount3, defaultCoins)
	accountBuilder.InitUser(KeyAccount4, defaultCoins)

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

// NewAppConstructor returns a new shareLedger AppConstructor
func NewAppConstructor() AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return simapp.New(val.Ctx.Config.RootDir)
	}
}
