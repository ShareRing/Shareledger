package messages

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	tags "github.com/sharering/shareledger/x/bank/tags"
)

//----------------------------------------------------------------
// Msg

// MsgBurn to load coins into an account
var _ sdk.Msg = MsgBurn{}

// Load coins to an account
type MsgBurn struct {
	Account sdk.AccAddress `json:"address"`
	Amount  types.Coin  `json:"amount"`
}

// NewMsgBurn
func NewMsgBurn(account sdk.AccAddress, amt types.Coin) MsgBurn {
	return MsgBurn{account, amt}
}

func (msg MsgBurn) Route() string { return constants.MESSAGE_BANK }

// Implement Msg
func (msg MsgBurn) Type() string { return constants.MESSAGE_BANK }

// Implement Msg. Load to ensure the address are good and the amount is positive
func (msg MsgBurn) ValidateBasic() sdk.Error {
	if len(msg.Account) == 0 {
		return sdk.ErrInvalidAddress("Account address is empty")
	}

	if !msg.Amount.IsPositive() {
		return sdk.ErrInvalidCoins("Amount is not positive")
	}

	if strings.Compare(msg.Amount.Denom, constants.BOOKING_DENOM) != 0 {
		return sdk.ErrInternal(fmt.Sprintf(constants.BANK_INVALID_BURNT_DENOM, constants.BOOKING_DENOM))
	}

	return nil
}

func (msg MsgBurn) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return bz
}

func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg MsgBurn) Tags() sdk.Tags {
	return sdk.NewTags(tags.AccountAddress, msg.Account.String()).
		AppendTag(tags.Amount, msg.Amount.String()).
		AppendTag(tags.Event, tags.Credit)

}
