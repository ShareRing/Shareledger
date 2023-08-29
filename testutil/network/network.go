package network

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/testutil/simapp"
	electoraltypes "github.com/sharering/shareledger/x/electoral/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

// AppConstructor defines a function which accepts a network configuration and
// creates an ABCI Application to provide to Tendermint.

var Accounts = map[string]sdk.Address{}

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
func New(t *testing.T, rootDir string, configs ...network.Config) *network.Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg network.Config
	if len(configs) == 0 {
		cfg = DefaultConfig()
	} else {
		cfg = configs[0]
	}
	// net, err := network.New(t, t.TempDir(), cfg)
	net, err := network.New(t, rootDir, cfg)
	require.NoError(t, err)
	t.Cleanup(net.Cleanup)
	return net
}

func CompileGenesis(t *testing.T, config *network.Config, genesisState map[string]json.RawMessage, au []authtypes.GenesisAccount, b []banktypes.Balance, elGen electoraltypes.GenesisState) map[string]json.RawMessage {
	var bankGenesis banktypes.GenesisState
	var authGenesis authtypes.GenesisState
	var stakingGenesis stakingtypes.GenesisState

	config.Codec.MustUnmarshalJSON(genesisState[banktypes.ModuleName], &bankGenesis)
	config.Codec.MustUnmarshalJSON(genesisState[authtypes.ModuleName], &authGenesis)
	config.Codec.MustUnmarshalJSON(genesisState[stakingtypes.ModuleName], &stakingGenesis)

	accounts, err := authtypes.PackAccounts(au)
	if err != nil {
		t.Errorf("int fails")
	}

	authGenesis.Accounts = append(authGenesis.Accounts, accounts...)
	bankGenesis.Balances = b
	stakingGenesis.Params.BondDenom = denom.Base

	genesisState[banktypes.ModuleName] = config.Codec.MustMarshalJSON(&bankGenesis)
	genesisState[authtypes.ModuleName] = config.Codec.MustMarshalJSON(&authGenesis)
	genesisState[electoraltypes.ModuleName] = config.Codec.MustMarshalJSON(&elGen)
	genesisState[stakingtypes.ModuleName] = config.Codec.MustMarshalJSON(&stakingGenesis)

	return genesisState
}

// SetTestingGenesis init the genesis state for testing in here
func SetTestingGenesis(t *testing.T, config *network.Config, nodedir, moniker string) (keyring.Keyring, string) {
	genesisState := config.GenesisState

	// kb := keyring.NewInMemory(config.Codec, config.KeyringOptions...)
	kb, err := keyring.New(moniker, keyring.BackendTest, nodedir, nil, config.Codec, config.KeyringOptions...)
	if err != nil {
		return nil, ""
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

	idSignerAcc := MustAddressFormKeyring(kb, KeyIDSigner)
	docIssuerAcc := MustAddressFormKeyring(kb, KeyDocIssuer)
	relayerAcc := MustAddressFormKeyring(kb, KeyApproverRelayer)
	genElectoral := electoraltypes.GenesisState{
		AccStateList: []electoraltypes.AccState{
			{
				Key:     fmt.Sprintf("idsigner%s", idSignerAcc),
				Address: idSignerAcc.String(),
				Status:  "active",
			},
			{
				Key:     fmt.Sprintf("docIssuer%s", docIssuerAcc),
				Address: docIssuerAcc.String(),
				Status:  "active",
			},
			{
				Key:     fmt.Sprintf("docIssuer%s", docIssuerAcc),
				Address: docIssuerAcc.String(),
				Status:  "active",
			},
			{
				Key:     string(electoraltypes.GenAccStateIndexKey(relayerAcc, electoraltypes.AccStateKeyRelayer)),
				Address: relayerAcc.String(),
				Status:  "active",
			},
			{
				Key:     string(electoraltypes.GenAccStateIndexKey(relayerAcc, electoraltypes.AccStateKeyApprover)),
				Address: relayerAcc.String(),
				Status:  "active",
			},
		},
		Authority: &electoraltypes.Authority{Address: Accounts[KeyAuthority].String()},
		Treasurer: &electoraltypes.Treasurer{Address: Accounts[KeyTreasurer].String()},
	}

	genesisState = CompileGenesis(t, config, genesisState, genAccounts, genBalances, genElectoral)
	config.GenesisState = genesisState
	return newKeyringService, ""
}

// ShareLedgerChainConstructor returns a new shareLedger AppConstructor
func ShareLedgerChainConstructor() AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return simapp.New(simtestutil.AppOptionsMap{
			app.FlagAppOptionSkipCheckVoter: true,
			flags.FlagHome:                  val.Ctx.Config.RootDir,
		})
	}
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
