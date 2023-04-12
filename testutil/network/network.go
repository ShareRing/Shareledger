package network

import (
	"bufio"
	"encoding/json"
	distributionxtypes "github.com/sharering/shareledger/x/distributionx/types"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/testutil/simapp"
	electoraltypes "github.com/sharering/shareledger/x/electoral/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/require"
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

type (
	Network = network.Network
	Config  = network.Config
)

// New creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func New(t *testing.T, configs ...network.Config) *network.Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg network.Config
	if len(configs) == 0 {
		cfg = DefaultConfig()
	} else {
		cfg = configs[0]
	}
	net, err := network.New(t, t.TempDir(), cfg)
	require.NoError(t, err)
	t.Cleanup(net.Cleanup)
	return net
}

func CompileGenesis(t *testing.T, config *network.Config, genesisState map[string]json.RawMessage, au []authtypes.GenesisAccount, b []banktypes.Balance, elGen electoraltypes.GenesisState, distributionXGen distributionxtypes.GenesisState) map[string]json.RawMessage {
	var bankGenesis types.GenesisState
	var authGenesis authtypes.GenesisState
	var stakingGenesis stakingtypes.GenesisState

	config.Codec.MustUnmarshalJSON(genesisState[types.ModuleName], &bankGenesis)
	config.Codec.MustUnmarshalJSON(genesisState[authtypes.ModuleName], &authGenesis)
	config.Codec.MustUnmarshalJSON(genesisState[stakingtypes.ModuleName], &stakingGenesis)

	accounts, err := authtypes.PackAccounts(au)
	if err != nil {
		t.Errorf("int fails")
	}

	authGenesis.Accounts = append(authGenesis.Accounts, accounts...)
	bankGenesis.Balances = b
	stakingGenesis.Params.BondDenom = denom.Base

	genesisState[types.ModuleName] = config.Codec.MustMarshalJSON(&bankGenesis)
	genesisState[authtypes.ModuleName] = config.Codec.MustMarshalJSON(&authGenesis)
	genesisState[electoraltypes.ModuleName] = config.Codec.MustMarshalJSON(&elGen)
	genesisState[stakingtypes.ModuleName] = config.Codec.MustMarshalJSON(&stakingGenesis)
	genesisState[distributionxtypes.ModuleName] = config.Codec.MustMarshalJSON(&distributionXGen)

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
	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, baseDir, buf, config.Codec, config.KeyringOptions...)
	if err != nil {
		t.Error("fail to create keyring")
	}
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

		{Key: KeyMasterBuilder1, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyMasterBuilder2, Balance: OneThousandSHROneHundredSHRPCoins},
		{Key: KeyDevPoolAccount, Balance: ZeroSHRSHRP},

		{Key: KeyApproverRelayer, Balance: OneThousandSHROneHundredSHRPCoins},

		{Key: KeyApproverRelayer, Balance: OneThousandSHROneHundredSHRPCoins},
	}

	for _, u := range users {
		accountBuilder.InitUser(u.Key, u.Balance)
	}

	accountBuilder.GenBSCSigner(KeyAccountSwapBSC, OneThousandSHROneHundredSHRPCoins)
	accountBuilder.GenETHSigner(KeyAccountSwapETH, OneThousandSHROneHundredSHRPCoins)
	accountBuilder.NewAccountToSign()

	newKeyringService, genAccounts, genBalances := accountBuilder.BuildGenesis()

	genElectoral := electoraltypes.GenesisState{
		Authority: &electoraltypes.Authority{
			Address: Accounts[KeyAuthority].String(),
		},
		Treasurer: &electoraltypes.Treasurer{
			Address: Accounts[KeyTreasurer].String(),
		},
	}

	disXGen := distributionxtypes.DefaultGenesis()

	devPoolAddr := MustAddressFormKeyring(newKeyringService, KeyDevPoolAccount)

	disXGen.Params.DevPoolAccount = devPoolAddr.String()

	disXGen.BuilderListList = []distributionxtypes.BuilderList{
		{
			Id:              0,
			ContractAddress: "",
		},
	}

	genesisState = CompileGenesis(t, config, genesisState, genAccounts, genBalances, genElectoral, *disXGen)
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

func MustAddressFormKeyring(kr keyring.Keyring, id string) sdk.AccAddress {
	r, err := kr.Key(id)
	if err != nil {
		panic(err)
	}
	p, err := r.GetPubKey()
	if err != nil {
		panic(err)
	}
	return sdk.AccAddress(p.Address())

}
