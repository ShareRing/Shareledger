package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLoadShr{}

func NewMsgLoadShr(creator string, address string, amount string) *MsgLoadShr {
	return &MsgLoadShr{
		Creator: creator,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgLoadShr) Route() string {
	return RouterKey
}

func (msg *MsgLoadShr) Type() string {
	return "LoadShr"
}

func (msg *MsgLoadShr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLoadShr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLoadShr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}
	v, e := sdk.NewDecFromStr(msg.Amount)
	if e != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid value. %s", e.Error())
	}
	if v.LTE(sdk.NewDec(0)) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid value, should large then 0")
	}

	return nil
}
