package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	ibcante "github.com/cosmos/ibc-go/v3/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
)

func NewHandler(
	gentlemintKeeper GentlemintKeeper,
	accountKeeper ante.AccountKeeper,
	bankKeeper authtypes.BankKeeper,
	signModeHandler authsigning.SignModeHandler,
	feegrantKeeper ante.FeegrantKeeper,
	sigGasConsumer func(meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params) error,
	roleKeeper RoleKeeper,
	idKeeper IDKeeper,
	ibcKeeper *ibckeeper.Keeper,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewLoadFeeDecorator(gentlemintKeeper),
		NewCheckFeeDecorator(gentlemintKeeper),
		NewCosmosAuthAnteDecorator(
			accountKeeper,
			bankKeeper,
			signModeHandler,
			feegrantKeeper,
			sigGasConsumer,
		),
		NewAuthDecorator(roleKeeper, idKeeper),
		ibcante.NewAnteDecorator(ibcKeeper),
		sdk.Terminator{},
	)
}

type RoleKeeperWithoutVoter struct {
	RoleKeeper
}

func (r RoleKeeperWithoutVoter) IsVoter(ctx sdk.Context, address sdk.AccAddress) bool {
	return true
}
