package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgOut{}, "swap/Out", nil)
	cdc.RegisterConcrete(&MsgApprove{}, "swap/Approve", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "swap/Deposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "swap/Withdraw", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "swap/Cancel", nil)
	cdc.RegisterConcrete(&MsgReject{}, "swap/Reject", nil)
	cdc.RegisterConcrete(&MsgIn{}, "swap/In", nil)
	cdc.RegisterConcrete(&MsgApproveIn{}, "swap/ApproveIn", nil)
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateFormat{}, "swap/CreateFormat", nil)
	cdc.RegisterConcrete(&MsgUpdateFormat{}, "swap/UpdateFormat", nil)
	cdc.RegisterConcrete(&MsgDeleteFormat{}, "swap/DeleteFormat", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOut{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApprove{},
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
		&MsgIn{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveIn{},
	)
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateFormat{},
		&MsgUpdateFormat{},
		&MsgDeleteFormat{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
