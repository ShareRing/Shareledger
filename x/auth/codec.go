package auth

import "github.com/tendermint/go-amino"

func RegisterCodec(cdc *amino.Codec) *amino.Codec {
	cdc.RegisterConcrete(MsgNonce{}, "shareledger/auth/MsgNonce", nil)
	return cdc
}
