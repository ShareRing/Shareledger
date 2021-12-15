package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateId{}

func NewMsgCreateId(issuerAddress string, backupAddress string, extraData string, id string, ownerAddress string) *MsgCreateId {
	return &MsgCreateId{
		IssuerAddress: issuerAddress,
		BackupAddress: backupAddress,
		ExtraData:     extraData,
		Id:            id,
		OwnerAddress:  ownerAddress,
	}
}

func (msg *MsgCreateId) Route() string {
	return RouterKey
}

func (msg *MsgCreateId) Type() string {
	return "CreateId"
}

func (msg *MsgCreateId) GetSigners() []sdk.AccAddress {
	issuerAddress, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{issuerAddress}
}

func (msg *MsgCreateId) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateId) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuerAddress address (%s)", err)
	}
	return nil
}
