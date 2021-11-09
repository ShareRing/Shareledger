package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgCreateDocBatch(issuer string, holderId, proof, data []string) *MsgCreateDocBatch {
	return &MsgCreateDocBatch{Issuer: issuer, Holder: holderId, Proof: proof, Data: data}
}

func (msg MsgCreateDocBatch) Route() string {
	return RouterKey
}

func (msg MsgCreateDocBatch) Type() string {
	return TypeMsgCreateDocInBatch
}

func (msg MsgCreateDocBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Issuer address (%s)", err)
	}

	// Check len
	maxLen := len(msg.Holder)
	if maxLen > MAX_LEN_BATCH || maxLen != len(msg.Proof) || maxLen != len(msg.Data) {
		return ErrDocInvalidData
	}
	for i := 0; i < maxLen; i++ {
		if len(msg.Holder[i]) > MAX_LEN || len(msg.Data) > MAX_LEN || len(msg.Proof) > MAX_LEN {
			return ErrDocInvalidData
		}
	}
	return nil
}

func (msg MsgCreateDocBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateDocBatch) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Issuer)
	return []sdk.AccAddress{signer}
}
