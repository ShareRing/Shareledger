package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

var _ sdk.Msg = &MsgSetLevelFee{}

func NewMsgSetLevelFee(
	creator string,
	level string,
	fee string,

) *MsgSetLevelFee {
	return &MsgSetLevelFee{
		Creator: creator,
		Level:   level,
		Fee:     fee,
	}
}

func (msg *MsgSetLevelFee) Route() string {
	return RouterKey
}

func (msg *MsgSetLevelFee) Type() string {
	return "CreateLevelFee"
}

func (msg *MsgSetLevelFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetLevelFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetLevelFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	dc, err := sdk.ParseDecCoin(msg.Fee)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid fee, %s. Format should be 2pshr or 2.3shrp", msg.Fee)
	}
	if dc.Denom != denom.PShr && dc.Denom != denom.ShrP {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid token type, %s. Only support pshr or shrp", dc.Denom)
	}
	if dc.Denom == denom.PShr && !dc.Amount.Equal(dc.Amount.RoundInt().ToDec()) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "pshr amount should be int, %v", dc)
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
