package network

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
)

const (
	KeyPathETH = `m/44'/60'/0'/0/0`
	KeyPathBSC = `m/44'/714'/0'/0`
)

type (
	KeyRingBuilder struct {
		t           *testing.T
		kb          keyring.Keyring
		accGens     []authtypes.GenesisAccount
		genBalances []banktypes.Balance
	}
)

func NewKeyringBuilder(t *testing.T, kr keyring.Keyring) *KeyRingBuilder {
	return &KeyRingBuilder{
		t:           t,
		kb:          kr,
		accGens:     []authtypes.GenesisAccount{},
		genBalances: []banktypes.Balance{},
	}
}

func (kb *KeyRingBuilder) BuildGenesis() (keyring.Keyring, []authtypes.GenesisAccount, []banktypes.Balance) {
	return kb.kb, kb.accGens, kb.genBalances
}

func (kb *KeyRingBuilder) InitUser(id string, coins sdk.Coins) {
	kb.InitUserByHDPath(id, coins, sdk.FullFundraiserPath)
}

func (kb *KeyRingBuilder) InitUserByHDPath(id string, coins sdk.Coins, path string) {
	info, _, err := kb.kb.NewMnemonic(id, keyring.English, path, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(kb.t, err, "init fail")
	Accounts[id] = mustNewAddr(info)

	kb.accGens = append(kb.accGens, authtypes.NewBaseAccount(mustNewAddr(info), mustNewPubKey(info), 0, 0))
	kb.genBalances = append(kb.genBalances, banktypes.Balance{
		Address: mustNewAddr(info).String(),
		Coins:   coins,
	})
}

func (kb *KeyRingBuilder) GenETHSigner(id string, coins sdk.Coins) {
	kb.InitUserByHDPath(id, coins, KeyPathETH) // ETH hd path
}

func (kb *KeyRingBuilder) GenBSCSigner(id string, coins sdk.Coins) {
	kb.InitUserByHDPath(id, coins, KeyPathBSC) // ETH hd path
}

func (kb *KeyRingBuilder) NewAccountToSign() {
	info, err := kb.kb.NewAccount(KeyAccountTestSign, SignMnemonic, keyring.DefaultBIP39Passphrase, KeyPathETH, hd.Secp256k1)
	require.NoError(kb.t, err, "init fail")
	Accounts[KeyAccountTestSign] = mustNewAddr(info)

	kb.accGens = append(kb.accGens, authtypes.NewBaseAccount(mustNewAddr(info), mustNewPubKey(info), 0, 0))
	kb.genBalances = append(kb.genBalances, banktypes.Balance{
		Address: mustNewAddr(info).String(),
		Coins:   OneThousandSHROneHundredSHRPCoins,
	})
}

func mustNewAddr(info *keyring.Record) sdk.AccAddress {
	addr, err := info.GetAddress()
	if err != nil {
		panic(err)
	}
	return addr
}

func mustNewPubKey(info *keyring.Record) cryptotypes.PubKey {
	pub, err := info.GetPubKey()
	if err != nil {
		panic(err)
	}
	return pub
}
