package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func EnableAutoLoadFee() error {
	if err := tx.SetAutoLoadFee(autoLoadFee); err != nil {
		return err
	}
	return nil
}

var autoLoadFee = func(clientCtx client.Context) (status tx.AutoLoadFeeStatus, err error) {
	queryClient := types.NewQueryClient(clientCtx)
	msg := &types.QueryCheckFeesRequest{}

	res, err := queryClient.CheckFees(context.Background(), msg)
	if err != nil {
		return
	}
	status.TotalFees, err = types.ParseShrCoinsStr(res.ShrFee)
	if err != nil {
		return
	}
	if !res.SufficientFee && res.SufficientFundForFee {
		status.MsgLoadFee = types.NewMsgLoadFee(clientCtx.GetFromAddress().String(), res.ShrpCostLoadingFee)
		fmt.Println("-----------", res.String())
	}
	return
}
