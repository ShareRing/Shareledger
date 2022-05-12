package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevokeSwapManagers = "revoke_swap_managers"

var _ sdk.Msg = &MsgRevokeSwapManagers{}

func NewMsgRevokeSwapManagers(creator string, addresses []string) *MsgRevokeSwapManagers {
	return &MsgRevokeSwapManagers{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgRevokeSwapManagers) Route() string {
	return RouterKey
}

func (msg *MsgRevokeSwapManagers) Type() string {
	return TypeMsgRevokeSwapManagers
}

func (msg *MsgRevokeSwapManagers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeSwapManagers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeSwapManagers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
