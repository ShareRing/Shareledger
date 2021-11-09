package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgCreateId{}, "id/MsgCreateId", nil)
	cdc.RegisterConcrete(MsgCreateIdBatch{}, "id/MsgCreateIdBatch", nil)
	cdc.RegisterConcrete(MsgUpdateId{}, "id/MsgUpdateId", nil)
	cdc.RegisterConcrete(MsgReplaceIdOwner{}, "id/MsgReplaceIdOwner", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateId{},
		&MsgCreateIdBatch{},
		&MsgUpdateId{},
		&MsgReplaceIdOwner{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
