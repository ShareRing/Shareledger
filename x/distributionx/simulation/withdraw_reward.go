package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/distributionx/keeper"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func SimulateMsgWithdrawReward(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgWithdrawReward{}

		moduleAddr := ak.GetModuleAccount(ctx, types.ModuleName).GetAddress()

		moduleBalance := bk.SpendableCoins(ctx, moduleAddr)

		if moduleBalance.AmountOf(denom.Base).IsZero() {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "WithdrawReward not available now"), nil, nil
		}

		acc := testutil.RandPick(r, accs)
		msg = types.NewMsgWithdrawReward(acc.Address)
		err := SimBroadcastTransaction(r, app, msg, ak, bk, ctx, chainID, []cryptotypes.PrivKey{acc.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "WithdrawReward user insufficient fund"), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "Success", nil), nil, nil
	}
}
