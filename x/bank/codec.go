package bank

import (
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	msg "github.com/sharering/shareledger/x/bank/messages"
)

func RegisterCodec(cdc *wire.Codec) *wire.Codec {
	cdc.RegisterConcrete(msg.MsgSend{}, "shareledger/bank/MsgSend", nil)
	cdc.RegisterConcrete(msg.MsgCheck{}, "shareledger/bank/MsgCheck", nil)
	cdc.RegisterConcrete(msg.MsgLoad{}, "shareledger/bank/MsgLoad", nil)
	return cdc
}
