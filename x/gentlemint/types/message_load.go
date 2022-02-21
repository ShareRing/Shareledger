package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

var _ sdk.Msg = &MsgLoad{}

func NewMsgLoad(creator string, address string, coins sdk.DecCoins) *MsgLoad {
	return &MsgLoad{
		Creator: creator,
		Address: address,
		Coins:   coins,
	}
}

func (msg *MsgLoad) Route() string {
	return RouterKey
}

func (msg *MsgLoad) Type() string {
	return "Load"
}

func (msg *MsgLoad) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLoad) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLoad) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid destination address (%s)", err)
	}

	if err := msg.Coins.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "coins string was not supported. Format should be {amount0}{denomination},...,{amountN}{denominationN}")
	}
	if err := denom.CheckSupportedCoins(msg.Coins, nil); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, err.Error())
	}
	return nil
}
