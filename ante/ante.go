package ante

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibcante "github.com/cosmos/ibc-go/v5/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	globalfeeante "github.com/sharering/shareledger/x/gentlemint/ante"
)

type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper         *ibckeeper.Keeper
	WasmConfig        *wasmTypes.WasmConfig
	TXCounterStoreKey storetypes.StoreKey

	// additional data
	GentlemintKeeper    GentlemintKeeper
	RoleKeeper          RoleKeeper
	IdKeeper            IDKeeper
	DistributionxKeeper DistributionxKeeper
	WasmKeeper          WasmKeeper

	BypassMinFeeMsgTypes []string
	GlobalFeeSubspace    paramtypes.Subspace
	StakingSubspace      paramtypes.Subspace
}

func NewHandler(opts HandlerOptions) (sdk.AnteHandler, error) {
	if opts.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	}

	if opts.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	}

	if opts.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	sigGasConsumer := opts.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(),
		wasmkeeper.NewLimitSimulationGasDecorator(opts.WasmConfig.SimulationGasLimit),
		wasmkeeper.NewCountTXDecorator(opts.TXCounterStoreKey),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(opts.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(opts.AccountKeeper),

		// ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		NewDeductFeeDecorator(opts.AccountKeeper, opts.BankKeeper,
			opts.FeegrantKeeper, opts.TxFeeChecker, opts.DistributionxKeeper, opts.WasmKeeper),

		// there 2 ante that check for transaction fee
		//NewCheckFeeDecorator(opts.GentlemintKeeper),
		globalfeeante.NewFeeDecorator(opts.BypassMinFeeMsgTypes, opts.GlobalFeeSubspace, opts.StakingSubspace),

		NewCountBuilderDecorator(opts.DistributionxKeeper),
		//NewAuthDecorator(opts.RoleKeeper, opts.IdKeeper),
		ante.NewSetPubKeyDecorator(opts.AccountKeeper),
		ante.NewValidateSigCountDecorator(opts.AccountKeeper),
		ante.NewSigGasConsumeDecorator(opts.AccountKeeper, sigGasConsumer),
		ante.NewSigVerificationDecorator(opts.AccountKeeper, opts.SignModeHandler),
		ante.NewIncrementSequenceDecorator(opts.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(opts.IBCKeeper),
		//NewLoadFeeDecorator(opts.GentlemintKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

type RoleKeeperWithoutVoter struct {
	RoleKeeper
}

func (r RoleKeeperWithoutVoter) IsVoter(ctx sdk.Context, address sdk.AccAddress) bool {
	return true
}
