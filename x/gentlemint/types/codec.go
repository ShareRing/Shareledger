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
	cdc.RegisterConcrete(MsgLoadSHR{}, "gentlemint/LoadSHR", nil)
	cdc.RegisterConcrete(MsgBurnSHRP{}, "gentlemint/BurnSHRP", nil)
	cdc.RegisterConcrete(MsgBurnSHR{}, "gentlemint/BurnSHR", nil)
	cdc.RegisterConcrete(MsgLoadSHRP{}, "gentlemint/LoadSHRP", nil)
	cdc.RegisterConcrete(MsgBuyCent{}, "gentlemint/BuyCent", nil)
	cdc.RegisterConcrete(MsgEnrollSHRPLoaders{}, "gentlemint/EnrollSHRPLoaders", nil)
	cdc.RegisterConcrete(MsgRevokeSHRPLoaders{}, "gentlemint/RevokeSHRPLoaders", nil)
	cdc.RegisterConcrete(MsgBuySHR{}, "gentlemint/BuySHR", nil)
	cdc.RegisterConcrete(MsgSetExchange{}, "gentlemint/SetExchange", nil)
	cdc.RegisterConcrete(MsgSendSHRP{}, "gentlemint/SendSHRP", nil)
	cdc.RegisterConcrete(MsgSendSHR{}, "gentlemint/SendSHR", nil)

	cdc.RegisterConcrete(MsgEnrollIDSigners{}, "gentlemint/MsgEnrollIDSigners", nil)
	cdc.RegisterConcrete(MsgRevokeIDSigners{}, "gentlemint/MsgRevokeIDSigners", nil)

	cdc.RegisterConcrete(MsgEnrollDocIssuers{}, "gentlemint/MsgEnrollDocIssuers", nil)
	cdc.RegisterConcrete(MsgRevokeDocIssuers{}, "gentlemint/MsgRevokeDocIssuers", nil)

	cdc.RegisterConcrete(MsgEnrollAccOperators{}, "gentlemint/MsgEnrollAccOperators", nil)
	cdc.RegisterConcrete(MsgRevokeAccOperators{}, "gentlemint/MsgRevokeAccOperators", nil)
}
