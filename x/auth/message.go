package auth

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"fmt"

	"github.com/sharering/shareledger/constants"
)

var _ sdk.Msg = MsgNonce{}

type MsgNonce struct {
	Address sdk.AccAddress `json:"address"`
}


func NewMsgNonce(account sdk.AccAddress) MsgNonce {
	return MsgNonce{account}
}

func (msg MsgNonce) Route() string {
	return constants.MESSAGE_AUTH
}

func (msg MsgNonce) Type() string {
	return constants.MESSAGE_AUTH
}

func (msg MsgNonce) ValidateBasic() sdk.Error {
	if len(msg.Address) == 0 {
		return sdk.ErrInvalidAddress("Invalid address")
	}
	return nil
}

func (msg MsgNonce) GetSignBytes() []byte {
	return []byte{}
}

func (msg MsgNonce) String() string {
	return fmt.Sprintf("Auth/MsgNonce{%s}", msg.Address)
}

func (msg MsgNonce) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg MsgNonce) Tags() sdk.Tags {
	return sdk.NewTags("checkNonce", []byte(msg.Address.String()))
}
