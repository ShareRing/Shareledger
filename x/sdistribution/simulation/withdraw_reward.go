package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/x/sdistribution/keeper"
	"github.com/sharering/shareledger/x/sdistribution/types"
)

func SimulateMsgWithdrawReward(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgWithdrawReward{}

		// TODO: Handling the WithdrawReward simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "WithdrawReward simulation not implemented"), nil, nil
	}
}
