package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k Keeper) In(ctx sdk.Context, msg *types.MsgIn) (rId uint64, err error) {

	req, err := k.AppendPendingRequest(ctx, types.Request{
		SrcAddr:     msg.SrcAddress,
		DestAddr:    msg.DestAddress,
		SrcNetwork:  msg.SrcNetwork,
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      msg.Amount,
		Fee:         msg.Fee,
		Status:      types.SwapStatusPending,
	})

	if err != nil {
		return 0, err
	}

	return req.Id, nil
}
