package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateDocument{}

func NewMsgUpdateDocument(creator string, data string, holder string, issuer string, proof string) *MsgUpdateDocument {
	return &MsgUpdateDocument{
		Creator: creator,
		Data:    data,
		Holder:  holder,
		Issuer:  issuer,
		Proof:   proof,
	}
}

func (msg *MsgUpdateDocument) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDocument) Type() string {
	return "UpdateDocument"
}

func (msg *MsgUpdateDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
