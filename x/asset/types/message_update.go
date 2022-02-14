package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateAsset{}

func NewMsgUpdate(creator string, hash []byte, UUID string, status bool, rate int64) *MsgUpdateAsset {
	return &MsgUpdateAsset{
		Creator: creator,
		Hash:    hash,
		UUID:    UUID,
		Status:  status,
		Rate:    rate,
	}
}

func (msg *MsgUpdateAsset) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAsset) Type() string {
	return TypeAssetUpdateMsg
}

func (msg *MsgUpdateAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if len(msg.Creator) == 0 || err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "UUID must not be empty")
	}

	return nil
}
