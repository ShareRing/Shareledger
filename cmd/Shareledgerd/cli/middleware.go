package cli

import "github.com/cosmos/cosmos-sdk/client/tx"

func InitMiddleWare() {
	tx.AddPreRunBroadcastTx(autoLoadFee)
	tx.AddPreRunBroadcastTx(autoConvertFee)
}
