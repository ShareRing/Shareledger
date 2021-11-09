package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgReplaceIdOwner(id, ownerAddr, backupAddr string) MsgReplaceIdOwner {
	return MsgReplaceIdOwner{
		Id:            id,
		BackupAddress: backupAddr,
		OwnerAddress:  ownerAddr,
	}
}

func (msg MsgReplaceIdOwner) Route() string {
	return RouterKey
}

func (msg MsgReplaceIdOwner) Type() string {
	return TypeMsgReplaceIdOwner
}

func (msg MsgReplaceIdOwner) ValidateBasic() error {
	if len(msg.Id) > MAX_ID_LEN || len(msg.Id) == 0 {
		return InvalidData
	}

	// Check address
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid OwnerAddress address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.BackupAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid BackupAddress address (%s)", err)
	}
	return nil
}

func (msg MsgReplaceIdOwner) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgReplaceIdOwner) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.BackupAddress)
	return []sdk.AccAddress{signer}
}
