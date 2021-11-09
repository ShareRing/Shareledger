package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgUpdateId(issuerAddr, id, extraData string) MsgUpdateId {
	return MsgUpdateId{
		ExtraData:     extraData,
		Id:            id,
		IssuerAddress: issuerAddr,
	}
}

func (msg MsgUpdateId) Route() string {
	return RouterKey
}

func (msg MsgUpdateId) Type() string {
	return TypeMsgUpdateID
}

func (msg MsgUpdateId) ValidateBasic() error {
	if len(msg.Id) > MAX_ID_LEN || len(msg.Id) == 0 || len(msg.ExtraData) > MAX_ID_LEN {
		return InvalidData
	}

	_, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid IssuerAddress address (%s)", err)
	}

	return nil
}

func (msg MsgUpdateId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdateId) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.IssuerAddress)
	return []sdk.AccAddress{signer}
}
