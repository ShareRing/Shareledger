package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDocument{}

func NewMsgCreateDocument(data string, holder string, issuer string, proof string) *MsgCreateDocument {
	return &MsgCreateDocument{
		Data:   data,
		Holder: holder,
		Issuer: issuer,
		Proof:  proof,
	}
}

func (msg *MsgCreateDocument) Route() string {
	return RouterKey
}

func (msg *MsgCreateDocument) Type() string {
	return TypeMsgCreateDoc
}

func (msg *MsgCreateDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDocument) ValidateBasic() error {
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
