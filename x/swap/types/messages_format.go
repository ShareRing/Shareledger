package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

const (
	TypeMsgCreateFormat = "create_format"
	TypeMsgUpdateFormat = "update_format"
	TypeMsgDeleteFormat = "delete_format"
)

var _ sdk.Msg = &MsgCreateSchema{}

func NewMsgCreateSchema(
	creator string,
	network string,
	dataType string,
	inFee, outFee sdk.DecCoin,
	contractExponent int32,
) *MsgCreateSchema {
	return &MsgCreateSchema{
		Creator:          creator,
		Network:          network,
		Schema:           dataType,
		In:               inFee,
		Out:              outFee,
		ContractExponent: contractExponent,
	}
}

func (msg *MsgCreateSchema) Route() string {
	return RouterKey
}

func (msg *MsgCreateSchema) Type() string {
	return TypeMsgCreateFormat
}

func (msg *MsgCreateSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Schema) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "data type format is required")
	}

	if err := msg.Out.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid swap out fee (%s)", err)
	}
	if err := msg.In.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid swap in fee (%s)", err)
	}
	if !denom.IsShrOrBase(msg.In) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid swap in fee")
	}
	if !denom.IsShrOrBase(msg.Out) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid swap out fee")
	}
	return nil

}

var _ sdk.Msg = &MsgUpdateSchema{}

func NewMsgUpdateSchema(
	creator string,
	network string,
	dataFormat string,
	in, out *sdk.DecCoin,
	exp int32,
) *MsgUpdateSchema {
	return &MsgUpdateSchema{
		Creator:          creator,
		Network:          network,
		Schema:           dataFormat,
		In:               in,
		Out:              out,
		ContractExponent: exp,
	}
}

func (msg *MsgUpdateSchema) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSchema) Type() string {
	return TypeMsgUpdateFormat
}

func (msg *MsgUpdateSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Network == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network name is required")
	}

	if (msg.In == nil || msg.In.IsZero()) &&
		(msg.Out == nil || msg.Out.IsZero()) &&
		msg.Schema == "" &&
		msg.ContractExponent == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "update schema should require at least one params on the list of [fee-in, fee-out, exp, schema]")
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteSchema{}

func NewMsgDeleteSchema(
	creator string,
	network string,

) *MsgDeleteSchema {
	return &MsgDeleteSchema{
		Creator: creator,
		Network: network,
	}
}
func (msg *MsgDeleteSchema) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSchema) Type() string {
	return TypeMsgDeleteFormat
}

func (msg *MsgDeleteSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Network == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network name is required")
	}
	return nil
}
