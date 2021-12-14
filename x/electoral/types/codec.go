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
	cdc.RegisterConcrete(&MsgCreateAccState{}, "electoral/CreateAccState", nil)
	cdc.RegisterConcrete(&MsgUpdateAccState{}, "electoral/UpdateAccState", nil)
	cdc.RegisterConcrete(&MsgDeleteAccState{}, "electoral/DeleteAccState", nil)
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
		&MsgCreateAccState{},
		&MsgUpdateAccState{},
		&MsgDeleteAccState{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
