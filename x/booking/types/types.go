package types

import (
	"encoding/json"
)

// type Booking struct {
// 	BookID      string         `json:"bookingId"`
// 	Booker      sdk.AccAddress `json:"renter"`
// 	UUID        string         `json:"uuid"`
// 	Duration    int64          `json:"duration"`
// 	IsCompleted bool           `json:"isCompleted`
// }

func (b Booking) GetString() (string, error) {
	js, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

func NewBooking() Booking {
	return Booking{}
}

func NewBookingFromMsgBook(msg MsgBook) Booking {
	return Booking{
		UUID:     msg.UUID,
		Booker:   msg.Booker,
		Duration: msg.Duration,
	}
}
