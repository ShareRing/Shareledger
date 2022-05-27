package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

const TypeMsgRequestIn = "in"

var _ sdk.Msg = &MsgRequestIn{}

func NewMsgRequestIn(creator string, desAddress string, srcNetwork string, txHashes []string, amount, fee sdk.DecCoin) *MsgRequestIn {
	return &MsgRequestIn{
		Creator:     creator,
		DestAddress: desAddress,
		Network:     srcNetwork,
		Amount:      &amount,
		TxHashes:    txHashes,
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
	if msg.Amount == nil || msg.Amount.Validate() != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid swap amount")
	}
	if msg.Fee == nil || msg.Fee.Validate() != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid swap fee")
	}

	if _, err := sdk.AccAddressFromBech32(msg.GetDestAddress()); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	if len(msg.TxHashes) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "tx hashes are required")
	}

	if strings.TrimSpace(msg.GetNetwork()) == NetworkNameShareLedger {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid network %s", msg.GetDestAddress())
	}

	return nil
}
