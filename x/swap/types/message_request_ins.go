package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestIns = "request_ins"

var _ sdk.Msg = &MsgRequestIns{}

func NewMsgRequestIns(creator string) *MsgRequestIns {
	return &MsgRequestIns{
		Creator: creator,
	}
}

func (msg *MsgRequestIns) Route() string {
	return RouterKey
}

func (msg *MsgRequestIns) Type() string {
	return TypeMsgRequestIns
}

func (msg *MsgRequestIns) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestIns) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestIns) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
