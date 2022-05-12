package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnrollSwapManagers = "enroll_swap_manager"

var _ sdk.Msg = &MsgEnrollSwapManagers{}

func NewMsgEnrollSwapManagers(creator string, addresses []string) *MsgEnrollSwapManagers {
	return &MsgEnrollSwapManagers{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgEnrollSwapManagers) Route() string {
	return RouterKey
}

func (msg *MsgEnrollSwapManagers) Type() string {
	return TypeMsgEnrollSwapManagers
}

func (msg *MsgEnrollSwapManagers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnrollSwapManagers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnrollSwapManagers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.GetAddresses()) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid enroll address can't empty")
	}
	return nil
}
