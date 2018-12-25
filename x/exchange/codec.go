package exchange

import (
	"github.com/sharering/shareledger/x/exchange/messages"
	"github.com/tendermint/go-amino"
)

// RegisterWire registers messages into the amino.codec
func RegisterCodec(cdc *amino.Codec) *amino.Codec {
	cdc.RegisterConcrete(messages.MsgCreate{}, "shareledger/exchange/MsgCreate", nil)
	cdc.RegisterConcrete(messages.MsgRetrieve{}, "shareledger/exchange/MsgRetrieve", nil)
	cdc.RegisterConcrete(messages.MsgUpdate{}, "shareledger/exchange/MsgUpdate", nil)
	cdc.RegisterConcrete(messages.MsgDelete{}, "shareledger/exchange/MsgDelete", nil)
	cdc.RegisterConcrete(messages.MsgExchange{}, "shareledger/exchange/MsgExchange", nil)
	return cdc
}
