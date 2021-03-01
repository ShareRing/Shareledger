package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeEnrollIDSignersMsg = "enroll_id_signers"
	TypeRevokeIDSignersMsg = "revoke_id_signers"
)

type MsgEnrollIDSigners struct {
	Approver  sdk.AccAddress   `json:"approver"`
	IDSigners []sdk.AccAddress `json:"id_signers"`
}

func NewMsgEnrollIDSigners(approver sdk.AccAddress, signers []sdk.AccAddress) MsgEnrollIDSigners {
	return MsgEnrollIDSigners{
		Approver:  approver,
		IDSigners: signers,
	}
}

func (msg MsgEnrollIDSigners) Route() string {
	return RouterKey
}

func (msg MsgEnrollIDSigners) Type() string {
	return TypeEnrollIDSignersMsg
}

func (msg MsgEnrollIDSigners) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if len(msg.IDSigners) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The signers list are empty")
	}
	for _, addr := range msg.IDSigners {
		if addr.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Signer address empty")
		}
	}
	return nil
}

func (msg MsgEnrollIDSigners) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgEnrollIDSigners) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

func (msg MsgEnrollIDSigners) String() string {
	rs := fmt.Sprintf("MsgEnrollIDSigners(%s:%v)", msg.Approver, msg.IDSigners)
	return rs
}

type MsgRevokeIDSigners struct {
	Approver  sdk.AccAddress   `json:"approver"`
	IDSigners []sdk.AccAddress `json:"id_signers"`
}

func NewMsgRevokeIDSigners(approver sdk.AccAddress, signers []sdk.AccAddress) MsgRevokeIDSigners {
	return MsgRevokeIDSigners{
		Approver:  approver,
		IDSigners: signers,
	}
}

func (msg MsgRevokeIDSigners) Route() string {
	return RouterKey
}

func (msg MsgRevokeIDSigners) Type() string {
	return TypeRevokeIDSignersMsg
}

func (msg MsgRevokeIDSigners) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if len(msg.IDSigners) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The signers list are empty")
	}
	for _, addr := range msg.IDSigners {
		if addr.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Signer address empty")
		}
	}
	return nil
}

func (msg MsgRevokeIDSigners) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgRevokeIDSigners) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

func (msg MsgRevokeIDSigners) String() string {
	return fmt.Sprintf("MsgRevokeIDSigners(%s, %v)", msg.Approver.String(), msg.IDSigners)
}
