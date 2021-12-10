package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurnShr{}

func NewMsgBurnShr(creator string, amount string) *MsgBurnShr {
	return &MsgBurnShr{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBurnShr) Route() string {
	return RouterKey
}

func (msg *MsgBurnShr) Type() string {
	return "BurnShr"
}

func (msg *MsgBurnShr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnShr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnShr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if v, ok := sdk.NewIntFromString(msg.Amount); !ok || v.LTE(sdk.NewInt(0)) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid amount, %v", msg.Amount)
	}
	return nil
}
