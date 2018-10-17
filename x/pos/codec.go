package pos

import (
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	msg "github.com/sharering/shareledger/x/pos/message"
)

func RegisterCodec(cdc *wire.Codec) *wire.Codec {
	cdc.RegisterConcrete(msg.MsgCreateValidator{}, "shareledger/pos/MsgCreateValidator", nil)
	cdc.RegisterConcrete(msg.MsgDelegate{}, "shareledger/pos/MsgDelegate", nil)
	//cdc.RegisterConcrete(msg.MsgLoad{}, "shareledger/bank/MsgLoad", nil)
	return cdc
}
