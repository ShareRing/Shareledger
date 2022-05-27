package services

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/database"
	event "github.com/sharering/shareledger/cmd/Shareledgerd/services/subscriber"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

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
			shrAmount := denom.ExponentToBase(sdk.NewIntFromBigInt(e.Amount.BigInt()), exponentNetwork)

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
				BaseAmount:  shrAmount.Amount.String(),
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
		if err := r.db.TryToBatchPendingSwapIn(network, request.DestAddress, schema.Schema.Fee.In.Amount.BigInt()); err != nil {
			return err
		}
	}
	return nil
}

func (r *Relayer) ApprovePendingBatchesIn(ctx context.Context) error {
	pendingBatchesIn, err := r.db.GetPendingBatchesIn(ctx)
	if err != nil {
		return err
	}
	for _ = range pendingBatchesIn {
		//batch.TxHashes
	}
	return nil
}
