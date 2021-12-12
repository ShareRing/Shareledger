package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRevokeDocument{}

func NewMsgRevokeDocument(creator string, holder string, issuer string, proof string) *MsgRevokeDocument {
	return &MsgRevokeDocument{
		Creator: creator,
		Holder:  holder,
		Issuer:  issuer,
		Proof:   proof,
	}
}

func (msg *MsgRevokeDocument) Route() string {
	return RouterKey
}

func (msg *MsgRevokeDocument) Type() string {
	return "RevokeDocument"
}

func (msg *MsgRevokeDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
