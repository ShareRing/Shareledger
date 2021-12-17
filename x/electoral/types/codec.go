package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgEnrollVoter{}, "electoral/EnrollVoter", nil)
	cdc.RegisterConcrete(&MsgRevokeVoter{}, "electoral/RevokeVoter", nil)
	cdc.RegisterConcrete(&MsgEnrollLoaders{}, "electoral/EnrollLoaders", nil)
	cdc.RegisterConcrete(&MsgRevokeLoaders{}, "electoral/RevokeLoaders", nil)
	cdc.RegisterConcrete(&MsgEnrollIdSigner{}, "electoral/EnrollIdSigner", nil)
	cdc.RegisterConcrete(&MsgRevokeIdSigner{}, "electoral/RevokeIdSigner", nil)
	cdc.RegisterConcrete(&MsgEnrollDocIssuer{}, "electoral/EnrollDocIssuer", nil)
	cdc.RegisterConcrete(&MsgRevokeDocIssuer{}, "electoral/RevokeDocIssuer", nil)
	cdc.RegisterConcrete(&MsgEnrollAccountOperator{}, "electoral/EnrollAccountOperator", nil)
	cdc.RegisterConcrete(&MsgRevokeAccountOperator{}, "electoral/RevokeAccountOperator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollVoter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeVoter{},
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
