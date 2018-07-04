package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/types"
	"encoding/json"
)

//----------------------------------------------------------------
// Msg


// MsgLoad to load coins into an account
var _ sdk.Msg = MsgLoad{}


// Load coins to an account
type MsgLoad struct {
	Nonce int64 `json:"nonce"`
	Account sdk.Address `json:"from"`
	Amount types.Coin `json:"amount"`
}


// NewMsgLoad
func NewMsgLoad(nonce int64, account sdk.Address, amt types.Coin) MsgLoad {
	return MsgLoad{nonce, account, amt}
}

// Implement Msg
func (msg MsgLoad) Type() string {return "load"}


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
	return []sdk.Address{msg.Account}
}


func (msg MsgLoad) Tags() sdk.Tags {
	return sdk.NewTags("account", []byte(msg.Account.String()))
}
