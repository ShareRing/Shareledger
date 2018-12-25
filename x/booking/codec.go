package booking

import (
	"github.com/tendermint/go-amino"
	msg "github.com/sharering/shareledger/x/booking/messages"
)

func RegisterCodec(cdc *amino.Codec) *amino.Codec {
	cdc.RegisterConcrete(msg.MsgBook{}, "shareledger/booking/MsgBook", nil)
	cdc.RegisterConcrete(msg.MsgComplete{}, "shareledger/booking/MsgComplete", nil)
	return cdc
}
