package message

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// MsgWithdraw - struct for bonding transactions
type MsgWithdraw struct {
	DelegatorAddr sdk.AccAddress `json:"delegatorAddress"`
	ValidatorAddr sdk.AccAddress `json:"validatorAddress"`
}

func NewMsgWithdraw(delAddr sdk.AccAddress, valAddr sdk.AccAddress) MsgWithdraw {
	return MsgWithdraw{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
	}
}

var _ sdk.Msg = MsgWithdraw{}

//nolint
func (msg MsgWithdraw) Type() string { return constants.MESSAGE_POS }

func (msg MsgWithdraw) Route() string { return constants.MESSAGE_POS }

func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.DelegatorAddr}
}

// get the bytes for the message signer to sign on
func (msg MsgWithdraw) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b //sdk.MustSortJSON(b)
}

// quick validity check
func (msg MsgWithdraw) ValidateBasic() sdk.Error {
	if msg.DelegatorAddr == nil {
		return posTypes.ErrNilDelegatorAddr(posTypes.DefaultCodespace)
	}
	if msg.ValidatorAddr == nil {
		return posTypes.ErrNilValidatorAddr(posTypes.DefaultCodespace)
	}
	return nil
}
