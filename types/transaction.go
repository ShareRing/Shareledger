package types

import (
	//"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
)


//------------------------------------------------------------------
// Tx

// Simple tx to wrap the Msg.
type app1Tx struct {
	sdk.Msg
}

// This tx only has one Msg.
func (tx app1Tx) GetMsgs() []sdk.Msg {
	return []sdk.Msg{tx.Msg}
}

// JSON decode MsgSend.
func TxDecoder(txBytes []byte) (sdk.Tx, sdk.Error) {
	var tx sdk.Msg
	cdc := MakeCodec()

	fmt.Println("TxDecoder:", txBytes)

	//err := json.Unmarshal(txBytes, &tx)
	err := cdc.UnmarshalJSON(txBytes, &tx)
	if err != nil {
		fmt.Println("Error in decoding")
		return nil, sdk.ErrTxDecode(err.Error())
	}
	fmt.Println("Decoded Tx:", tx)
	return app1Tx{tx}, nil
}

func MakeCodec() *wire.Codec {
	cdc := wire.NewCodec()
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterConcrete(MsgSend{}, "shareledger/MsgSend", nil)
	cdc.RegisterConcrete(MsgCheck{}, "shareledger/MsgCheck", nil)
	return cdc
}