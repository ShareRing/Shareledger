package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	types2 "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/thanhpk/randstr"

	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
)

func SimulateRevokeAccountOperator(
	k keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgRevokeAccountOperators{Creator: creator.Address.String()}

		acc := fmt.Sprintf("shareledger%s", randstr.Base62(39))
		k.ActiveAccState(ctx, []byte(acc), types.AccStateKeyAccOp)

		msg.Addresses = []string{acc}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.Address.String()))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		err = makeTransaction(r, app, msg, ak, ctx, chainID, []types2.PrivKey{a.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

func SimulateRevokeApprover(
	k keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgRevokeApprovers{Creator: creator.Address.String()}

		acc := fmt.Sprintf("shareledger%s", randstr.Base62(39))
		k.ActiveAccState(ctx, []byte(acc), types.AccStateKeyApprover)

		msg.Addresses = []string{acc}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.Address.String()))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		err = makeTransaction(r, app, msg, ak, ctx, chainID, []types2.PrivKey{a.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

func SimulateRevokeDocIssuer(
	k keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgRevokeDocIssuers{Creator: creator.Address.String()}

		acc := fmt.Sprintf("shareledger%s", randstr.Base62(39))
		k.ActiveAccState(ctx, []byte(acc), types.AccStateKeyDocIssuer)

		msg.Addresses = []string{acc}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.Address.String()))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		err = makeTransaction(r, app, msg, ak, ctx, chainID, []types2.PrivKey{a.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

func SimulateRevokeIdSigner(
	k keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgRevokeIdSigners{Creator: creator.Address.String()}

		acc := fmt.Sprintf("shareledger%s", randstr.Base62(39))
		k.ActiveAccState(ctx, []byte(acc), types.AccStateKeyIdSigner)

		msg.Addresses = []string{acc}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.Address.String()))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		err = makeTransaction(r, app, msg, ak, ctx, chainID, []types2.PrivKey{a.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

func SimulateRevokeLoader(
	k keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgRevokeLoaders{Creator: creator.Address.String()}
		acc := fmt.Sprintf("shareledger%s", randstr.Base62(39))
		k.ActiveAccState(ctx, []byte(acc), types.AccStateKeyShrpLoaders)

		msg.Addresses = []string{acc}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.Address.String()))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		err = makeTransaction(r, app, msg, ak, ctx, chainID, []types2.PrivKey{a.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

func SimulateRevokeRelayer(
	k keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgRevokeRelayers{Creator: creator.Address.String()}

		acc := fmt.Sprintf("shareledger%s", randstr.Base62(39))
		k.ActiveAccState(ctx, []byte(acc), types.AccStateKeyRelayer)

		msg.Addresses = []string{acc}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.Address.String()))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		err = makeTransaction(r, app, msg, ak, ctx, chainID, []types2.PrivKey{a.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
func SimulateRevokeSwapManager(
	k keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgRevokeRelayers{Creator: creator.Address.String()}

		acc := fmt.Sprintf("shareledger%s", randstr.Base62(39))
		k.ActiveAccState(ctx, []byte(acc), types.AccStateKeySwapManager)

		msg.Addresses = []string{acc}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.Address.String()))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		err = makeTransaction(r, app, msg, ak, ctx, chainID, []types2.PrivKey{a.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
