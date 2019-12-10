package messages

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
)

//------------------------------------------------------------------
// Msg

// MsgTransferShr implements sdk.Msg
var _ sdk.Msg = MsgTransferShr{}

// MsgTransferShr to send coins from Input to Output
type MsgTransferShr struct {
	To     sdk.AccAddress `json:"to"`
	Amount types.Coin     `json:"amount"`
}

// NewMsgTransferShr
func NewMsgTransferShr(to sdk.AccAddress, amt types.Coin) MsgTransferShr {
	return MsgTransferShr{to, amt}
}

func (msg MsgTransferShr) Route() string { return constants.MESSAGE_BANK }

// Implements Msg.
func (msg MsgTransferShr) Type() string { return constants.MESSAGE_BANK }

// Implements Msg. Ensure the addresses are good and the
// amount is positive.
func (msg MsgTransferShr) ValidateBasic() sdk.Error {
	if len(msg.To) == 0 {
		return sdk.ErrInvalidAddress("To address is empty")
	}
	if !msg.Amount.IsPositive() {
		return sdk.ErrInvalidCoins("Amount is not positive")
	}
	return nil
}

// Implements Msg. JSON encode the message.
func (msg MsgTransferShr) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return bz
}

// Implements Msg. Return the signer.
func (msg MsgTransferShr) GetSigners() []sdk.AccAddress {
	//return []sdk.AccAddress{msg.From}
	return []sdk.AccAddress{}
}
