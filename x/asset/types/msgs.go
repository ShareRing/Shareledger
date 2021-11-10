package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeAssetCreateMsg = "asset_create"
	TypeAssetUpdateMsg = "asset_update"
	TypeAssetDeleteMsg = "asset_delete"
)

func NewMsgCreate(creator string, hash []byte, uuid string, status bool, rate int64) MsgCreate {
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if len(msg.Creator) == 0 || err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Address of creator must not be empty")
	}
	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	return nil
}

func (msg MsgCreate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreate) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{signer}
}

func NewMsgUpdate(creator string, hash []byte, uuid string, status bool, rate int64) MsgUpdate {
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if len(msg.Creator) == 0 || err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Address of creator must not be empty")
	}
	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	return nil
}

func (msg MsgUpdate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdate) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{signer}
}

func NewMsgDelete(owner, uuid string) MsgDelete {
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
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if len(msg.Owner) == 0 || err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Owner is invalid")
	}
	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	return nil
}

func (msg MsgDelete) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDelete) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{signer}
}
