package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

const TypeMsgIn = "in"

var _ sdk.Msg = &MsgIn{}

func NewMsgIn(creator string, srcAddress string, desAddress string, srcNetwork string, amount, fee sdk.DecCoin) *MsgIn {
	return &MsgIn{
		Creator:     creator,
		SrcAddress:  srcAddress,
		DestAddress: desAddress,
		SrcNetwork:  srcNetwork,
		Amount:      &amount,
		Fee:         &fee,
	}
}

func (msg *MsgIn) Route() string {
	return RouterKey
}

func (msg *MsgIn) Type() string {
	return TypeMsgIn
}

func (msg *MsgIn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIn) ValidateBasic() error {
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
	return nil
}
