package cli

import "github.com/cosmos/cosmos-sdk/client/tx"

func InitMiddleWare() {
	tx.AddPreRunBroadcastTx(autoLoadFee)
	tx.AddPreRunBroadcastTx(autoConvertFee)
}

type MiddleWareCli interface {
	AddPreRunBroadcastTx(tx.PreRunBroadcastTx)
}

func AddMiddleWare(cli MiddleWareCli) {
	cli.AddPreRunBroadcastTx(autoLoadFee)
	cli.AddPreRunBroadcastTx(autoConvertFee)
}
