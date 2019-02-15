package messages

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
)

type MsgUpdate struct {
	Creator sdk.AccAddress `json:"creator"`
	Hash    []byte      `json:"hash"`
	UUID    string      `json:"uuid"`
	Status  bool        `json:"status"`
	Fee     int64       `json:"fee"`
}

// enforce the msg type at compile time
var _ sdk.Msg = MsgUpdate{}

func NewMsgUpdate(creator sdk.AccAddress, hash []byte, uuid string, status bool, fee int64) MsgUpdate {
	return MsgUpdate{
		Creator: creator,
		Hash:    hash,
		UUID:    uuid,
		Fee:     fee,
		Status:  status,
	}
}

// Type Implements Msg
func (msg MsgUpdate) Route() string {
	return constants.MESSAGE_ASSET
}

// Type Implements Msg
func (msg MsgUpdate) Type() string {
	return constants.MESSAGE_ASSET
}

// ValidateBasic Implements Msg
func (msg MsgUpdate) ValidateBasic() sdk.Error {
	if len(msg.Creator) == 0 {
		return sdk.ErrInvalidAddress("Invalid address")
	}

	return nil
}

func (msg MsgUpdate) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg MsgUpdate) Get(key interface{}) (value interface{}) { return nil }

func (msg MsgUpdate) String() string {
	return fmt.Sprintf("Asset/MsgUpdate{%s}", msg.UUID)
}

func (msg MsgUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

func (msg MsgUpdate) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", "asset").
		AppendTag("msg.action", "update").
		AppendTag("asset.creator", msg.Creator.String()).
		AppendTag("asset.UUID", msg.UUID).
		AppendTag("asset.Hash", fmt.Sprintf("%X", msg.Hash)).
		AppendTag("asset.Status", strconv.FormatBool(msg.Status)).
		AppendTag("asset.Fee", strconv.Itoa(int(msg.Fee)))
}
