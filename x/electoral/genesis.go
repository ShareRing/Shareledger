package electoral

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

type GenesisState struct {
	Addresses []sdk.AccAddress `json:"Addresses"`
}

func NewGenesisState(addrs []sdk.AccAddress) GenesisState {
	return GenesisState{
		Addresses: addrs,
	}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, addr := range data.Addresses {
		voterID := VoterPrefix + addr.String()
		keeper.SetVoterAddress(ctx, voterID, addr)
		keeper.SetVoterStatus(ctx, voterID, types.StatusVoterEnrolled)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var voterAddresses []sdk.AccAddress
	cb := func(voter types.Voter) bool {
		if voter.Status == types.StatusVoterEnrolled {
			voterAddresses = append(voterAddresses, voter.Address)
		}
		return false
	}
	k.IterateVoters(ctx, cb)
	return GenesisState{
		Addresses: voterAddresses,
	}
}

func GetGenesisStateFromAppState(cdc *codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}
