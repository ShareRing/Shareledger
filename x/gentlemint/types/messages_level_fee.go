package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateLevelFee{}

func NewMsgCreateLevelFee(
	creator string,
	level string,
	fee string,

) *MsgCreateLevelFee {
	return &MsgCreateLevelFee{
		Creator: creator,
		Level:   level,
		Fee:     fee,
	}
}

func (msg *MsgCreateLevelFee) Route() string {
	return RouterKey
}

func (msg *MsgCreateLevelFee) Type() string {
	return "CreateLevelFee"
}

func (msg *MsgCreateLevelFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLevelFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLevelFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	dc, err := sdk.ParseDecCoin(msg.Fee)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid fee, %s. Format should be 2shr or 2.3shrp", msg.Fee)
	}
	if dc.Denom != DenomSHR && dc.Denom != DenomSHRP {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid token type, %s. Only support shr or shrp", dc.Denom)
	}
	if dc.Denom == DenomSHR && !dc.Amount.Equal(dc.Amount.RoundInt().ToDec()) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "shr amount should be int, %v", dc)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateLevelFee{}

func NewMsgUpdateLevelFee(
	creator string,
	level string,
	fee string,

) *MsgUpdateLevelFee {
	return &MsgUpdateLevelFee{
		Creator: creator,
		Level:   level,
		Fee:     fee,
	}
}

func (msg *MsgUpdateLevelFee) Route() string {
	return RouterKey
}

func (msg *MsgUpdateLevelFee) Type() string {
	return "UpdateLevelFee"
}

func (msg *MsgUpdateLevelFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateLevelFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateLevelFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	dc, err := sdk.ParseDecCoin(msg.Fee)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid fee, %s. Format should be <amount><denom>", msg.Fee)
	}
	if dc.Denom != DenomSHR && dc.Denom != DenomSHRP {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid token type, %s. Only support shr or shrp", dc.Denom)
	}
	if dc.Denom == DenomSHR && !dc.Amount.Equal(dc.Amount.RoundInt().ToDec()) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "shr amount should be int, %v", dc)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteLevelFee{}

func NewMsgDeleteLevelFee(
	creator string,
	level string,

) *MsgDeleteLevelFee {
	return &MsgDeleteLevelFee{
		Creator: creator,
		Level:   level,
	}
}
func (msg *MsgDeleteLevelFee) Route() string {
	return RouterKey
}

func (msg *MsgDeleteLevelFee) Type() string {
	return "DeleteLevelFee"
}

func (msg *MsgDeleteLevelFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteLevelFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteLevelFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
