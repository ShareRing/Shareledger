package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEnrollDocIssuer{}

func NewMsgEnrollDocIssuer(creator string, addresses string) *MsgEnrollDocIssuer {
	return &MsgEnrollDocIssuer{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgEnrollDocIssuer) Route() string {
	return RouterKey
}

func (msg *MsgEnrollDocIssuer) Type() string {
	return "EnrollDocIssuer"
}

func (msg *MsgEnrollDocIssuer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnrollDocIssuer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnrollDocIssuer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
