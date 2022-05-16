package database

import "context"

type DBRelayer interface {
	ConnectDB(ctx context.Context) error
	Disconnect(ctx context.Context) error
	SearchBatchByType(shareledgerID uint64, requestType Type) (*Batch, error)
	GetBatchByTxHash(txHash string) (Batch, error)
	SetBatch(request Batch) error
	UpdateBatchesOut(shareledgerIDs []uint64, status Status) error
	SetLastScannedBlockNumber(lastScannedBlockNumer uint64) error
	GetSLP3Address(erc20Addr, network string) (string, error)
	GetNextPendingBatchOut(network string) (*Batch, error)
	SetLog(batchId uint64, msg string) error
}
