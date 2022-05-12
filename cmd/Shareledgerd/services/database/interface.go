package database

import "context"

type DBRelayer interface {
	ConnectDB(ctx context.Context) error
	Disconnect(ctx context.Context) error
	GetBatchByType(shareledgerID, requestType string) (Batch, error)
	GetBatchByTxHash(txHash string) (Batch, error)
	SetBatch(request Batch) (interface{}, error)
	UpdateBatches(shareledgerIDs []uint64, status Status) error
	SetLastScannedBlockNumber(lastScannedBlockNumer uint32) error
}
