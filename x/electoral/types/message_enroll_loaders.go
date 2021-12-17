package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEnrollLoaders{}

func NewMsgEnrollLoaders(creator string, addresses []string) *MsgEnrollLoaders {
	return &MsgEnrollLoaders{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgEnrollLoaders) Route() string {
	return RouterKey
}

func (msg *MsgEnrollLoaders) Type() string {
	return "EnrollLoaders"
}

func (msg *MsgEnrollLoaders) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnrollLoaders) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnrollLoaders) ValidateBasic() error {
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
