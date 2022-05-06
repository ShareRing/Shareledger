package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevokeRelayers = "revoke_relayers"

var _ sdk.Msg = &MsgRevokeRelayers{}

func NewMsgRevokeRelayers(creator string, addresses string) *MsgRevokeRelayers {
	return &MsgRevokeRelayers{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgRevokeRelayers) Route() string {
	return RouterKey
}

func (msg *MsgRevokeRelayers) Type() string {
	return TypeMsgRevokeRelayers
}

func (msg *MsgRevokeRelayers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeRelayers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeRelayers) ValidateBasic() error {
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
