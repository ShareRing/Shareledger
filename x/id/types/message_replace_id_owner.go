package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgReplaceIdOwner{}

func NewMsgReplaceIdOwner(backupAddress string, id string, ownerAddress string) *MsgReplaceIdOwner {
	return &MsgReplaceIdOwner{
		BackupAddress: backupAddress,
		Id:            id,
		OwnerAddress:  ownerAddress,
	}
}

func (msg *MsgReplaceIdOwner) Route() string {
	return RouterKey
}

func (msg *MsgReplaceIdOwner) Type() string {
	return TypeMsgReplaceIdOwner
}

func (msg *MsgReplaceIdOwner) GetSigners() []sdk.AccAddress {
	backupAddress, err := sdk.AccAddressFromBech32(msg.BackupAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{backupAddress}
}

func (msg *MsgReplaceIdOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReplaceIdOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.BackupAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid backupAddress address (%s)", err)
	}

	if len(msg.Id) > MAX_ID_LEN || len(msg.Id) == 0 {
		return sdkerrors.Wrap(InvalidData, msg.Id)
	}

	_, err = sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid OwnerAddress address (%s)", err)
	}
	return nil
}
