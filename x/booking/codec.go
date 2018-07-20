package booking
import (
	"github.com/cosmos/cosmos-sdk/wire"
	msg "github.com/sharering/shareledger/x/booking/messages"
)

func RegisterCodec(cdc *wire.Codec) *wire.Codec {
	cdc.RegisterConcrete(msg.MsgBook{}, "shareledger/booking/MsgBook", nil)
	cdc.RegisterConcrete(msg.MsgComplete{}, "shareledger/booking/MsgComplete", nil)
	return cdc
}