package identity

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	IDSigners []string
	IDs       map[string]string
}

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, signerKey := range data.IDSigners {
		keeper.SetIdSignerStatus(ctx, signerKey, types.IdSignerActive)
	}
	for idKey, hash := range data.IDs {
		keeper.SetId(ctx, idKey, hash)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var idSignerKeys []string
	cb := func(idSignerKey string, s types.IdSigner) bool {
		if s.Status == types.IdSignerActive {
			idSignerKeys = append(idSignerKeys, idSignerKey)
		}
		return false
	}
	k.IterateIdSigners(ctx, cb)

	ids := make(map[string]string)
	cb2 := func(idKey, hash string) bool {
		ids[idKey] = hash
		return false
	}
	k.IterateIds(ctx, cb2)

	return GenesisState{
		IDSigners: idSignerKeys,
		IDs:       ids,
	}
}
