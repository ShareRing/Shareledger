package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

var _ sdk.Msg = &MsgSetLevelFee{}

func NewMsgSetLevelFee(
	creator string,
	level string,
	fee sdk.DecCoin,

) *MsgSetLevelFee {
	return &MsgSetLevelFee{
		Creator: creator,
		Level:   level,
		Fee:     fee,
	}
}

func (msg *MsgSetLevelFee) Route() string {
	return RouterKey
}

func (msg *MsgSetLevelFee) Type() string {
	return "SetLevelFee"
}

func (msg *MsgSetLevelFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetLevelFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetLevelFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if err := msg.Fee.Validate(); err != nil {
		return err
	}
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, err.Error())
	}
	if err := denom.CheckFeeSupportedCoin(msg.Fee); err != nil {
		return err
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteLevelFee{}

func NewMsgDeleteLevelFee(
	creator string,
	level string,

) *MsgDeleteLevelFee {
	return &MsgDeleteLevelFee{
		Creator: creator,
		Level:   level,
	}
}
func (msg *MsgDeleteLevelFee) Route() string {
	return RouterKey
}

func (msg *MsgDeleteLevelFee) Type() string {
	return "DeleteLevelFee"
}

func (msg *MsgDeleteLevelFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteLevelFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteLevelFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
