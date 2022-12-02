package ante

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CountBuilderDecorator struct {
	sdistributionKeeper SDistributionKeeper
}

// count tx by contract address
func NewCountBuilderDecorator(sk SDistributionKeeper) CountBuilderDecorator {
	return CountBuilderDecorator{
		sdistributionKeeper: sk,
	}
}

func (cbd CountBuilderDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if execMsg, ok := msg.(*wasmtypes.MsgExecuteContract); ok {
			cbd.sdistributionKeeper.IncBuilderCount(ctx, execMsg.Contract)
		}
	}
	return next(ctx, tx, simulate)
}
