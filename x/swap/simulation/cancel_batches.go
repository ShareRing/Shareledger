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
)

func SimulateMsgCancelBatches(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		//TODO random take address from list of authorized account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCancelBatches{
			Creator: simAccount.Address.String(),
		}
		batches, err := k.Batches(ctx, &types.QueryBatchesRequest{
			Network: testutil.RandNetwork(r),
		})
		if err != nil || len(batches.GetBatches()) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "not valid batch now"), nil, err
		}
		b := testutil.RandPick(r, batches.GetBatches())
		msg.Ids = []uint64{b.GetId()}

		err = makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
