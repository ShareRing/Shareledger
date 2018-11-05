package pos

import (
	"fmt"
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/pos/keeper"
	"github.com/sharering/shareledger/x/pos/message"
	"github.com/sharering/shareledger/x/pos/tags"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		// NOTE msg already has validate basic run
		switch msg := msg.(type) {
		case message.MsgCreateValidator:
			return handleMsgCreateValidator(ctx, msg, k)
		case message.MsgDelegate:
			return handleMsgDelegate(ctx, msg, k)
		case message.MsgBeginUnbonding:
			return handleMsgBeginUnbonding(ctx, msg, k)
		case message.MsgCompleteUnbonding:
			return handleMsgCompleteUnbonding(ctx, msg, k)
		case message.MsgWithdraw:
			return handleMsgWithdraw(ctx, msg, k)

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
		_, found = k.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(msg.PubKey))
		if found {
			return ErrValidatorPubKeyExists(k.Codespace()).Result()
		}

		if msg.Delegation.Denom != k.GetParams(ctx).BondDenom {
			return ErrBadDenom(k.Codespace()).Result()
		}
	*/
	validator := posTypes.NewValidator(msg.ValidatorAddr, msg.PubKey, msg.Description)
	/*	commission := NewCommissionWithTime(
			msg.Commission.Rate, msg.Commission.MaxChangeRate,
			msg.Commission.MaxChangeRate, ctx.BlockHeader().Time,
		)
	*/
	// Todo: commission

	k.SetValidator(ctx, validator)
	//k.SetValidatorByConsAddr(ctx, validator)
	//k.SetNewValidatorByPowerIndex(ctx, validator)

	// move coins from the msg.Address account to a (self-delegation) delegator account
	// the validator account and global shares are updated within here

	_, err := k.Delegate(ctx, msg.DelegatorAddr, msg.Delegation, validator, true)
	if err != nil {
		return err.Result()
	}

	//	k.OnValidatorCreated(ctx, validator.OperatorAddr)
	//accAddr := sdk.AccAddress(validator.OperatorAddr)
	//k.OnDelegationCreated(ctx, accAddr, validator.OperatorAddr)

	tags := sdk.NewTags(
		tags.Event, tags.ValidatorCreated,
		tags.Validator, []byte(msg.ValidatorAddr.String()),
		tags.Moniker, []byte(msg.Description.Moniker),
		tags.Identity, []byte(msg.Description.Identity),
	)

	return sdk.Result{
		Tags: tags,
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

	tags := sdk.NewTags(
		tags.Event, tags.Delegated,
		tags.Delegator, []byte(msg.DelegatorAddr.String()),
		tags.Validator, []byte(msg.ValidatorAddr.String()),
	)

	return sdk.Result{
		Tags: tags,
	}
}

func handleMsgBeginUnbonding(ctx sdk.Context, msg message.MsgBeginUnbonding, k keeper.Keeper) sdk.Result {
	err := k.BeginUnbonding(ctx, msg.DelegatorAddr, msg.ValidatorAddr, msg.SharesAmount)
	if err != nil {
		return err.Result()
	}

	tags := sdk.NewTags(
		tags.Event, tags.BeginUnbonding,
		tags.Delegator, []byte(msg.DelegatorAddr.String()),
		tags.Validator, []byte(msg.ValidatorAddr.String()),
	)
	return sdk.Result{Tags: tags}
}

func handleMsgCompleteUnbonding(ctx sdk.Context, msg message.MsgCompleteUnbonding, k keeper.Keeper) sdk.Result {

	err := k.CompleteUnbonding(ctx, msg.DelegatorAddr, msg.ValidatorAddr)
	if err != nil {
		return err.Result()
	}

	tags := sdk.NewTags(
		tags.Event, tags.CompleteUnbonding,
		tags.Delegator, []byte(msg.DelegatorAddr.String()),
		tags.Validator, []byte(msg.ValidatorAddr.String()),
	)

	return sdk.Result{Tags: tags}
}

func handleMsgWithdraw(
	ctx sdk.Context,
	msg message.MsgWithdraw,
	k keeper.Keeper,
) sdk.Result {

	vdi, amount, err := k.WithdrawDelReward(ctx, msg.ValidatorAddr, msg.DelegatorAddr)

	fmt.Printf("validator dist info: %v\n", vdi)

	if err != nil {
		return err.Result()
	}

	tags := sdk.NewTags(
		tags.Event, tags.CompleteUnbonding,
		tags.Delegator, []byte(msg.DelegatorAddr.String()),
		tags.Validator, []byte(msg.ValidatorAddr.String()),
	)

	return sdk.Result{
		Tags: tags,
		Log: fmt.Sprintf("%s", amount),
	}
}
