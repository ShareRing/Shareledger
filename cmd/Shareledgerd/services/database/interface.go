package database

import "context"

type DBRelayer interface {
	ConnectDB(ctx context.Context) error
	Disconnect(ctx context.Context) error
	GetBatchByType(shareledgerID, requestType string) (Batch, error)
	GetBatchByTxHash(txHash string) (Batch, error)
	SetBatch(request Batch) error
	UpdateBatches(shareledgerIDs []uint64, status Status) error
	SetLastScannedBlockNumber(contractAddress string, lastScannedBlockNumber int64) error
	GetSLP3Address(erc20Addr, network string) (string, error)
}
