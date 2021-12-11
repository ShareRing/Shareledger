package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBook{}

func NewMsgBook(booker string, uUID string, duration int64) *MsgBook {
	return &MsgBook{
		Booker:   booker,
		UUID:     uUID,
		Duration: duration,
	}
}

func (msg *MsgBook) Route() string {
	return RouterKey
}

func (msg *MsgBook) Type() string {
	return "Book"
}

func (msg *MsgBook) GetSigners() []sdk.AccAddress {
	booker, err := sdk.AccAddressFromBech32(msg.Booker)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{booker}
}

func (msg *MsgBook) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBook) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Booker)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid booker address (%s)", err)
	}
	return nil
}
