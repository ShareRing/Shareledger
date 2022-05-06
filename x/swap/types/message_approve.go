package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApprove = "approve"

var _ sdk.Msg = &MsgApprove{}

func NewMsgApprove(creator string, signedHash string, ids []uint64) *MsgApprove {
	return &MsgApprove{
		Creator:   creator,
		Signature: signedHash,
		Ids:       ids,
	}
}

func (msg *MsgApprove) Route() string {
	return RouterKey
}

func (msg *MsgApprove) Type() string {
	return TypeMsgApprove
}

func (msg *MsgApprove) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApprove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApprove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Ids) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction ids are required. Supported format input: ID1,ID2,ID3")
	}

	return nil
}
