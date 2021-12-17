package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgLoadShr{}, "gentlemint/LoadShr", nil)
	cdc.RegisterConcrete(&MsgLoadShrp{}, "gentlemint/LoadShrp", nil)
	cdc.RegisterConcrete(&MsgBuyShr{}, "gentlemint/BuyShr", nil)
	cdc.RegisterConcrete(&MsgSendShr{}, "gentlemint/SendShr", nil)
	cdc.RegisterConcrete(&MsgBuyCent{}, "gentlemint/BuyCent", nil)
	cdc.RegisterConcrete(&MsgBurnShrp{}, "gentlemint/BurnShrp", nil)
	cdc.RegisterConcrete(&MsgSendShrp{}, "gentlemint/SendShrp", nil)
	cdc.RegisterConcrete(&MsgBurnShr{}, "gentlemint/BurnShr", nil)
	cdc.RegisterConcrete(&MsgSetExchange{}, "gentlemint/SetExchange", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLoadShr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLoadShrp{},
	)
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
		&MsgSendShr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyCent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnShrp{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendShrp{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnShr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetExchange{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
