package messages

import (
	"encoding/json"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	tags "github.com/sharering/shareledger/x/bank/tags"
)

//----------------------------------------------------------------
// Msg

// MsgLoad to load coins into an account
var _ sdk.Msg = MsgLoad{}

// Load coins to an account
type MsgLoad struct {
	Account sdk.Address `json:"address"`
	Amount  types.Coin  `json:"amount"`
}

// NewMsgLoad
func NewMsgLoad(account sdk.Address, amt types.Coin) MsgLoad {
	return MsgLoad{account, amt}
}

// Implement Msg
func (msg MsgLoad) Type() string { return constants.MESSAGE_BANK }

// Implement Msg. Load to ensure the address are good and the amount is positive
func (msg MsgLoad) ValidateBasic() sdk.Error {
	if len(msg.Account) == 0 {
		return sdk.ErrInvalidAddress("Account address is empty")
	}

	if !msg.Amount.IsPositive() {
		return sdk.ErrInvalidCoins("Amount is not positive")
	}
	return nil
}

func (msg MsgLoad) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return bz
}

func (msg MsgLoad) GetSigners() []sdk.Address {
	return []sdk.Address{}
	//return []sdk.Address{msg.Account}
}

func (msg MsgLoad) Tags() sdk.Tags {
	return sdk.NewTags(tags.AccountAddress, []byte(msg.Account.String())).
		AppendTag(tags.Amount, []byte(msg.Amount.String())).
		AppendTag(tags.Event, tags.Credit)

}
