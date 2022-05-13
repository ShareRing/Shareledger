package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelBatches = "cancel_batches"

var _ sdk.Msg = &MsgCancelBatches{}

func NewMsgCancelBatches(creator string, ids []uint64) *MsgCancelBatches {
	return &MsgCancelBatches{
		Creator: creator,
		Ids:     ids,
	}
}

func (msg *MsgCancelBatches) Route() string {
	return RouterKey
}

func (msg *MsgCancelBatches) Type() string {
	return TypeMsgCancelBatches
}

func (msg *MsgCancelBatches) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelBatches) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelBatches) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
