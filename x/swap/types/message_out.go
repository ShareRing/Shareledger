package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

const TypeMsgSwapOut = "out"

var _ sdk.Msg = &MsgSwapOut{}

func NewMsgSwapOut(creator string, destAddr string, network string, amount sdk.DecCoin, fee sdk.DecCoin) *MsgSwapOut {
	return &MsgSwapOut{
		Creator:     creator,
		DestAddress: destAddr,
		Network:     network,
		Amount:      &amount,
		Fee:         &fee,
	}
}

func (msg *MsgSwapOut) Route() string {
	return RouterKey
}

func (msg *MsgSwapOut) Type() string {
	return TypeMsgSwapOut
}

func (msg *MsgSwapOut) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSwapOut) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSwapOut) ValidateBasic() error {
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
