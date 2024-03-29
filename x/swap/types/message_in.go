package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestIn = "in"

var _ sdk.Msg = &MsgRequestIn{}

func NewMsgRequestIn(creator, desAddress, srcNetwork string, txEvents []*TxEvent, amount, fee sdk.DecCoin) *MsgRequestIn {
	return &MsgRequestIn{
		Creator:     creator,
		SrcAddress:  creator,
		DestAddress: desAddress,
		Network:     srcNetwork,
		Amount:      &amount,
		TxEvents:    txEvents,
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
	if msg.Amount == nil || msg.Amount.Validate() != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid swap amount")
	}

	if _, err := sdk.AccAddressFromBech32(msg.GetDestAddress()); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	if len(msg.TxEvents) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "tx hashes are required")
	}

	var checkMap = make(map[string]bool)
	for _, h := range msg.TxEvents {
		key := fmt.Sprintf("%s/%d", h.TxHash, h.LogIndex)
		_, found := checkMap[key]
		if !found {
			checkMap[key] = true
		} else {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "tx events has duplicate request: %+v", h)
		}

		if h.TxHash == "" {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "tx hashes are required")
		}
		if h.Sender == "" {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "sender is required")
		}

	}

	if strings.TrimSpace(msg.GetNetwork()) == NetworkNameShareLedger {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid network %s", msg.GetDestAddress())
	}

	return nil
}
