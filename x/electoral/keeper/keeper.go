package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		gmKeeper types.GentlemintKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	gmKeeper types.GentlemintKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		gmKeeper: gmKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetVoter(ctx sdk.Context, voterID string) types.Voter {
	store := ctx.KVStore(k.storeKey)

	if !k.IsVoterPresent(ctx, voterID) {
		return types.NewVoter()
	}

	bz := store.Get([]byte(voterID))

	var result types.Voter

	k.cdc.MustUnmarshalLengthPrefixed(bz, &result)

	return result
}

func (k Keeper) SetVoter(ctx sdk.Context, voterID string, v types.Voter) {
	// TODO: return error?
	if len(v.Address) == 0 {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(voterID), k.cdc.MustMarshalLengthPrefixed(&v))
}

func (k Keeper) GetVoterStatus(ctx sdk.Context, voterID string) string {
	voter := k.GetVoter(ctx, voterID)
	return voter.Status
}

func (k Keeper) SetVoterAddress(ctx sdk.Context, voterID string, addr string) {
	voter := k.GetVoter(ctx, voterID)
	voter.Address = addr
	k.SetVoter(ctx, voterID, voter)
}

func (k Keeper) SetVoterStatus(ctx sdk.Context, voterID string, status string) {
	voter := k.GetVoter(ctx, voterID)
	voter.Status = status
	k.SetVoter(ctx, voterID, voter)
}

func (k Keeper) DeleteVoter(ctx sdk.Context, voterID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(voterID))
}

func (k Keeper) GetVotersIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) IterateVoters(ctx sdk.Context, cb func(voter types.Voter) (stop bool)) {
	iterator := k.GetVotersIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var voter types.Voter
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &voter)

		if cb(voter) {
			break
		}
	}
}

func (k Keeper) IsVoterPresent(ctx sdk.Context, voterID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(voterID))
}

func (k Keeper) IsAuthority(ctx sdk.Context, addr sdk.AccAddress) bool {
	authAddr := k.gmKeeper.GetAuthorityAccount(ctx)

	return authAddr == addr.String()
}
