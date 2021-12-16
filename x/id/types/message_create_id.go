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
	return TypeMsgCreateID
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

	if len(msg.Id) > MAX_ID_LEN || len(msg.Id) == 0 || len(msg.ExtraData) > MAX_ID_LEN {
		return sdkerrors.Wrap(InvalidData, msg.String())
	}

	_, err = sdk.AccAddressFromBech32(msg.BackupAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Backup address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid OwnerAddress address (%s)", err)
	}

	return nil
}

func (msg *MsgCreateId) ToBaseID() BaseID {
	baseID := BaseID{
		IssuerAddress: msg.IssuerAddress,
		BackupAddress: msg.BackupAddress,
		OwnerAddress:  msg.OwnerAddress,
		ExtraData:     msg.ExtraData,
	}

	return baseID
}

func (msg *MsgCreateId) ToID() Id {
	baseID := msg.ToBaseID()
	id := Id{
		Id:   msg.Id,
		Data: &baseID,
	}

	return id
}
