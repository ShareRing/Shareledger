package pos

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/pos/keeper"
	"github.com/sharering/shareledger/x/pos/message"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		// NOTE msg already has validate basic run
		switch msg := msg.(type) {
		case message.MsgCreateValidator:
			return handleMsgCreateValidator(ctx, msg, k)

		default:
			return sdk.ErrTxDecode("invalid message parse in staking module").Result()
		}
	}
}

func handleMsgCreateValidator(ctx sdk.Context, msg message.MsgCreateValidator, k keeper.Keeper) sdk.Result {
	// check to see if the pubkey or sender has been registered before
	_, found := k.GetValidator(ctx, msg.ValidatorAddr)
	if found {
		return sdk.Result{} //return posTypes.ErrValidatorOwnerExists(k.codespace().Result()
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
	//todo: delegate
	/*
		_, err := k.Delegate(ctx, msg.DelegatorAddr, msg.Delegation, validator, true)
		if err != nil {
			return err.Result()
		}
	*/

	//	k.OnValidatorCreated(ctx, validator.OperatorAddr)
	//accAddr := sdk.AccAddress(validator.OperatorAddr)
	//k.OnDelegationCreated(ctx, accAddr, validator.OperatorAddr)

	tags := sdk.NewTags(
	/* 	tags.Action, tags.ActionCreateValidator,
	tags.DstValidator, []byte(msg.ValidatorAddr.String()),
	tags.Moniker, []byte(msg.Description.Moniker),
	tags.Identity, []byte(msg.Description.Identity), */
	)

	return sdk.Result{
		Tags: tags,
	}
}
