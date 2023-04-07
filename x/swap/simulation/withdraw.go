package simulation

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
)

func SimulateMsgWithdraw(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		receiver, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgWithdraw{
			Creator:  simAccount.Address.String(),
			Receiver: receiver.Address.String(),
		}

		moduleAddr := ak.GetModuleAddress(types.ModuleName)

		availableCoins := bk.SpendableCoins(ctx, moduleAddr)
		if availableCoins.Empty() {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Withdraw not available now"), nil, nil
		}

		msg.Amount = sdk.NewDecCoin(denom.Base, simtypes.RandomAmount(r, availableCoins.AmountOf(denom.Base)))
		err := makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Withdraw not available now"), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
