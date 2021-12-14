package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeEnrollVoterMsg = "enroll_voter"
	TypeRevokeVoterMsg = "revoke_voter"
)

func NewMsgEnrollVoter(approver, voter string) MsgEnrollVoter {
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
	_, err := sdk.AccAddressFromBech32(msg.Approver)
	if len(msg.Approver) == 0 || err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid Approver %s", msg.Approver))
	}

	_, err = sdk.AccAddressFromBech32(msg.Voter)
	if len(msg.Approver) == 0 || err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid Voter %s", msg.Voter))
	}
	return nil
}

func (msg MsgEnrollVoter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgEnrollVoter) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Approver)
	return []sdk.AccAddress{signer}
}

func NewMsgRevokeVoter(approver, voter string) MsgRevokeVoter {
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
	_, err := sdk.AccAddressFromBech32(msg.Approver)
	if len(msg.Approver) == 0 || err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid Approver %s", msg.Approver))
	}

	_, err = sdk.AccAddressFromBech32(msg.Voter)
	if len(msg.Approver) == 0 || err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid Voter %s", msg.Voter))
	}
	return nil
}

func (msg MsgRevokeVoter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRevokeVoter) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Approver)
	return []sdk.AccAddress{signer}
}
