package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"strings"
)

const TypeMsgWithdraw = "withdraw"

var _ sdk.Msg = &MsgWithdraw{}

func NewMsgWithdraw(creator string, receiver string, amount *sdk.DecCoin) *MsgWithdraw {
	return &MsgWithdraw{
		Creator:  creator,
		Receiver: receiver,
		Amount:   amount,
	}
}

func (msg *MsgWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgWithdraw) Type() string {
	return TypeMsgWithdraw
}

func (msg *MsgWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	if strings.ToLower(msg.Amount.GetDenom()) != denom.Shr {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "withdraw message invalid must be shr (%s)", err)
	}
	if msg.GetAmount().Amount.IsZero() || msg.GetAmount().Amount.IsNegative() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "withdraw message amount must be greater than 0", err)
	}
	return nil
}
