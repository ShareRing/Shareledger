package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateFormat = "create_format"
	TypeMsgUpdateFormat = "update_format"
	TypeMsgDeleteFormat = "delete_format"
)

var _ sdk.Msg = &MsgCreateSignSchema{}

func NewMsgCreateFormat(
	creator string,
	network string,
	dataType string,
) *MsgCreateSignSchema {
	return &MsgCreateSignSchema{
		Creator: creator,
		Network: network,
		Schema:  dataType,
	}
}

func (msg *MsgCreateSignSchema) Route() string {
	return RouterKey
}

func (msg *MsgCreateSignSchema) Type() string {
	return TypeMsgCreateFormat
}

func (msg *MsgCreateSignSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSignSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSignSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Schema) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "data type format is required")
	}
	return nil
	//return validateEIP712Data(msg.Network, msg.DataType)
}

//func validateEIP712Data(network string, data *EIP712DataType) error {
//	if len(network) == 0 {
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network is required")
//	}
//	if data == nil {
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "eip712 data is required")
//	}
//	if len(data.PrimaryType) == 0 {
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "eip712 data primary type is required")
//	}
//	if data.Domain == nil {
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "eip712 data domain is required")
//	}
//	if len(data.Domain.Name) == 0 {
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "eip712 data domain name is required")
//	}
//	if len(data.Domain.Version) == 0 {
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "eip712 data domain version is required")
//	}
//	if len(data.Domain.ChainId) == 0 {
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "eip712 data domain chainId is required")
//	}
//	var hexOrDecimal math.HexOrDecimal256
//	if err := hexOrDecimal.UnmarshalText([]byte(data.Domain.ChainId)); err != nil {
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "eip712 data domain chainId should be hex or decimal string. %+v", err)
//	}
//	if len(data.Domain.VerifyingContract) == 0 {
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "eip712 data domain verifying contract is required")
//	}
//	return nil
//}

var _ sdk.Msg = &MsgUpdateSignSchema{}

func NewMsgUpdateFormat(
	creator string,
	network string,
	dataFormat string,
) *MsgUpdateSignSchema {
	return &MsgUpdateSignSchema{
		Creator: creator,
		Network: network,
		Schema:  dataFormat,
	}
}

func (msg *MsgUpdateSignSchema) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSignSchema) Type() string {
	return TypeMsgUpdateFormat
}

func (msg *MsgUpdateSignSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSignSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSignSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Schema) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "data type format is required")
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSignSchema{}

func NewMsgDeleteFormat(
	creator string,
	network string,

) *MsgDeleteSignSchema {
	return &MsgDeleteSignSchema{
		Creator: creator,
		Network: network,
	}
}
func (msg *MsgDeleteSignSchema) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSignSchema) Type() string {
	return TypeMsgDeleteFormat
}

func (msg *MsgDeleteSignSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSignSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSignSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
