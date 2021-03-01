package electoral

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/electoral/types"
)

const (
	VoterPrefix = "voter"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgEnrollVoter:
			return handleMsgEnrollVoter(ctx, keeper, msg)
		case MsgRevokeVoter:
			return handleMsgRevokeVoter(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized booking Msg type: %v", msg.Type()))
		}
	}
}

func handleMsgEnrollVoter(ctx sdk.Context, keeper Keeper, msg MsgEnrollVoter) (*sdk.Result, error) {
	if !keeper.IsAuthority(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not authority")
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	voterKey := VoterPrefix + msg.Voter.String()
	keeper.SetVoterAddress(ctx, voterKey, msg.Voter)
	keeper.SetVoterStatus(ctx, voterKey, types.StatusVoterEnrolled)
	log := fmt.Sprintf("Successfully enroll voter % s", msg.Voter.String())
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgRevokeVoter(ctx sdk.Context, keeper Keeper, msg MsgRevokeVoter) (*sdk.Result, error) {
	if !keeper.IsAuthority(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not authority")
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	voterKey := VoterPrefix + msg.Voter.String()
	if !keeper.IsVoterPresent(ctx, voterKey) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Voter is not enrolled")
	}
	keeper.DeleteVoter(ctx, voterKey)
	log := fmt.Sprintf("Successfully delete voter %s", msg.Voter.String())
	return &sdk.Result{
		Log: log,
	}, nil
}

func IsEnrolledVoter(ctx sdk.Context, address sdk.AccAddress, k Keeper) bool {
	addr := VoterPrefix + address.String()
	status := k.GetVoterStatus(ctx, addr)
	return status == types.StatusVoterEnrolled
}
