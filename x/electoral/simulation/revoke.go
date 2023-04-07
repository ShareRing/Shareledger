package simulation

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	types2 "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
	"math/rand"
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
		msg := &types.MsgRevokeAccountOperators{Creator: creator.String()}

		operator, err := k.AccountOperators(ctx, &types.QueryAccountOperatorsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		msg.Addresses = []string{operator.AccStates[0].Address}

		err = gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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
		msg := &types.MsgRevokeApprovers{Creator: creator.String()}

		approver, err := k.Approvers(ctx, &types.QueryApproversRequest{})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		msg.Addresses = []string{approver.Approvers[0].Address}

		err = gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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
		msg := &types.MsgRevokeDocIssuers{Creator: creator.String()}

		docIssuer, err := k.DocumentIssuers(ctx, &types.QueryDocumentIssuersRequest{})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		msg.Addresses = []string{docIssuer.AccStates[0].Address}

		err = gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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
		msg := &types.MsgRevokeIdSigners{Creator: creator.String()}

		idSigner, err := k.IdSigners(ctx, &types.QueryIdSignersRequest{})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		msg.Addresses = []string{idSigner.AccStates[0].Address}

		err = gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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
		msg := &types.MsgRevokeLoaders{Creator: creator.String()}

		loader, err := k.Loaders(ctx, &types.QueryLoadersRequest{})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		msg.Addresses = []string{loader.Loaders[0].Address}

		err = gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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
		msg := &types.MsgRevokeRelayers{Creator: creator.String()}

		relayers, err := k.Relayers(ctx, &types.QueryRelayersRequest{})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		msg.Addresses = []string{relayers.Relayers[0].Address}

		err = gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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
		msg := &types.MsgRevokeRelayers{Creator: creator.String()}

		swapManager, err := k.SwapManagers(ctx, &types.QuerySwapManagersRequest{})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		msg.Addresses = []string{swapManager.SwapManagers[0].Address}

		err = gk.LoadAllowanceLoader(ctx, sdk.MustAccAddressFromBech32(creator.String()))
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
