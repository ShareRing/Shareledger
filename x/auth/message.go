package auth

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"fmt"

	"github.com/sharering/shareledger/constants"
)

var _ sdk.Msg = MsgNonce{}

type MsgNonce struct {
	Address sdk.Address `json:"address"`
}


func NewMsgNonce(account sdk.Address) MsgNonce {
	return MsgNonce{account}
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

func (msg MsgNonce) GetSigners() []sdk.Address {
	return []sdk.Address{}
}

func (msg MsgNonce) Tags() sdk.Tags {
	return sdk.NewTags("checkNonce", []byte(msg.Address.String()))
}
