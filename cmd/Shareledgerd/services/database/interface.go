package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
)

type DBRelayer interface {
	ConnectDB(ctx context.Context) error
	Disconnect(ctx context.Context) error
	InsertBatchesOut([]BatchOut) error
	UpdateLatestScannedBatchId(id uint64, network string) error
	//SearchBatchByType(shareledgerID uint64, requestType BatchType) (*BatchOut, error)
	//SearchBatchByStatus(networks string, status BatchStatus) ([]BatchOut, error)
	SearchUnSyncedBatchOutByStatus(network string, status BatchStatus) ([]BatchOut, error)
	GetBatchOutByTxHash(txHash string) (*BatchOut, error)
	SetBatch(request IBatch) error
	SetBatches(batches []IBatch) error
	UpdateBatchesOut(shareledgerIDs []uint64, status BatchStatus) error
	SetBatchesOutFailed(nonceNumber uint64) error
	GetSLP3Address(erc20Addr, network string) (string, error)
	GetNextUnfinishedBatchOut(network string, offset int64) (*BatchOut, error)
	SetLog(batchId uint64, msg string) error
	SetLastScannedBlockNumber(network string, contractAddress string, lastScannedBlockNumber int64) error
	GetLastScannedBatch(network string) (uint64, error)
	GetLastScannedBlockNumber(network string, contractAddr string) (uint64, error)
	MarkBatchToSynced(sIDs []uint64) error
	InsertRequestIn(request RequestsIn) error
	GetPendingRequestsIn(network string, destAddress string) ([]RequestsIn, error)
	TryToBatchPendingSwapIn(network string, destAddress string, minFee *big.Int) error
	GetRequestIn(txHash string) (*RequestsIn, error)
	GetPendingBatchesIn(ctx context.Context) ([]BatchIn, error)
	GetRequestsInByBatchID(batchID primitive.ObjectID) ([]RequestsIn, error)
}
