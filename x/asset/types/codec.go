package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAsset{}, "asset/CreateAsset", nil)
	cdc.RegisterConcrete(&MsgUpdateAsset{}, "asset/UpdateAsset", nil)
	cdc.RegisterConcrete(&MsgDeleteAsset{}, "asset/DeleteAsset", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAsset{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateAsset{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteAsset{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
