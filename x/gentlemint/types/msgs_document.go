package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeEnrollDocIssuerMsg = "enroll_doc_issuer"
	TypeRevokeDocIssuerMsg = "revoke_doc_issuer"
)

type MsgEnrollDocIssuers struct {
	Approver sdk.AccAddress   `json:"approver"`
	Issuers  []sdk.AccAddress `json:"issuers"`
}

func NewMsgEnrollDocIssuers(approver sdk.AccAddress, issuers []sdk.AccAddress) MsgEnrollDocIssuers {
	return MsgEnrollDocIssuers{
		Approver: approver,
		Issuers:  issuers,
	}
}

func (msg MsgEnrollDocIssuers) Route() string {
	return RouterKey
}

func (msg MsgEnrollDocIssuers) Type() string {
	return TypeEnrollDocIssuerMsg
}

func (msg MsgEnrollDocIssuers) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if len(msg.Issuers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The issuer list are empty")
	}
	for _, addr := range msg.Issuers {
		if addr.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Issuer address empty")
		}
	}
	return nil
}

func (msg MsgEnrollDocIssuers) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgEnrollDocIssuers) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

func (msg MsgEnrollDocIssuers) String() string {
	rs := fmt.Sprintf("MsgEnrollDocIssuers(%s:%v)", msg.Approver, msg.Issuers)
	return rs
}

type MsgRevokeDocIssuers struct {
	Approver sdk.AccAddress   `json:"approver"`
	Issuers  []sdk.AccAddress `json:"issuers"`
}

func NewMsgRevokeDocIssuers(approver sdk.AccAddress, signers []sdk.AccAddress) MsgRevokeDocIssuers {
	return MsgRevokeDocIssuers{
		Approver: approver,
		Issuers:  signers,
	}
}

func (msg MsgRevokeDocIssuers) Route() string {
	return RouterKey
}

func (msg MsgRevokeDocIssuers) Type() string {
	return TypeRevokeDocIssuerMsg
}

func (msg MsgRevokeDocIssuers) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if len(msg.Issuers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The signers list are empty")
	}
	for _, addr := range msg.Issuers {
		if addr.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Signer address empty")
		}
	}
	return nil
}

func (msg MsgRevokeDocIssuers) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgRevokeDocIssuers) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

func (msg MsgRevokeDocIssuers) String() string {
	return fmt.Sprintf("MsgRevokeDocIssuers(%s, %v)", msg.Approver.String(), msg.Issuers)
}
