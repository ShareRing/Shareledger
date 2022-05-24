package network

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
	"testing"
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
	Accounts[id] = info.GetAddress()

	kb.accGens = append(kb.accGens, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	kb.genBalances = append(kb.genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   coins,
	})
}

func (kb *KeyRingBuilder) GenETHSigner(id string, coins sdk.Coins) {
	kb.InitUserByHDPath(id, coins, KeyPathETH) //ETH hd path
}
func (kb *KeyRingBuilder) GenBSCSigner(id string, coins sdk.Coins) {
	kb.InitUserByHDPath(id, coins, KeyPathBSC) //ETH hd path
}

func (kb *KeyRingBuilder) NewAccountToSign() {

	info, err := kb.kb.NewAccount(KeyAccountTestSign, SignMnemonic, keyring.DefaultBIP39Passphrase, KeyPathETH, hd.Secp256k1)
	require.NoError(kb.t, err, "init fail")
	Accounts[KeyAccountTestSign] = info.GetAddress()

	kb.accGens = append(kb.accGens, authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0))
	kb.genBalances = append(kb.genBalances, banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   OneThousandSHROneHundredSHRPCoins,
	})
}
