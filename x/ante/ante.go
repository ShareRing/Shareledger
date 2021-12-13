package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func NewHandler(
	gentlemintKeeper GentlemintKeeper,
	accountKeeper ante.AccountKeeper,
	bankKeeper authtypes.BankKeeper,
	signModeHandler authsigning.SignModeHandler,
	feegrantKeeper ante.FeegrantKeeper,
	sigGasConsumer func(meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params) error,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewCheckFeeDecorator(gentlemintKeeper),
		NewAuthAnteDecorator(
			accountKeeper,
			bankKeeper,
			signModeHandler,
			feegrantKeeper,
			sigGasConsumer,
		),
		sdk.Terminator{},
	)
}
