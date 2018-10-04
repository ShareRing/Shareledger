package app

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/pos"
)

// State to Unmarshal
type GenesisState struct {
	//	Accounts  []GenesisAccount `json:"accounts"`
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
	Address sdk.Address `json:"address"`
	Coins   types.Coins `json:"coins"`
}

func GenerateGenesisState(pubKey types.PubKeySecp256k1) GenesisState {
	return GenesisState{
		StakeData: pos.GenerateGenesis(pubKey),
	}
}
