package message

import (
	"bytes"
	"encoding/json"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// MsgCreateValidator - struct for unbonding transactions
type MsgCreateValidator struct {
	Description   posTypes.Description
	DelegatorAddr sdk.Address  `json:"delegator_address"`
	ValidatorAddr sdk.Address  `json:"validator_address"`
	PubKey        types.PubKey `json:"pubkey"`
	Delegation    types.Coin   `json:"delegation"`
}

var _ sdk.Msg = MsgCreateValidator{}

// Type Implements Msg
func (msg MsgCreateValidator) Type() string {
	return constants.MESSAGE_POS
}

// Return address(es) that must sign over msg.GetSignBytes()
func (msg MsgCreateValidator) GetSigners() []sdk.Address {
	// delegator is first signer so delegator pays fees
	addrs := []sdk.Address{msg.DelegatorAddr}

	if !bytes.Equal(msg.DelegatorAddr.Bytes(), msg.ValidatorAddr.Bytes()) {
		// if validator addr is not same as delegator addr, validator must sign
		// msg as well
		addrs = append(addrs, sdk.Address(msg.ValidatorAddr))
	}
	return addrs
}

// get the bytes for the message signer to sign on
func (msg MsgCreateValidator) GetSignBytes() []byte {
	b, err := json.Marshal(struct {
		Description   posTypes.Description `json:"description"`
		DelegatorAddr sdk.Address          `json:"delegatorAddress"`
		ValidatorAddr sdk.Address          `json:"validatorAddress"`
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
