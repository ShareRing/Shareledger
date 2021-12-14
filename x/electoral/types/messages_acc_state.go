package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAccState{}

func NewMsgCreateAccState(
	creator string,
	key string,
	address string,
	status string,

) *MsgCreateAccState {
	return &MsgCreateAccState{
		Creator: creator,
		Key:     key,
		Address: address,
		Status:  status,
	}
}

func (msg *MsgCreateAccState) Route() string {
	return RouterKey
}

func (msg *MsgCreateAccState) Type() string {
	return "CreateAccState"
}

func (msg *MsgCreateAccState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAccState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAccState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAccState{}

func NewMsgUpdateAccState(
	creator string,
	key string,
	address string,
	status string,

) *MsgUpdateAccState {
	return &MsgUpdateAccState{
		Creator: creator,
		Key:     key,
		Address: address,
		Status:  status,
	}
}

func (msg *MsgUpdateAccState) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAccState) Type() string {
	return "UpdateAccState"
}

func (msg *MsgUpdateAccState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAccState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAccState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAccState{}

func NewMsgDeleteAccState(
	creator string,
	key string,

) *MsgDeleteAccState {
	return &MsgDeleteAccState{
		Creator: creator,
		Key:     key,
	}
}
func (msg *MsgDeleteAccState) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAccState) Type() string {
	return "DeleteAccState"
}

func (msg *MsgDeleteAccState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAccState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAccState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
