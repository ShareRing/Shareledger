package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgRevokeDoc(issuer, holder, proof string) MsgRevokeDoc {
	return MsgRevokeDoc{Issuer: issuer, Holder: holder, Proof: proof}
}

func (msg MsgRevokeDoc) Route() string {
	return RouterKey
}

func (msg MsgRevokeDoc) Type() string {
	return TypeMsgRevokeDoc
}

func (msg MsgRevokeDoc) ValidateBasic() error {
	if len(msg.Holder) > MAX_LEN || len(msg.Holder) == 0 {
		return ErrDocInvalidData
	}

	if len(msg.Proof) > MAX_LEN || len(msg.Proof) == 0 {
		return ErrDocInvalidData
	}

	_, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Issuer address (%s)", err)
	}
	return nil
}

func (msg MsgRevokeDoc) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRevokeDoc) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Issuer)
	return []sdk.AccAddress{signer}
}
