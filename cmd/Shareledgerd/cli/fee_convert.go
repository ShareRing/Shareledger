package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

var autoConvertFee tx.PreRunBroadcastTx = func(clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg) (nClientCtx client.Context, nTxf tx.Factory, nMsgs []sdk.Msg, err error) {
	// Not autoload fee if there is a flag fee on cli
	if txf.Fees().IsZero() {
		return clientCtx, txf, msgs, nil
	}

	nClientCtx = clientCtx
	nTxf = txf
	nMsgs = msgs[:]

	baseCoins, err := denom.NormalizeToBaseCoins(sdk.NewDecCoinsFromCoins(txf.Fees()...), true)
	if err != nil {
		return
	}

	nTxf = nTxf.WithFees(baseCoins.String())
	return
}
