package database

import (
	"context"
	"math/big"
)

type DBRelayer interface {
	ConnectDB(ctx context.Context) error
	Disconnect(ctx context.Context) error
	InsertBatches([]Batch) error
	UpdateLatestScannedBatchId(id uint64, network string) error
	SearchBatchByType(shareledgerID uint64, requestType BatchType) (*Batch, error)
	SearchBatchByStatus(networks string, status BatchStatus) ([]Batch, error)
	SearchUnSyncedBatchByStatus(network string, status BatchStatus) ([]Batch, error)
	GetBatchByTxHash(txHash string) (*Batch, error)
	SetBatch(request Batch) error
	SetBatches(batches []Batch) error
	UpdateBatchesOut(shareledgerIDs []uint64, status BatchStatus) error
	SetBatchesOutFailed(nonceNumber uint64) error
	GetSLP3Address(erc20Addr, network string) (string, error)
	GetNextUnfinishedBatchOut(network string, offset int64) (*Batch, error)
	SetLog(batchId uint64, msg string) error
	SetLastScannedBlockNumber(network string, contractAddress string, lastScannedBlockNumber int64) error
	GetLastScannedBatch(network string) (uint64, error)
	GetLastScannedBlockNumber(network string, contractAddr string) (uint64, error)
	MarkBatchToSynced(sIDs []uint64) error
	InsertRequestIn(request RequestsIn) error
	GetPendingRequestsIn(network string, destAddress string) ([]RequestsIn, error)
	TryToBatchPendingSwapIn(network string, destAddress string, minFee *big.Int) error
	GetRequestIn(txHash string) (*RequestsIn, error)
	GetPendingBatchesIn(ctx context.Context) ([]Batch, error)
}
