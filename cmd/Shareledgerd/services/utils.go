package services

import (
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	"math/big"
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
const msgUnderPrice = "replacement transaction underpriced"
const msgAlreadyKnown = "already known"

func IsErrBatchProcessed(err error) bool {
	return strings.Contains(err.Error(), msgBatchProcessed)
}

func IsErrRequestProcessed(err error) bool {
	return strings.Contains(err.Error(), msgRequestDuplicated)
}

func IsErrUnderPrice(err error) bool {
	return strings.Contains(err.Error(), msgUnderPrice)
}

func IsErrAlreadyKnown(err error) bool {
	return strings.Contains(err.Error(), msgAlreadyKnown)
}

func CalculateNextFee(current *big.Int, percentage float64) *big.Int {
	if current == nil {
		return nil
	}
	nextPrice := percentage*float64(current.Int64())/100 + float64(current.Int64())
	v, _ := big.NewFloat(nextPrice).Int(nil)
	return v
}
