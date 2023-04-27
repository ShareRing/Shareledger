package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func RandomizedGenState(simState *module.SimulationState) {
	// var wasmMasterBuilder sdk.Dec
	// simState.AppParams.GetOrGenerate(simState.Cdc, string(types.WasmMasterBuilderKey), &wasmMasterBuilder,
	// 	simState.Rand, func(r *rand.Rand) {
	// 		wasmMasterBuilder = sdk.NewDec(r.Int63n(40))
	// 	})
	//
	// var wasmContractAdmin sdk.Dec
	// simState.AppParams.GetOrGenerate(simState.Cdc, string(types.WasmContractAdminKey), &wasmContractAdmin,
	// 	simState.Rand, func(r *rand.Rand) {
	// 		wasmContractAdmin = sdk.NewDec(r.Int63n(40))
	// 	})
	//
	// var wasmDevelopment sdk.Dec
	// simState.AppParams.GetOrGenerate(simState.Cdc, string(types.WasmDevelopmentKey), &wasmDevelopment,
	// 	simState.Rand, func(r *rand.Rand) {
	// 		wasmDevelopment = sdk.NewDec(r.Int63n(40))
	// 	})
	//
	// var wasmValidator sdk.Dec
	// simState.AppParams.GetOrGenerate(simState.Cdc, string(types.WasmValidatorKey), &wasmValidator,
	// 	simState.Rand, func(r *rand.Rand) {
	// 		wasmValidator = sdk.NewDec(r.Int63n(40))
	// 	})
	//
	// var nativeValidator sdk.Dec
	// simState.AppParams.GetOrGenerate(simState.Cdc, string(types.NativeValidatorKey), &nativeValidator,
	// 	simState.Rand, func(r *rand.Rand) {
	// 		nativeValidator = sdk.NewDec(r.Int63n(40))
	// 	})
	//
	// var nativeDevelopment sdk.Dec
	// simState.AppParams.GetOrGenerate(simState.Cdc, string(types.NativeDevelopmentKey), &nativeDevelopment,
	// 	simState.Rand, func(r *rand.Rand) {
	// 		nativeDevelopment = sdk.NewDec(r.Int63n(40))
	// 	})

	var wasmBuilderWindow uint32
	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.BuilderWindowsKey), &wasmBuilderWindow,
		simState.Rand, func(r *rand.Rand) {
			wasmBuilderWindow = uint32(r.Int31n(8-1)+1) * 1000
		})

	var txThreshold uint32
	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.TxThresholdKey), &txThreshold,
		simState.Rand, func(r *rand.Rand) {
			txThreshold = uint32(r.Int31n(8-1)+1) * 100
		})

	var devPoolAccount string
	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.DevPoolAccountKey), &devPoolAccount,
		simState.Rand, func(r *rand.Rand) {
			devPoolAccount = testutil.RandPick(r, simState.Accounts).Address.String()
		})

	builderList := mustRandBuilderList(simState.Rand, simState.Accounts)
	distributionXGenesis := &types.GenesisState{
		Params: types.Params{
			ConfigPercent:  types.DefaultParams().ConfigPercent,
			BuilderWindows: wasmBuilderWindow,
			TxThreshold:    txThreshold,
			DevPoolAccount: devPoolAccount,
		},
		RewardList:       mustRandReward(simState.Rand, simState.Accounts),
		BuilderCountList: mustRandBuilderCount(simState.Rand, simState.Accounts),
		BuilderListList:  builderList,
		BuilderListCount: uint64(len(builderList)),
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(distributionXGenesis)
}

func mustRandReward(r *rand.Rand, simAcc []simulation.Account) []types.Reward {
	rNum := r.Intn(20)
	rewards := make([]types.Reward, rNum)

	for i := 0; i < rNum; i++ {

		rAmount := simulation.RandIntBetween(r, 10, 100000000)
		amount := sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(int64(rAmount))))
		rewards[i] = types.Reward{
			Index:  testutil.RandPick(r, simAcc).Address.String(),
			Amount: amount,
		}
	}
	return rewards
}

func mustRandBuilderCount(r *rand.Rand, simAcc []simulation.Account) []types.BuilderCount {
	bNum := r.Intn(20)
	builderCount := make([]types.BuilderCount, bNum)
	for i := 0; i < bNum; i++ {
		builderCount[i] = types.BuilderCount{
			Index: testutil.RandPick(r, simAcc).Address.String(),
			Count: uint32(simulation.RandIntBetween(r, 1, 32)),
		}
	}
	return builderCount
}

func mustRandBuilderList(r *rand.Rand, simAcc []simulation.Account) []types.BuilderList {
	bNum := r.Intn(20)
	builderList := make([]types.BuilderList, bNum)
	for i := 0; i < bNum; i++ {
		builderList[i] = types.BuilderList{
			Id:              uint64(i + 1),
			ContractAddress: testutil.RandPick(r, simAcc).Address.String(),
		}
	}
	return builderList
}
