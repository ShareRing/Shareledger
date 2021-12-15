package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateIdBatch{}

func NewMsgCreateIdBatch(issuerAddress string, backupAddress string, extraData string, id string, ownerAddress string) *MsgCreateIdBatch {
	return &MsgCreateIdBatch{
		IssuerAddress: issuerAddress,
		BackupAddress: backupAddress,
		ExtraData:     extraData,
		Id:            id,
		OwnerAddress:  ownerAddress,
	}
}

func (msg *MsgCreateIdBatch) Route() string {
	return RouterKey
}

func (msg *MsgCreateIdBatch) Type() string {
	return "CreateIdBatch"
}

func (msg *MsgCreateIdBatch) GetSigners() []sdk.AccAddress {
	issuerAddress, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{issuerAddress}
}

func (msg *MsgCreateIdBatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIdBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.IssuerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuerAddress address (%s)", err)
	}
	return nil
}
