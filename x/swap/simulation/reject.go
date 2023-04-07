package simulation

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/sharering/shareledger/testutil"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
)

func SimulateMsgReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgReject{
			Creator: simAccount.Address.String(),
		}

		res, err := k.Swap(ctx, &types.QuerySwapRequest{
			Status:     types.SwapStatusPending,
			SrcNetwork: types.NetworkNameShareLedger,
		})

		if err != nil || len(res.GetSwaps()) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "pending swap out not available now"), nil, nil
		}

		swapOut := testutil.RandPick(r, res.GetSwaps())
		msg.Ids = []uint64{swapOut.Id}

		err = makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
