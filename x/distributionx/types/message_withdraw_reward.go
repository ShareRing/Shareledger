package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgWithdrawReward = "withdraw_reward"

var _ sdk.Msg = &MsgWithdrawReward{}

func NewMsgWithdrawReward(creator string) *MsgWithdrawReward {
	return &MsgWithdrawReward{}
}

func (msg *MsgWithdrawReward) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawReward) Type() string {
	return TypeMsgWithdrawReward
}

func (msg *MsgWithdrawReward) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgWithdrawReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawReward) ValidateBasic() error {
	return nil
}
