package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateBatch = "update_batch"

var _ sdk.Msg = &MsgSetBatchDone{}

func NewMsgUpdateBatch(creator string, batchId uint64) *MsgSetBatchDone {
	return &MsgSetBatchDone{
		Creator: creator,
		BatchId: batchId,
	}
}

func (msg *MsgSetBatchDone) Route() string {
	return RouterKey
}

func (msg *MsgSetBatchDone) Type() string {
	return TypeMsgUpdateBatch
}

func (msg *MsgSetBatchDone) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetBatchDone) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetBatchDone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
