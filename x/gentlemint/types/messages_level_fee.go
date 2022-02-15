package types

import (
	"fmt"
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid fee, %s. Format should be 2%v or 2.3%v", msg.Fee, denom.Base, denom.ShrP)
	}

	switch dc.Denom {
	case denom.Cent, denom.Base, denom.ShrP, denom.Shr:
	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, fmt.Sprint("invalid denominations. supported ", denom.Base, denom.Shr, denom.ShrP, denom.Cent))
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
