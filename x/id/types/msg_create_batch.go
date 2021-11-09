package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// type MsgCreateIdBatch struct {
// 	BackupAddr []sdk.AccAddress `json:"backup_address"`
// 	ExtraData  []string         `json:"extra_data"`
// 	Id         []string         `json:"id"`
// 	IssuerAddr sdk.AccAddress   `json:"issuer_address"`
// 	OwnerAddr  []sdk.AccAddress `json:"owner_address"`
// }

func NewMsgCreateIdBatch(issuerAddr string, backupAddr, ownerAddr, id, extraData []string) MsgCreateIdBatch {
	return MsgCreateIdBatch{
		IssuerAddress: issuerAddr,
		BackupAddress: backupAddr,
		OwnerAddress:  ownerAddr,
		ExtraData:     extraData,
		Id:            id,
	}
}

func (msg MsgCreateIdBatch) Route() string {
	return RouterKey
}

func (msg MsgCreateIdBatch) Type() string {
	return TypeMsgCreateIDBatch
}

func (msg MsgCreateIdBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Issuer address (%s)", err)
	}

	// Check len
	if len(msg.Id) == 0 || len(msg.BackupAddress) == 0 || len(msg.OwnerAddress) == 0 || len(msg.ExtraData) == 0 {
		return InvalidData
	}

	maxLen := len(msg.Id)
	if maxLen > MAX_ID_IN_BATCH || len(msg.BackupAddress) != maxLen || len(msg.OwnerAddress) != maxLen || len(msg.ExtraData) != maxLen {
		return InvalidData
	}

	// Check id compose
	for i := 0; i < maxLen; i++ {
		// Check len
		if len(msg.Id[i]) > MAX_ID_LEN || len(msg.Id[i]) == 0 || len(msg.ExtraData[i]) > MAX_ID_LEN {
			return InvalidData
		}

		// Check address
		_, err := sdk.AccAddressFromBech32(msg.OwnerAddress[i])
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid OwnerAddress address (%s)", err)
		}
		_, err = sdk.AccAddressFromBech32(msg.BackupAddress[i])
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid BackupAddress address (%s)", err)
		}

		// Check duplicate
		for j := i + 1; j < maxLen; j++ {
			if msg.Id[j] == msg.Id[i] || msg.OwnerAddress[j] == msg.OwnerAddress[i] {
				return InvalidData
			}
		}
	}

	return nil
}

func (msg MsgCreateIdBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateIdBatch) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.IssuerAddress)
	return []sdk.AccAddress{signer}
}
