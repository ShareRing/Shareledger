package database

type DBRelayer interface {
	GetRequestByType(shareledgerID, requestType string) (Request, error)
	GetRequestByTxHash(txHash string) (Request, error)
	SetRequest(request Request)
	UpdateRequests(shareledgerIDs []uint64, status Status) error
}
