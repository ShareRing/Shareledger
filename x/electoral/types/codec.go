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
	cdc.RegisterConcrete(&MsgEnrollIdSigners{}, "electoral/EnrollIdSigners", nil)
	cdc.RegisterConcrete(&MsgRevokeIdSigners{}, "electoral/RevokeIdSigners", nil)
	cdc.RegisterConcrete(&MsgEnrollDocIssuers{}, "electoral/EnrollDocIssuers", nil)
	cdc.RegisterConcrete(&MsgRevokeDocIssuers{}, "electoral/RevokeDocIssuers", nil)
	cdc.RegisterConcrete(&MsgEnrollAccountOperators{}, "electoral/EnrollAccountOperators", nil)
	cdc.RegisterConcrete(&MsgRevokeAccountOperators{}, "electoral/RevokeAccountOperators", nil)
	cdc.RegisterConcrete(&MsgEnrollRelayers{}, "electoral/EnrollRelayers", nil)
	cdc.RegisterConcrete(&MsgRevokeRelayers{}, "electoral/RevokeRelayers", nil)
	cdc.RegisterConcrete(&MsgEnrollApprovers{}, "electoral/EnrollApprovers", nil)
	cdc.RegisterConcrete(&MsgRevokeApprovers{}, "electoral/RevokeApprovers", nil)
	cdc.RegisterConcrete(&MsgEnrollSwapManagers{}, "electoral/EnrollSwapManagers", nil)
	cdc.RegisterConcrete(&MsgRevokeSwapManagers{}, "electoral/RevokeSwapManagers", nil)
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
		&MsgEnrollIdSigners{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeIdSigners{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollDocIssuers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeDocIssuers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollAccountOperators{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeAccountOperators{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollRelayers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeRelayers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollApprovers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeApprovers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnrollSwapManagers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeSwapManagers{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
