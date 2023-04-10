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

func SimulateMsgApproveIn(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount := elecSim.GetElectoralAddress(r, "approver")
		msg := &types.MsgApproveIn{
			Creator: simAccount.Address.String(),
		}
		reqIn, err := k.Swap(ctx, &types.QuerySwapRequest{
			Status:     types.SwapStatusPending,
			SrcNetwork: testutil.RandNetwork(r),
		})
		if err != nil || len(reqIn.GetSwaps()) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no pending request out found "), nil, nil
		}

		rqs := testutil.RandPick(r, reqIn.GetSwaps())

		msg.Ids = []uint64{rqs.Id}

		err = makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
