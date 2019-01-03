package message

import (
	"bytes"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

var _ sdk.Msg = MsgCreateValidator{}
var _ sdk.Msg = MsgEditValidator{}

// MsgCreateValidator - struct for unbonding transactions
type MsgCreateValidator struct {
	Description   posTypes.Description
	DelegatorAddr sdk.AccAddress `json:"delegator_address"`
	ValidatorAddr sdk.AccAddress `json:"validator_address"`
	PubKey        types.PubKey   `json:"pubkey"`
	Delegation    types.Coin     `json:"delegation"`
}

// Type Implements Msg
func (msg MsgCreateValidator) Type() string {
	return constants.MESSAGE_POS
}
func (msg MsgCreateValidator) Route() string { return constants.MESSAGE_POS }

// Return address(es) that must sign over msg.GetSignBytes()
func (msg MsgCreateValidator) GetSigners() []sdk.AccAddress {
	// delegator is first signer so delegator pays fees
	addrs := []sdk.AccAddress{msg.DelegatorAddr}

	if !bytes.Equal(msg.DelegatorAddr.Bytes(), msg.ValidatorAddr.Bytes()) {
		// if validator addr is not same as delegator addr, validator must sign
		// msg as well
		addrs = append(addrs, sdk.AccAddress(msg.ValidatorAddr))
	}
	return addrs
}

// get the bytes for the message signer to sign on
func (msg MsgCreateValidator) GetSignBytes() []byte {
	b, err := json.Marshal(struct {
		Description   posTypes.Description `json:"description"`
		DelegatorAddr sdk.AccAddress       `json:"delegatorAddress"`
		ValidatorAddr sdk.AccAddress       `json:"validatorAddress"`
		PubKey        types.PubKey         `json:"pubKey"`
		Delegation    types.Coin           `json:"delegation"`
	}{
		Description:   msg.Description,
		DelegatorAddr: msg.DelegatorAddr,
		ValidatorAddr: msg.ValidatorAddr,
		PubKey:        msg.PubKey,
		Delegation:    msg.Delegation,
	})
	if err != nil {
		panic(err)
	}

	return b //sdk.MustSortJSON(b)
}

// quick validity check
func (msg MsgCreateValidator) ValidateBasic() sdk.Error {
	if msg.DelegatorAddr == nil {
		return posTypes.ErrNilDelegatorAddr(posTypes.DefaultCodespace)
	}
	if msg.ValidatorAddr == nil {
		return posTypes.ErrNilValidatorAddr(posTypes.DefaultCodespace)
	}
	if !(msg.Delegation.IsPositive()) {
		return posTypes.ErrBadDelegationAmount(posTypes.DefaultCodespace)
	}
	if msg.Description == (posTypes.Description{}) {
		return sdk.NewError(posTypes.DefaultCodespace, posTypes.CodeInvalidInput, "description must be included")
	}

	return nil
}

type MsgEditValidator struct {
	posTypes.Description
	ValidatorAddr sdk.AccAddress `json:"address"`
}

func NewMsgEditValidator(valAddr sdk.AccAddress, description posTypes.Description) MsgEditValidator {
	return MsgEditValidator{
		Description:   description,
		ValidatorAddr: valAddr,
	}
}

//nolint

func (msg MsgEditValidator) Type() string  { return constants.MESSAGE_POS }
func (msg MsgEditValidator) Route() string { return constants.MESSAGE_POS }
func (msg MsgEditValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

// get the bytes for the message signer to sign on
func (msg MsgEditValidator) GetSignBytes() []byte {
	b, err := json.Marshal(struct {
		posTypes.Description
		ValidatorAddr sdk.AccAddress `json:"validatorAddress"`
	}{
		Description:   msg.Description,
		ValidatorAddr: msg.ValidatorAddr,
	})
	if err != nil {
		panic(err)
	}
	return b //sdk.MustSortJSON(b)
}

// quick validity check
func (msg MsgEditValidator) ValidateBasic() sdk.Error {
	if msg.ValidatorAddr == nil {
		return sdk.NewError(posTypes.DefaultCodespace, posTypes.CodeInvalidInput, "nil validator address")
	}

	if msg.Description == (posTypes.Description{}) {
		return sdk.NewError(posTypes.DefaultCodespace, posTypes.CodeInvalidInput, "transaction must include some information to modify")
	}

	return nil
}
