package auth

import "bitbucket.org/shareringvn/cosmos-sdk/wire"

func RegisterCodec(cdc *wire.Codec) *wire.Codec {
	cdc.RegisterConcrete(MsgNonce{}, "shareledger/auth/MsgNonce", nil)
	return cdc
}
