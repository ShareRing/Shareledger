package messages

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
)

type MsgDelete struct {
	UUID string `json:"uuid"`
}

// enforce the msg type at compile time
var _ sdk.Msg = MsgDelete{}

func NewMsgDelete(uuid string) MsgDelete {
	return MsgDelete{
		UUID: uuid,
	}
}

// Type Implements Msg
func (msg MsgDelete) Route() string {
	return constants.MESSAGE_ASSET
}

// Type Implements Msg
func (msg MsgDelete) Type() string {
	return constants.MESSAGE_ASSET
}

// ValidateBasic Implements Msg
func (msg MsgDelete) ValidateBasic() sdk.Error {
	if len(msg.UUID) == 0 {
		return sdk.ErrInvalidAddress("Invalid address")
	}

	return nil
}

func (msg MsgDelete) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg MsgDelete) Get(key interface{}) (value interface{}) { return nil }

func (msg MsgDelete) String() string {
	return fmt.Sprintf("Asset/MsgDelete{%s}", msg.UUID)
}

func (msg MsgDelete) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg MsgDelete) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", []byte("asset")).
		AppendTag("msg.action", []byte("delete")).
		AppendTag("asset.UUID", []byte(msg.UUID))
}
