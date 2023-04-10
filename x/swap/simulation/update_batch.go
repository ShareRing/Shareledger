package simulation

import (
	"math/rand"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/sharering/shareledger/testutil"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	elecSim "github.com/sharering/shareledger/x/electoral/simulation"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
)

func SimulateMsgCompleteBatch(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount := elecSim.GetElectoralAddress(r, "relayer")

		msg := &types.MsgCompleteBatch{
			Creator: simAccount.Address.String(),
		}

		res, err := k.Batches(ctx, &types.QueryBatchesRequest{
			Network: testutil.RandNetwork(r),
		})

		if err != nil || len(res.GetBatches()) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "the batch not available now"), nil, nil
		}
		rBatch := testutil.RandPick(r, res.GetBatches())
		msg.BatchId = rBatch.GetId()

		err = makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
