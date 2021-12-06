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
	cdc.RegisterConcrete(&MsgBuyShr{}, "gentlemint/BuyShr", nil)
	cdc.RegisterConcrete(&MsgBuyShr{}, "gentlemint/BuyShr", nil)
	cdc.RegisterConcrete(&MsgSendShr{}, "gentlemint/SendShr", nil)
	cdc.RegisterConcrete(&MsgBuyCent{}, "gentlemint/BuyCent", nil)
	cdc.RegisterConcrete(&MsgBurnShrp{}, "gentlemint/BurnShrp", nil)
	cdc.RegisterConcrete(&MsgSendShrp{}, "gentlemint/SendShrp", nil)
	cdc.RegisterConcrete(&MsgBurnShr{}, "gentlemint/BurnShr", nil)
	cdc.RegisterConcrete(&MsgSetExchange{}, "gentlemint/SetExchange", nil)
	cdc.RegisterConcrete(&MsgEnrollLoaders{}, "gentlemint/EnrollLoaders", nil)
	cdc.RegisterConcrete(&MsgRevokeLoaders{}, "gentlemint/RevokeLoaders", nil)
	cdc.RegisterConcrete(&MsgEnrollIdSigner{}, "gentlemint/EnrollIdSigner", nil)
	cdc.RegisterConcrete(&MsgRevokeIdSigner{}, "gentlemint/RevokeIdSigner", nil)
	cdc.RegisterConcrete(&MsgEnrollDocIssuer{}, "gentlemint/EnrollDocIssuer", nil)
	cdc.RegisterConcrete(&MsgRevokeDocIssuer{}, "gentlemint/RevokeDocIssuer", nil)
	cdc.RegisterConcrete(&MsgEnrollAccountOperator{}, "gentlemint/EnrollAccountOperator", nil)
	cdc.RegisterConcrete(&MsgRevokeAccountOperator{}, "gentlemint/RevokeAccountOperator", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollLoaders{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeLoaders{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollIdSigner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeIdSigner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollDocIssuer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeDocIssuer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollAccountOperator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeAccountOperator{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
