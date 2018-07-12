package bank

import (
	"github.com/cosmos/cosmos-sdk/wire"
	msg "github.com/sharering/shareledger/x/bank/messages"
)

func RegisterCodec(cdc *wire.Codec) *wire.Codec {
	cdc.RegisterConcrete(msg.MsgSend{}, "shareledger/MsgSend", nil)
	cdc.RegisterConcrete(msg.MsgCheck{}, "shareledger/MsgCheck", nil)
	cdc.RegisterConcrete(msg.MsgLoad{}, "shareledger/MsgLoad", nil)
	return cdc
}
