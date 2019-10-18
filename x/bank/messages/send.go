package messages

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
)

//------------------------------------------------------------------
// Msg

// MsgSend implements sdk.Msg
var _ sdk.Msg = MsgSend{}

// MsgSend to send coins from Input to Output
type MsgSend struct {
	To     sdk.AccAddress `json:"to"`
	Amount types.Coin     `json:"amount"`
}

// NewMsgSend
func NewMsgSend(to sdk.AccAddress, amt types.Coin) MsgSend {
	return MsgSend{to, amt}
}

func (msg MsgSend) Route() string { return constants.MESSAGE_BANK }

// Implements Msg.
func (msg MsgSend) Type() string { return constants.MESSAGE_BANK }

// Implements Msg. Ensure the addresses are good and the
// amount is positive.
func (msg MsgSend) ValidateBasic() sdk.Error {
	if len(msg.To) == 0 {
		return sdk.ErrInvalidAddress("To address is empty")
	}
	if !msg.Amount.IsPositive() {
		return sdk.ErrInvalidCoins("Amount is not positive")
	}
	return nil
}

// Implements Msg. JSON encode the message.
func (msg MsgSend) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return bz
}

// Implements Msg. Return the signer.
func (msg MsgSend) GetSigners() []sdk.AccAddress {
	//return []sdk.AccAddress{msg.From}
	return []sdk.AccAddress{}
}
