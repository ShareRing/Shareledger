package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateDocument{}

func NewMsgUpdateDocument(data string, holder string, issuer string, proof string) *MsgUpdateDocument {
	return &MsgUpdateDocument{
		Data:   data,
		Holder: holder,
		Issuer: issuer,
		Proof:  proof,
	}
}

func (msg *MsgUpdateDocument) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDocument) Type() string {
	return TypeMsgUpdateDoc
}

func (msg *MsgUpdateDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Holder) > MAX_LEN || len(msg.Holder) == 0 {
		return ErrDocInvalidData
	}

	if len(msg.Proof) > MAX_LEN || len(msg.Proof) == 0 {
		return ErrDocInvalidData
	}

	if len(msg.Data) > MAX_LEN {
		return ErrDocInvalidData
	}

	return nil
}
