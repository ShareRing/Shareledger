package asset

import (
	"github.com/tendermint/go-amino"
	"github.com/sharering/shareledger/x/asset/messages"
)

// RegisterWire registers messages into the amino.codec
func RegisterCodec(cdc *amino.Codec) *amino.Codec {
	cdc.RegisterConcrete(messages.MsgCreate{}, "shareledger/asset/MsgCreate", nil)
	cdc.RegisterConcrete(messages.MsgRetrieve{}, "shareledger/asset/MsgRetrieve", nil)
	cdc.RegisterConcrete(messages.MsgUpdate{}, "shareledger/asset/MsgUpdate", nil)
	cdc.RegisterConcrete(messages.MsgDelete{}, "shareledger/asset/MsgDelete", nil)
	return cdc
}
