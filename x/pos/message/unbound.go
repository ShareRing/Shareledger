package message

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// MsgBeginUnbonding - struct for unbonding transactions
type MsgBeginUnbonding struct {
	DelegatorAddr sdk.AccAddress `json:"delegatorAddress"`
	ValidatorAddr sdk.AccAddress `json:"validatorAddress"`
	SharesAmount  types.Dec      `json:"sharesAmount"`
}

func NewMsgBeginUnbonding(delAddr sdk.AccAddress, valAddr sdk.AccAddress, sharesAmount types.Dec) MsgBeginUnbonding {
	return MsgBeginUnbonding{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
		SharesAmount:  sharesAmount,
	}
}

//nolint
func (msg MsgBeginUnbonding) Type() string                 { return constants.MESSAGE_POS }
func (msg MsgBeginUnbonding) Route() string                { return constants.MESSAGE_POS }
func (msg MsgBeginUnbonding) Name() string                 { return "begin_unbonding" }
func (msg MsgBeginUnbonding) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{msg.DelegatorAddr} }

// get the bytes for the message signer to sign on
func (msg MsgBeginUnbonding) GetSignBytes() []byte {
	shareAmount, _ := msg.SharesAmount.MarshalAmino()
	b, err := json.Marshal(struct {
		DelegatorAddr sdk.AccAddress `json:"delegatorAddress"`
		ValidatorAddr sdk.AccAddress `json:"validatorAddress"`
		SharesAmount  string         `json:"sharesAmount"`
	}{
		DelegatorAddr: msg.DelegatorAddr,
		ValidatorAddr: msg.ValidatorAddr,
		SharesAmount:  shareAmount,
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
	DelegatorAddr sdk.AccAddress `json:"delegatorAddress"`
	ValidatorAddr sdk.AccAddress `json:"validatorAddress"`
}

func NewMsgCompleteUnbonding(delAddr sdk.AccAddress, valAddr sdk.AccAddress) MsgCompleteUnbonding {
	return MsgCompleteUnbonding{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
	}
}

//nolint
func (msg MsgCompleteUnbonding) Type() string  { return constants.MESSAGE_POS }
func (msg MsgCompleteUnbonding) Route() string { return constants.MESSAGE_POS }
func (msg MsgCompleteUnbonding) Name() string  { return "complete_unbonding" }
func (msg MsgCompleteUnbonding) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.DelegatorAddr}
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
var _ sdk.Msg = MsgBeginUnbonding{}
