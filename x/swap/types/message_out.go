package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

const TypeMsgRequestOut = "out"

var _ sdk.Msg = &MsgRequestOut{}

func NewMsgRequestOut(creator string, destAddr string, network string, amount sdk.DecCoin, fee sdk.DecCoin) *MsgRequestOut {
	return &MsgRequestOut{
		Creator:     creator,
		DestAddress: destAddr,
		Network:     network,
		Amount:      &amount,
		Fee:         &fee,
	}
}

func (msg *MsgRequestOut) Route() string {
	return RouterKey
}

func (msg *MsgRequestOut) Type() string {
	return TypeMsgRequestOut
}

func (msg *MsgRequestOut) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestOut) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestOut) ValidateBasic() error {
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
	coin, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.GetAmount()), false)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if coin.AmountOf(denom.Base).IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "only nshr or shr is supported")
	}

	if _, f := SupportedSwapOutNetwork[msg.GetNetwork()]; !f {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "network is unsupported")
	}

	return nil
}
