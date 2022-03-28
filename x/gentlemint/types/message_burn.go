package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

var _ sdk.Msg = &MsgBurn{}

func NewMsgBurn(creator string, coins sdk.DecCoins) *MsgBurn {
	return &MsgBurn{
		Creator: creator,
		Coins:   coins,
	}
}

func (msg *MsgBurn) Route() string {
	return RouterKey
}

func (msg *MsgBurn) Type() string {
	return "Burn"
}

func (msg *MsgBurn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := msg.Coins.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "coins string was not supported. Format should be {amount0}{denomination},...,{amountN}{denominationN}")
	}
	if err := denom.CheckSupportedCoins(msg.Coins, nil); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, err.Error())
	}
	return nil
}
