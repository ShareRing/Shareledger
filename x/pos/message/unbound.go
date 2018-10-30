package message

import (
	"encoding/json"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// MsgBeginUnbonding - struct for unbonding transactions
type MsgBeginUnbonding struct {
	DelegatorAddr sdk.Address `json:"delegator_addr"`
	ValidatorAddr sdk.Address `json:"validator_addr"`
	SharesAmount  types.Dec   `json:"shares_amount"`
}

func NewMsgBeginUnbonding(delAddr sdk.Address, valAddr sdk.Address, sharesAmount types.Dec) MsgBeginUnbonding {
	return MsgBeginUnbonding{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
		SharesAmount:  sharesAmount,
	}
}

//nolint
func (msg MsgBeginUnbonding) Type() string              { return constants.MESSAGE_POS }
func (msg MsgBeginUnbonding) Name() string              { return "begin_unbonding" }
func (msg MsgBeginUnbonding) GetSigners() []sdk.Address { return []sdk.Address{msg.DelegatorAddr} }

// get the bytes for the message signer to sign on
func (msg MsgBeginUnbonding) GetSignBytes() []byte {
	b, err := json.Marshal(struct {
		DelegatorAddr sdk.Address `json:"delegator_addr"`
		ValidatorAddr sdk.Address `json:"validator_addr"`
		SharesAmount  string      `json:"shares_amount"`
	}{
		DelegatorAddr: msg.DelegatorAddr,
		ValidatorAddr: msg.ValidatorAddr,
		SharesAmount:  msg.SharesAmount.String(),
	})
	if err != nil {
		panic(err)
	}
	return b //sdk.MustSortJSON(b)
}

// quick validity check
func (msg MsgBeginUnbonding) ValidateBasic() sdk.Error {
	if msg.DelegatorAddr == nil {
		return posTypes.ErrNilDelegatorAddr(posTypes.DefaultCodespace)
	}
	if msg.ValidatorAddr == nil {
		return posTypes.ErrNilValidatorAddr(posTypes.DefaultCodespace)
	}
	if msg.SharesAmount.LTE(types.ZeroDec()) {
		return posTypes.ErrBadSharesAmount(posTypes.DefaultCodespace)
	}
	return nil
}

var _ sdk.Msg = MsgBeginUnbonding{}

// MsgCompleteUnbonding - struct for unbonding transactions
type MsgCompleteUnbonding struct {
	DelegatorAddr sdk.Address `json:"delegator_addr"`
	ValidatorAddr sdk.Address `json:"validator_addr"`
}

func NewMsgCompleteUnbonding(delAddr sdk.Address, valAddr sdk.Address) MsgCompleteUnbonding {
	return MsgCompleteUnbonding{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
	}
}

//nolint
func (msg MsgCompleteUnbonding) Type() string { return constants.MESSAGE_POS }
func (msg MsgCompleteUnbonding) Name() string { return "complete_unbonding" }
func (msg MsgCompleteUnbonding) GetSigners() []sdk.Address {
	return []sdk.Address{msg.DelegatorAddr}
}

// get the bytes for the message signer to sign on
func (msg MsgCompleteUnbonding) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b //sdk.MustSortJSON(b)
}

// quick validity check
func (msg MsgCompleteUnbonding) ValidateBasic() sdk.Error {
	if msg.DelegatorAddr == nil {
		return posTypes.ErrNilDelegatorAddr(posTypes.DefaultCodespace)
	}
	if msg.ValidatorAddr == nil {
		return posTypes.ErrNilValidatorAddr(posTypes.DefaultCodespace)
	}
	return nil
}

var _ sdk.Msg = MsgCompleteUnbonding{}
