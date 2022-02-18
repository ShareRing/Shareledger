package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetActionLevelFee{}

func NewMsgSetActionLevelFee(
	creator string,
	action string,
	level string,

) *MsgSetActionLevelFee {
	return &MsgSetActionLevelFee{
		Creator: creator,
		Action:  action,
		Level:   level,
	}
}

func (msg *MsgSetActionLevelFee) Route() string {
	return RouterKey
}

func (msg *MsgSetActionLevelFee) Type() string {
	return "CreateActionLevelFee"
}

func (msg *MsgSetActionLevelFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetActionLevelFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetActionLevelFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteActionLevelFee{}

func NewMsgDeleteActionLevelFee(
	creator string,
	action string,

) *MsgDeleteActionLevelFee {
	return &MsgDeleteActionLevelFee{
		Creator: creator,
		Action:  action,
	}
}
func (msg *MsgDeleteActionLevelFee) Route() string {
	return RouterKey
}

func (msg *MsgDeleteActionLevelFee) Type() string {
	return "DeleteActionLevelFee"
}

func (msg *MsgDeleteActionLevelFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteActionLevelFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteActionLevelFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
