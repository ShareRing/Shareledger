package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgCreateDoc(issuer, holderId, proof, data string) *MsgCreateDoc {
	return &MsgCreateDoc{Issuer: issuer, Holder: holderId, Proof: proof, Data: data}
}

func (msg MsgCreateDoc) Route() string {
	return RouterKey
}

func (msg MsgCreateDoc) Type() string {
	return TypeMsgCreateDoc
}

func (msg MsgCreateDoc) ValidateBasic() error {
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

func (msg MsgCreateDoc) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateDoc) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Issuer)
	return []sdk.AccAddress{addr}
}
