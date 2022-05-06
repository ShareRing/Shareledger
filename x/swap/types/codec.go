package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRequestOut{}, "swap/MsgRequestOut", nil)
	cdc.RegisterConcrete(&MsgApproveOut{}, "swap/MsgApproveOut", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "swap/Deposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "swap/Withdraw", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "swap/Cancel", nil)
	cdc.RegisterConcrete(&MsgReject{}, "swap/Reject", nil)
	cdc.RegisterConcrete(&MsgRequestIn{}, "swap/MsgRequestIn", nil)
	cdc.RegisterConcrete(&MsgApproveIn{}, "swap/MsgApproveIn", nil)
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateSignSchema{}, "swap/CreateSignSchema", nil)
	cdc.RegisterConcrete(&MsgUpdateSignSchema{}, "swap/UpdateSignSchema", nil)
	cdc.RegisterConcrete(&MsgDeleteSignSchema{}, "swap/DeleteSignSchema", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestOut{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveOut{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdraw{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestIn{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveIn{},
	)
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSignSchema{},
		&MsgUpdateSignSchema{},
		&MsgDeleteSignSchema{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
