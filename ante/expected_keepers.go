package ante

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdistributiontypes "github.com/sharering/shareledger/x/sdistribution/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetParams(ctx sdk.Context) sdistributiontypes.Params
}

type WasmKeeper interface {
	GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
}

type FeegrantKeeper interface{}

type SDistributionKeeper interface {
	GetParams(ctx sdk.Context) sdistributiontypes.Params
	IncReward(ctx sdk.Context, address string, coins sdk.Coins)
}
