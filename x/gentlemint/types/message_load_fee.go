package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLoadFee{}

func NewMsgLoadFee(creator string, shrp sdk.DecCoin) *MsgLoadFee {
	return &MsgLoadFee{
		Creator: creator,
		Shrp:    &shrp,
	}
}

func (msg *MsgLoadFee) Route() string {
	return RouterKey
}

func (msg *MsgLoadFee) Type() string {
	return "LoadFee"
}

func (msg *MsgLoadFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLoadFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLoadFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
