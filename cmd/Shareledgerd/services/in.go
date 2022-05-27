package services

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/database"
	event "github.com/sharering/shareledger/cmd/Shareledgerd/services/subscriber"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

func (r *Relayer) ProcessApproveSubmittedBatchesIn(ctx context.Context, network string) error {
	var logData []interface{}
	defer func() {
		log.Infow("process approving submitted batches in", logData)
	}()

	res, err := r.qClient.Swap(ctx, &types.QuerySwapRequest{
		Status:      types.SwapStatusPending,
		DestNetwork: types.NetworkNameShareLedger,
	})
	if err != nil {
		return err
	}
	logData = append(logData, "number pending swap in", len(res.Swaps))

	batches, err := r.db.GetSubmittedBatchesIn(network)
	logData = append(logData, "submitted_batches", len(batches))
	if err != nil {
		logData = append(logData, "get submitted batches in", err.Error())
	}
	for _, swap := range res.Swaps {
		logData = append(logData, "swap_id", swap.Id)
		res, err := r.qClient.Balance(ctx, &types.QueryBalanceRequest{})
		if err != nil {
			return errors.Wrapf(err, "check balances of swap module")
		}
		moduleBalances, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*res.Balance), false)
		if err != nil {
			log.Errorw("parse coin swap module", "err", errors.Wrapf(err, "normalize %s", res.Balance.String()))
			continue
		}

		if !moduleBalances.IsAllGTE(sdk.NewCoins(*swap.Amount)) {
			logData = append(logData, "skip approve since lacking swap balances", fmt.Sprintf("swap in amount, %s, module balance, %s", swap.Amount, moduleBalances.String()))
			continue
		}

		batchLog := database.BatchIn{
			Batch: database.Batch{
				ShareledgerID: swap.Id,
				Status:        database.BatchStatusPending,
				Type:          database.BatchTypeIn,
				TxHashes:      swap.TxHashes,
				Network:       swap.SrcNetwork,
			},
			BaseAmount: swap.Amount.String(),
			BaseFee:    swap.Fee.String(),
			DestAddr:   swap.DestAddr,
		}

		fullBatchDone := true
		var txAmount sdk.Coin
		for _, txHash := range swap.TxHashes {
			_, amount, err := r.getConfirmedTXTransfer(ctx, network, common.HexToHash(txHash))
			if err != nil {
				fullBatchDone = false
				log.Errorw("check tx hash", "err", errors.Wrapf(err, "swap_id, %s, txHash, %s", swap.Id, txHash))
				r.db.SetLog(batchLog, err.Error())
				//should not happen this case in real life...
				break
			}
			txAmount.Add(*amount)
		}
		// cover rounding number between chains.
		if !fullBatchDone || !txAmount.Amount.Sub(swap.Amount.Amount.Add(swap.Fee.Amount)).LTE(sdk.NewInt(1)) {
			err := errors.Errorf("amount batched requests, %s, is not match with contracts data, %s", swap.Amount.Add(*swap.Fee).String(), txAmount.String())
			r.db.SetLog(batchLog, err.Error())
			continue
		}

		approveMsg := types.NewMsgApproveIn(r.clientTx.GetFromAddress().String(), []uint64{swap.Id})
		if err := approveMsg.ValidateBasic(); err != nil {
			return errors.Wrap(err, "message approve in is invalid")
		}
		err = tx.GenerateOrBroadcastTxCLI(r.clientTx, r.cmd.Flags(), approveMsg)
		if err != nil {
			log.Errorw("approve swap in", "error", err.Error())
			r.db.SetLog(batchLog, err.Error())
			logData = append(logData, "error", err)
			continue
		}
	}
	return nil
}

func (r *Relayer) aProcessIn(ctx context.Context, network string) error {
	eventService, found := r.events[network]
	if !found {
		return fmt.Errorf("%s does not have event subcriber", network)
	}
	handledRequests := make([]database.RequestsIn, 0)
	errEvent := eventService.HandlerTransferEvent(ctx, func(events []event.EventTransferOutput) error {
		for _, e := range events {
			request, err := r.db.GetRequestIn(e.TxHash)
			if err != nil {
				return errors.Wrapf(err, fmt.Sprintf("get request by txHash, %s", e.TxHash))
			}
			// already handler
			if request != nil {
				// some case, the relayer was restarted, so this will scan again. In order to re-trigger processing request.
				handledRequests = append(handledRequests, *request)
				log.Warnw("request was already handled into db.", "request", request)
				continue
			}

			exponentNetwork := r.Config.Network[network].Exponent
			if exponentNetwork == 0 {
				return errors.New(fmt.Sprintf("network %s does not have exponent config", network))
			}
			baseAmount := denom.ExponentToBase(sdk.NewIntFromBigInt(e.Amount.BigInt()), exponentNetwork)

			slp3, err := r.db.GetSLP3Address(e.ToAddress, network)
			if err != nil {
				// log error and skip process this request
				log.Errorw("get slp3 address", "errEvent", err, "event", e)
				continue
			}
			if slp3 == "" {
				log.Errorw("slp3 address is not found", "eth_address", e.ToAddress, "event", e)
				continue
			}

			ri := database.RequestsIn{
				Status:      database.RequestInPending,
				TxHash:      e.TxHash,
				DestAddress: slp3,
				SrcAddress:  e.ToAddress, // ToAddress is user's ETH/BSC wallet
				BaseAmount:  baseAmount.String(),
				BatchID:     nil,
				Network:     network,
			}
			handledRequests = append(handledRequests, ri)

			if err := r.db.InsertRequestIn(ri); err != nil {
				return err
			}
		}
		return nil
	})
	if errEvent != nil {
		if len(handledRequests) == 0 {
			return errEvent
		}
		log.Errorw("handle event ", "error", errEvent, "network", network)
	}
	var errProcessing error
	if len(handledRequests) > 0 {
		errProcessing = r.ProcessPendingRequestsIn(ctx, network, handledRequests)
	}
	return errProcessing
}

