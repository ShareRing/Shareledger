package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRevokeDocIssuer{}

func NewMsgRevokeDocIssuer(creator string, addresses string) *MsgRevokeDocIssuer {
	return &MsgRevokeDocIssuer{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgRevokeDocIssuer) Route() string {
	return RouterKey
}

func (msg *MsgRevokeDocIssuer) Type() string {
	return "RevokeDocIssuer"
}

func (msg *MsgRevokeDocIssuer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeDocIssuer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeDocIssuer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
