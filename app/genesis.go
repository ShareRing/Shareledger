package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/pos"
)

// State to Unmarshal
type GenesisState struct {
	Accounts  []GenesisAccount `json:"accounts"`
	StakeData pos.GenesisState `json:"stake"`
}

func (gs *GenesisState) ToJSON() []byte {

	cdc := MakeCodec()

	jsonBytes, err := cdc.MarshalJSONIndent(gs, "", "  ")
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// GenesisAccount doesn't need pubkey or sequence
type GenesisAccount struct {
	Address sdk.AccAddress `json:"address"`
	Coins   types.Coins `json:"coins"`
}

func NewGenesisAccount(acc *auth.SHRAccount) GenesisAccount {
	return GenesisAccount{
		Address: acc.Address,
		Coins:   acc.Coins,
	}
}

// convert GenesisAccount to auth.BaseAccount
func (ga *GenesisAccount) ToSHRAccount() (acc *auth.SHRAccount) {
	return &auth.SHRAccount{
		Address: ga.Address,
		Coins:   ga.Coins,
	}
}

func GenerateGenesisState(pubKey types.PubKeySecp256k1) GenesisState {
	return GenesisState{
		StakeData: pos.GenerateGenesis(pubKey),
	}
}
