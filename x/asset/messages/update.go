package messages

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgUpdate struct {
	Creator sdk.Address `json:"creator"`
	Hash    []byte      `json:"hash"`
	UUID    string      `json:"uuid"`
	Fee     int64       `json:"fee"`
	Status  bool        `json:"status"`
}

// enforce the msg type at compile time
var _ sdk.Msg = MsgUpdate{}

func NewMsgUpdate(creator sdk.Address, hash []byte, uuid string, status bool, fee int64) MsgUpdate {
	return MsgUpdate{
		Creator: creator,
		Hash:    hash,
		UUID:    uuid,
		Fee:     fee,
		Status:  status,
	}
}

// Type Implements Msg
func (msg MsgUpdate) Type() string {
	return "asset"
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
	return fmt.Sprintf("Asset/MsgCreation{%s}", msg.UUID)
}

func (msg MsgUpdate) GetSigners() []sdk.Address {
	return []sdk.Address{msg.Creator}
}

func (msg MsgUpdate) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", []byte("asset")).
		AppendTag("msg.action", []byte("create")).
		AppendTag("asset.creator", []byte(msg.Creator.String())).
		AppendTag("asset.UUID", []byte(msg.UUID)).
		AppendTag("asset.Hash", []byte(msg.Hash)).
		AppendTag("asset.Status", []byte(strconv.FormatBool(msg.Status))).
		AppendTag("asset.Fee", []byte(strconv.Itoa(int(msg.Fee))))
}
