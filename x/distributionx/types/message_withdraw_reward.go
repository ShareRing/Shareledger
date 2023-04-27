package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawReward = "withdraw_reward"

var _ sdk.Msg = &MsgWithdrawReward{}

func NewMsgWithdrawReward(creator sdk.AccAddress) *MsgWithdrawReward {
	return &MsgWithdrawReward{
		Creator: creator.String(),
	}
}

func (msg *MsgWithdrawReward) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawReward) Type() string {
	return TypeMsgWithdrawReward
}

func (msg *MsgWithdrawReward) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", msg.Creator)
	}
	return nil
}
