package services

import (
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	"strings"
)

type (
	BatchSortByIDAscending []swapmoduletypes.Batch
)

func (v BatchSortByIDAscending) Len() int           { return len(v) }
func (v BatchSortByIDAscending) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v BatchSortByIDAscending) Less(i, j int) bool { return v[i].Id < v[j].Id }

const msgBatchProcessed = "batch already exists"
const msgRequestDuplicated = "request already exists"

func IsErrBatchProcessed(err error) bool {
	return strings.Contains(err.Error(), msgBatchProcessed)
}

func IsErrRequestProcessed(err error) bool {
	return strings.Contains(err.Error(), msgRequestDuplicated)
}
