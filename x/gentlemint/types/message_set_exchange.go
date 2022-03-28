package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetExchange{}

func NewMsgSetExchange(creator string, rate string) *MsgSetExchange {
	return &MsgSetExchange{
		Creator: creator,
		Rate:    rate,
	}
}

func (msg *MsgSetExchange) Route() string {
	return RouterKey
}

func (msg *MsgSetExchange) Type() string {
	return "SetExchange"
}

func (msg *MsgSetExchange) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetExchange) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.NewDecFromStr(msg.Rate); err != nil {
		return sdkerrors.Wrapf(err, "invalid number format %v", msg.Rate)
	}
	return nil
}
