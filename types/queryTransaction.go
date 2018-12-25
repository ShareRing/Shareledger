package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ SHRTx = QueryTx{}

type QueryTx struct {
	sdk.Msg `json:"message"`
}

func NewQueryTx(msg sdk.Msg) QueryTx {
	return QueryTx{
		Msg: msg,
	}
}

func (tx QueryTx) GetMsgs() []sdk.Msg {
	return []sdk.Msg{tx.Msg}
}

func (tx QueryTx) GetMsg() sdk.Msg {
	return tx.Msg
}

func (tx QueryTx) GetSignature() SHRSignature {
	return nil
}

func (tx QueryTx) VerifySignature() bool {
	return true
}

func (tx QueryTx) GetSignBytes() []byte {
	return []byte{}
}
