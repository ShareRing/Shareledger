package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateId{}

func NewMsgUpdateId(issuerAddress string, id string, extraData string) *MsgUpdateId {
	return &MsgUpdateId{
		IssuerAddress: issuerAddress,
		Id:            id,
		ExtraData:     extraData,
	}
}

func (msg *MsgUpdateId) Route() string {
	return RouterKey
}

func (msg *MsgUpdateId) Type() string {
	return TypeMsgUpdateID
}

func (msg *MsgUpdateId) GetSigners() []sdk.AccAddress {
	issuerAddress, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{issuerAddress}
}

func (msg *MsgUpdateId) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateId) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuerAddress address (%s)", err)
	}

	if len(msg.Id) > MAX_ID_LEN || len(msg.Id) == 0 || len(msg.ExtraData) > MAX_ID_LEN {
		return sdkerrors.Wrap(InvalidData, msg.String())
	}

	return nil
}
