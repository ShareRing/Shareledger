package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeEnrollAccOperatorMsg = "enroll_account_operator"
	TypeRevokeAccOperatorMsg = "revoke_account_operator"
)

type MsgEnrollAccOperators struct {
	Approver  sdk.AccAddress   `json:"approver"`
	Operators []sdk.AccAddress `json:"operators"`
}

func NewMsgEnrollAccOperators(approver sdk.AccAddress, operators []sdk.AccAddress) MsgEnrollAccOperators {
	return MsgEnrollAccOperators{
		Approver:  approver,
		Operators: operators,
	}
}

func (msg MsgEnrollAccOperators) Route() string {
	return RouterKey
}

func (msg MsgEnrollAccOperators) Type() string {
	return TypeEnrollAccOperatorMsg
}

func (msg MsgEnrollAccOperators) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if len(msg.Operators) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The account list are empty")
	}
	for _, addr := range msg.Operators {
		if addr.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Account address empty")
		}
	}
	return nil
}

func (msg MsgEnrollAccOperators) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgEnrollAccOperators) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

func (msg MsgEnrollAccOperators) String() string {
	rs := fmt.Sprintf("MsgEnrollAccOperators(%s:%v)", msg.Approver, msg.Operators)
	return rs
}

type MsgRevokeAccOperators struct {
	Approver  sdk.AccAddress   `json:"approver"`
	Operators []sdk.AccAddress `json:"operators"`
}

func NewMsgRevokeAccOperators(approver sdk.AccAddress, accs []sdk.AccAddress) MsgRevokeAccOperators {
	return MsgRevokeAccOperators{
		Approver:  approver,
		Operators: accs,
	}
}

func (msg MsgRevokeAccOperators) Route() string {
	return RouterKey
}

func (msg MsgRevokeAccOperators) Type() string {
	return TypeRevokeAccOperatorMsg
}

func (msg MsgRevokeAccOperators) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if len(msg.Operators) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The account list are empty")
	}

	for _, addr := range msg.Operators {
		if addr.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Account address empty")
		}
	}
	return nil
}

func (msg MsgRevokeAccOperators) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgRevokeAccOperators) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

func (msg MsgRevokeAccOperators) String() string {
	return fmt.Sprintf("MsgRevokeAccOperators(%s, %v)", msg.Approver.String(), msg.Operators)
}
