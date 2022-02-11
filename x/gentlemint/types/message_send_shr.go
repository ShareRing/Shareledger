package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendShr{}

func NewMsgSendShr(creator string, address string, amount string) *MsgSendShr {
	return &MsgSendShr{
		Creator: creator,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgSendShr) Route() string {
	return RouterKey
}

func (msg *MsgSendShr) Type() string {
	return "SendShr"
}

func (msg *MsgSendShr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendShr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendShr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}
	if d, err := sdk.NewDecFromStr(msg.Amount); err != nil || d.LTE(sdk.NewDec(0)) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, msg.Amount)
	}
	return nil
}
