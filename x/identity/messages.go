package identity

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
)

//------- CREATE -----

type MsgIDCreate struct {
	Address sdk.AccAddress `json:"address"`
	Hash    string         `json:"hash"`
}

var _ sdk.Msg = MsgIDCreate{}

func NewMsgIDCreate(address sdk.AccAddress, hash string) MsgIDCreate {
	return MsgIDCreate{
		Address: address,
		Hash:    hash,
	}
}

// Route ...
func (msg MsgIDCreate) Route() string {
	return constants.MESSAGE_IDENTITY
}

// Type ...
func (msg MsgIDCreate) Type() string {
	return constants.MESSAGE_IDENTITY
}

// ValidateBasic ...
func (msg MsgIDCreate) ValidateBasic() sdk.Error {
	if len(msg.Address) == 0 {
		return sdk.ErrInvalidAddress("Invalid Address")
	}
	return nil
}

// GetSignBytes ...
func (msg MsgIDCreate) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// Get ...
func (msg MsgIDCreate) Get(key interface{}) (value interface{}) { return nil }

// String ...
func (msg MsgIDCreate) String() string {
	return fmt.Sprintf("Identity/MsgIDCreate{Address:%s,Hash:%s}", msg.Address, msg.Hash)
}

// GetSigners ...
func (msg MsgIDCreate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

// Tags ....
func (msg MsgIDCreate) Tags() sdk.Tags {
	return sdk.NewTags(Address, msg.Address.String()).
		AppendTag(Hash, msg.Hash).
		AppendTag(Event, CreatedEvent)
}

//------- UPDATE -----

type MsgIDUpdate struct {
	Address sdk.AccAddress `json:"address"`
	Hash    string         `json:"hash"`
}

var _ sdk.Msg = MsgIDUpdate{}

func NewMsgIDUpdate(address sdk.AccAddress, hash string) MsgIDUpdate {
	return MsgIDUpdate{
		Address: address,
		Hash:    hash,
	}
}

// Route ...
func (msg MsgIDUpdate) Route() string {
	return constants.MESSAGE_IDENTITY
}

// Type ...
func (msg MsgIDUpdate) Type() string {
	return constants.MESSAGE_IDENTITY
}

// ValidateBasic ...
func (msg MsgIDUpdate) ValidateBasic() sdk.Error {
	if len(msg.Address) == 0 {
		return sdk.ErrInvalidAddress("Invalid Address")
	}
	return nil
}

// GetSignBytes ...
func (msg MsgIDUpdate) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// Get ...
func (msg MsgIDUpdate) Get(key interface{}) (value interface{}) { return nil }

// String ...
func (msg MsgIDUpdate) String() string {
	return fmt.Sprintf("Identity/MsgIDUpdate{Address:%s,Hash:%s}", msg.Address, msg.Hash)
}

// GetSigners ...
func (msg MsgIDUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

// Tags ....
func (msg MsgIDUpdate) Tags() sdk.Tags {
	return sdk.NewTags(Address, msg.Address.String()).
		AppendTag(Hash, msg.Hash).
		AppendTag(Event, UpdatedEvent)
}

//-------- DELETED -----

type MsgIDDelete struct {
	Address sdk.AccAddress `json:"address"`
}

var _ sdk.Msg = MsgIDDelete{}

func NewMsgIDDelete(address sdk.AccAddress) MsgIDDelete {
	return MsgIDDelete{
		Address: address,
	}
}

// Route ...
func (msg MsgIDDelete) Route() string {
	return constants.MESSAGE_IDENTITY
}

// Type ...
func (msg MsgIDDelete) Type() string {
	return constants.MESSAGE_IDENTITY
}

// ValidateBasic ...
func (msg MsgIDDelete) ValidateBasic() sdk.Error {
	if len(msg.Address) == 0 {
		return sdk.ErrInvalidAddress("Invalid Address")
	}
	return nil
}

// GetSignBytes ...
func (msg MsgIDDelete) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// Get ...
func (msg MsgIDDelete) Get(key interface{}) (value interface{}) { return nil }

// String ...
func (msg MsgIDDelete) String() string {
	return fmt.Sprintf("Identity/MsgIDDelete{Address:%s}", msg.Address)
}

// GetSigners ...
func (msg MsgIDDelete) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

// Tags ....
func (msg MsgIDDelete) Tags() sdk.Tags {
	return sdk.NewTags(Address, msg.Address.String()).
		AppendTag(Event, DeletedEvent)
}
