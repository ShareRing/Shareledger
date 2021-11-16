package electoral

import (
	"encoding/json"

	"github.com/ShareRing/Shareledger/x/electoral/keeper"
	"github.com/ShareRing/Shareledger/x/electoral/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, addr := range genState.Addresses {
		voterID := types.VoterPrefix + addr
		k.SetVoterAddress(ctx, voterID, addr)
		k.SetVoterStatus(ctx, voterID, types.StatusVoterEnrolled)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var voterAddresses []string
	cb := func(voter types.Voter) bool {
		if voter.Status == types.StatusVoterEnrolled {
			voterAddresses = append(voterAddresses, voter.Address)
		}
		return false
	}
	k.IterateVoters(ctx, cb)
	return &types.GenesisState{
		Addresses: voterAddresses,
	}
}

func NewGenesisState(addrs []string) types.GenesisState {
	return types.GenesisState{
		Addresses: addrs,
	}
}

// TODO
func ValidateGenesis(data types.GenesisState) error {
	return nil
}

func DefaultGenesisState() types.GenesisState {
	return types.GenesisState{}
}

func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) types.GenesisState {
	var genesisState types.GenesisState
	if appState[types.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[types.ModuleName], &genesisState)
	}

	return genesisState
}
