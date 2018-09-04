package messages

import (
	"encoding/json"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
)

type MsgBook struct {
	UUID     string `json:"uuid"`
	Duration int64  `json:"duration"`
}

var _ sdk.Msg = MsgBook{}

func NewMsgBook(uuid string, duration int64) MsgBook {
	return MsgBook{
		UUID:     uuid,
		Duration: duration,
	}
}

func (msg MsgBook) Type() string {
	return constants.MESSAGE_BOOKING
}

func (msg MsgBook) ValidateBasic() sdk.Error {
	//if len(msg.Renter) == 0 {
	//return sdk.ErrInvalidAddress("Invalid address")
	//}

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
	return fmt.Sprintf("Booking/MsgBook{UUID: %s}", msg.UUID)
}

func (msg MsgBook) GetSigners() []sdk.Address {
	//return []sdk.Address{msg.Renter}
	return []sdk.Address{}
}

func (msg MsgBook) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", []byte("booking")).
		AppendTag("msg.action", []byte("book"))
}
