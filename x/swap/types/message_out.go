package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

const TypeMsgOut = "out"

var _ sdk.Msg = &MsgOut{}

func NewMsgOut(creator string, destAddr string, network string, amount sdk.DecCoin, fee sdk.DecCoin) *MsgOut {
	return &MsgOut{
		Creator:  creator,
		DestAddr: destAddr,
		Network:  network,
		Amount:   &amount,
		Fee:      &fee,
	}
}

func (msg *MsgOut) Route() string {
	return RouterKey
}

func (msg *MsgOut) Type() string {
	return TypeMsgOut
}

func (msg *MsgOut) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOut) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOut) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Amount == nil || msg.Fee == nil || len(msg.Network) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid valid amount,%s, or fee, %s, or network, %s", msg.Amount, msg.Fee, msg.Network)
	}

	if err := denom.CheckSupportedCoins(sdk.NewDecCoins(*msg.Amount), nil); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if err := denom.CheckSupportedCoins(sdk.NewDecCoins(*msg.Fee), nil); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}
