package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AuthAnteDecorator map 1:1 with "github.com/cosmos/cosmos-sdk/x/auth/ante"
// In order to wrap in Shareledger ante chain
type AuthAnteDecorator struct {
	accountKeeper   ante.AccountKeeper
	bankKeeper      types.BankKeeper
	signModeHandler authsigning.SignModeHandler
	feegrantKeeper  ante.FeegrantKeeper
	sigGasConsumer  func(meter sdk.GasMeter, sig signing.SignatureV2, params types.Params) error
}

func NewCosmosAuthAnteDecorator(
	accountKeeper ante.AccountKeeper,
	bankKeeper types.BankKeeper,
	signModeHandler authsigning.SignModeHandler,
	feegrantKeeper ante.FeegrantKeeper,
	sigGasConsumer func(meter sdk.GasMeter, sig signing.SignatureV2, params types.Params) error,
) AuthAnteDecorator {
	return AuthAnteDecorator{
		accountKeeper:   accountKeeper,
		bankKeeper:      bankKeeper,
		signModeHandler: signModeHandler,
		feegrantKeeper:  feegrantKeeper,
		sigGasConsumer:  sigGasConsumer,
	}
}

func (a AuthAnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	authAnteHandler, err := ante.NewAnteHandler(ante.HandlerOptions{
		AccountKeeper:   a.accountKeeper,
		BankKeeper:      a.bankKeeper,
		SignModeHandler: a.signModeHandler,
		FeegrantKeeper:  a.feegrantKeeper,
		SigGasConsumer:  a.sigGasConsumer,
	})
	if err != nil {
		return
	}
	newCtx, err = authAnteHandler(ctx, tx, simulate)
	if err != nil {
		return newCtx, err
	}
	return next(newCtx, tx, simulate)
}
