package messages

import (
	"encoding/json"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
)

//------------------------------------------------------------------
// Msg

// MsgSend implements sdk.Msg
var _ sdk.Msg = MsgSend{}

// MsgSend to send coins from Input to Output
type MsgSend struct {
	To     sdk.Address `json:"to"`
	Amount types.Coin  `json:"amount"`
}

// NewMsgSend
func NewMsgSend(to sdk.Address, amt types.Coin) MsgSend {
	return MsgSend{to, amt}
}

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
func (msg MsgSend) GetSigners() []sdk.Address {
	//return []sdk.Address{msg.From}
	return []sdk.Address{}
}

// Returns the sdk.Tags for the message
func (msg MsgSend) Tags() sdk.Tags {
	return sdk.NewTags("msg.To", []byte(msg.To.String())).
		AppendTag("msg.type", []byte(msg.Type())).
		AppendTag("msg.Amount", []byte(msg.Amount.String())).
		AppendTag("ShareLedgerEvt", []byte("BalanceChanged"))
}
