package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateActionLevelFee{}

func NewMsgCreateActionLevelFee(
    creator string,
    action string,
    level string,
    
) *MsgCreateActionLevelFee {
  return &MsgCreateActionLevelFee{
		Creator : creator,
		Action: action,
		Level: level,
        
	}
}

func (msg *MsgCreateActionLevelFee) Route() string {
  return RouterKey
}

func (msg *MsgCreateActionLevelFee) Type() string {
  return "CreateActionLevelFee"
}

func (msg *MsgCreateActionLevelFee) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateActionLevelFee) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateActionLevelFee) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateActionLevelFee{}

func NewMsgUpdateActionLevelFee(
    creator string,
    action string,
    level string,
    
) *MsgUpdateActionLevelFee {
  return &MsgUpdateActionLevelFee{
		Creator: creator,
        Action: action,
        Level: level,
        
	}
}

func (msg *MsgUpdateActionLevelFee) Route() string {
  return RouterKey
}

func (msg *MsgUpdateActionLevelFee) Type() string {
  return "UpdateActionLevelFee"
}

func (msg *MsgUpdateActionLevelFee) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateActionLevelFee) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateActionLevelFee) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgDeleteActionLevelFee{}

func NewMsgDeleteActionLevelFee(
    creator string,
    action string,
    
) *MsgDeleteActionLevelFee {
  return &MsgDeleteActionLevelFee{
		Creator: creator,
		Action: action,
        
	}
}
func (msg *MsgDeleteActionLevelFee) Route() string {
  return RouterKey
}

func (msg *MsgDeleteActionLevelFee) Type() string {
  return "DeleteActionLevelFee"
}

func (msg *MsgDeleteActionLevelFee) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteActionLevelFee) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteActionLevelFee) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}