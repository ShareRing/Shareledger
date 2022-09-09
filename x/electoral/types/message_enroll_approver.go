package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnrollApprovers = "enroll_approver"

var _ sdk.Msg = &MsgEnrollApprovers{}

func NewMsgEnrollApprovers(creator string, addresses []string) *MsgEnrollApprovers {
	return &MsgEnrollApprovers{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgEnrollApprovers) Route() string {
	return RouterKey
}

func (msg *MsgEnrollApprovers) Type() string {
	return TypeMsgEnrollApprovers
}

func (msg *MsgEnrollApprovers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnrollApprovers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnrollApprovers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Addresses) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "addresses should not be empty")
	}
	for _, a := range msg.Addresses {
		if _, err := sdk.AccAddressFromBech32(a); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
		}
	}
	return nil
}
