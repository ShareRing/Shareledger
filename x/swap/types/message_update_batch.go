package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateBatch = "update_batch"

var _ sdk.Msg = &MsgUpdateBatch{}

func NewMsgUpdateBatch(creator string, batchId uint64, txHash, network, status string) *MsgUpdateBatch {
	return &MsgUpdateBatch{
		Creator: creator,
		BatchId: batchId,
		Status:  status,
		TxHash:  txHash,
		Network: network,
	}
}

func (msg *MsgUpdateBatch) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBatch) Type() string {
	return TypeMsgUpdateBatch
}

func (msg *MsgUpdateBatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Status == BatchStatusProcessing && msg.TxHash == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "the nonce number can't empty or zero in case %s status", msg.Status)
	}
	return nil
}
