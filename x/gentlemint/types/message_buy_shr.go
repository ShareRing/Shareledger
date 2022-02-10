package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuyPShr{}

func NewMsgBuyPShr(creator string, amount string) *MsgBuyPShr {
	return &MsgBuyPShr{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBuyPShr) Route() string {
	return RouterKey
}

func (msg *MsgBuyPShr) Type() string {
	return "BuyPShr"
}

func (msg *MsgBuyPShr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyPShr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyPShr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := ParsePShrCoinsStr(msg.Amount); err != nil {
		return err
	}
	return nil
}
