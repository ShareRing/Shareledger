package messages

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
)

type MsgRetrieve struct {
	UUID string `json:"uuid"`
}

// enforce the msg type at compile time
var _ sdk.Msg = MsgRetrieve{}

func NewMsgRetrieve(uuid string) MsgRetrieve {
	return MsgRetrieve{
		UUID: uuid,
	}
}

// Type Implements Msg
func (msg MsgRetrieve) Route() string {
	return constants.MESSAGE_ASSET
}

// Type Implements Msg
func (msg MsgRetrieve) Type() string {
	return constants.MESSAGE_ASSET
}

// ValidateBasic Implements Msg
func (msg MsgRetrieve) ValidateBasic() sdk.Error {
	if len(msg.UUID) == 0 {
		return sdk.ErrInvalidAddress("Invalid address")
	}

	return nil
}

func (msg MsgRetrieve) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg MsgRetrieve) Get(key interface{}) (value interface{}) { return nil }

func (msg MsgRetrieve) String() string {
	return fmt.Sprintf("Asset/MsgRetrieve{%s}", msg.UUID)
}

func (msg MsgRetrieve) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg MsgRetrieve) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", "asset").
		AppendTag("msg.action", "retrieve").
		AppendTag("asset.UUID", msg.UUID)
}
