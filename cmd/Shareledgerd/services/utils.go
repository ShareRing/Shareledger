package services

import swapmoduletypes "github.com/sharering/shareledger/x/swap/types"

type (
	BatchSortByIDAscending []swapmoduletypes.Batch
)

func (v BatchSortByIDAscending) Len() int           { return len(v) }
func (v BatchSortByIDAscending) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v BatchSortByIDAscending) Less(i, j int) bool { return v[i].Id < v[j].Id }
