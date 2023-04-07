package simulation

import (
	types2 "github.com/cosmos/cosmos-sdk/crypto/types"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
)

func SimulateEnrollAccountOperator(
	_ keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgEnrollAccountOperators{Creator: creator.String(), Addresses: []string{a.Address.String()}}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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

func SimulateEnrollApprover(
	_ keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgEnrollApprovers{Creator: creator.String(), Addresses: []string{a.Address.String()}}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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

func SimulateEnrollDocIssuer(
	_ keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgEnrollDocIssuers{Creator: creator.String(), Addresses: []string{a.Address.String()}}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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
func SimulateEnrollIdSigner(
	_ keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgEnrollIdSigners{Creator: creator.String(), Addresses: []string{a.Address.String()}}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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

func SimulateEnrollLoader(
	_ keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgEnrollLoaders{Creator: creator.String(), Addresses: []string{a.Address.String()}}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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

func SimulateEnrollRelayer(
	_ keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		creator := GetElectoralAddress(r, "authority")

		a := testutil.RandPick(r, accs)

		msg := &types.MsgEnrollRelayers{Creator: creator.String(), Addresses: []string{a.Address.String()}}
		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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

func SimulateEnrollSwapManager(
	_ keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgEnrollSwapManagers{Creator: creator.String(), Addresses: []string{a.Address.String()}}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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

func SimulateEnrollVoter(
	_ keeper.Keeper,
	gk keeper.GentlemintKeeper,
	ak types.AccountKeeper,
	_ types.BankKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator := GetElectoralAddress(r, "authority")
		a := testutil.RandPick(r, accs)
		msg := &types.MsgEnrollVoter{Creator: creator.String(), Address: a.Address.String()}

		err := gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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
