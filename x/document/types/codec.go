package types

import (
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgCreateDoc{}, fmt.Sprintf("%s/%s", ModuleName, TypeMsgCreateDoc), nil)
	cdc.RegisterConcrete(MsgCreateDocBatch{}, fmt.Sprintf("%s/%s", ModuleName, TypeMsgCreateDocInBatch), nil)
	cdc.RegisterConcrete(MsgUpdateDoc{}, fmt.Sprintf("%s/%s", ModuleName, TypeMsgUpdateDoc), nil)
	cdc.RegisterConcrete(MsgRevokeDoc{}, fmt.Sprintf("%s/%s", ModuleName, TypeMsgRevokeDoc), nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDoc{},
		&MsgCreateDocBatch{},
		&MsgUpdateDoc{},
		&MsgRevokeDoc{},
	)

	// registry.RegisterInterface(
	// 	"cosmos.bank.v1beta1.SupplyI",
	// 	(*exported.SupplyI)(nil),
	// 	&Supply{},
	// )

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
