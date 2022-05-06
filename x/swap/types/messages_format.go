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

var _ sdk.Msg = &MsgCreateFormat{}

func NewMsgCreateFormat(
	creator string,
	network string,
	dataType string,
) *MsgCreateFormat {
	return &MsgCreateFormat{
		Creator:        creator,
		Network:        network,
		DataTypeFormat: dataType,
	}
}

func (msg *MsgCreateFormat) Route() string {
	return RouterKey
}

func (msg *MsgCreateFormat) Type() string {
	return TypeMsgCreateFormat
}

func (msg *MsgCreateFormat) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateFormat) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateFormat) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.DataTypeFormat) == 0 {
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

var _ sdk.Msg = &MsgUpdateFormat{}

func NewMsgUpdateFormat(
	creator string,
	network string,
	dataFormat string,
) *MsgUpdateFormat {
	return &MsgUpdateFormat{
		Creator:        creator,
		Network:        network,
		DataTypeFormat: dataFormat,
	}
}

func (msg *MsgUpdateFormat) Route() string {
	return RouterKey
}

func (msg *MsgUpdateFormat) Type() string {
	return TypeMsgUpdateFormat
}

func (msg *MsgUpdateFormat) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateFormat) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateFormat) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.DataTypeFormat) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "data type format is required")
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteFormat{}

func NewMsgDeleteFormat(
	creator string,
	network string,

) *MsgDeleteFormat {
	return &MsgDeleteFormat{
		Creator: creator,
		Network: network,
	}
}
func (msg *MsgDeleteFormat) Route() string {
	return RouterKey
}

func (msg *MsgDeleteFormat) Type() string {
	return TypeMsgDeleteFormat
}

func (msg *MsgDeleteFormat) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteFormat) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteFormat) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
