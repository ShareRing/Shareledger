package asset

import (
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/sharering/shareledger/x/asset/messages"
)

// RegisterWire registers messages into the wire codec
func RegisterCodec(cdc *wire.Codec) *wire.Codec {
	cdc.RegisterConcrete(messages.MsgCreate{}, "shareledger/asset/MsgCreate", nil)
	cdc.RegisterConcrete(messages.MsgRetrieve{}, "shareledger/asset/MsgRetrieve", nil)
	cdc.RegisterConcrete(messages.MsgUpdate{}, "shareledger/asset/MsgUpdate", nil)
	cdc.RegisterConcrete(messages.MsgDelete{}, "shareledger/asset/MsgDelete", nil)
	return cdc
}