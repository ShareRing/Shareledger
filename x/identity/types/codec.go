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
	cdc.RegisterConcrete(MsgCreateId{}, "identity/CreateId", nil)
	cdc.RegisterConcrete(MsgUpdateId{}, "identity/UpdateId", nil)
	cdc.RegisterConcrete(MsgDeleteId{}, "identity/DeleteId", nil)
	cdc.RegisterConcrete(MsgEnrollIDSigners{}, "identity/EnrollIdSigners", nil)
	cdc.RegisterConcrete(MsgRevokeIDSigners{}, "identity/RevokeIdSigners", nil)
}
