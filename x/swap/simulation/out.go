package simulation

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/utils/denom"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
)

func SimulateMsgOut(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	gk types.GentlemintKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		shrRand := rand.Int63n(100000000000-100000000) + 100000000000
		network := testutil.RandNetwork(r)

		amount := sdk.NewDecCoinFromCoin(sdk.NewCoin(denom.Base, sdk.NewInt(shrRand)))
		msg := &types.MsgRequestOut{
			Creator:     simAccount.Address.String(),
			SrcAddress:  simAccount.Address.String(),
			DestAddress: testutil.RandEthAddress(),
			Network:     network,
			Amount:      &amount,
		}
		shrRand = rand.Int63n(10000000000000-100000000) + 10000000000000
		lC := sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(shrRand)))

		err := gk.LoadCoins(ctx, sdk.MustAccAddressFromBech32(simAccount.Address.String()), lC)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		err = makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
