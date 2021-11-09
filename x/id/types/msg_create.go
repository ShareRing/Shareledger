package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgCreateId(issuerAddr, backupAddr, ownerAddr sdk.AccAddress, id, extraData string) *MsgCreateId {
	return &MsgCreateId{
		IssuerAddress: issuerAddr.String(),
		BackupAddress: backupAddr.String(),
		OwnerAddress:  ownerAddr.String(),
		ExtraData:     extraData,
		Id:            id,
	}
}

func (msg MsgCreateId) Route() string {
	return RouterKey
}

func (msg MsgCreateId) Type() string {
	return TypeMsgCreateID
}

func (msg MsgCreateId) ValidateBasic() error {
	if len(msg.Id) > MAX_ID_LEN || len(msg.Id) == 0 || len(msg.ExtraData) > MAX_ID_LEN {
		return InvalidData
	}

	_, err := sdk.AccAddressFromBech32(msg.BackupAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Backup address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Issuer address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid OwnerAddress address (%s)", err)
	}

	return nil
}

func (msg MsgCreateId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateId) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.IssuerAddress)
	return []sdk.AccAddress{addr}
}

func (msg MsgCreateId) ToBaseID() BaseID {
	baseId := BaseID{IssuerAddress: msg.IssuerAddress, BackupAddress: msg.BackupAddress, OwnerAddress: msg.OwnerAddress, ExtraData: msg.ExtraData}
	return baseId
}

func (msg MsgCreateId) ToID() ID {
	baseId := msg.ToBaseID()
	id := ID{Id: msg.Id, Data: &baseId}
	return id
}
