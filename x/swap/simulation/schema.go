package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/sharering/shareledger/testutil"
	elecSim "github.com/sharering/shareledger/x/electoral/simulation"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/thanhpk/randstr"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "cosmossdk.io/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateFormat(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount := elecSim.GetElectoralAddress(r, "authority")
		msg := &types.MsgCreateSchema{
			Creator: simAccount.Address.String(),
			Network: randstr.String(4),
		}
		shrRandIn := rand.Int63n(10000000000000-1000000000) + 10000000000000
		amountIn := sdk.NewDecCoin(denom.Base, sdk.NewInt(shrRandIn))

		shrRandOut := rand.Int63n(10000000000000-1000000000) + 10000000000000
		amountOut := sdk.NewDecCoin(denom.Base, sdk.NewInt(shrRandOut))

		randSchemaS := apitypes.TypedData{
			Domain: apitypes.TypedDataDomain{
				VerifyingContract: fmt.Sprintf("0x%s", randstr.Hex(40)),
			},
		}

		bz, err := json.Marshal(randSchemaS)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "fail make random schemas"), nil, nil
		}
		msg.Schema = string(bz)
		msg.In = amountIn
		msg.Out = amountOut

		_, found := k.GetSchema(ctx, msg.Network)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Format already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateFormat(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			format     = types.Schema{}
			msg        = &types.MsgUpdateSchema{}
			allFormat  = k.GetAllSchema(ctx)
		)
		simAccount = elecSim.GetElectoralAddress(r, "authority")

		if len(allFormat) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no no schema for update"), nil, nil
		}
		schema := testutil.RandPick(r, allFormat)
		msg.Schema = schema.Schema
		msg.In = testutil.PtrOf(sdk.NewDecCoinFromCoin(*schema.Fee.In))
		msg.Out = testutil.PtrOf(sdk.NewDecCoinFromCoin(*schema.Fee.Out))
		msg.Creator = simAccount.Address.String()
		msg.Network = format.Network
		err := makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

func SimulateMsgDeleteFormat(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			msg        = &types.MsgDeleteSchema{}
			allFormat  = k.GetAllSchema(ctx)
		)
		if len(allFormat) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no no schema for delete"), nil, nil
		}
		c := types.Schema{}
		for {
			c := testutil.RandPick(r, allFormat)
			if c.Network != "bsc" && c.Schema != "eth" {
				break
			}
		}

		simAccount = elecSim.GetElectoralAddress(r, "authority")

		msg.Creator = simAccount.Address.String()
		msg.Network = c.GetNetwork()

		err := makeTransaction(r, app, msg, ak, bk, k, ctx, chainID, []cryptotypes.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), ""), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
