package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateId{}, "id/CreateId", nil)
	cdc.RegisterConcrete(&MsgCreateIds{}, "id/CreateIds", nil)
	cdc.RegisterConcrete(&MsgUpdateId{}, "id/UpdateId", nil)
	cdc.RegisterConcrete(&MsgReplaceIdOwner{}, "id/ReplaceIdOwner", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateId{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateIds{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateId{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReplaceIdOwner{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
