package keeper

import (
	"cosmossdk.io/math"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/types"
)

// AllocateTokens handles distribution of the collected fees
func (k *Keeper) AllocateTokens(ctx sdk.Context) {
	logger := k.Logger(ctx)
	logger.Debug("AllocateTokens")

	params := k.GetParams(ctx)

	// 1. Allocate tokens from wasm-pool to master_builder_list and dev_pool
	feeWasmCollector := k.authKeeper.GetModuleAccount(ctx, types.FeeWasmName)
	feeWasmCollected := k.bankKeeper.GetAllBalances(ctx, feeWasmCollector.GetAddress())
	k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.FeeWasmName, types.ModuleName, feeWasmCollected)

	feeBuilderRate := params.WasmMasterBuilder.Quo(params.WasmMasterBuilder.Add(params.WasmDevelopment))

	builderList := k.GetAllBuilderList(ctx)
	if len(builderList) > 0 {
		rate := feeBuilderRate.Quo(sdk.NewDec(int64(len(builderList))))
		rewardAmount := getFeeRounded(feeWasmCollected, rate)
		feeWasmCollected.Sub(rewardAmount.MulInt(math.NewInt(int64(len(builderList))))...)
		for _, builder := range builderList {
			addr, err := sdk.AccAddressFromBech32(builder.ContractAddress)
			if err != nil {
				panic(err)
			}
			contractInfo := k.wasmKeeper.GetContractInfo(ctx, addr)
			k.IncReward(ctx, contractInfo.Creator, rewardAmount)
		}
	}

	// distribution for dev_pool
	k.IncReward(ctx, params.DevPoolAccount, feeWasmCollected)

	// 2. Allocate tokens from native-pool
	feeNativeCollector := k.authKeeper.GetModuleAccount(ctx, types.FeeNativeName)
	feeNativeCollected := k.bankKeeper.GetAllBalances(ctx, feeNativeCollector.GetAddress())
	k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.FeeNativeName, types.ModuleName, feeNativeCollected)

	k.IncReward(ctx, params.DevPoolAccount, feeNativeCollected)

}

// TODO: make this logic cleaner
func getFeeRounded(fee sdk.Coins, rate sdk.Dec) sdk.Coins {
	rateFloat := rate.MustFloat64()
	const ROUND_FACTOR = 10000
	tp := sdkmath.NewInt(int64(rateFloat * ROUND_FACTOR))
	return fee.MulInt(tp).QuoInt(sdkmath.NewInt(ROUND_FACTOR))
}
