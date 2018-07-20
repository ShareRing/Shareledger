package messages

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgBook struct {
	Nonce    int64       `json:"nonce"`
	Renter   sdk.Address `json:"renter"`
	UUID     string      `json:"uuid"`
	Duration int64       `json:"duration"`
}

var _ sdk.Msg = MsgBook{}

func NewMsgBook(nonce int64, renter sdk.Address, uuid string, duration int64) MsgBook {
	return MsgBook{
		Nonce:    nonce,
		Renter:   renter,
		UUID:     uuid,
		Duration: duration,
	}
}

func (msg MsgBook) Type() string {
	return "booking"
}

func (msg MsgBook) ValidateBasic() sdk.Error {
	if len(msg.Renter) == 0 {
		return sdk.ErrInvalidAddress("Invalid address")
	}

	return nil
}

func (msg MsgBook) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return b
}

func (msg MsgBook) Get(key interface{}) (value interface{}) { return nil }

func (msg MsgBook) String() string {
	return fmt.Sprintf("Booking/MsgBook{Renter: %s, UUID: %s}", msg.Renter, msg.UUID)
}

func (msg MsgBook) GetSigners() []sdk.Address {
	return []sdk.Address{msg.Renter}
}

func (msg MsgBook) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", []byte("booking")).
		AppendTag("msg.action", []byte("book"))
}
