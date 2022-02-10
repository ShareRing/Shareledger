package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLoadPShr{}

func NewMsgLoadShr(creator string, address string, amount string) *MsgLoadPShr {
	return &MsgLoadPShr{
		Creator: creator,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgLoadPShr) Route() string {
	return RouterKey
}

func (msg *MsgLoadPShr) Type() string {
	return "LoadPShr"
}

func (msg *MsgLoadPShr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLoadPShr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLoadPShr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	if _, err := ParsePShrCoinsStr(msg.Amount); err != nil {
		return err
	}

	return nil
}
