package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdate{}

func NewMsgUpdate(creator string, hash string, uUID string, status string, rate string) *MsgUpdate {
	return &MsgUpdate{
		Creator: creator,
		Hash:    hash,
		UUID:    uUID,
		Status:  status,
		Rate:    rate,
	}
}

func (msg *MsgUpdate) Route() string {
	return RouterKey
}

func (msg *MsgUpdate) Type() string {
	return "Update"
}

func (msg *MsgUpdate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
