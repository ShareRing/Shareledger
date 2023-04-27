package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/types"
)

// AllocateTokens handles distribution of the collected fees
func (k *Keeper) AllocateTokens(ctx sdk.Context) {
	logger := k.Logger(ctx)
	logger.Debug("Distributionx AllocateTokens")

	params := k.GetParams(ctx)

	// 1. Allocate tokens from wasm-pool to master_builder_list and dev_pool
	feeWasmCollector := k.authKeeper.GetModuleAccount(ctx, types.FeeWasmName)
	feeWasmCollected := k.bankKeeper.GetAllBalances(ctx, feeWasmCollector.GetAddress())
	if !feeWasmCollected.IsZero() {
		err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.FeeWasmName, types.ModuleName, feeWasmCollected)
		if err != nil {
			panic(fmt.Errorf("SendCoinsFromModuleToModule allocateTokens in distributionx: %w", err))
		}
		config := params.ConfigPercent
		totalRate := config.WasmMasterBuilder.Add(config.WasmDevelopment).Add(config.WasmContractAdmin)
		adminAmount := getFeeRounded(feeWasmCollected, config.WasmContractAdmin.Quo(totalRate))

		builderList := k.GetAllBuilderList(ctx)
		if len(builderList) > 0 {
			feeBuilderRate := config.WasmMasterBuilder.Quo(totalRate)
			rate := feeBuilderRate.Quo(sdk.NewDec(int64(len(builderList))))
			rewardAmount := getFeeRounded(feeWasmCollected, rate)
			feeWasmCollected = feeWasmCollected.Sub(rewardAmount.MulInt(sdkmath.NewInt(int64(len(builderList))))...)
			for _, builder := range builderList {
				addr, err2 := sdk.AccAddressFromBech32(builder.ContractAddress)
				if err2 != nil {
					panic(err2)
				}
				contractInfo := k.wasmKeeper.GetContractInfo(ctx, addr)
				k.IncReward(ctx, contractInfo.Creator, rewardAmount)
			}
		}

		feeWasmCollected = feeWasmCollected.Sub(adminAmount...)

		// distribution for dev_pool
		k.IncReward(ctx, params.DevPoolAccount, feeWasmCollected)
	}

	// 2. Allocate tokens from native-pool
	feeNativeCollector := k.authKeeper.GetModuleAccount(ctx, types.FeeNativeName)
	feeNativeCollected := k.bankKeeper.GetAllBalances(ctx, feeNativeCollector.GetAddress())
	if !feeNativeCollected.IsZero() {
		err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.FeeNativeName, types.ModuleName, feeNativeCollected)
		if err != nil {
			panic(fmt.Errorf("SendCoinsFromModuleToModule allocateTokens in distributionx: %w", err))
		}
		k.IncReward(ctx, params.DevPoolAccount, feeNativeCollected)
	}
}

// TODO: make this logic cleaner
func getFeeRounded(fee sdk.Coins, rate sdk.Dec) sdk.Coins {
	rateFloat := rate.MustFloat64()
	const rountFactor = 10000
	tp := sdkmath.NewInt(int64(rateFloat * rountFactor))
	return fee.MulInt(tp).QuoInt(sdkmath.NewInt(rountFactor))
}
