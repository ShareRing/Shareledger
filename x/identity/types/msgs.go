package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeCreateIdMsg        = "create_id"
	TypeUpdateIdMsg        = "update_id"
	TypeDeleteIdMsg        = "delete_id"
	TypeEnrollIDSignersMsg = "enroll_id_signers"
	TypeRevokeIDSignersMsg = "revoke_id_signers"
)

type MsgCreateId struct {
	Approver sdk.AccAddress `json:"approver"`
	Owner    sdk.AccAddress `json:"owner"`
	Hash     string         `json:"hash"`
}

func NewMsgCreateId(approver, owner sdk.AccAddress, hash string) MsgCreateId {
	return MsgCreateId{
		Approver: approver,
		Owner:    owner,
		Hash:     hash,
	}
}

func (msg MsgCreateId) Route() string {
	return RouterKey
}

func (msg MsgCreateId) Type() string {
	return TypeCreateIdMsg
}

func (msg MsgCreateId) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Approver address must not be empty")
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Owner address must not be empty")
	}
	return nil
}

func (msg MsgCreateId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateId) String() string {
	return fmt.Sprintf("Identity/MsgCreateId{Address:%s,Hash:%s}", msg.Owner.String(), msg.Hash)
}

func (msg MsgCreateId) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgUpdateId struct {
	Approver sdk.AccAddress `json:"approver"`
	Owner    sdk.AccAddress `json:"owner"`
	Hash     string         `json:"hash"`
}

func NewMsgUpdateId(approver, owner sdk.AccAddress, hash string) MsgUpdateId {
	return MsgUpdateId{
		Approver: approver,
		Owner:    owner,
		Hash:     hash,
	}
}

func (msg MsgUpdateId) Route() string {
	return RouterKey
}

func (msg MsgUpdateId) Type() string {
	return TypeUpdateIdMsg
}

func (msg MsgUpdateId) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Approver address must not be empty")
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Owner address must not be empty")
	}
	return nil
}

func (msg MsgUpdateId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgUpdateId) String() string {
	return fmt.Sprintf("Identity/MsgUpdateId{Address:%s,Hash:%s}", msg.Owner.String(), msg.Hash)
}

func (msg MsgUpdateId) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgDeleteId struct {
	Approver sdk.AccAddress `json:"approver"`
	Owner    sdk.AccAddress `json:"owner"`
}

func NewMsgDeleteId(approver, owner sdk.AccAddress) MsgDeleteId {
	return MsgDeleteId{
		Approver: approver,
		Owner:    owner,
	}
}

func (msg MsgDeleteId) Route() string {
	return RouterKey
}

func (msg MsgDeleteId) Type() string {
	return TypeDeleteIdMsg
}

func (msg MsgDeleteId) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Approver address must not be empty")
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Owner address must not be empty")
	}
	return nil
}

func (msg MsgDeleteId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDeleteId) String() string {
	return fmt.Sprintf("Identity/MsgDeleteId{Address:%s}", msg.Owner.String())
}

func (msg MsgDeleteId) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgEnrollIDSigners struct {
	Approver  sdk.AccAddress   `json:"approver"`
	IDSigners []sdk.AccAddress `json:"receiver"`
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

type MsgRevokeIDSigners struct {
	Approver  sdk.AccAddress   `json:"approver"`
	IDSigners []sdk.AccAddress `json:"receiver"`
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
