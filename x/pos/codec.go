package pos

import (
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	msg "github.com/sharering/shareledger/x/pos/message"
)

func RegisterCodec(cdc *wire.Codec) *wire.Codec {
	cdc.RegisterConcrete(msg.MsgCreateValidator{}, "shareledger/pos/MsgCreateValidator", nil)
	cdc.RegisterConcrete(msg.MsgEditValidator{}, "shareledger/pos/EditValidator", nil)
	cdc.RegisterConcrete(msg.MsgDelegate{}, "shareledger/pos/MsgDelegate", nil)
	cdc.RegisterConcrete(msg.MsgBeginUnbonding{}, "shareledger/pos/MsgBeginUnbonding", nil)
	cdc.RegisterConcrete(msg.MsgCompleteUnbonding{}, "shareledger/pos/MsgCompleteUnbonding", nil)
	cdc.RegisterConcrete(msg.MsgBeginRedelegate{}, "shareledger/pos/MsgBeginRedelegate", nil)
	cdc.RegisterConcrete(msg.MsgCompleteRedelegate{}, "shareledger/pos/MsgCompleteRedelegate", nil)

	cdc.RegisterConcrete(msg.MsgWithdraw{}, "shareledger/pos/MsgWithdraw", nil)
	return cdc
}
