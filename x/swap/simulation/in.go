package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/thanhpk/randstr"

	"github.com/sharering/shareledger/testutil"
	elecSim "github.com/sharering/shareledger/x/electoral/simulation"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

var msg types.MsgRequestIn

func SimulateMsgIn(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount := elecSim.GetElectoralAddress(r, "relayer")
		shrRand := rand.Int63n(10000000000000-1000000000) + 10000000000000
		net := testutil.RandNetwork(r)
		amount := sdk.NewDecCoinFromCoin(sdk.NewCoin(denom.Base, sdk.NewInt(shrRand)))
		msg := &types.MsgRequestIn{
			Creator:     simAccount.Address.String(),
			SrcAddress:  fmt.Sprintf("0x%s", randstr.Hex(40)),
			DestAddress: simAccount.Address.String(),
			Network:     net,
			Amount:      &amount,
			TxEvents:    testutil.RandERC20Event(r),
		}
		err := makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
		//return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no"), nil, nil
	}
}
