package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuyShr{}

func NewMsgBuyShr(creator string, amount string) *MsgBuyShr {
	return &MsgBuyShr{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBuyShr) Route() string {
	return RouterKey
}

func (msg *MsgBuyShr) Type() string {
	return "BuyShr"
}

func (msg *MsgBuyShr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyShr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyShr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := ParseShrCoinsStr(msg.Amount); err != nil {
		return err
	}
	return nil
}
