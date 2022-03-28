package types

import "encoding/json"

func (b Booking) GetString() (string, error) {
	js, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

func NewBookingFromMsgBook(msg MsgCreateBooking) Booking {
	return Booking{
		UUID:     msg.UUID,
		Booker:   msg.Booker,
		Duration: msg.Duration,
	}
}
