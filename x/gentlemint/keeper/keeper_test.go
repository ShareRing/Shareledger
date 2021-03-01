/**
 * Based on https://github.com/cosmos/cosmos-sdk/blob/master/x/distribution/keeper
 */
package keeper

import (
	"testing"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/gentlemint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	log "github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

// returns context and app with params set on account keeper
func createOutput(t *testing.T) (sdk.Context, Keeper) {

	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)
	keySupply := sdk.NewKVStoreKey(supply.StoreKey)

	keyStoreGentlemint := sdk.NewKVStoreKey(types.StoreKey)
	cdc := MakeTestCodec()

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyStoreGentlemint, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	ms.MountStoreWithDB(keySupply, sdk.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	ctx := sdk.NewContext(ms, abci.Header{ChainID: "foochainid"}, true, log.NewNopLogger())

	pk := params.NewKeeper(cdc, keyParams, tkeyParams)

	accountKeeper := auth.NewAccountKeeper(
		cdc,    // amino codec
		keyAcc, // target store
		pk.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount, // prototype
	)

	blacklistedAddrs := make(map[string]bool)

	bk := bank.NewBaseKeeper(
		accountKeeper,
		pk.Subspace(bank.DefaultParamspace),
		blacklistedAddrs,
	)

	maccPerms := map[string][]string{
		auth.FeeCollectorName: nil,
		// types.NotBondedPoolName: {supply.Burner, supply.Staking},
		// types.BondedPoolName:    {supply.Burner, supply.Staking},
	}

	supplyKeeper := supply.NewKeeper(cdc, keySupply, accountKeeper, bk, maccPerms)

	k := NewKeeper(cdc, keyStoreGentlemint, accountKeeper, supplyKeeper, bk)

	return ctx, k
}

func TestAuthoritySet(t *testing.T) {
	ctx, k := createOutput(t)

	authorityAddr := "123"

	k.SetAuthorityAccount(ctx, authorityAddr)

	getAcc := k.GetAuthorityAccount(ctx)

	require.Equal(t, authorityAddr, getAcc)
}
