package simulation

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v5/testing/simapp/helpers"
	"github.com/sharering/shareledger/x/electoral/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"math/rand"
)

func makeTransaction(_ *rand.Rand,
	app *baseapp.BaseApp, msg sdk.Msg, ak types.AccountKeeper, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey) error {
	var (
		err error
	)

	from, err := sdk.AccAddressFromBech32(msg.GetSigners()[0].String())
	if err != nil {
		return err
	}
	account := ak.GetAccount(ctx, from)
	txGen := simappparams.MakeTestEncodingConfig().TxConfig

	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(10000000000))), //10shr
		helpers.DefaultGenTxGas,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)

	if err != nil {
		return err
	}
	_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
	if err != nil {
		return err
	}
	return nil
}
