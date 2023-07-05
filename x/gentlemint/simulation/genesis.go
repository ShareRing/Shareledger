package simulation

import (
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/fee"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

const (
	MinGasPrice = "min_gas_price"
)

func MustGenerateGentlemintGenesis(simSate *module.SimulationState) {

	var exchangeRate int

	for exchangeRate > 0 {
		exchangeRate = simSate.Rand.Intn(200)
	}

	var lowInt int
	for lowInt > 0 {
		lowInt = simSate.Rand.Intn(50000000)
	}

	low := types.LevelFee{
		Level:   "low",
		Fee:     sdk.NewDecCoin(denom.Base, sdk.NewInt(int64(lowInt))),
		Creator: testutil.RandPick(simSate.Rand, simSate.Accounts).Address.String(),
	}

	var mediumInt int
	for mediumInt > 0 {
		mediumInt = simSate.Rand.Intn(50000000)
	}

	medium := types.LevelFee{
		Level:   "medium",
		Fee:     sdk.NewDecCoin(denom.Base, sdk.NewInt(int64(mediumInt))),
		Creator: testutil.RandPick(simSate.Rand, simSate.Accounts).Address.String(),
	}

	var highInt int
	for highInt > 0 {
		highInt = simSate.Rand.Intn(50000000)
	}

	high := types.LevelFee{
		Level:   "high",
		Fee:     sdk.NewDecCoin(denom.Base, sdk.NewInt(int64(highInt))),
		Creator: testutil.RandPick(simSate.Rand, simSate.Accounts).Address.String(),
	}

	feeDefault := fee.GetListActionsWithDefaultLevel()
	actionLevelFees := make([]types.ActionLevelFee, 0, len(feeDefault))
	for a, l := range feeDefault {
		actionLevelFees = append(actionLevelFees, types.ActionLevelFee{
			Action: a,
			Level:  l,
		})
	}

	minimumGasPriceDec := sdk.DecCoins{}
	simSate.AppParams.GetOrGenerate(simSate.Cdc, MinGasPrice, &minimumGasPriceDec, simSate.Rand,
		func(r *rand.Rand) {
			var minimumGasPrice int
			for minimumGasPrice > 0 {
				minimumGasPrice = r.Intn(5000)
			}
			minimumGasPriceDec = sdk.NewDecCoins(sdk.NewDecCoin(denom.Base, sdk.NewInt(int64(minimumGasPrice))))
		})

	gentlemintGen := &types.GenesisState{
		ExchangeRate:       &types.ExchangeRate{Rate: fmt.Sprintf("%d", exchangeRate)},
		LevelFeeList:       []types.LevelFee{low, medium, high},
		ActionLevelFeeList: actionLevelFees,
		Params:             types.Params{MinimumGasPrices: minimumGasPriceDec},
	}

	genBz := simSate.Cdc.MustMarshalJSON(gentlemintGen)
	simSate.GenState[types.ModuleName] = genBz
	return

}
