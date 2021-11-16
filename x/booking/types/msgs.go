package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeBookMsg         = "booking"
	TypeBookCompleteMsg = "book_complete"
)

// type MsgBook struct {
// 	Booker   sdk.AccAddress `json:"booker"`
// 	UUID     string         `json:"uuid"`
// 	Duration int64          `json:"duration"`
// }

func NewMsgBook(booker, uuid string, duration int64) MsgBook {
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
	_, err := sdk.AccAddressFromBech32(msg.Booker)
	if len(msg.Booker) == 0 || err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid booker address")
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
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgBook) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Booker)
	return []sdk.AccAddress{signer}
}

// type MsgComplete struct {
// 	Booker sdk.AccAddress `json:"booker"`
// 	BookID string         `json:"uuid"`
// }

func NewMsgComplete(booker, bookid string) MsgComplete {
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
	_, err := sdk.AccAddressFromBech32(msg.Booker)
	if len(msg.Booker) == 0 || err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid booker address")
	}

	if len(msg.BookID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "UUID must not be empty")
	}
	return nil
}

func (msg MsgComplete) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgComplete) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Booker)
	return []sdk.AccAddress{signer}
}
