package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// TODO emit event
func (k msgServer) EnrollVoter(goCtx context.Context, msg *types.MsgEnrollVoter) (*types.MsgEnrollVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAuthority(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not authority")
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	voterKey := types.VoterPrefix + msg.Voter
	k.SetVoterAddress(ctx, voterKey, msg.Voter)
	k.SetVoterStatus(ctx, voterKey, types.StatusVoterEnrolled)

	return &types.MsgEnrollVoterResponse{}, nil
}

// TODO emit event
func (k msgServer) RevokeVoter(goCtx context.Context, msg *types.MsgRevokeVoter) (*types.MsgRevokeVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAuthority(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not authority")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	voterKey := types.VoterPrefix + msg.Voter

	if !k.IsVoterPresent(ctx, voterKey) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Voter is not enrolled")
	}

	k.DeleteVoter(ctx, voterKey)
	// log := fmt.Sprintf("Successfully delete voter %s", msg.Voter)
	// return &sdk.Result{
	// 	Log: log,
	// }, nil

	return &types.MsgRevokeVoterResponse{}, nil

}

func IsEnrolledVoter(ctx sdk.Context, address sdk.AccAddress, k Keeper) bool {
	addr := types.VoterPrefix + address.String()
	status := k.GetVoterStatus(ctx, addr)
	return status == types.StatusVoterEnrolled
}
