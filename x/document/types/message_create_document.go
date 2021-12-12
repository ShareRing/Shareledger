package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDocument{}

func NewMsgCreateDocument(creator string, data string, holder string, issuer string, proof string) *MsgCreateDocument {
	return &MsgCreateDocument{
		Creator: creator,
		Data:    data,
		Holder:  holder,
		Issuer:  issuer,
		Proof:   proof,
	}
}

func (msg *MsgCreateDocument) Route() string {
	return RouterKey
}

func (msg *MsgCreateDocument) Type() string {
	return "CreateDocument"
}

func (msg *MsgCreateDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
