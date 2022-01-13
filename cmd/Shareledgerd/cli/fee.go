package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/fee"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func EnableAutoLoadFee() error {
	if err := tx.SetAutoLoadFee(autoLoadFee); err != nil {
		return err
	}
	return nil
}

var autoLoadFee = func(clientCtx client.Context, msgs []sdk.Msg) (status tx.AutoLoadFeeStatus, err error) {
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

	decCoin, err := sdk.ParseDecCoin(res.ShrFee)
	if err != nil {
		return
	}
	status.TotalFees = sdk.NewCoins(sdk.NewCoin(types.DenomSHR, decCoin.Amount.RoundInt()))
	if !res.SufficientFee && res.SufficientFundForFee {
		status.MsgLoadFee = types.NewMsgLoadFee(clientCtx.GetFromAddress().String(), res.ShrpCostLoadingFee)
	}
	return
}
