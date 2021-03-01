package gatecheck

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AuthAnteDecorator struct {
	ak             keeper.AccountKeeper
	supplyKeeper   types.SupplyKeeper
	sigGasConsumer ante.SignatureVerificationGasConsumer
}

func NewAuthAnteDecorator(ak keeper.AccountKeeper, sk types.SupplyKeeper, sgc ante.SignatureVerificationGasConsumer) AuthAnteDecorator {
	return AuthAnteDecorator{
		ak:             ak,
		supplyKeeper:   sk,
		sigGasConsumer: sgc,
	}
}

func (aad AuthAnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	authAnteHandler := auth.NewAnteHandler(aad.ak, aad.supplyKeeper, aad.sigGasConsumer)
	newCtx, err = authAnteHandler(ctx, tx, simulate)
	if err != nil {
		return newCtx, err
	}
	return next(newCtx, tx, simulate)
}
