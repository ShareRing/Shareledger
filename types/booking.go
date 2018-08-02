package types

import (
	"fmt"
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"encoding/json"
)

// Simple Booking struct
type Booking struct {
	BookingID   string      `json:"bookingId"`
	Renter      sdk.Address `json:"renter"`
	UUID        string      `json:"uuid"`
	Duration    int64       `json:"duration"`
	IsCompleted bool        `json:"is_completed"`
}

func NewBooking(_bid string, _acc sdk.Address, _uuid string, _dur int64, _isCompleted bool) Booking {
	return Booking{
		BookingID: _bid,
		Renter:    _acc,
		UUID:      _uuid,
		Duration:  _dur,
		IsCompleted: _isCompleted,
	}
}

func (b Booking) String() string {
	//return fmt.Sprintf("{BookingID: %s, Renter: %s, UUID: %x, Duration: %d, IsCompleted: %t}",
	//	b.BookingID, b.Renter, b.UUID, b.Duration, b.IsCompleted)
	bookingBytes, _ := json.Marshal(b)
	return fmt.Sprintf("%s", bookingBytes)
}
