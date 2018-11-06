package exchange

import (
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	"github.com/sharering/shareledger/x/exchange/messages"
)

// RegisterWire registers messages into the wire codec
func RegisterCodec(cdc *wire.Codec) *wire.Codec {
	cdc.RegisterConcrete(messages.MsgCreate{}, "shareledger/exchange/MsgCreate", nil)
	cdc.RegisterConcrete(messages.MsgRetrieve{}, "shareledger/exchange/MsgRetrieve", nil)
	cdc.RegisterConcrete(messages.MsgUpdate{}, "shareledger/exchange/MsgUpdate", nil)
	cdc.RegisterConcrete(messages.MsgDelete{}, "shareledger/exchange/MsgDelete", nil)
	cdc.RegisterConcrete(messages.MsgExchange{}, "shareledger/exchange/MsgExchange", nil)
	return cdc
}
