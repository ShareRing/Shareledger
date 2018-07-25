package messages

import (
	"encoding/json"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
)

type MsgComplete struct {
	Nonce     int64       `json:"nonce"`
	Renter    sdk.Address `json:"renter"`
	BookingID string      `json:"bookingId"`
}

var _ sdk.Msg = MsgComplete{}

func NewMsgComplete(nonce int64, renter sdk.Address, bookingId string) MsgComplete {
	return MsgComplete{
		Nonce:     nonce,
		Renter:    renter,
		BookingID: bookingId,
	}
}

func (msg MsgComplete) Type() string {
	return "booking"
}

func (msg MsgComplete) ValidateBasic() sdk.Error {
	if len(msg.Renter) == 0 {
		return sdk.ErrInvalidAddress("Invalid address")
	}

	return nil
}

func (msg MsgComplete) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return b
}

func (msg MsgComplete) Get(key interface{}) (value interface{}) { return nil }

func (msg MsgComplete) String() string {
	return fmt.Sprintf("Booking/MsgComplete{Renter: %s, BookingID: %s}", msg.Renter, msg.BookingID)
}

func (msg MsgComplete) GetSigners() []sdk.Address {
	return []sdk.Address{msg.Renter}
}

func (msg MsgComplete) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", []byte("booking")).
		AppendTag("msg.action", []byte("complete"))
}
