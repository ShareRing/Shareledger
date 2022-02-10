package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendPShr{}

func NewMsgSendShr(creator string, address string, amount string) *MsgSendPShr {
	return &MsgSendPShr{
		Creator: creator,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgSendPShr) Route() string {
	return RouterKey
}

func (msg *MsgSendPShr) Type() string {
	return "SendPShr"
}

func (msg *MsgSendPShr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendPShr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendPShr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}
	if _, ok := sdk.NewIntFromString(msg.Amount); !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, msg.Amount)
	}
	return nil
}
