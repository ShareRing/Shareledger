package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateBatch = "update_batch"

var _ sdk.Msg = &MsgCompleteBatch{}

func NewMsgCompleteBatch(creator string, batchId uint64) *MsgCompleteBatch {
	return &MsgCompleteBatch{
		Creator: creator,
		BatchId: batchId,
	}
}

func (msg *MsgCompleteBatch) Route() string {
	return RouterKey
}

func (msg *MsgCompleteBatch) Type() string {
	return TypeMsgUpdateBatch
}

func (msg *MsgCompleteBatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCompleteBatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCompleteBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
