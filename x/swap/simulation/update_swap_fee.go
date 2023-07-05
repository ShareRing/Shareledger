package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/sharering/shareledger/testutil"
	elecSim "github.com/sharering/shareledger/x/electoral/simulation"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func SimulateMsgUpdateSwapFee(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount := elecSim.GetElectoralAddress(r, "swapManager")
		msg := &types.MsgUpdateSwapFee{
			Creator: simAccount.Address.String(),
			In:      &sdk.DecCoin{},
			Out:     &sdk.DecCoin{},
		}

		res, err := k.Schemas(ctx, &types.QuerySchemasRequest{})
		if err != nil || len(res.GetSchemas()) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no schema here"), nil, nil
		}
		s := testutil.RandPick(r, res.GetSchemas())

		msg.Network = s.GetNetwork()
		shrRandIn := rand.Int63n(10000000000000-1000000000) + 10000000000000
		amountIn := sdk.NewDecCoin(denom.Base, sdk.NewInt(shrRandIn))

		shrRandOut := rand.Int63n(10000000000000-1000000000) + 10000000000000
		amountOut := sdk.NewDecCoin(denom.Base, sdk.NewInt(shrRandOut))

		*msg.In = amountIn
		*msg.Out = amountOut

		err = makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no transaction"), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
