package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEnrollIdSigner{}

func NewMsgEnrollIdSigner(creator string, addresses []string) *MsgEnrollIdSigner {
	return &MsgEnrollIdSigner{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgEnrollIdSigner) Route() string {
	return RouterKey
}

func (msg *MsgEnrollIdSigner) Type() string {
	return "EnrollIdSigner"
}

func (msg *MsgEnrollIdSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnrollIdSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnrollIdSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Addresses) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "addresses should not be empty")
	}
	for _, a := range msg.Addresses {
		_, err := sdk.AccAddressFromBech32(a)
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid id signer address (%s)", err)
	}

	return nil
}
