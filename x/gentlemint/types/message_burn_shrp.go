package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurnShrp{}

func NewMsgBurnShrp(creator string, amount string) *MsgBurnShrp {
	return &MsgBurnShrp{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBurnShrp) Route() string {
	return RouterKey
}

func (msg *MsgBurnShrp) Type() string {
	return "BurnShrp"
}

func (msg *MsgBurnShrp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnShrp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnShrp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := ParseShrpCoinsStr(msg.Amount); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "parsing amount %s", err.Error())
	}
	return nil
}
