package types

import (
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLoadShrp{}

func NewMsgLoadShrp(creator string, address string, amount string) *MsgLoadShrp {
	return &MsgLoadShrp{
		Creator: creator,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgLoadShrp) Route() string {
	return RouterKey
}

func (msg *MsgLoadShrp) Type() string {
	return "LoadShrp"
}

func (msg *MsgLoadShrp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLoadShrp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLoadShrp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	f, err := strconv.ParseFloat(msg.Amount, 64)
	if err != nil {
		return err
	}
	if f < 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "amount should be larger than 0")
	}

	numbers := strings.Split(msg.Amount, ".")
	if len(numbers) == 2 {
		decimalNumber, err := strconv.ParseInt(numbers[1], 10, 64)
		if err != nil {
			return err
		}
		if decimalNumber > 99 {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "cent amount should be less than or equal to 99")
		}
	}

	return nil
}
