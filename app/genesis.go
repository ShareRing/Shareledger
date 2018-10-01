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

// GenesisAccount doesn't need pubkey or sequence
type GenesisAccount struct {
	Address sdk.Address `json:"address"`
	Coins   types.Coins `json:"coins"`
}
