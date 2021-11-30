package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuyCent{}

func NewMsgBuyCent(creator string, amount string) *MsgBuyCent {
	return &MsgBuyCent{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBuyCent) Route() string {
	return RouterKey
}

func (msg *MsgBuyCent) Type() string {
	return "BuyCent"
}

func (msg *MsgBuyCent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyCent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyCent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, ok := sdk.NewIntFromString(msg.Amount); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount)
	}
	return nil
}
