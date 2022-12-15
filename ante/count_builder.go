package ante

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CountBuilderDecorator struct {
	distributionxKeeper DistributionxKeeper
}

// count tx by contract address
func NewCountBuilderDecorator(sk DistributionxKeeper) CountBuilderDecorator {
	return CountBuilderDecorator{
		distributionxKeeper: sk,
	}
}

func (cbd CountBuilderDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if execMsg, ok := msg.(*wasmtypes.MsgExecuteContract); ok {
			cbd.distributionxKeeper.IncBuilderCount(ctx, execMsg.Contract)
		}
	}
	return next(ctx, tx, simulate)
}
