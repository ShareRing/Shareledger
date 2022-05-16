package database

import (
	"context"
)

type DBRelayer interface {
	ConnectDB(ctx context.Context) error
	Disconnect(ctx context.Context) error
	InsertBatches([]Batch) error
	UpdateLatestScannedBatchId(id uint64) error
	SearchBatchByType(shareledgerID uint64, requestType Type) (*Batch, error)
	SearchBatchByStatus(networks string, status Status, nonce uint64) ([]Batch, error)
	GetBatchByTxHash(txHash string) (Batch, error)
	SetBatch(request Batch) error
	UpdateBatchesOut(shareledgerIDs []uint64, status Status) error
	SetLastScannedBlockNumber(contractAddress string, lastScannedBlockNumber int64) error
	GetSLP3Address(erc20Addr, network string) (string, error)
	GetNextPendingBatchOut(network string) (*Batch, error)
	SetLog(batchId uint64, msg string) error
	GetLastScannedBatch() (uint64, error)
}
