package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReject = "reject"

var _ sdk.Msg = &MsgReject{}

func NewMsgReject(creator string, txnIDs []uint64) *MsgReject {
	return &MsgReject{
		Creator: creator,
		TxnIDs:  txnIDs,
	}
}

func (msg *MsgReject) Route() string {
	return RouterKey
}

func (msg *MsgReject) Type() string {
	return TypeMsgReject
}

func (msg *MsgReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.GetTxnIDs()) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "swap transaction list must not empty (%s)", err)
	}
	return nil
}
