package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/fee"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

var autoLoadFee tx.PreRunBroadcastTx = func(clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg) (nClientCtx client.Context, nTxf tx.Factory, nMsgs []sdk.Msg, err error) {
	// Not autoload fee if there is a flag fee on cli
	if !txf.Fees().IsZero() {
		return clientCtx, txf, msgs, nil
	}

	nClientCtx = clientCtx
	nTxf = txf
	nMsgs = msgs[:]

	queryClient := types.NewQueryClient(clientCtx)

	actions := make([]string, 0, len(msgs))
	for _, m := range msgs {
		actions = append(actions, fee.GetActionKey(m))
	}

	msg := &types.QueryCheckFeesRequest{
		Address: clientCtx.GetFromAddress().String(),
		Actions: actions,
	}

	res, err := queryClient.CheckFees(context.Background(), msg)

	if err != nil {
		return
	}
	if !res.SufficientFee && res.SufficientFundForFee {
		nMsgs = append([]sdk.Msg{types.NewMsgLoadFee(clientCtx.GetFromAddress().String(), *res.CostLoadingFee)}, nMsgs...)
	}
	nTxf = nTxf.WithFees(res.ConvertedFee.String())
	return
}
