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

func SimulateMsgReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount := elecSim.GetElectoralAddress(r, "approver")
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
