package bank

import (
	msg "github.com/sharering/shareledger/x/bank/messages"
	"github.com/tendermint/go-amino"
)

func RegisterCodec(cdc *amino.Codec) *amino.Codec {
	cdc.RegisterConcrete(msg.MsgSend{}, "shareledger/bank/MsgSend", nil)
	cdc.RegisterConcrete(msg.MsgCheck{}, "shareledger/bank/MsgCheck", nil)
	cdc.RegisterConcrete(msg.MsgLoad{}, "shareledger/bank/MsgLoad", nil)
	cdc.RegisterConcrete(msg.MsgBurn{}, "shareledger/bank/MsgBurn", nil)
	return cdc
}
