package messages

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//------------------------------------------------------------------
// Msg

// MsgCheck implements sdk.Msg
var _ sdk.Msg = MsgCheck{}

// MsgCheck to send coins from Input to Output
type MsgCheck struct {
	Account   sdk.Address `json:"account"`
	Denom 	  string   `json:"denom"`
}

// NewMsgCheck
func NewMsgCheck(account sdk.Address, denom string) MsgCheck {
	return MsgCheck{account, denom}
}

// Implements Msg.
func (msg MsgCheck) Type() string { return "bank" }

// Implements Msg. Ensure the addresses are good and the
// amount is positive.
func (msg MsgCheck) ValidateBasic() sdk.Error {
	if len(msg.Account) == 0 {
		return sdk.ErrInvalidAddress("Account address is empty")
	}
	if len(msg.Denom) == 0 {
		return sdk.ErrInvalidAddress("Denomination is empty")
	}
	return nil
}

// Implements Msg. JSON encode the message.
func (msg MsgCheck) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return bz
}

// Implements Msg. Return the signer.
func (msg MsgCheck) GetSigners() []sdk.Address {
	return []sdk.Address{msg.Account}
}

// Returns the sdk.Tags for the message
func (msg MsgCheck) Tags() sdk.Tags {
	return sdk.NewTags("check", []byte(msg.Account.String()))
}