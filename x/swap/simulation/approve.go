package simulation

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/ibc-go/v5/testing/simapp/helpers"
	"github.com/sharering/shareledger/testutil"
	"github.com/thanhpk/randstr"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
)

func SimulateMsgApprove(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgApproveOut{
			Signature: randstr.Hex(50),
			Creator:   simAccount.Address.String(),
		}
		reqIn, err := k.Swap(ctx, &types.QuerySwapRequest{
			Status:      types.SwapStatusPending,
			DestNetwork: testutil.RandNetwork(r),
		})
		if err != nil || len(reqIn.GetSwaps()) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no pending request out found "), nil, nil
		}

		reqs := testutil.RandPick(r, reqIn.GetSwaps())
		msg.Ids = []uint64{reqs.Id}
		err = makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

func makeTransaction(r *rand.Rand,
	app *baseapp.BaseApp, msg sdk.Msg, ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey) error {
	var (
		fees sdk.Coins
		err  error
	)

	from, err := sdk.AccAddressFromBech32(msg.GetSigners()[0].String())
	if err != nil {
		return err
	}
	account := ak.GetAccount(ctx, from)
	balances := bk.GetBalance(ctx, from, "nshr")
	fees, err = simtypes.RandomFees(r, ctx, sdk.NewCoins(balances))
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
