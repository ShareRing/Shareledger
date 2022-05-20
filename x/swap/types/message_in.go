package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

const TypeMsgRequestIn = "in"

var _ sdk.Msg = &MsgRequestIn{}

func NewMsgRequestIn(creator string, srcAddress string, desAddress string, srcNetwork string, amount, fee sdk.DecCoin) *MsgRequestIn {
	return &MsgRequestIn{
		Creator:     creator,
		SrcAddress:  srcAddress,
		DestAddress: desAddress,
		Network:     srcNetwork,
		Amount:      &amount,
		Fee:         &fee,
	}
}

func (msg *MsgRequestIn) Route() string {
	return RouterKey
}

func (msg *MsgRequestIn) Type() string {
	return TypeMsgRequestIn
}

func (msg *MsgRequestIn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestIn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestIn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if err := msg.GetAmount().Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid swap amount (%s)", err)
	}
	if err := msg.GetFee().Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid swap fee (%s)", err)
	}
	if strings.TrimSpace(msg.GetDestAddress()) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "destination address can't empty")
	}

	if strings.TrimSpace(msg.GetSrcAddress()) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "source address empty")
	}
	if strings.TrimSpace(msg.GetDestAddress()) != NetworkNameShareLedger {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid network %s", msg.GetDestAddress())
	}

	return nil
}
