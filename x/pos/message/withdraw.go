package message

import (
	"encoding/json"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// MsgWithdraw - struct for bonding transactions
type MsgWithdraw struct {
	DelegatorAddr sdk.Address `json:"delegatorAddress"`
	ValidatorAddr sdk.Address `json:"validatorAddress"`
}

func NewMsgWithdraw(delAddr sdk.Address, valAddr sdk.Address) MsgWithdraw {
	return MsgWithdraw{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
	}
}

var _ sdk.Msg = MsgWithdraw{}

//nolint
func (msg MsgWithdraw) Type() string { return constants.MESSAGE_POS }

func (msg MsgWithdraw) GetSigners() []sdk.Address {
	return []sdk.Address{msg.DelegatorAddr}
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
