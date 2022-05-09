package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveIn = "approve_in"

var _ sdk.Msg = &MsgApproveIn{}

func NewMsgApproveIn(creator string, ids []uint64) *MsgApproveIn {
	return &MsgApproveIn{
		Creator: creator,
		Ids:     ids,
	}
}

func (msg *MsgApproveIn) Route() string {
	return RouterKey
}

func (msg *MsgApproveIn) Type() string {
	return TypeMsgApproveIn
}

func (msg *MsgApproveIn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveIn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveIn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.GetIds()) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "transaction ids list can't empty")
	}
	return nil
}
