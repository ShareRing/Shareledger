package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgComplete{}

func NewMsgComplete(booker string, bookID string) *MsgComplete {
	return &MsgComplete{
		Booker: booker,
		BookID: bookID,
	}
}

func (msg *MsgComplete) Route() string {
	return RouterKey
}

func (msg *MsgComplete) Type() string {
	return "Complete"
}

func (msg *MsgComplete) GetSigners() []sdk.AccAddress {
	booker, err := sdk.AccAddressFromBech32(msg.Booker)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{booker}
}

func (msg *MsgComplete) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgComplete) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Booker)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid booker address (%s)", err)
	}
	return nil
}
