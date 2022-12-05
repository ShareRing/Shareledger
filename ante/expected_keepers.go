package ante

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdistributiontypes "github.com/sharering/shareledger/x/sdistribution/types"
)

type WasmKeeper interface {
	GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
}

type SDistributionKeeper interface {
	GetParams(ctx sdk.Context) sdistributiontypes.Params
	IncReward(ctx sdk.Context, address string, coins sdk.Coins)
	IncBuilderCount(ctx sdk.Context, address string)
}
