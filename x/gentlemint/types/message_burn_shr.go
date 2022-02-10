package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurnPShr{}

func NewMsgBurnShr(creator string, amount string) *MsgBurnPShr {
	return &MsgBurnPShr{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBurnPShr) Route() string {
	return RouterKey
}

func (msg *MsgBurnPShr) Type() string {
	return "BurnPShr"
}

func (msg *MsgBurnPShr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnPShr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnPShr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if v, ok := sdk.NewIntFromString(msg.Amount); !ok || v.LTE(sdk.NewInt(0)) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid amount, %v", msg.Amount)
	}
	return nil
}
