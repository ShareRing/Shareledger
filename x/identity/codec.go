package identity

import (
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(cdc *amino.Codec) *amino.Codec {
	cdc.RegisterConcrete(MsgIDCreate{}, "shareledger/identity/MsgIDCreate", nil)
	cdc.RegisterConcrete(MsgIDUpdate{}, "shareledger/identity/MsgIDUpdate", nil)
	cdc.RegisterConcrete(MsgIDDelete{}, "shareledger/identity/MsgIDDelete", nil)
	return cdc
}
