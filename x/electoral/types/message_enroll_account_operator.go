package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEnrollAccountOperator{}

func NewMsgEnrollAccountOperator(creator string, addresses string) *MsgEnrollAccountOperator {
	return &MsgEnrollAccountOperator{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgEnrollAccountOperator) Route() string {
	return RouterKey
}

func (msg *MsgEnrollAccountOperator) Type() string {
	return "EnrollAccountOperator"
}

func (msg *MsgEnrollAccountOperator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnrollAccountOperator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnrollAccountOperator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
