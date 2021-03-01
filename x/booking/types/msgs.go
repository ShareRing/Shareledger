package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeBookMsg         = "booking"
	TypeBookCompleteMsg = "book_complete"
)

type MsgBook struct {
	Booker   sdk.AccAddress `json:"booker"`
	UUID     string         `json:"uuid"`
	Duration int64          `json:"duration"`
}

func NewMsgBook(booker sdk.AccAddress, uuid string, duration int64) MsgBook {
	return MsgBook{
		Booker:   booker,
		UUID:     uuid,
		Duration: duration,
	}
}

func (msg MsgBook) Route() string {
	return RouterKey
}

func (msg MsgBook) Type() string {
	return TypeBookMsg
}

func (msg MsgBook) ValidateBasic() error {
	if len(msg.Booker) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Address of creator must not be empty")
	}
	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	if msg.Duration <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "duration must be positive")
	}
	return nil
}

func (msg MsgBook) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBook) String() string {
	return fmt.Sprintf("Booking/MsgBook{UUID: %s}", msg.UUID)
}

func (msg MsgBook) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Booker}
}

type MsgComplete struct {
	Booker sdk.AccAddress `json:"booker"`
	BookID string         `json:"uuid"`
}

func NewMsgComplete(booker sdk.AccAddress, bookid string) MsgComplete {
	return MsgComplete{
		Booker: booker,
		BookID: bookid,
	}
}

func (msg MsgComplete) Route() string {
	return RouterKey
}

func (msg MsgComplete) Type() string {
	return TypeBookCompleteMsg
}

func (msg MsgComplete) ValidateBasic() error {
	if len(msg.Booker) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Address of creator must not be empty")
	}
	if len(msg.BookID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	return nil
}

func (msg MsgComplete) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgComplete) String() string {
	return fmt.Sprintf("Booking/MsgComplete{BookID: %s}", msg.BookID)
}

func (msg MsgComplete) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Booker}
}
