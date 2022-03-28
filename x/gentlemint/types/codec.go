package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBuyShr{}, "gentlemint/BuyShr", nil)
	cdc.RegisterConcrete(&MsgSetExchange{}, "gentlemint/SetExchange", nil)
	cdc.RegisterConcrete(&MsgSetLevelFee{}, "gentlemint/SetLevelFee", nil)
	cdc.RegisterConcrete(&MsgDeleteLevelFee{}, "gentlemint/DeleteLevelFee", nil)
	cdc.RegisterConcrete(&MsgSetActionLevelFee{}, "gentlemint/SetActionLevelFee", nil)
	cdc.RegisterConcrete(&MsgDeleteActionLevelFee{}, "gentlemint/DeleteActionLevelFee", nil)
	cdc.RegisterConcrete(&MsgLoadFee{}, "gentlemint/LoadFee", nil)
	cdc.RegisterConcrete(&MsgLoad{}, "gentlemint/Load", nil)
	cdc.RegisterConcrete(&MsgSend{}, "gentlemint/Send", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "gentlemint/Burn", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyShr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyShr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyShr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetExchange{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetLevelFee{},
		&MsgDeleteLevelFee{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetActionLevelFee{},
		&MsgDeleteActionLevelFee{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLoadFee{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLoad{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurn{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
