package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateSwapFee = "update_swap_fee"

var _ sdk.Msg = &MsgUpdateSwapFee{}

func NewMsgUpdateSwapFee(creator, network string, in, out *sdk.DecCoin) *MsgUpdateSwapFee {
	return &MsgUpdateSwapFee{
		Creator: creator,
		In:      in,
		Out:     out,
		Network: network,
	}
}

func (msg *MsgUpdateSwapFee) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSwapFee) Type() string {
	return TypeMsgUpdateSwapFee
}

func (msg *MsgUpdateSwapFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSwapFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSwapFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
