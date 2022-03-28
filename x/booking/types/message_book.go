package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBooking{}

func NewMsgBook(booker string, UUID string, duration int64) *MsgCreateBooking {
	return &MsgCreateBooking{
		Booker:   booker,
		UUID:     UUID,
		Duration: duration,
	}
}

func (msg *MsgCreateBooking) Route() string {
	return RouterKey
}

func (msg *MsgCreateBooking) Type() string {
	return TypeBookMsg
}

func (msg *MsgCreateBooking) GetSigners() []sdk.AccAddress {
	booker, err := sdk.AccAddressFromBech32(msg.Booker)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{booker}
}

func (msg *MsgCreateBooking) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBooking) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Booker)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid booker address (%s)", err)
	}

	if len(msg.UUID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "UUID must not be empty")
	}
	if msg.Duration <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "duration must be positive")
	}

	return nil
}
