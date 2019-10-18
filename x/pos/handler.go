package pos

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/pos/keeper"
	"github.com/sharering/shareledger/x/pos/message"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

var (
	ValidatorChanged = false
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx.WithEventManager(sdk.NewEventManager())
		// NOTE msg already has validate basic run
		switch msg := msg.(type) {
		case message.MsgCreateValidator:
			return handleMsgCreateValidator(ctx, msg, k)
		case message.MsgEditValidator:
			return handleMsgEditValidator(ctx, msg, k)
		case message.MsgDelegate:
			return handleMsgDelegate(ctx, msg, k)
		case message.MsgBeginUnbonding:
			return handleMsgBeginUnbonding(ctx, msg, k)
		case message.MsgCompleteUnbonding:
			return handleMsgCompleteUnbonding(ctx, msg, k)
		case message.MsgWithdraw:
			return handleMsgWithdraw(ctx, msg, k)
		case message.MsgBeginRedelegate:
			return handleMsgBeginRedelegate(ctx, msg, k)

		case message.MsgCompleteRedelegate:
			return handleMsgCompleteRedelegate(ctx, msg, k)

		default:
			return sdk.ErrTxDecode("invalid message parse in staking module").Result()
		}
	}
}

func handleMsgCreateValidator(ctx sdk.Context, msg message.MsgCreateValidator, k keeper.Keeper) sdk.Result {
	// check to see if the pubkey or sender has been registered before
	_, found := k.GetValidator(ctx, msg.ValidatorAddr)
	if found {
		return posTypes.ErrValidatorOwnerExists(k.Codespace()).Result()
	}
	/*
		_, found = k.GetValidatorByConsAddr(ctx, sdk.AccAddress(msg.PubKey.Address()))
		if found {
			return posTypes.ErrValidatorPubKeyExists(k.Codespace()).Result()
		}*/

	if msg.Delegation.Denom != k.GetParams(ctx).BondDenom {
		return posTypes.ErrBadDenom(k.Codespace()).Result()
	}
	if msg.Delegation.Amount.LT(types.NewDec(constants.MIN_MASTER_NODE_TOKEN)) {
		return posTypes.ErrInSufficientMasterNodeToken(k.Codespace()).Result()
	}
	validator := posTypes.NewValidator(msg.ValidatorAddr, msg.PubKey, msg.Description)
	// Update delegator shares = 0
	validator.DelegatorShares = types.ZeroDec()
	/*	commission := NewCommissionWithTime(
			msg.Commission.Rate, msg.Commission.MaxChangeRate,
			msg.Commission.MaxChangeRate, ctx.BlockHeader().Time,
		)
	*/
	// Todo: commission

	k.SetValidator(ctx, validator)
	//k.SetValidatorByConsAddr(ctx, validator)
	k.SetNewValidatorByPowerIndex(ctx, validator)

	vdi := posTypes.NewValidatorDistInfo(validator.Owner, ctx.BlockHeight())

	k.SetValidatorDistInfo(ctx, vdi)

	// move coins from the msg.Address account to a (self-delegation) delegator account
	// the validator account and global shares are updated within here

	_, err := k.Delegate(ctx, msg.DelegatorAddr, msg.Delegation, validator, true)
	if err != nil {
		return err.Result()
	}

	//	k.OnValidatorCreated(ctx, validator.OperatorAddr)
	//accAddr := sdk.AccAddress(validator.OperatorAddr)
	//k.OnDelegationCreated(ctx, accAddr, validator.OperatorAddr)

	// Update ValidatorChanged
	// This variable is reset at the beginning of a block
	ValidatorChanged = true

	event := sdk.NewEvent(
		EventTypeValidatorCreated,
		sdk.NewAttribute(AttributeValidator, msg.ValidatorAddr.String()),
		sdk.NewAttribute(AttributeMoniker, msg.Description.Moniker),
		sdk.NewAttribute(AttributeIdentity, msg.Description.Identity),
	)

	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgDelegate(ctx sdk.Context, msg message.MsgDelegate, k keeper.Keeper) sdk.Result {
	validator, found := k.GetValidator(ctx, msg.ValidatorAddr)
	if !found {
		return posTypes.ErrNoValidatorFound(k.Codespace()).Result()
	}

	if msg.Delegation.Denom != k.GetParams(ctx).BondDenom {
		return posTypes.ErrBadDenom(k.Codespace()).Result()
	}
	/*
		if validator.Jailed && !bytes.Equal(validator.OperatorAddr, msg.DelegatorAddr) {
			return posTypes.ErrValidatorJailed(k.Codespace()).Result()
		}*/

	_, err := k.Delegate(ctx, msg.DelegatorAddr, msg.Delegation, validator, true)
	if err != nil {
		return err.Result()
	}

	// call the hook if present
	//k.OnDelegationCreated(ctx, msg.DelegatorAddr, validator.OperatorAddr)
	event := sdk.NewEvent(
		EventTypeDelegated,
		sdk.NewAttribute(AttributeDelegator, msg.DelegatorAddr.String()),
		sdk.NewAttribute(AttributeValidator, msg.ValidatorAddr.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgBeginUnbonding(ctx sdk.Context, msg message.MsgBeginUnbonding, k keeper.Keeper) sdk.Result {
	err := k.BeginUnbonding(ctx, msg.DelegatorAddr, msg.ValidatorAddr, msg.SharesAmount)
	if err != nil {
		return err.Result()
	}
	event := sdk.NewEvent(
		EventTypeBeginUnbonding,
		sdk.NewAttribute(AttributeDelegator, msg.DelegatorAddr.String()),
		sdk.NewAttribute(AttributeValidator, msg.ValidatorAddr.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}
}
func handleMsgCompleteUnbonding(ctx sdk.Context, msg message.MsgCompleteUnbonding, k keeper.Keeper) sdk.Result {

	err := k.CompleteUnbonding(ctx, msg.DelegatorAddr, msg.ValidatorAddr)
	if err != nil {
		return err.Result()
	}

	event := sdk.NewEvent(
		EventTypeCompleteUnbonding,
		sdk.NewAttribute(AttributeDelegator, msg.DelegatorAddr.String()),
		sdk.NewAttribute(AttributeValidator, msg.ValidatorAddr.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgBeginRedelegate(ctx sdk.Context, msg message.MsgBeginRedelegate, k keeper.Keeper) sdk.Result {
	err := k.BeginRedelegation(ctx, msg.DelegatorAddr, msg.ValidatorSrcAddr,
		msg.ValidatorDstAddr, msg.SharesAmount)
	if err != nil {
		return err.Result()
	}

	event := sdk.NewEvent(
		EventTypeBeginRedelegation,
		sdk.NewAttribute(AttributeDelegator, msg.DelegatorAddr.String()),
		sdk.NewAttribute(AttributeSrcValidator, msg.ValidatorSrcAddr.String()),
		sdk.NewAttribute(AttributeDstValidator, msg.ValidatorDstAddr.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}

}

func handleMsgCompleteRedelegate(ctx sdk.Context, msg message.MsgCompleteRedelegate, k keeper.Keeper) sdk.Result {
	err := k.CompleteRedelegation(ctx, msg.DelegatorAddr, msg.ValidatorSrcAddr, msg.ValidatorDstAddr)
	if err != nil {
		return err.Result()
	}

	event := sdk.NewEvent(
		EventTypeCompleteRedelegation,
		sdk.NewAttribute(AttributeDelegator, msg.DelegatorAddr.String()),
		sdk.NewAttribute(AttributeSrcValidator, msg.ValidatorSrcAddr.String()),
		sdk.NewAttribute(AttributeDstValidator, msg.ValidatorDstAddr.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}

}

func handleMsgWithdraw(
	ctx sdk.Context,
	msg message.MsgWithdraw,
	k keeper.Keeper,
) sdk.Result {

	_, amount, err := k.WithdrawDelReward(ctx, msg.ValidatorAddr, msg.DelegatorAddr)

	// fmt.Printf("validator dist info: %v\n", vdi)

	if err != nil {
		return err.Result()
	}

	event := sdk.NewEvent(
		EventTypeWithdraw,
		sdk.NewAttribute(AttributeDelegator, msg.DelegatorAddr.String()),
		sdk.NewAttribute(AttributeValidator, msg.ValidatorAddr.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Log:    fmt.Sprintf("%s", amount),
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgEditValidator(ctx sdk.Context, msg message.MsgEditValidator, k keeper.Keeper) sdk.Result {
	// validator must already be registered
	validator, found := k.GetValidator(ctx, msg.ValidatorAddr)
	if !found {
		return posTypes.ErrNoValidatorFound(k.Codespace()).Result()
	}

	// replace all editable fields (clients should autofill existing values)
	description, err := validator.Description.UpdateDescription(msg.Description)
	if err != nil {
		return err.Result()
	}

	validator.Description = description

	/*if msg.CommissionRate != nil {
		commission, err := k.UpdateValidatorCommission(ctx, validator, *msg.CommissionRate)
		if err != nil {
			return err.Result()
		}
		validator.Commission = commission
		k.OnValidatorModified(ctx, msg.ValidatorAddr)
	}*/

	k.SetValidator(ctx, validator)

	event := sdk.NewEvent(
		EventTypeEditValidator,
		sdk.NewAttribute(AttributeDstValidator, msg.ValidatorAddr.String()),
		sdk.NewAttribute(AttributeMoniker, description.Moniker),
		sdk.NewAttribute(AttributeIdentity, description.Identity),
	)

	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}
}
