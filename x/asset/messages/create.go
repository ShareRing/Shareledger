package messages

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
)

type MsgCreate struct {
	Creator sdk.Address `json:"creator"`
	Hash    []byte      `json:"hash"`
	UUID    string      `json:"uuid"`
	Status  bool        `json:"status"`
	Fee     int64       `json:"fee"`
}

// enforce the msg type at compile time
var _ sdk.Msg = MsgCreate{}

func NewMsgCreate(creator sdk.Address, hash []byte, uuid string, status bool, fee int64) MsgCreate {
	return MsgCreate{
		Creator: creator,
		Hash:    hash,
		UUID:    uuid,
		Fee:     fee,
		Status:  status,
	}
}

// Type Implements Msg
func (msg MsgCreate) Type() string {
	return constants.MESSAGE_ASSET
}

// ValidateBasic Implements Msg
func (msg MsgCreate) ValidateBasic() sdk.Error {
	if len(msg.Creator) == 0 {
		return sdk.ErrInvalidAddress("Invalid address")
	}

	return nil
}

func (msg MsgCreate) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg MsgCreate) Get(key interface{}) (value interface{}) { return nil }

func (msg MsgCreate) String() string {
	return fmt.Sprintf("Asset/MsgCreation{%s}", msg.UUID)
}

func (msg MsgCreate) GetSigners() []sdk.Address {
	return []sdk.Address{msg.Creator}
}

func (msg MsgCreate) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", []byte("asset")).
		AppendTag("msg.action", []byte("create")).
		AppendTag("asset.creator", []byte(msg.Creator.String())).
		AppendTag("asset.UUID", []byte(msg.UUID)).
		AppendTag("asset.Hash", []byte(msg.Hash)).
		AppendTag("asset.Status", []byte(strconv.FormatBool(msg.Status))).
		AppendTag("asset.Fee", []byte(strconv.Itoa(int(msg.Fee))))
}

//------------------------------------------
// Testing purpose

func GetMsgCreate() MsgCreate {
	pubKey, _ := types.GenerateKeyPair()

	address := pubKey.Address()
	hash := []byte("111111")
	fee := int64(1)
	status := true
	uuid := "112233"

	msgCreate := NewMsgCreate(
		address,
		hash,
		uuid,
		status,
		fee,
	)
	return msgCreate
}
