package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func SimulateMsgDeposit(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDeposit{
			Creator: simAccount.Address.String(),
		}

		availableCoins := bk.SpendableCoins(ctx, simAccount.Address)

		if availableCoins.Empty() {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "account balance isn't enough"), nil, nil
		}
		msg.Amount = testutil.PtrOf(sdk.NewDecCoin(denom.Base, simtypes.RandomAmount(r, availableCoins.AmountOf(denom.Base))))

		err := makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
