package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	uuid2 "github.com/google/uuid"
	"github.com/thanhpk/randstr"

	"github.com/sharering/shareledger/testutil"
	assetTypes "github.com/sharering/shareledger/x/asset/types"
)

func RandomizedGenState(simState *module.SimulationState) {
	assets := MustRandAssets(simState.Rand, 10, simState)

	assetGenesis := assetTypes.GenesisState{Assets: assets}

	simState.GenState[assetTypes.ModuleName] = simState.Cdc.MustMarshalJSON(&assetGenesis)

}

// MustRandAssets generate random assets information panic if any error
func MustRandAssets(r *rand.Rand, numAsset int, simState *module.SimulationState) []*assetTypes.Asset {
	assets := make([]*assetTypes.Asset, numAsset)
	for i := 0; i < numAsset; i++ {

		randCreator, _ := simtypes.RandomAcc(r, simState.Accounts)
		randHash := randstr.Hex(16)
		uuid, err := uuid2.NewRandom()
		if err != nil {
			panic(err)
		}

		asset := &assetTypes.Asset{
			Creator: randCreator.Address.String(),
			Hash:    []byte(randHash),
			UUID:    uuid.String(),
			Status:  testutil.RandBool(r),
			Rate:    0,
		}

		assets[i] = asset
	}
	return assets
}
