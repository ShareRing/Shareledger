package types

import (
	//"encoding/json"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
)

//------------------------------------------------------------------
// Tx

// Simple tx to wrap the Msg.
type SHRTx struct {
	sdk.Msg
}

// This tx only has one Msg.
func (tx SHRTx) GetMsgs() []sdk.Msg {
	return []sdk.Msg{tx.Msg}
}

// JSON decode MsgSend.
func GetTxDecoder(cdc *wire.Codec) func([]byte) (sdk.Tx, sdk.Error) {
	return func(txBytes []byte) (sdk.Tx, sdk.Error) {
		var tx sdk.Msg

		fmt.Println("TxDecoder:", txBytes)

		//err := json.Unmarshal(txBytes, &tx)
		err := cdc.UnmarshalJSON(txBytes, &tx)

		if err != nil {
			fmt.Println("Error in decoding")
			return nil, sdk.ErrTxDecode(err.Error())
		}
		fmt.Println("Decoded Tx:", tx)
		return SHRTx{tx}, nil
	}
}
