package simulation

import (
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func MustGenRandGenesis(simState module.SimulationState) {
	// randRate := testutil.RandRate(input.Rand)
	miniMumGasPrice := sdk.DecCoins{}
	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(types.ParamStoreKeyMinGasPrices), &miniMumGasPrice, simState.Rand,
		func(r *rand.Rand) {
			miniMumGasPrice = sdk.NewDecCoinsFromCoins(sdk.NewCoin(denom.Base, sdk.NewInt(0)))
		},
	)

	gentlemintGenesis := &types.GenesisState{
		ExchangeRate: &types.ExchangeRate{
			Rate: fmt.Sprintf("%d", simState.Rand.Int31n(40)),
		},
		LevelFeeList:       []types.LevelFee{},
		ActionLevelFeeList: []types.ActionLevelFee{},
		Params: types.Params{
			MinimumGasPrices: miniMumGasPrice,
		},
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(gentlemintGenesis)
}
