package ante

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	distributionxtypes "github.com/sharering/shareledger/x/distributionx/types"
)

type WasmKeeper interface {
	GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
}

type DistributionxKeeper interface {
	GetParams(ctx sdk.Context) distributionxtypes.Params
	IncReward(ctx sdk.Context, address string, coins sdk.Coins)
	IncBuilderCount(ctx sdk.Context, address string)
}
