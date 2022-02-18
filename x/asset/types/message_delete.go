package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteAsset{}

func NewMsgDelete(owner, UUID string) *MsgDeleteAsset {
	return &MsgDeleteAsset{
		Owner: owner,
		UUID:  UUID,
	}
}

func (msg *MsgDeleteAsset) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAsset) Type() string {
	return TypeAssetDeleteMsg
}

func (msg *MsgDeleteAsset) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgDeleteAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if len(msg.Owner) == 0 || err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	if len(msg.UUID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "UUID must not be empty")
	}

	return nil
}
