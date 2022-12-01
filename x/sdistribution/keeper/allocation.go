package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/sdistribution/types"
)

// AllocateTokens handles distribution of the collected fees
func (k *Keeper) AllocateTokens(ctx sdk.Context) {
	logger := k.Logger(ctx)
	logger.Debug("Start")

	params := k.GetParams(ctx)

	// allocate tokens from wasm-pool
	feeWasmCollector := k.authKeeper.GetModuleAccount(ctx, types.FeeWasmName)
	feeWasmCollected := k.bankKeeper.GetAllBalances(ctx, feeWasmCollector.GetAddress())
	k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.FeeWasmName, types.ModuleName, feeWasmCollected)

	builderList := k.GetAllBuilderList(ctx)
	for _, builder := range builderList {
		addr, err := sdk.AccAddressFromBech32(builder.ContractAddress)
		if err != nil {
			panic(err)
		}
		contractInfo := k.wasmKeeper.GetContractInfo(ctx, addr)
		rate := params.WasmMasterBuilder / float64(len(builderList))
		k.IncReward(ctx, contractInfo.Creator, getFeeRounded(feeWasmCollected, rate))
	}

	// distribution for dev_pool
	k.IncReward(ctx, params.DevPoolAccount, getFeeRounded(feeWasmCollected, params.WasmDevelopment))

	// allocate tokens from wasm-pool
	feeNativeCollector := k.authKeeper.GetModuleAccount(ctx, types.FeeNativeName)
	feeNativeCollected := k.bankKeeper.GetAllBalances(ctx, feeNativeCollector.GetAddress())
	k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.FeeNativeName, types.ModuleName, feeNativeCollected)

	// allocate tokens from native-pool
	k.IncReward(ctx, params.DevPoolAccount, getFeeRounded(feeWasmCollected, params.NativeDevelopment))

}

func getFeeRounded(fee sdk.Coins, rate float64) sdk.Coins {
	// round params to 4 decimals
	const ROUND_FACTOR = 10000
	tp := sdkmath.NewInt(int64(rate * ROUND_FACTOR))
	return fee.MulInt(tp).QuoInt(sdkmath.NewInt(ROUND_FACTOR))
}
