package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateIds{}

func NewMsgCreateIds(issuerAddress string, backupAddress []string, extraData []string, id []string, ownerAddress []string) *MsgCreateIds {
	return &MsgCreateIds{
		IssuerAddress: issuerAddress,
		BackupAddress: backupAddress,
		ExtraData:     extraData,
		Id:            id,
		OwnerAddress:  ownerAddress,
	}
}

func (msg *MsgCreateIds) Route() string {
	return RouterKey
}

func (msg *MsgCreateIds) Type() string {
	return TypeMsgCreateIDs
}

func (msg *MsgCreateIds) GetSigners() []sdk.AccAddress {
	issuerAddress, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{issuerAddress}
}

func (msg *MsgCreateIds) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIds) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuerAddress address (%s)", err)
	}

	// check len
	if len(msg.Id) == 0 || len(msg.BackupAddress) == 0 || len(msg.OwnerAddress) == 0 || len(msg.ExtraData) == 0 {
		return sdkerrors.Wrap(InvalidData, msg.String())
	}

	maxLen := len(msg.Id)
	if maxLen > MAX_ID_IN_BATCH || len(msg.BackupAddress) != maxLen || len(msg.OwnerAddress) != maxLen || len(msg.ExtraData) != maxLen {
		return sdkerrors.Wrap(InvalidData, msg.String())
	}

	// Check id compose
	for i := 0; i < maxLen; i++ {
		// Check len
		if len(msg.Id[i]) > MAX_ID_LEN || len(msg.Id[i]) == 0 || len(msg.ExtraData[i]) > MAX_ID_LEN {
			return sdkerrors.Wrap(InvalidData, msg.String())
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
				return sdkerrors.Wrap(InvalidData, msg.String())
			}
		}
	}

	return nil
}
