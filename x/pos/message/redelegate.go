package message

import (
	"encoding/json"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	types "github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

//______________________________________________________________________

// MsgDelegate - struct for bonding transactions
type MsgBeginRedelegate struct {
	DelegatorAddr    sdk.Address `json:"delegatorAddress"`
	ValidatorSrcAddr sdk.Address `json:"validatorSrcAddress"`
	ValidatorDstAddr sdk.Address `json:"validatorDstAddress"`
	SharesAmount     types.Dec   `json:"shareAmount"`
}

func NewMsgBeginRedelegate(delAddr sdk.Address, valSrcAddr,
	valDstAddr sdk.Address, sharesAmount types.Dec) MsgBeginRedelegate {

	return MsgBeginRedelegate{
		DelegatorAddr:    delAddr,
		ValidatorSrcAddr: valSrcAddr,
		ValidatorDstAddr: valDstAddr,
		SharesAmount:     sharesAmount,
	}
}

//nolint

func (msg MsgBeginRedelegate) Type() string { return "BeginRedelegate" }
func (msg MsgBeginRedelegate) GetSigners() []sdk.Address {
	return []sdk.Address{msg.DelegatorAddr}
}

// get the bytes for the message signer to sign on
func (msg MsgBeginRedelegate) GetSignBytes() []byte {
	b, err := /*MsgCdc.MarshalJSON*/ json.Marshal(struct {
		DelegatorAddr    sdk.Address `json:"delegatorAddress"`
		ValidatorSrcAddr sdk.Address `json:"validatorSrcAddress"`
		ValidatorDstAddr sdk.Address `json:"validatorDstAddress"`
		SharesAmount     string      `json:"shareAmount"`
	}{
		DelegatorAddr:    msg.DelegatorAddr,
		ValidatorSrcAddr: msg.ValidatorSrcAddr,
		ValidatorDstAddr: msg.ValidatorDstAddr,
		SharesAmount:     msg.SharesAmount.String(),
	})
	if err != nil {
		panic(err)
	}
	return b //sdk.MustSortJSON(b)
}

// quick validity check
func (msg MsgBeginRedelegate) ValidateBasic() sdk.Error {
	if msg.DelegatorAddr == nil {
		return posTypes.ErrNilDelegatorAddr(posTypes.DefaultCodespace)
	}
	if msg.ValidatorSrcAddr == nil {
		return posTypes.ErrNilValidatorAddr(posTypes.DefaultCodespace)
	}
	if msg.ValidatorDstAddr == nil {
		return posTypes.ErrNilValidatorAddr(posTypes.DefaultCodespace)
	}
	if msg.SharesAmount.LTE(types.ZeroDec()) {
		return posTypes.ErrBadSharesAmount(posTypes.DefaultCodespace)
	}
	return nil
}

// MsgDelegate - struct for bonding transactions
type MsgCompleteRedelegate struct {
	DelegatorAddr    sdk.Address `json:"delegatorAddress"`
	ValidatorSrcAddr sdk.Address `json:"validatorSrcAddress"`
	ValidatorDstAddr sdk.Address `json:"validatorDstAddress"`
}

func NewMsgCompleteRedelegate(delegatorAddr, validatorSrcAddr,
	validatorDstAddr sdk.Address) MsgCompleteRedelegate {

	return MsgCompleteRedelegate{
		DelegatorAddr:    delegatorAddr,
		ValidatorSrcAddr: validatorSrcAddr,
		ValidatorDstAddr: validatorDstAddr,
	}
}

//nolint
func (msg MsgCompleteRedelegate) Type() string { return "CompleteRedelegate" }
func (msg MsgCompleteRedelegate) GetSigners() []sdk.Address {
	return []sdk.Address{msg.DelegatorAddr}
}

// get the bytes for the message signer to sign on
func (msg MsgCompleteRedelegate) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b //sdk.MustSortJSON(b)
}

// quick validity check
func (msg MsgCompleteRedelegate) ValidateBasic() sdk.Error {
	if msg.DelegatorAddr == nil {
		return posTypes.ErrNilDelegatorAddr(posTypes.DefaultCodespace)
	}
	if msg.ValidatorSrcAddr == nil {
		return posTypes.ErrNilValidatorAddr(posTypes.DefaultCodespace)
	}
	if msg.ValidatorDstAddr == nil {
		return posTypes.ErrNilValidatorAddr(posTypes.DefaultCodespace)
	}
	return nil
}
