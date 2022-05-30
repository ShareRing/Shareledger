package services

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pkg/errors"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	"sync"
)

var txLock sync.Mutex

func (r *Relayer) txApproveIn(swapIds []uint64) error {
	txLock.Lock()
	defer txLock.Unlock()

	approveMsg := swapmoduletypes.NewMsgApproveIn(r.clientTx.GetFromAddress().String(), swapIds)
	if err := approveMsg.ValidateBasic(); err != nil {
		return errors.Wrap(err, "message approve in is invalid")
	}
	return tx.GenerateOrBroadcastTxCLI(r.clientTx, r.cmd.Flags(), approveMsg)
}

func (r *Relayer) txCancelBatches(ids []uint64) error {
	txLock.Lock()
	defer txLock.Unlock()

	msg := &swapmoduletypes.MsgCancelBatches{
		Creator: r.clientTx.GetFromAddress().String(),
		Ids:     ids,
	}
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	return tx.GenerateOrBroadcastTxCLI(r.clientTx, r.cmd.Flags(), msg)
}

func (r *Relayer) txSubmitRequestIn(msg swapmoduletypes.MsgRequestIn) error {
	txLock.Lock()
	defer txLock.Unlock()
	msg.Creator = r.clientTx.GetFromAddress().String()
	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	err := tx.GenerateOrBroadcastTxCLI(r.clientTx, r.cmd.Flags(), &msg)
	return err
}

// txUpdateBatch thread safe to avoid running in multiple go routine for multiple network
func (r *Relayer) txUpdateBatch(msg *swapmoduletypes.MsgUpdateBatch) error {
	txLock.Lock()
	defer txLock.Unlock()

	msg.Creator = r.clientTx.GetFromAddress().String()
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	err := tx.GenerateOrBroadcastTxCLI(r.clientTx, r.cmd.Flags(), msg)
	if err != nil {
		return err
	}
	batchRes, err := r.qClient.Batches(context.Background(), &swapmoduletypes.QueryBatchesRequest{Ids: []uint64{msg.GetBatchId()}})
	if err != nil || len(batchRes.GetBatches()) == 0 {
		return errors.Wrapf(err, "recheck the batch id %d fail", msg.GetBatchId())
	}

	if batchRes.GetBatches()[0].GetStatus() != msg.GetStatus() {
		return errors.New("update the batch status fail")
	}

	return nil
}
