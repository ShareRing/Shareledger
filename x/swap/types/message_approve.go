package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApprove = "approve"

var _ sdk.Msg = &MsgApproveOut{}

func NewMsgApprove(creator string, signedHash string, ids []uint64) *MsgApproveOut {
	return &MsgApproveOut{
		Creator:   creator,
		Signature: signedHash,
		Ids:       ids,
	}
}

func (msg *MsgApproveOut) Route() string {
	return RouterKey
}

func (msg *MsgApproveOut) Type() string {
	return TypeMsgApprove
}

func (msg *MsgApproveOut) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveOut) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveOut) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Ids) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction ids are required. Supported format input: ID1,ID2,ID3")
	}

	return nil
}
