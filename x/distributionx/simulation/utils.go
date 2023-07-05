package simulation

import (
	"errors"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/ibc-go/v5/testing/simapp/helpers"

	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func SimBroadcastTransaction(r *rand.Rand,
	app *baseapp.BaseApp, msg sdk.Msg, ak types.AccountKeeper, bk types.BankKeeper, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey) error {

	var (
		fees sdk.Coins
		err  error
	)

	from, err := sdk.AccAddressFromBech32(msg.GetSigners()[0].String())
	if err != nil {
		return err
	}
	account := ak.GetAccount(ctx, from)
	balances := bk.GetAllBalances(ctx, from)

	var baseBalance sdk.Coin

	for _, c := range balances {
		if c.Denom == denom.Base {
			baseBalance = c
		}
	}
	if baseBalance.IsZero() {
		return errors.New("balance empty")
	}

	fees, err = simtypes.RandomFees(r, ctx, sdk.NewCoins(baseBalance))
	if err != nil {
		return err
	}
	txGen := simappparams.MakeTestEncodingConfig().TxConfig

	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		fees,
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
