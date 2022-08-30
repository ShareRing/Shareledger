package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

const TypeMsgDeposit = "deposit"

var _ sdk.Msg = &MsgDeposit{}

func NewMsgDeposit(creator string, amount *sdk.DecCoin) *MsgDeposit {
	return &MsgDeposit{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgDeposit) Route() string {
	return RouterKey
}

func (msg *MsgDeposit) Type() string {
	return TypeMsgDeposit
}

func (msg *MsgDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !denom.IsShrOrBase(*msg.Amount) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "the deposit must be shr or nshr (%s)", err)
	}
	if msg.GetAmount().Amount.IsZero() || msg.GetAmount().Amount.IsNegative() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "the deposit must be greater than 0", err)
	}
	return nil
}
