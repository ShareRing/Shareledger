package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreate{}, "asset/Create", nil)
	cdc.RegisterConcrete(MsgUpdate{}, "asset/Update", nil)
	cdc.RegisterConcrete(MsgDelete{}, "asset/Delete", nil)
}
