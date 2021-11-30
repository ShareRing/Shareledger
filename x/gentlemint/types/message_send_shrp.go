package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendShrp{}

func NewMsgSendShrp(creator string, address string, amount string) *MsgSendShrp {
	return &MsgSendShrp{
		Creator: creator,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgSendShrp) Route() string {
	return RouterKey
}

func (msg *MsgSendShrp) Type() string {
	return "SendShrp"
}

func (msg *MsgSendShrp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendShrp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendShrp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
