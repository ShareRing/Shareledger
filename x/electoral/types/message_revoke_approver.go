package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevokeApprover = "revoke_approver"

var _ sdk.Msg = &MsgRevokeApprover{}

func NewMsgRevokeApprover(creator string, addresses string) *MsgRevokeApprover {
	return &MsgRevokeApprover{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgRevokeApprover) Route() string {
	return RouterKey
}

func (msg *MsgRevokeApprover) Type() string {
	return TypeMsgRevokeApprover
}

func (msg *MsgRevokeApprover) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeApprover) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeApprover) ValidateBasic() error {
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
