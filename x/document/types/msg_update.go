package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgUpdateDoc(issuer, holderId, proof, data string) MsgUpdateDoc {
	return MsgUpdateDoc{Issuer: issuer, Holder: holderId, Proof: proof, Data: data}
}

func (msg MsgUpdateDoc) Route() string {
	return RouterKey
}

func (msg MsgUpdateDoc) Type() string {
	return TypeMsgUpdateDoc
}

func (msg MsgUpdateDoc) ValidateBasic() error {
	if len(msg.Holder) > MAX_LEN || len(msg.Holder) == 0 {
		return ErrDocInvalidData
	}

	if len(msg.Proof) > MAX_LEN || len(msg.Proof) == 0 {
		return ErrDocInvalidData
	}

	if len(msg.Data) > MAX_LEN || len(msg.Data) == 0 {
		return ErrDocInvalidData
	}

	_, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Issuer address (%s)", err)
	}
	return nil
}

func (msg MsgUpdateDoc) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdateDoc) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Issuer)
	return []sdk.AccAddress{signer}
}