func (r *Relayer) ProcessPendingRequestsIn(ctx context.Context, network string, requests []database.RequestsIn) error {
	if len(requests) == 0 {
		return nil
	}
	schema, err := r.qClient.Schema(ctx, &types.QueryGetSchemaRequest{Network: network})
	if err != nil {
		return err
	}
	if schema.Schema.Fee == nil || schema.Schema.Fee.In == nil {
		return errors.New(fmt.Sprintf("schema swap fee is nil. Network %s", network))
	}
	fee := schema.Schema.Fee.In
	if fee.Amount.LTE(sdk.NewInt(0)) {
		return errors.New(fmt.Sprintf("swap in fee should not be less than or equal to 0"))
	}
	processedMap := make(map[string]struct{})
	for _, request := range requests {
		if _, found := processedMap[request.DestAddress]; found {
		} else {
			processedMap[request.DestAddress] = struct{}{}
		}
		if err := r.db.TryToBatchPendingSwapIn(network, request.DestAddress, *schema.Schema.Fee.In); err != nil {
			return err
		}
	}
	return nil
}

// IsSubmitted check request in is submitted or not
// 0 not yet
// 1 processed partial
// 2 processed full
func (r *Relayer) IsSubmitted(ctx context.Context, batch database.BatchIn) (status int, submittedTxHash []string, err error) {
	processedRequests, err := r.qClient.RequestedIns(ctx, &types.QueryRequestedInsRequest{
		Address: batch.DestAddr,
	})
	if err != nil {
		return 0, nil, err
	}
	submittedTxHash = make([]string, 0, len(batch.TxHashes))
	for _, txHash := range batch.TxHashes {
		if _, found := processedRequests.RequestedIn.TxHashes[txHash]; found {
			submittedTxHash = append(submittedTxHash, txHash)
		}
	}
	if len(submittedTxHash) > 0 {
		status = 1
	}
	if len(submittedTxHash) == len(batch.TxHashes) {
		status = 2
	}
	return status, submittedTxHash, nil
}

func (r *Relayer) SubmitPendingBatchesIn(ctx context.Context) error {
	pendingBatchesIn, err := r.db.GetPendingBatchesIn(ctx)
	if err != nil {
		return err
	}
	for _, req := range pendingBatchesIn {
		status, submittedTxHash, err := r.IsSubmitted(ctx, req)
		fmt.Println("submittedTxHash", submittedTxHash)
		if err != nil {
			log.Errorw("check submitted batch in", "err", err.Error())
			continue
		}
		switch status {
		case 1: // there are some txHash already submitted
			// Unbatch the request
			// Update requests have submitted txHash to Batched ( already)
			// Update requests haven't submitted to pending
			// TODO: Khang
			continue
		case 2: //already submitted this full batch.
			req.Status = database.BatchStatusSubmitted
			r.db.SetBatch(req)
			continue
		}

		bAmount, err := sdk.ParseCoinNormalized(req.BaseAmount)
		if err != nil {
			return err
		}
		dAmount := sdk.NewDecCoinFromCoin(bAmount)

		bFee, err := sdk.ParseCoinNormalized(req.BaseFee)
		if err != nil {
			return err
		}
		dFee := sdk.NewDecCoinFromCoin(bFee)

		err = r.txSubmitRequestIn(types.MsgRequestIn{
			DestAddress: req.DestAddr,
			Network:     req.Network,
			Amount:      &dAmount,
			Fee:         &dFee,
			TxHashes:    req.TxHashes,
		})
		if err != nil {
			if e := r.db.SetLog(req, err.Error()); e != nil {
				log.Errorw("set log error", "logerr", e, "error", err)
			} else {
				log.Errorw("submit request has error", "error", err)
			}
			continue
		}
		req.Status = database.BatchStatusSubmitted
		r.db.SetBatch(req)
	}
	return nil
}
