package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeAssetCreateMsg = "asset_create"
	TypeAssetUpdateMsg = "asset_update"
	TypeAssetDeleteMsg = "asset_delete"
)

type MsgCreate struct {
	Creator sdk.AccAddress `json:"creator"`
	Hash    []byte         `json:"hash"`
	UUID    string         `json:"uuid"`
	Status  bool           `json:"status"`
	Rate    int64          `json:"rate"`
}

func NewMsgCreate(creator sdk.AccAddress, hash []byte, uuid string, status bool, rate int64) MsgCreate {
	return MsgCreate{
		Creator: creator,
		Hash:    hash,
		UUID:    uuid,
		Rate:    rate,
		Status:  status,
	}
}

func (msg MsgCreate) Route() string {
	return RouterKey
}

// Type Implements Msg
func (msg MsgCreate) Type() string {
	return TypeAssetCreateMsg
}

// ValidateBasic Implements Msg
func (msg MsgCreate) ValidateBasic() error {
	if len(msg.Creator) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Address of creator must not be empty")
	}
	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	return nil
}

func (msg MsgCreate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreate) String() string {
	return fmt.Sprintf("Asset/MsgCreation{%s}", msg.UUID)
}

func (msg MsgCreate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

type MsgUpdate struct {
	Creator sdk.AccAddress `json:"creator"`
	Hash    []byte         `json:"hash"`
	UUID    string         `json:"uuid"`
	Status  bool           `json:"status"`
	Rate    int64          `json:"rate"`
}

func NewMsgUpdate(creator sdk.AccAddress, hash []byte, uuid string, status bool, rate int64) MsgUpdate {
	return MsgUpdate{
		Creator: creator,
		Hash:    hash,
		UUID:    uuid,
		Rate:    rate,
		Status:  status,
	}
}

func (msg MsgUpdate) Route() string {
	return RouterKey
}

func (msg MsgUpdate) Type() string {
	return TypeAssetUpdateMsg
}

func (msg MsgUpdate) ValidateBasic() error {
	if len(msg.Creator) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Address of creator must not be empty")
	}
	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	return nil
}

func (msg MsgUpdate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

type MsgDelete struct {
	Owner sdk.AccAddress `json:"owner"`
	UUID  string         `json:"uuid"`
}

func NewMsgDelete(owner sdk.AccAddress, uuid string) MsgDelete {
	return MsgDelete{
		Owner: owner,
		UUID:  uuid,
	}
}

func (msg MsgDelete) Route() string {
	return RouterKey
}

func (msg MsgDelete) Type() string {
	return TypeAssetDeleteMsg
}

func (msg MsgDelete) ValidateBasic() error {
	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	return nil
}

func (msg MsgDelete) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDelete) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
