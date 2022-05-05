package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancel = "cancel"

var _ sdk.Msg = &MsgCancel{}

func NewMsgCancel(creator string, ids []uint64) *MsgCancel {
	return &MsgCancel{
		Creator: creator,
		Ids:     ids,
	}
}

func (msg *MsgCancel) Route() string {
	return RouterKey
}

func (msg *MsgCancel) Type() string {
	return TypeMsgCancel
}

func (msg *MsgCancel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
