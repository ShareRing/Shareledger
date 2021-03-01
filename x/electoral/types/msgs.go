package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeEnrollVoterMsg = "enroll_voter"
	TypeRevokeVoterMsg = "revoke_voter"
)

type MsgEnrollVoter struct {
	Approver sdk.AccAddress `json:"approver"`
	Voter    sdk.AccAddress `json:"voters"`
}

func NewMsgEnrollVoter(approver sdk.AccAddress, voter sdk.AccAddress) MsgEnrollVoter {
	return MsgEnrollVoter{
		Approver: approver,
		Voter:    voter,
	}
}

func (msg MsgEnrollVoter) Route() string {
	return RouterKey
}

func (msg MsgEnrollVoter) Type() string {
	return TypeEnrollVoterMsg
}

func (msg MsgEnrollVoter) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if msg.Voter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	return nil
}

func (msg MsgEnrollVoter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgEnrollVoter) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgRevokeVoter struct {
	Approver sdk.AccAddress `json:"approver"`
	Voter    sdk.AccAddress `json:"voters"`
}

func NewMsgRevokeVoter(approver sdk.AccAddress, voter sdk.AccAddress) MsgRevokeVoter {
	return MsgRevokeVoter{
		Approver: approver,
		Voter:    voter,
	}
}

func (msg MsgRevokeVoter) Route() string {
	return RouterKey
}

func (msg MsgRevokeVoter) Type() string {
	return TypeRevokeVoterMsg
}

func (msg MsgRevokeVoter) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if msg.Voter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	return nil
}

func (msg MsgRevokeVoter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgRevokeVoter) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}
