package gatecheck

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/electoral"
	"bitbucket.org/shareringvietnam/shareledger-fix/x/gentlemint"
	"bitbucket.org/shareringvietnam/shareledger-modules/id"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

func NewAnteHandler(ak keeper.AccountKeeper, ek electoral.Keeper, supplyKeeper types.SupplyKeeper, gmKeeper gentlemint.Keeper, idKeeper id.Keeper, sigGasConsumer ante.SignatureVerificationGasConsumer) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewCheckValDecorator(ek),
		NewCheckFeeDecorator(gmKeeper),
		NewFeeEnsureDecorator(gmKeeper),
		NewAuthAnteDecorator(ak, supplyKeeper, sigGasConsumer),
		NewCheckAuthDecorator(gmKeeper, idKeeper),
	)
}
